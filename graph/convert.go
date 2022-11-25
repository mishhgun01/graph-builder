package graph

// ConvWeightedToUnweighted Создает невзвешенный граф из взвешенного
func ConvWeightedToUnweighted(graph *Weighted) *Unweighted {
	var output = make(map[string][]string)
	for k, v := range graph.Graph {
		for val := range v {
			output[k] = append(output[k], val)
		}
	}
	return &Unweighted{Graph: output}
}

// ConvUnweightedToWeighted Создает взвешенный граф из невзвешенного. Веса равны единице.
func ConvUnweightedToWeighted(graph *Unweighted) *Weighted {
	var output = make(map[string]map[string]int)
	for k, v := range graph.Graph {
		temp := make(map[string]int)
		for _, val := range v {
			temp[val] = 1
		}
		output[k] = temp
	}
	return &Weighted{Graph: output}
}
