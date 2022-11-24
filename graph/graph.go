package graph

// Undirected Ненаправленный невзвешенный граф
type Undirected struct {
	Graph map[string][]string
}

// WeighedUndirected Ненаправленный взвешенный граф
type WeighedUndirected struct {
	Graph map[string]map[string]int
}

// NewUGraph Создает ненаправленный невзвешенный граф из ненаправленного взвешанного
func NewUGraph(graph *WeighedUndirected) *Undirected {
	var output = make(map[string][]string)
	for k, v := range graph.Graph {
		for val := range v {
			output[k] = append(output[k], val)
		}
	}
	return &Undirected{Graph: output}
}

// NewWGraph Создает ненаправленный взвешенный граф из ненаправленного невзвешенного. Веса равны единице.
func NewWGraph(graph *Undirected) *WeighedUndirected {
	var output = make(map[string]map[string]int)
	for k, v := range graph.Graph {
		temp := make(map[string]int)
		for _, val := range v {
			temp[val] = 1
		}
		output[k] = temp
	}
	return &WeighedUndirected{Graph: output}
}
