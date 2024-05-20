package algo

type heapNode[vType Vertex] struct {
	v        vType
	priority float64
}

type GraphHeap[vType Vertex] []*heapNode[vType]

func (m *GraphHeap[vType]) Len() int {
	return len(*m)
}

func (m *GraphHeap[vType]) Less(i, j int) bool {
	return (*m)[i].priority < (*m)[j].priority
}

func (m *GraphHeap[vType]) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *GraphHeap[vType]) Push(x any) {
	*m = append(*m, x.(*heapNode[vType]))
}

func (m *GraphHeap[vType]) Pop() any {
	n := m.Len()
	returned := (*m)[n-1]
	*m = (*m)[:n-1]
	return returned
}

func (m *GraphHeap[vType]) Update() {}
