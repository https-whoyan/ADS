package graph

type Edge[vType Vertex] struct {
	U      vType
	Weight float64
}

func NewEdge[vType Vertex](u vType, weight float64) *Edge[vType] {
	return &Edge[vType]{
		U:      u,
		Weight: weight,
	}
}
