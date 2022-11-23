package algorythms

// Ненаправленный невзвешенный граф
type UndirectedGraph struct {
	Graph map[string][]string
}

// Ненаправленный взвешенный граф
type WeighedUndirectedGraph struct {
	Graph map[string]map[string]int
}

// Создает ненаправленный невзвешенный граф из ненаправленного взвешанного
func NewUGraph(graph *WeighedUndirectedGraph) *UndirectedGraph {
	var output = make(map[string][]string)
	for k, v := range graph.Graph {
		for val, _ := range v {
			output[k] = append(output[k], val)
		}
	}
	return &UndirectedGraph{Graph: output}
}

// Создает ненаправленный взвешенный граф из ненаправленного невзвешенного. Веса равны единице.
func NewWGraph(graph *UndirectedGraph) *WeighedUndirectedGraph {
	var output = make(map[string]map[string]int)
	for k, v := range graph.Graph {
		temp := make(map[string]int)
		for _, val := range v {
			temp[val] = 1
		}
		output[k] = temp
	}
	return &WeighedUndirectedGraph{Graph: output}
}

// Поиск в глубину для невзвешенного графа
func (g *UndirectedGraph) DFS(start string, compare func(want string) bool) (bool, string) {
	var searchStack []string
	searchStack = append(searchStack, g.Graph[start]...)
	var searched []string

	for len(searchStack) != 0 {
		var vertex = searchStack[len(searchStack)-1]
		searchStack = searchStack[:len(searchStack)-1]
		vertexSearched := false
		for _, v := range searched {
			if v == vertex {
				vertexSearched = true
			}
			if !vertexSearched {
				if compare(vertex) {
					return true, vertex
				}
				searchStack = append(searchStack, g.Graph[vertex]...)
				searched = append(searched, vertex)
			}
		}
	}
	return false, ""
}

// Поиск в ширину для невзвешенного графа
func (g *UndirectedGraph) BFS(start string, compare func(want string) bool) (bool, string) {
	var searchQueue []string
	searchQueue = append(searchQueue, g.Graph[start]...)
	var searched []string

	for len(searchQueue) != 0 {
		var vertex = searchQueue[0]
		searchQueue = searchQueue[1:]
		vertexSearched := false
		for _, v := range searched {
			if v == vertex {
				vertexSearched = true
			}
			if !vertexSearched {
				if compare(vertex) {
					return true, vertex
				}
				searchQueue = append(searchQueue, g.Graph[vertex]...)
				searched = append(searched, vertex)
			}
		}
	}
	return false, ""
}
