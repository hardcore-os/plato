package bizflow

type FlowName string
type Flow interface {
	Name() FlowName
	BuildGraph(e *Engine) *Graph
}
