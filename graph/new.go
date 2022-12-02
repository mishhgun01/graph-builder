package graph

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

func New[T comparable](weighted byte, graph map[T]map[T]int) *AbstractGraph[T] {
	output := make(map[*Node[T]]map[*Node[T]]int, len(graph))
	vertices := make([]*Node[T], len(graph))
	i := 0
	for vert := range graph {
		n := &Node[T]{Name: vert, Mark: 0, Power: len(graph[vert])}
		vertices[i] = n
		i++
	}
	for vert, list := range graph {
		var parentNode *Node[T]
		childList := make(map[*Node[T]]int, len(list))
		for _, v := range vertices {
			if v.Name == vert {
				parentNode = v
				output[v] = childList
				break
			}
		}
		for vertex, weight := range list {
			for _, n := range vertices {
				if n.Name == vertex {
					childList[n] = weight
				}
			}
		}
		output[parentNode] = childList
	}
	_type := GData{}
	return &AbstractGraph[T]{Graph: output, Properties: &_type}
}
