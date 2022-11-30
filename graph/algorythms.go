package graph

import "fmt"

// DFS Поиск в глубину для невзвешенного графа
func (g *AbstractGraph) DFS(start Node, compare func(want string) bool) (bool, *Node) {
	var searchStack []*Node
	for vert := range g.Graph {
		if vert.Name == start.Name {
			searchStack = append(searchStack, vert)
			break
		}
	}
	for len(searchStack) != 0 {
		var vertex = searchStack[len(searchStack)-1]
		if vertex.Mark != 1 {
			if compare(vertex.Name) {
				fmt.Println("found")
				g.Clean()
				return true, vertex
			}
			vertex.Mark = 1
			for v := range g.Graph[vertex] {
				searchStack = append(searchStack, v)
			}
		} else {
			searchStack = searchStack[:len(searchStack)-1]
		}
	}
	g.Clean()
	return false, nil
}

//
//// BFS Поиск в ширину для невзвешенного графа
//func (g *Unweighted) BFS(start Node, compare func(want string) bool) (bool, *Node) {
//	var searchQueue []Node
//	searchQueue = append(searchQueue, g.Graph[start]...)
//	for len(searchQueue) != 0 {
//		var vertex = searchQueue[0]
//		searchQueue = searchQueue[1:]
//		if vertex.Mark != 1 {
//			if compare(vertex.Name) {
//				g.Clean()
//				return true, &vertex
//			}
//			vertex.Mark = 1
//			searchQueue = append(searchQueue, g.Graph[vertex]...)
//		}
//
//	}
//	g.Clean()
//	return false, nil
//}
