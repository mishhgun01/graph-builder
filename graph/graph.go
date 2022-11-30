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

// Abstract Взвешенный граф
type Abstract struct {
	Graph      map[*Node]map[*Node]int
	Properties *GData
}

type Node struct {
	Name string
	Mark byte
}

type AdjacencyMatrix struct {
	matrix [][]int
}
