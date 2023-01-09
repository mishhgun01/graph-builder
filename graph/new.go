package graph

import (
	"constraints"
	"sort"
)

// Инициализация из списка смежности.
func New[T comparable](graph map[T]map[T]float64) *AbstractGraph[T] {
	output := make(map[*Node[T]]map[*Node[T]]float64, len(graph))
	vertexes := make([]*Node[T], len(graph))
	g := &AbstractGraph[T]{Graph: output}
	i := 0
	for vert := range graph {
		n := &Node[T]{Name: vert, Mark: 0, Power: len(graph[vert])}
		vertexes[i] = n
		i++
	}
	for vert, list := range graph {
		var parentNode *Node[T]
		childList := make(map[*Node[T]]float64, len(list))
		for _, v := range vertexes {
			if v.Name == vert {
				parentNode = v
				g.Graph[v] = childList
				break
			}
		}
		for vertex, weight := range list {
			for _, n := range vertexes {
				if n.Name == vertex {
					childList[n] = weight
				}
			}
		}
		g.Graph[parentNode] = childList
	}
	return g
}

// Создание матрицы смежности из списка смежности.
func NewAdjMatrix[T constraints.Ordered](graph map[T]map[T]float64) *AdjacencyMatrix[T] {
	var output = make([][]float64, len(graph))
	for i := range output {
		output[i] = make([]float64, len(graph))
	}
	listOfVerts := make([]T, len(graph))
	i := 0
	for vert := range graph {
		listOfVerts[i] = vert
		i++
	}
	sort.Slice(listOfVerts, func(i, j int) bool {
		return listOfVerts[i] < listOfVerts[j]
	})
	i = 0
	for _, vert := range listOfVerts {
		list := graph[vert]
		for subvert, weight := range list {
			for j, name := range listOfVerts {
				if name == subvert {
					output[i][j] = weight
					break
				}
			}
		}
		i++
	}
	return &AdjacencyMatrix[T]{output, listOfVerts}
}

// Создание списка смежности из матрицы смежности
func OutputFromAdjMatrix[T constraints.Ordered](matrix *AdjacencyMatrix[T]) map[T]map[T]float64 {
	output := make(map[T]map[T]float64, len(matrix.Verts))
	temp := make(map[int]T, len(matrix.Verts))
	for i := 0; i < len(matrix.Verts); i++ {
		temp[i] = matrix.Verts[i]
	}
	i := 0
	for _, name := range matrix.Verts {
		listOfVerts := map[T]float64{}
		for j, weight := range matrix.Matrix[i] {
			if int(weight) != 0 {
				listOfVerts[matrix.Verts[j]] = weight
			}
		}
		output[name] = listOfVerts
		i++
	}
	return output
}
