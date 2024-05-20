package algo

import (
	"Assigment4/pkg/graph"
	"container/heap"
	"math"
)

type V Vertex

func initDistances[V Vertex, G *graph.WeightedGraph[V]](pG G, start V) map[V]float64 {
	distances := make(map[V]float64)
	var g = &graph.WeightedGraph[V]{}
	*g = *pG
	keys := BFS(g, start)
	for _, key := range keys {
		distances[key] = math.MaxFloat32
	}

	distances[start] = 0
	return distances
}

func Dijkstra[V graph.Vertex, G *graph.WeightedGraph[V]](g G, startV V) map[V]float64 {
	distances := initDistances(g, startV)

	rowPq := make(GraphHeap[V], 0, (*g).GetVerticesCount())
	pq := &rowPq
	heap.Init(pq)
	heap.Push(pq, &heapNode[V]{
		v:        startV,
		priority: 0,
	})

	for pq.Len() > 0 {
		minNode := pq.Pop().(*heapNode[V])

		iterableV := minNode.v
		checkedNextNodes := (*g).AdjacencyList(iterableV)
		for _, nextEdge := range checkedNextNodes {
			updatedV := nextEdge.U
			nextDest := distances[iterableV] + nextEdge.Weight

			if nextDest < distances[updatedV] {
				distances[updatedV] = nextDest
				heap.Push(pq, &heapNode[V]{
					v:        updatedV,
					priority: nextDest,
				})
			}
		}
		//...

	}
	return distances
}
