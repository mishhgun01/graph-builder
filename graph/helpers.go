package graph

import (
	"math"
	"math/rand"
	"time"
)

func (g *AbstractGraph[T]) GetAdjacentVertices(n *Node[T]) []*Node[T] {
	keys := make([]*Node[T], len(g.Graph[n]))
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
func checkIfIn[T comparable](list map[*Node[T]]float64, node *Node[T]) bool {
	for vert := range list {
		if node.Name == vert.Name {
			return true
		}
	}
	return false
}

func findMinimalCostNode[T comparable](list map[*Node[T]]float64, processed []*Node[T]) *Node[T] {
	var minimal *Node[T]
	minWeight := math.MaxFloat64
	for vert, weight := range list {
		if weight < minWeight && !contains(processed, vert) {
			minWeight = weight
			minimal = vert
		}
	}
	return minimal
}

func contains[T comparable](list []*Node[T], el *Node[T]) bool {
	for _, node := range list {
		if el == node {
			return true
		}
	}
	return false
}
