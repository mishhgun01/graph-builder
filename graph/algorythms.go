package graph

import (
	"sort"
)

// DFS Поиск в глубину для невзвешенного графа
func (g *AbstractGraph[T]) DFS(start T, compare func(want T) bool) ([]*Node[T], bool) {
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
				g.Unmark()
				g.Vertexes = append(g.Vertexes, vertex)
				return g.Vertexes, true
			}
			vertex.Mark = 1
			g.Vertexes = append(g.Vertexes, vertex)
			searchStack = append(searchStack, g.GetAdjacentVertices(vertex)...)
		}
	}
	g.Unmark()
	return nil, false
}

// BFS Поиск в ширину
func (g *AbstractGraph[T]) BFS(start T, compare func(want T) bool) ([]*Node[T], bool) {
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
				g.Unmark()
				g.Vertexes = append(g.Vertexes, vertex)
				return g.Vertexes, true
			}
			vertex.Mark = 1
			g.Vertexes = append(g.Vertexes, vertex)
			searchQueue = append(searchQueue, g.GetAdjacentVertices(vertex)...)
		}

	}
	g.Unmark()
	return nil, false
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
	g.Unmark()
	return output
}

// Cycle - поиск цикла в графе на основе поиска в глубину (TODO НЕ ДОДЕЛАНО, нужна помощь)
func (g *AbstractGraph[T]) Cycle(start T) []*Node[T] {
	var searchStack, output []*Node[T]
	var startNode *Node[T]
	for vert := range g.Graph {
		if vert.Name == start {
			searchStack = append(searchStack, vert)
			startNode = vert
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
			if checkIfIn(g.Graph[vertex], startNode) && len(output) > 2 {
				output = append(output, startNode)
				return output
			}
		}
	}
	return nil
}

// Dijkstra Алгоритм Дейкстры поиска кратчайших путей из вершины в графе.
func (g *AbstractGraph[T]) Dijkstra(start, end T) map[*Node[T]]float64 {
	var processed []*Node[T]
	var endNode, _temp *Node[T]
	var nodeParents, vertCosts = make(map[*Node[T]]*Node[T], len(g.Graph)), make(map[*Node[T]]float64, len(g.Graph))
	for vert := range g.Graph {
		if vert.Name == start {
			for kid := range g.Graph[vert] {
				nodeParents[kid] = vert
				vertCosts[kid] = g.Graph[vert][kid]
			}
		}
		if vert.Name == end {
			endNode = vert
		}
	}
	nodeParents[endNode] = nil
	_temp = findMinimalCostNode(vertCosts, processed)
	for _temp != nil {
		cost := vertCosts[_temp]
		neighbours := g.Graph[_temp]
		for node := range neighbours {
			newCost := cost + neighbours[node]
			if vertCosts[node] > newCost {
				vertCosts[node] = newCost
				nodeParents[node] = node
			}
		}
		processed = append(processed, _temp)
		_temp = findMinimalCostNode(vertCosts, processed)
	}
	return vertCosts
}
