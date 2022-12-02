package graph

import (
	"sort"
)

// DFS Поиск в глубину для невзвешенного графа
func (g *AbstractGraph[T]) DFS(start T, compare func(want T) bool) (bool, *Node[T]) {
	var searchStack []*Node[T]
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
func (g *AbstractGraph[T]) BFS(start T, compare func(want T) bool) (bool, *Node[T]) {
	var searchQueue []*Node[T]
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
func (g *AbstractGraph[T]) Painting() map[T][]int {
	colors := randomColorsInRGB(len(g.Graph))
	i := 0
	var vList = make([]*Node[T], len(g.Graph))
	for vert := range g.Graph {
		vList[i] = vert
		i++
	}
	// сортируем список вершин по убыванию степени
	sort.Slice(vList, func(i, j int) bool {
		return vList[i].Power > vList[j].Power
	})
	i = 0
	var output = make(map[T][]int, len(g.Graph))
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
