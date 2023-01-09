package graph

import (
	"fmt"
	"testing"
)

func makeGraph[T comparable]() AbstractGraph[T] {
	var g AbstractGraph[T]
	g.Graph = make(map[*Node[T]]map[*Node[T]]float64, 3)
	var a, b, c, d, e, f Node[T]
	a.Name = "a"
	b.Name = "b"
	c.Name = "c"
	d.Name = "d"
	e.Name = "e"
	f.Name = "f"

	g.Graph[&a] = make(map[*Node[T]]float64, 2)
	g.Graph[&a][&b] = 1
	g.Graph[&a][&c] = 1
	g.Graph[&b] = make(map[*Node[T]]float64, 2)
	g.Graph[&b][&d] = 1
	g.Graph[&b][&e] = 1
	g.Graph[&c] = make(map[*Node[T]]float64, 1)
	g.Graph[&c][&f] = 1
	return g
}

func TestBFS(t *testing.T) {
	g := makeGraph()
	fmt.Println(g.BFS("a", func(want interface{}) bool {
		return false
	}))

}
func TestDFS(t *testing.T) {
	var g AbstractGraph
	g.Graph = make(map[*Node]map[*Node]int, 3)
	var a, b, c, d, e, f Node
	a.Name = "a"
	b.Name = "b"
	c.Name = "c"
	d.Name = "d"
	e.Name = "e"
	f.Name = "f"

	g.Graph[&a] = make(map[*Node]int, 2)
	g.Graph[&a][&b] = 1
	g.Graph[&a][&c] = 1
	g.Graph[&b] = make(map[*Node]int, 2)
	g.Graph[&b][&d] = 1
	g.Graph[&b][&e] = 1
	g.Graph[&c] = make(map[*Node]int, 1)
	g.Graph[&c][&f] = 1

	fmt.Println(g.DFS("a", func(want interface{}) bool {
		return false
	}))

}
