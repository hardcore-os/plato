package bizflow

import (
	"github.com/panjf2000/ants"
)

type Engine struct {
	flows           map[FlowName]Flow
	flowParallelNum int
	workPool        *ants.Pool
}

func NewEngine(workSize int) *Engine {
	e := &Engine{flows: map[FlowName]Flow{}}
	if wPool, err := ants.NewPool(workSize); err != nil {
		panic(err)
	} else {
		e.workPool = wPool
	}
	return e
}

func (e *Engine) AddFlow(f Flow) *Engine {
	e.flows[f.Name()] = f
	// 这里构建一次图是为了在进程初始化时就检查DAG是否合法
	f.BuildGraph(e)
	return e
}

func (e *Engine) FlowParallelNum(num int) *Engine {
	e.flowParallelNum = num
	return e
}
func (e *Engine) CreateDAG(name FlowName) *Graph {
	if f := e.getFlow(name); f != nil {
		g := f.BuildGraph(e)
		return g
	}
	return nil
}

func (e *Engine) InitGraph() *Graph {
	return &Graph{e: e, nodes: map[NodeName]*item{}, eventChan: make(chan *item, e.flowParallelNum)}
}

func (e *Engine) getFlow(name FlowName) Flow {
	if flow, ok := e.flows[name]; ok {
		return flow
	}
	return nil
}
