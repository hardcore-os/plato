package bizflow

import (
	"context"
	"testing"
)

type testFlow struct {
}

func (f *testFlow) Name() FlowName {
	return "test_flow"
}
func (f *testFlow) BuildGraph(e *Engine) *Graph {
	g := e.InitGraph()
	tt := &testNode{}
	ttB := &testNodeB{}
	ttC := &testNodeC{}
	ttD := &testNodeD{}
	g.AddNode(tt).AddNode(ttB).AddNode(ttC).AddNode(ttD)
	return g
}

type testNode struct {
	in string
}

func (t *testNode) Name() NodeName {
	return "testA"
}
func (t *testNode) Deps() []NodeName {
	return nil
}
func (t *testNode) Meta() *Meta {
	return &Meta{}
}
func (t *testNode) Run(f *Graph) error {
	return nil
}

type testNodeB struct {
}

func (t *testNodeB) Name() NodeName {
	return "testB"
}
func (t *testNodeB) Deps() []NodeName {
	return []NodeName{"testA"}
}
func (t *testNodeB) Meta() *Meta {
	return &Meta{}
}
func (t *testNodeB) Run(f *Graph) error {
	return nil
}

type testNodeC struct {
}

func (t *testNodeC) Name() NodeName {
	return "testC"
}
func (t *testNodeC) Deps() []NodeName {
	return []NodeName{"testA"}
}
func (t *testNodeC) Meta() *Meta {
	return &Meta{}
}
func (t *testNodeC) Run(f *Graph) error {
	return nil
}

type testNodeD struct {
	out string
}

func (t *testNodeD) Name() NodeName {
	return "testD"
}
func (t *testNodeD) Deps() []NodeName {
	return []NodeName{"testB", "testC"}
}
func (t *testNodeD) Meta() *Meta {
	return &Meta{}
}
func (t *testNodeD) Run(g *Graph) error {
	if node := g.GetNode("testA"); node != nil {
		testA, _ := node.(*testNode)
		t.out = testA.in
	}
	return nil
}
func TestBizflow(t *testing.T) {
	// 进程初始化
	eng := NewEngine(64)
	eng.AddFlow(&testFlow{})

	// 请求到来时
	g := eng.CreateDAG("test_flow")
	testA := g.Input().(*testNode)
	testA.in = "hello bizflow"
	g.Run(context.TODO())
	testD := g.Output().(*testNodeD)
	t.Logf("TestBizflow %+v", testD.out)
}
