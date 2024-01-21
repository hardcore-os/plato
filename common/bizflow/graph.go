package bizflow

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/bytedance/gopkg/util/logger"
)

type Graph struct {
	nodes         map[NodeName]*item
	lastNodeName  NodeName
	firstNodeName NodeName
	eventChan     chan *item
	e             *Engine
}

func (g *Graph) AddNode(node Node) *Graph {
	it := &item{node: node, nexts: map[NodeName]*item{}}
	nodeName := node.Name()
	deps := node.Deps()
	meta := node.Meta()
	g.nodes[nodeName] = it
	if len(deps) == 0 {
		if len(g.firstNodeName) == 0 {
			g.firstNodeName = nodeName
		} else {
			panic("There cannot be any independent stages except for the first node.")
		}
	}
	g.lastNodeName = nodeName
	it.depsNum = len(deps)
	it.retryNum = meta.RetryNum
	for _, parent := range deps {
		if parentItem := g.getItem(parent); parentItem != nil {
			if parent == nodeName {
				panic("The node has a dependency on itself.")
			}
			parentItem.nexts[nodeName] = it
		} else {
			panic(fmt.Sprintf("The dependent %s is not initialized. There is a cyclic dependency:[%s<-%s] "+
				"Please ensure that the order of adding nodes during buildgraph follows the precedence order.", parent, nodeName, parent))
		}
	}
	return g
}

func (g *Graph) GetNode(name NodeName) Node {
	if it := g.getItem(name); it != nil {
		event := it.getEvent()
		if event == execOK {
			return it.node
		} else {
			panic("The node you obtained has not been completed yet. Please follow the declared DAG dependency rules.")
		}
	}
	return nil
}

func (g *Graph) Input() Node {
	return g.getItem(g.firstNodeName).node
}

func (g *Graph) Output() Node {
	return g.GetNode(g.lastNodeName)
}

func (g *Graph) getItem(name NodeName) *item {
	if it, ok := g.nodes[name]; ok {
		return it
	}
	return nil
}

func (g *Graph) Run(ctx context.Context) error {
	if len(g.firstNodeName) == 0 {
		return errors.New("Graph has not been created")
	}
	firstIt := g.getItem(g.firstNodeName)
	if g.firstNodeName == g.lastNodeName {
		return firstIt.node.Run(g)
	}
	g.e.workPool.Submit(func() { g.work(firstIt) })
	for it := range g.eventChan {
		switch it.event {
		case execOK:
			g.execNext(it)
			g.tryStop(it)
		case RetryableError:
			if it.retryNum > 0 {
				it.retryNum -= 1
				g.e.workPool.Submit(func() { g.work(it) })
			}
		case NonRetryable:
			logger.CtxErrorf(ctx, "bizflow.masterRun NonRetryable.err=%+v", it.err)
			g.tryStop(it)
			return it.err
		case AbortErr:
			close(g.eventChan)
			return nil
		}
	}
	return nil
}

func (g *Graph) tryStop(it *item) {
	if g.lastNodeName == it.node.Name() {
		close(g.eventChan)
	}
}

func (g *Graph) execNext(it *item) {
	for _, itt := range it.nexts {
		itt.depsNum -= 1
		if itt.depsNum <= 0 {
			g.e.workPool.Submit(func() { g.work(itt) })
		}
	}
}
func (g *Graph) masterChanClosed() bool {
	select {
	case _, ok := <-g.eventChan:
		return !ok
	default:
		return false
	}
}
func (g *Graph) work(it *item) {
	// 先检查一下
	if g.masterChanClosed() {
		return
	}
	defer func() {
		if !g.masterChanClosed() {
			g.eventChan <- it
		}
	}()
	err := it.node.Run(g)
	if err == nil {
		it.setEvent(execOK)
		return
	}
	it.err = err
	meta := it.node.Meta()
	if meta.IsRetryErr[err] {
		it.setEvent(RetryableError)
		return
	}
	if meta.IsNonRetryErr[err] {
		it.setEvent(NonRetryable)
		return
	}
	if meta.AbortErr[err] {
		it.setEvent(AbortErr)
		return
	}
}

type eventName string

const (
	execOK         eventName = "execOK"
	RetryableError eventName = "Retryable error"
	NonRetryable   eventName = "NonRetryable error"
	AbortErr       eventName = "abort error"
)

type item struct {
	node     Node
	event    eventName
	nexts    map[NodeName]*item
	err      error
	depsNum  int
	retryNum int
	sync.RWMutex
}

func (it *item) getEvent() eventName {
	it.RLock()
	defer it.RUnlock()
	return it.event
}

func (it *item) setEvent(e eventName) {
	it.Lock()
	defer it.Unlock()
	it.event = e
}
