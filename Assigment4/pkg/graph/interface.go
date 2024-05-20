package graph

import "cmp"

type Vertex interface{ cmp.Ordered }

type StoragedV interface{ GetV() }

type AdjacencyListInterface[vertex Vertex] interface {
	Iterate() <-chan vertex
	initNewEdge(v vertex, opt ...float64)
	containsEdge(v vertex, opt ...float64) bool
}

type Interface[vertex Vertex] interface {
	AddEdge(v vertex, u vertex, opt ...float64)
	GetVerticesCount() int
	GetEdgesCount() int
	HasVertex(v vertex) bool
	HasEdge(v vertex, u vertex, opt ...float64) bool
	InterfaceAdjacencyList(v vertex) AdjacencyListInterface[vertex]
}
