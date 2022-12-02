package graph

type GData struct {
	//  0 - нет
	//  1 - да
	// -1 - не посчитано
	Directed    byte
	Eulerian    byte
	Hamiltonian byte
	Bipartite   byte
}

// T  comparable - имя ноды может быть только типом, поддерживающим операции сравнения
type Node[T comparable] struct {
	Name  T
	Mark  byte
	Power int
}

// AbstractGraph Абстрактное представление графа. Информация хранится в GData, граф задан списком смежности в виде отображения (map) вершин. Вершины заданы структурами Node.
type AbstractGraph[T comparable] struct {
	Graph      map[*Node[T]]map[*Node[T]]int
	Vertices   []*Node[T]
	Properties *GData
}

// AdjacencyMatrix матрица смежности.
type AdjacencyMatrix struct {
	matrix [][]int
}

func (g *AbstractGraph[T]) Clean() {
	for _, v := range g.Vertices {
		v.Mark = 0
	}
}
