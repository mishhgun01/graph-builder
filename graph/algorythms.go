package graph

import (
	"sort"
)

// DFS Поиск в глубину для невзвешенного графа
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
	g.Clean()
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
				g.Clean()
				return true, vertex
			}
			vertex.Mark = 1
			searchQueue = append(searchQueue, g.GetAdjacentVertices(vertex)...)
		}

	}
	g.Clean()
	return false, nil
}

// Painting Раскраска графа
func (g *AbstractGraph) Painting() map[interface{}][]int {
	colors := randomColorsInRGB(len(g.Graph))
	i := 0
	var vList = make([]*Node, len(g.Graph))
	for vert := range g.Graph {
		vList[i] = vert
		i++
	}
	// сортируем список вершин по убыванию степени
	sort.Slice(vList, func(i, j int) bool {
		return vList[i].Power > vList[j].Power
	})
	i = 0
	var output = make(map[interface{}][]int, len(g.Graph))
	// проходим по списку вершин
	for _, vert := range vList {
		if vert.Mark == 1 {
			continue
		}
		output[vert.Name] = colors[i]
		vert.Mark = 1
		subVerts := g.Graph[vert]
		// и еще раз, чтобы покрасить все несмежные
		for _, vertex := range vList {
			if !checkIfIn(subVerts, vertex) && vertex.Mark != 1 {
				output[vertex.Name] = colors[i]
				vertex.Mark = 1
			}
		}
		i++
	}
	g.Clean()
	return output
}

// Cycle - поиск цикла в графе на основе поиска в глубину
func (g *AbstractGraph) Cycle(start interface{}) []*Node {
	var searchStack, output []*Node
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
			vertex.Mark = 1
			output = append(output, vertex)
			searchStack = append(searchStack, g.GetAdjacentVertices(vertex)...)
		}
		if vertex.Mark == 1 && output[0].Name == vertex.Name && len(output) > 2 {
			output = append(output, vertex)
			return output
		}
	}
	return nil
}
