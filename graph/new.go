package graph

import (
	"constraints"
	"sort"
)

//func AutoFromAdjMatrix(matrix AdjacencyMatrix) Abstract {
//	var res Abstract
//	var weighted bool
//	for i := 0; i < len(matrix.matrix[0]); i++ {
//		for j := i + 1; j < len(matrix.matrix[0]); j++ {
//			if matrix.matrix[i][j] != 0 {
//				if matrix.matrix[i][j] != 1 {
//					weighted = true
//				}
//				res.Graph[strconv.Itoa(i)][strconv.Itoa(j)] = matrix.matrix[i][j]
//			}
//		}
//	}
//	if !weighted {
//		return ConvWeightedToUnweighted(&res)
//	}
//	return res
//}

func New[T constraints.Ordered](graph map[T]map[T]float64) *AbstractGraph[T] {
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

func NewAdjMatrix[T constraints.Ordered](graph map[T]map[T]float64) *AdjacencyMatrix {
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
	return &AdjacencyMatrix{output}
}
