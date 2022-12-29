package main

import (
	"fmt"
	"graph-builder/graph"
)

// Пример задания графа
func main() {
	graph1 := make(map[float64]map[float64]float64)
	graph1[2] = map[float64]float64{}
	graph1[1] = map[float64]float64{2: 10, 3: 30, 4: 50, 5: 10}
	graph1[3] = map[float64]float64{5: 10}
	graph1[4] = map[float64]float64{2: 40, 3: 20}
	graph1[5] = map[float64]float64{1: 10, 3: 10, 4: 30}
	g := graph.NewAdjMatrix(graph1)
	for _, w := range g.Matrix {
		fmt.Println(w)
	}
}
