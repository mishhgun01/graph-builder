package graph

import (
	"math/rand"
	"time"
)

func (g *AbstractGraph) GetAdjacentVertices(n *Node) []*Node {
	keys := make([]*Node, len(g.Graph[n]))
	i := 0
	for k := range g.Graph[n] {
		keys[i] = k
		i++
	}

	return keys
}

// randomColorsInRGB - заполняет слайс предоставленной длины случайным цветом в формате RGB
func randomColorsInRGB(len int) [][]int {
	var colors = make([][]int, len)
	for i := 0; i < len; i++ {
		colors[i] = make([]int, 3)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		for j := range colors[i] {
			colors[i][j] = rand.Intn(255)
		}
	}
	return colors
}

// checkIfIn - проверяет, есть ли данная вершина в списке смежных
func checkIfIn(list map[*Node]int, node *Node) bool {
	for vert := range list {
		if node.Name == vert.Name {
			return true
		}
	}
	return false
}
