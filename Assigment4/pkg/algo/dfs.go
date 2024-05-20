package algo

import (
	"Assigment4/pkg/graph"
)

func dfs[V Vertex, G interface{ graph.Interface[V] }](v V, g G, used *map[V]bool) []V {
	// слижком просто, не хочу обьяснять
	if (*used)[v] {
		return []V{}
	}

	var ans = []V{v}
	(*used)[v] = true

	nextVs := g.InterfaceAdjacencyList(v)
	for nextV := range nextVs.Iterate() {
		nextDFSVs := dfs(nextV, g, used)
		ans = append(ans, nextDFSVs...)
	}

	return ans
}

func DFS[V Vertex, G interface{ graph.Interface[V] }](g G, startV V) []V {
	used := make(map[V]bool)
	return dfs(startV, g, &used)
}
