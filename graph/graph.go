package graph

type GData struct {
	//  0 - нет
	//  1 - да
	// -1 - не посчитано
	Directed    byte
	Eulerian    byte
	Hamiltonian byte
	Bipartite   byte
	Weighted    byte
}

type Node struct {
	Name interface{}
	Mark byte
}

// AbstractGraph Абстрактное представление графа. Информация хранится в GData, граф задан списком смежности в виде отображения (map) вершин. Вершины заданы структурами Node.
type AbstractGraph struct {
	Graph      map[*Node]map[*Node]int
	Vertices   []*Node
	Properties *GData
}

// AdjacencyMatrix матрица смежности.
type AdjacencyMatrix struct {
	matrix [][]int
}

func (g *AbstractGraph) Clean() {
	for _, v := range g.Vertices {
		v.Mark = 0
	}
}

func New(weighted byte, graph map[string]map[string]int) *AbstractGraph {
	output := make(map[*Node]map[*Node]int, len(graph))
	vertices := make([]*Node, len(graph))
	i := 0
	for vert := range graph {
		n := &Node{Name: vert, Mark: 0}
		vertices[i] = n
		i++
	}
	for vert, list := range graph {
		var parentNode *Node
		childList := make(map[*Node]int, len(list))
		for _, v := range vertices {
			if v.Name == vert {
				parentNode = v
				output[v] = childList
				break
			}
		}
		for vertex, weight := range list {
			for _, n := range vertices {
				if n.Name == vertex {
					childList[n] = weight
				}
			}
		}
		output[parentNode] = childList
	}
	_type := GData{Weighted: weighted}
	return &AbstractGraph{Graph: output, Properties: &_type}
}
