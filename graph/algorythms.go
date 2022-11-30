package graph

//DFS Поиск в глубину для невзвешенного графа
func (g *AbstractGraph) DFS(start string, compare func(want interface{}) bool) (bool, *Node) {
	var searchStack []*Node
	for vert := range g.Graph {
		if vert.Name == start {
			searchStack = append(searchStack, vert)
			break
		}
	}
	for len(searchStack) != 0 {
		var vertex = searchStack[len(searchStack)-1]
		searchStack = searchStack[:len(searchStack)-1]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				g.Clean()
				return true, vertex
			}
			vertex.Mark = 1
			searchStack = append(searchStack, g.GetAdjacentVertices(vertex)...)
		}
	}
	return false, nil
}

// BFS Поиск в ширину
func (g *AbstractGraph) BFS(start string, compare func(want interface{}) bool) (bool, *Node) {
	var searchQueue []*Node
	for vert := range g.Graph {
		if vert.Name == start {
			searchQueue = append(searchQueue, vert)
			break
		}
	}
	for len(searchQueue) != 0 {
		var vertex = searchQueue[0]
		searchQueue = searchQueue[1:]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				return true, vertex
			}
			vertex.Mark = 1
			searchQueue = append(searchQueue, g.GetAdjacentVertices(vertex)...)
		}

	}
	return false, nil
}
