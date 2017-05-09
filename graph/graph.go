package graph

import (
	"bytes"
	"strconv"
)

// A Graph is a simple undirected graph
type Graph struct {
	Nodes []*Node
}

// A Node is a single element of the graph.  It might contain
// some value. Also contains a list of edges.
type Node struct {
	Value int
	Edges []int
}

// New constructs a new zeroed graph with n nodes and no edges
func New(n int) *Graph {
	g := &Graph{make([]*Node, n)}
	for i := range g.Nodes {
		g.Nodes[i] = &Node{Value: i}
	}
	return g
}

// AddEdge returns true if the node was successfully added to the graph.
// returns false if the edge already exists in the graph
func (g *Graph) AddEdge(a, b int) bool {
	return (g.Nodes[a].addEdge(b) && g.Nodes[b].addEdge(a))
}

func (n *Node) addEdge(e int) bool {
	for _, v := range n.Edges {
		if v == e {
			return false
		}
	}
	n.Edges = append(n.Edges, e)
	return true
}

// RemoveNode cuts a node out of the graph by removing all edges to/from it.
func (g *Graph) RemoveNode(n int) {
	if n < 0 || n > len(g.Nodes) {
		return
	}
	for _, v := range g.Nodes[n].Edges {
		g.removeEdge(n, v)
	}
}

func (g *Graph) removeEdge(a, b int) {
	g.Nodes[a].removeEdge(b)
	g.Nodes[b].removeEdge(a)
}

func (n *Node) removeEdge(e int) {
	for i, v := range n.Edges {
		if v == e {
			n.Edges = append(n.Edges[:i], n.Edges[i+1:]...)
			return
		}
	}
}

func (g *Graph) String() string {
	var buf bytes.Buffer
	buf.WriteString("strict graph {\n")
	for _, n := range g.Nodes {
		buf.WriteString(n.String())
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (n Node) String() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(n.Value) + "; ")
	for _, e := range n.Edges {
		if e > n.Value {
			buf.WriteString(strconv.Itoa(n.Value) + " -- " + strconv.Itoa(e) + "; ")
		}
	}
	buf.WriteString("\n")
	return buf.String()
}
