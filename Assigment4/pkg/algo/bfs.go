package algo

import (
	"Assigment4/pkg/graph"
)

func BFS[V Vertex, G interface{ graph.Interface[V] }](g G, startV V) []V {
	if !g.HasVertex(startV) {
		return []V{}
	}

	n := g.GetVerticesCount()
	walkQueue := make([]V, 0, n)

	used := make(map[V]bool)
	stack, nextStack := make(map[V]bool), make(map[V]bool)
	stack[startV] = true
	nextStack = make(map[V]bool)

	for len(stack) != 0 {
		// помечаю, что в стаке вершины трогать нельзя
		for v, _ := range stack {
			used[v] = true
		}
		// получаю следующую пачку вершин
		for v, _ := range stack {
			walkQueue = append(walkQueue, v)
			nextVs := g.InterfaceAdjacencyList(v)
			for nextV := range nextVs.Iterate() {
				isUsed := used[nextV]
				if isUsed {
					continue
				}
				nextStack[nextV] = true
			}
		}

		// заменяю ее
		stack = nextStack
		nextStack = make(map[V]bool)
	}

	return walkQueue
}
