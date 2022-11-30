package main

import (
	"fmt"
	"graph-builder/graph"
)

// Пример задания графа
func main() {
	graph1 := make(map[string]map[string]int)
	graph1["A"] = map[string]int{"B": 1, "C": 1, "D": 1}
	graph1["B"] = map[string]int{"A": 1, "F": 1}
	graph1["C"] = map[string]int{"A": 1, "E": 1}
	graph1["D"] = map[string]int{"A": 1, "F": 1}
	graph1["E"] = map[string]int{"C": 1}
	graph1["F"] = map[string]int{"B": 1, "D": 1}
	g := graph.New(0, graph1)
	_, l := g.DFS("C", func(want interface{}) bool {
		return want == "F"
	})
	fmt.Println(l)
}
