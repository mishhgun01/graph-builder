package graph

import "constraints"

// T  comparable - имя ноды может быть только типом, поддерживающим операции сравнения
type Node[T constraints.Ordered] struct {
	Name  T
	Mark  byte
	Power int
}

// AbstractGraph Абстрактное представление графа. Информация хранится в GData, граф задан списком смежности в виде отображения (map) вершин. Вершины заданы структурами Node.
type AbstractGraph[T constraints.Ordered] struct {
	Graph    map[*Node[T]]map[*Node[T]]float64
	Vertexes []*Node[T]
}

// AdjacencyMatrix матрица смежности.
type AdjacencyMatrix struct {
	Matrix [][]float64
}

func (g *AbstractGraph[T]) Unmark() {
	for _, v := range g.Vertexes {
		v.Mark = 0
	}
}
