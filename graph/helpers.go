package graph

func (g *Abstract) GetAdjacentVertices(n *Node) []*Node {
	keys := make([]*Node, len(g.Graph[n]))
	i := 0
	for k := range g.Graph[n] {
		keys[i] = k
		i++
	}

	return keys
}
