package graph

type Graph struct {
	Vertex int
	List   []*DoubleLinkedList
}

func NewGraph(v int) *Graph {
	list := make([]*DoubleLinkedList, v)

	for i := 0; i < v; i++ {
		list[i] = NewDoubleLinkedList()
	}

	return &Graph{
		Vertex: v,
		List:   list,
	}
}

func (g *Graph) AddEdge(from, to, distance int) {
	g.List[from].AddFirst(to, distance)
	// g.List[to].AddFirst(from, distance)
}

func (g *Graph) OutDegree(from int) int {
	return g.List[from].GetSize()
}

func (g *Graph) InDegree(to int) int {
	inDegree := 0
	for i := 0; i < len(g.List); i++ {
		current := g.List[i].Head
		for current != nil {
			if current.Data == to {
				inDegree++
			}
			current = current.Next
		}
	}

	return inDegree
}

func (g *Graph) HasEdge(from,to int) bool {
	list := g.List[from]
	current := list.Head

	for current != nil {
		if current.Data == to {
			return true
		}
		current = current.Next
	}

	return false
}

func (g *Graph) TotalDegree(vertex int) int {
	return g.OutDegree(vertex) + g.InDegree(vertex)
}

func (g *Graph) GetNeighboor(vertex int) []int {
	current := g.List[vertex].Head

	var neighboor []int
	for current != nil {
		neighboor = append(neighboor, current.Data)
		current = current.Next
	}

	return neighboor
}

func (g *Graph) RemoveEdge(from, to int) {
	list := g.List[from]

	current := list.Head
	i := 0

	for current != nil {
		if current.Data == to {
			list.Remove(i)
			return
		}
		current = current.Next
		i++
	}
}