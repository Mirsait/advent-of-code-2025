package main

type Node struct {
	Value Point
	Edges []*Node
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddNode(value Point) *Node {
	node := g.FindNode(value)
	if node != nil {
		return node
	}
	newNode := &Node{Value: value}
	g.Nodes = append(g.Nodes, newNode)
	return newNode
}

func (g *Graph) AddEdge(from, to *Node) {
	from.Edges = append(from.Edges, to)
}

func (g *Graph) FindNode(value Point) *Node {
	for _, node := range g.Nodes {
		if node.Value == value {
			return node
		}
	}
	return nil
}
