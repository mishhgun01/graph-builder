package graph

// ConvWeightedToUnweighted Создает невзвешенный граф из взвешенного
func ConvWeightedToUnweighted(graph *AbstractGraph) *AbstractGraph {
	var output = make(map[*Node]map[*Node]int)
	for verts, list := range graph.Graph {
		var nodes = make(map[*Node]int, len(list))
		for vert := range list {
			nodes[vert] = 1
		}
		output[verts] = nodes
	}
	graph.Properties.Weighted = 0
	return &AbstractGraph{Graph: output, Properties: graph.Properties}
}
