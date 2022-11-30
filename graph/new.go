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

func New(weighted byte, graph map[string]map[string]int) *AbstractGraph {
	output := make(map[*Node]map[*Node]int, len(graph))
	vertices := make([]*Node, len(graph))
	i := 0
	for vert := range graph {
		n := &Node{Name: vert, Mark: 0}
		vertices[i] = n
		i++
	}
	for vert, list := range graph {
		var parentNode *Node
		childList := make(map[*Node]int, len(list))
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
	_type := GData{Weighted: weighted}
	return &AbstractGraph{Graph: output, Properties: &_type}
}
