package graph

type GData struct {
	//  0 - не посчитано
	//  1 - да
	// -1 - нет
	Directed    byte
	Eulerian    byte
	Hamiltonian byte
	Bipartite   byte
}

// Unweighted Невзвешенный граф
type Unweighted struct {
	Graph      map[Node][]Node
	Vertices   []*Node
	Properties *GData
}

// Weighted Взвешенный граф
type Weighted struct {
	Graph      map[Node]map[Node]int
	Vertices   []*Node
	Properties *GData
}

type Node struct {
	Name string
	Mark byte
}

type Abstract interface {
}

type AdjacencyMatrix struct {
	matrix [][]int
}

func (g *Unweighted) Clean() {
	for _, v := range g.Vertices {
		v.Mark = 0
	}
}
