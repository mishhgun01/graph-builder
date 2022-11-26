package graph

// DFS Поиск в глубину для невзвешенного графа
func (g *Unweighted) DFS(start Node, compare func(want string) bool) (bool, *Node) {
	var searchStack []Node
	searchStack = append(searchStack, g.Graph[start]...)
	for len(searchStack) != 0 {
		var vertex = searchStack[len(searchStack)-1]
		searchStack = searchStack[:len(searchStack)-1]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				return true, &vertex
			}
			vertex.Mark = 1
			searchStack = append(searchStack, g.Graph[vertex]...)
		}

	}
	return false, nil
}

// BFS Поиск в ширину для невзвешенного графа
func (g *Unweighted) BFS(start Node, compare func(want string) bool) (bool, *Node) {
	var searchQueue []Node
	searchQueue = append(searchQueue, g.Graph[start]...)
	for len(searchQueue) != 0 {
		var vertex = searchQueue[0]
		searchQueue = searchQueue[1:]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				return true, &vertex
			}
			vertex.Mark = 1
			searchQueue = append(searchQueue, g.Graph[vertex]...)
		}

	}
	return false, nil
}
