package graph

type AdjacencyListInterface interface {
	initNewEdge(v Vertex, opt ...float64)
}

type GraphInterface interface {
	AddEdge(v Vertex, u Vertex, opt ...float64)
	GetVerticesCount() int
	GetEdgesCount() int
	HasVertex(v Vertex) bool
	HasEdge(v Vertex, u Vertex, opt ...float64) bool
	AdjacencyList(v Vertex) AdjacencyListInterface
}
