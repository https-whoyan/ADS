package main

import (
	"Assigment4/pkg/algo"
	"Assigment4/pkg/graph"
	"fmt"
	"math"
	"time"
)

func main() {
	// тестов не будет
	// автор сильно занят
	/// прошу простить кто читает.

	// а вообще тесты это битнуе 180+ задач на литкоде

	//.. зачем я вообще этот ассаймент писал....

	testingCap := 10

	g := graph.NewWeightedGraph[int](true)

	for i := 1; i <= testingCap; i++ {
		// что бы избежать замыкания цикла
		time.Sleep(25 * time.Millisecond)

		from1, to1 := i, i*3-2
		newWeight1 := abs(float64(i*(i-5)) - math.Sqrt(float64(i+1)))
		g.AddEdge(from1, to1, newWeight1)
		fmt.Printf("Add edge: (%v, %v, %v)\n", from1, to1, newWeight1)

		from2, to2 := i*3-2, i-6
		newWeight2 := abs(float64(i*3+18) - 188*math.Sqrt(float64(i*3)) + 1000)
		g.AddEdge(from2, to2, newWeight2)
		fmt.Printf("Add edge: (%v, %v, %v)\n", from2, to2, newWeight2)
	}

	fmt.Println("DFS: ", algo.DFS(g, 1))
	fmt.Println("BFS: ", algo.BFS(g, 1))
	fmt.Println("Dijkstra: ", algo.Dijkstra(g, 1))
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
