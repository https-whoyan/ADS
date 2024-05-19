package graph

type Vertex any

type edge struct {
	to Vertex
}

func newEdge(v Vertex) *edge {
	return &edge{
		to: v,
	}
}

type weighedEdge struct {
	*edge
	weight float64
}

func newWeighedEdge(v Vertex, weight float64) *weighedEdge {
	return &weighedEdge{
		edge: &edge{
			to: v,
		},
		weight: weight,
	}
}
