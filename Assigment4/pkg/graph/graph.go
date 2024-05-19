package graph

import (
	"fmt"
	"log"
)

type graphAdjacencyList map[Vertex]bool

func (l *graphAdjacencyList) initNewEdge(v Vertex, _ ...float64) {
	(*l)[v] = true
}

type Graph struct {
	isNoDirection bool
	g             map[Vertex]graphAdjacencyList
}

func NewGraph(isNoDirection bool) *Graph {
	return &Graph{
		isNoDirection: isNoDirection,
		g:             make(map[Vertex]graphAdjacencyList),
	}
}

func (g *Graph) AddEdge(v Vertex, u Vertex, _ ...float64) {
	g.g[v][u] = true
	if g.isNoDirection {
		g.g[u][v] = true
	}
}

func (g *Graph) GetVerticesCount() int {
	return len(g.g)
}

func (g *Graph) GetEdgesCount() int {
	ans := 0
	for _, edges := range g.g {
		ans += len(edges)
	}
	if g.isNoDirection {
		return ans / 2
	}
	return ans
}

func (g *Graph) HasVertex(v Vertex) bool {
	_, isContains := g.g[v]
	return isContains
}

func (g *Graph) HasEdge(v Vertex, u Vertex, _ ...float64) bool {
	if !g.HasVertex(v) || !g.HasVertex(u) {
		return false
	}
	_, isContains := g.g[v][u]
	return isContains
}

func (g *Graph) AdjacencyList(v Vertex) AdjacencyListInterface {
	if !g.HasVertex(v) {
		return nil
	}
	return g.g[v]
}

type weightedGraphAdjacencyList []*weighedEdge

func (l *weightedGraphAdjacencyList) initNewEdge(v Vertex, opt ...float64) {
	*l = append(*l, newWeighedEdge(v, opt[0]))
}

type WeightedGraph struct {
	Graph
	g map[Vertex]weightedGraphAdjacencyList
}

func NewWeightedGraph(isNoDirection bool) *WeightedGraph {
	return &WeightedGraph{
		Graph: Graph{
			isNoDirection: isNoDirection,
		},
		g: make(map[Vertex]weightedGraphAdjacencyList),
	}
}

func (g *WeightedGraph) AddEdge(v Vertex, u Vertex, opt ...float64) {
	if len(opt) == 0 {
		log.Println("The weight of the rib is not specified. weight 0 is set")
	}
	g.g[v] = append(g.g[v], &weighedEdge{
		edge: edge{
			to: u,
		},
		weight: opt[0],
	})
	if g.isNoDirection {
		g.g[u] = append(g.g[u], &weighedEdge{
			edge: edge{
				to: u,
			},
			weight: opt[0],
		})
	}
}

func (g *WeightedGraph) PrintG() {
	fmt.Println(g.g)
}
