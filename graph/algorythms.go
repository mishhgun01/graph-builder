package graph

// DFS Поиск в глубину для невзвешенного графа
func (g *Unweighted) DFS(start string, compare func(want string) bool) (bool, string) {
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

// BFS Поиск в ширину для невзвешенного графа
func (g *Unweighted) BFS(start string, compare func(want string) bool) (bool, string) {
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
