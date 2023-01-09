package graph

// ConvWeightedToUnweighted Создает невзвешенный граф из взвешенного
func ConvWeightedToUnweighted[T comparable](graph *AbstractGraph[T]) *AbstractGraph[T] {
	var output = make(map[*Node[T]]map[*Node[T]]float64)
	for verts, list := range graph.Graph {
		var nodes = make(map[*Node[T]]float64, len(list))
		for vert := range list {
			nodes[vert] = 1
		}
		output[verts] = nodes
	}
	return &AbstractGraph[T]{Graph: output}
}
