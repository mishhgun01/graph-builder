package graph

// ConvWeightedToUnweighted Создает невзвешенный граф из взвешенного
func ConvWeightedToUnweighted[T comparable](graph *AbstractGraph[T]) *AbstractGraph[T] {
	var output = make(map[*Node[T]]map[*Node[T]]int)
	for verts, list := range graph.Graph {
		var nodes = make(map[*Node[T]]int, len(list))
		for vert := range list {
			nodes[vert] = 1
		}
		output[verts] = nodes
	}
	graph.Properties.Weighted = 0
	return &AbstractGraph[T]{Graph: output, Properties: graph.Properties}
}
