package graph

type GData struct {
	//  0 - не посчитано
	//  1 - да
	// -1 - нет
	Directed    byte
	Eulerian    byte
	Hamiltonian byte
}

// Unweighted Невзвешенный граф
type Unweighted struct {
	Graph      map[string][]string
	Properties *GData
}

// Weighted Взвешенный граф
type Weighted struct {
	Graph      map[string]map[string]int
	Properties *GData
}

type Abstract interface {
}

type AdjacencyMatrix struct {
	matrix [][]int
}
