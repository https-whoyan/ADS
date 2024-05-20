package graph

import (
	"fmt"
	"log"
	"slices"
)

type graphAdjacencyList[vertex Vertex] []vertex

func (l *graphAdjacencyList[vertex]) initNewEdge(v vertex, _ ...float64) {
	*l = append(*l, v)
}

func (l *graphAdjacencyList[vertex]) containsEdge(v vertex, _ ...float64) bool {
	return slices.Contains(*l, v)
}

func (l *graphAdjacencyList[vertex]) Iterate() <-chan vertex {
	ch := make(chan vertex)

	go func(ch chan vertex) {
		defer close(ch)
		for _, num := range *l {
			ch <- num
		}
	}(ch)

	return ch
}

func (l *graphAdjacencyList[vertex]) GetV(index int) vertex {
	return (*l)[index]
}

type Graph[vertex Vertex] struct {
	isNoDirection bool
	g             map[vertex]*graphAdjacencyList[vertex]
}

func NewGraph[vertex Vertex](isNoDirection bool) *Graph[vertex] {
	return &Graph[vertex]{
		isNoDirection: isNoDirection,
		g:             make(map[vertex]*graphAdjacencyList[vertex]),
	}
}

func (g *Graph[Vertex]) AddEdge(v Vertex, u Vertex, _ ...float64) {
	if !g.HasVertex(v) {
		var connectedVs []Vertex
		g.g[v] = (*graphAdjacencyList[Vertex])(&connectedVs)
	}
	if !g.HasVertex(u) {
		var connectedVs []Vertex
		g.g[u] = (*graphAdjacencyList[Vertex])(&connectedVs)
	}
	g.g[v].initNewEdge(u)
	if g.isNoDirection {
		g.g[v].initNewEdge(u)
	}
}

func (g *Graph[Vertex]) GetVerticesCount() int {
	return len(g.g)
}

func (g *Graph[Vertex]) GetEdgesCount() int {
	ans := 0
	for _, edges := range g.g {
		ans += len(*edges)
	}
	if g.isNoDirection {
		return ans / 2
	}
	return ans
}

func (g *Graph[Vertex]) HasVertex(v Vertex) bool {
	_, isContains := g.g[v]
	return isContains
}

func (g *Graph[Vertex]) HasEdge(v Vertex, u Vertex, _ ...float64) bool {
	if !g.HasVertex(v) || !g.HasVertex(u) {
		return false
	}
	return g.g[v].containsEdge(u)
}

func (g *Graph[Vertex]) InterfaceAdjacencyList(v Vertex) AdjacencyListInterface[Vertex] {
	if !g.HasVertex(v) {
		return nil
	}
	return g.g[v]
}

type weightedGraphAdjacencyList[vertex Vertex] []*Edge[vertex]

func (l *weightedGraphAdjacencyList[Vertex]) initNewEdge(v Vertex, opt ...float64) {
	if len(opt) == 0 {
		*l = append(*l, NewEdge(v, 0))
		return
	}
	*l = append(*l, NewEdge(v, opt[0]))
}

func (l *weightedGraphAdjacencyList[vertex]) containsEdge(v vertex, opt ...float64) bool {
	return slices.Contains(*l, NewEdge(v, opt[0]))
}

func (l *weightedGraphAdjacencyList[vertex]) Iterate() <-chan vertex {
	ch := make(chan vertex)

	go func(ch chan vertex) {
		defer close(ch)
		for _, num := range *l {
			ch <- num.U
		}
	}(ch)

	return ch
}

func (g *Graph[Vertex]) PrintG() {
	fmt.Println(g.g)
}

type WeightedGraph[vertex Vertex] struct {
	isNoDirection bool
	g             map[vertex]*weightedGraphAdjacencyList[vertex]
}

func NewWeightedGraph[vertex Vertex](isNoDirection bool) *WeightedGraph[vertex] {
	return &WeightedGraph[vertex]{
		isNoDirection: isNoDirection,
		g:             make(map[vertex]*weightedGraphAdjacencyList[vertex]),
	}
}

func (g *WeightedGraph[Vertex]) AddEdge(v Vertex, u Vertex, opt ...float64) {
	if !g.HasVertex(v) {
		var edges []*Edge[Vertex]
		g.g[v] = (*weightedGraphAdjacencyList[Vertex])(&edges)
	}
	if !g.HasVertex(u) {
		var edges []*Edge[Vertex]
		g.g[u] = (*weightedGraphAdjacencyList[Vertex])(&edges)
	}
	if len(opt) == 0 {
		edgeStr := fmt.Sprintf("(%v, %v)", v, u)
		log.Printf("The weight of the edge %v is not specified. weight 0 is set", edgeStr)
	}
	g.g[v].initNewEdge(u, opt...)
	if g.isNoDirection {
		g.g[u].initNewEdge(v, opt...)
	}
}

func (g *WeightedGraph[Vertex]) GetVerticesCount() int {
	return len(g.g)
}

func (g *WeightedGraph[Vertex]) GetEdgesCount() int {
	ans := 0
	for _, edges := range g.g {
		ans += len(*edges)
	}
	if g.isNoDirection {
		return ans / 2
	}
	return ans
}

func (g *WeightedGraph[Vertex]) HasVertex(v Vertex) bool {
	_, isContains := g.g[v]
	return isContains
}

func (g *WeightedGraph[Vertex]) HasEdge(v Vertex, u Vertex, _ ...float64) bool {
	if !g.HasVertex(v) || !g.HasVertex(u) {
		return false
	}
	return g.g[v].containsEdge(u)
}

func (g *WeightedGraph[Vertex]) InterfaceAdjacencyList(v Vertex) AdjacencyListInterface[Vertex] {
	if !g.HasVertex(v) {
		return nil
	}
	return g.g[v]
}

func (g *WeightedGraph[vertex]) AdjacencyList(v vertex) []*Edge[vertex] {
	if !g.HasVertex(v) {
		return []*Edge[vertex]{}
	}
	return *g.g[v]
}

func (g *WeightedGraph[Vertex]) PrintG() {
	fmt.Println(g.g)
}
