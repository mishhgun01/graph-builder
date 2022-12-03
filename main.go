package main

import (
	"fmt"
	"graph-builder/graph"
)

// Пример задания графа
func main() {
	graph1 := make(map[int]map[int]int)
	graph1[2] = map[int]int{}
	graph1[1] = map[int]int{2: 10, 3: 30, 4: 50, 5: 10}
	graph1[3] = map[int]int{5: 10}
	graph1[4] = map[int]int{2: 40, 3: 20}
	graph1[5] = map[int]int{1: 10, 3: 10, 4: 30}
	g := graph.New(graph1)
	res := g.Dijkstra(1, 3)
	for vert, w := range res {
		fmt.Println(vert, ":", w)
	}
}
