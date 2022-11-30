package graph

// BFSCompare Поиск в ширину
func (g *Abstract) BFSCompare(start *Node, compare func(want string) bool) (bool, *Node) {
	var searchQueue []*Node
	searchQueue = append(searchQueue, g.GetAdjacentVertices(start)...)
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

// BFS Обход в ширину
func (g *Abstract) BFS(start *Node) []string {
	var searchQueue []*Node
	var res []string
	res = append(res, start.Name)
	searchQueue = append(searchQueue, g.GetAdjacentVertices(start)...)
	for len(searchQueue) != 0 {
		var vertex = searchQueue[0]
		searchQueue = searchQueue[1:]
		if vertex.Mark != 1 {
			vertex.Mark = 1
			res = append(res, vertex.Name)
			searchQueue = append(searchQueue, g.GetAdjacentVertices(vertex)...)
		}

	}
	return res
}
