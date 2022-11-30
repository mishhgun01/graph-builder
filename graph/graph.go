package graph

type GData struct {
	//  0 - нет
	//  1 - да
	// -1 - не посчитано
	Directed    byte
	Eulerian    byte
	Hamiltonian byte
	Bipartite   byte
	Weighted    byte
}

type Node struct {
	Name  interface{}
	Mark  byte
	Power int
}

// AbstractGraph Абстрактное представление графа. Информация хранится в GData, граф задан списком смежности в виде отображения (map) вершин. Вершины заданы структурами Node.
type AbstractGraph struct {
	Graph      map[*Node]map[*Node]int
	Vertices   []*Node
	Properties *GData
}

// AdjacencyMatrix матрица смежности.
type AdjacencyMatrix struct {
	matrix [][]int
}

func (g *AbstractGraph) Clean() {
	for _, v := range g.Vertices {
		v.Mark = 0
	}
}
