package graph

// T  comparable - имя ноды может быть только типом, поддерживающим операции сравнения
type Node[T comparable] struct {
	Name  T
	Mark  byte
	Power int
}

// AbstractGraph Абстрактное представление графа. Информация хранится в GData, граф задан списком смежности в виде отображения (map) вершин. Вершины заданы структурами Node.
type AbstractGraph[T comparable] struct {
	Graph    map[*Node[T]]map[*Node[T]]int
	Vertexes []*Node[T]
}

// AdjacencyMatrix матрица смежности.
type AdjacencyMatrix struct {
	matrix [][]int
}

func (g *AbstractGraph[T]) Unmark() {
	for _, v := range g.Vertexes {
		v.Mark = 0
	}
}
