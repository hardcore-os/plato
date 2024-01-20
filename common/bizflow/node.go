package bizflow

type NodeName string
type Node interface {
	Name() NodeName
	Deps() []NodeName
	Meta() *Meta
	Run(g *Graph) error
}
