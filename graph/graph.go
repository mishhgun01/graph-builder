package graph

type GData struct {
	//  0 - не посчитано
	//  1 - да
	// -1 - нет
	Directed    byte
	Eulerian    byte
	Hamiltonian byte
}

// Unweighted Невзвешенный граф
type Unweighted struct {
	Graph      map[string][]string
	Properties *GData
}

// Weighed Взвешенный граф
type Weighed struct {
	Graph      map[string]map[string]int
	Properties *GData
}

// NewUGraph Создает ненаправленный невзвешенный граф из ненаправленного взвешанного
func NewUGraph(graph *Weighed) *Unweighted {
	var output = make(map[string][]string)
	for k, v := range graph.Graph {
		for val := range v {
			output[k] = append(output[k], val)
		}
	}
	return &Unweighted{Graph: output}
}

// NewWGraph Создает ненаправленный взвешенный граф из ненаправленного невзвешенного. Веса равны единице.
func NewWGraph(graph *Unweighted) *Weighed {
	var output = make(map[string]map[string]int)
	for k, v := range graph.Graph {
		temp := make(map[string]int)
		for _, val := range v {
			temp[val] = 1
		}
		output[k] = temp
	}
	return &Weighed{Graph: output}
}
