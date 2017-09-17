package graph

import (
	"bytes"
	"strconv"
)

// A Graph is a simple undirected graph
type Graph struct {
	nodes     []*Node
	sortOrder []int
}

// A Node is a single element of the graph.  It might contain
// some value. Also contains a list of edges.
type Node struct {
	Value int
	Edges map[int]struct{}
}

// New constructs a new zeroed graph with n nodes and no edges
func New(n int) *Graph {
	g := &Graph{make([]*Node, n), make([]int, n)}
	for i := range g.nodes {
		g.nodes[i] = &Node{Value: i, Edges: make(map[int]struct{})}
		g.sortOrder[i] = i
	}
	return g
}

// GetNode provides an accessor for a node without exposing the entire node
// list.
func (g *Graph) GetNode(n int) *Node {
	return g.nodes[n]
}

// Iter provides a method for iterating over the nodes without directly
// accessing the underlying structure.
// usage:
//	for i := range g.Iter() {
//		n := g.GetNode(i)
//		... do stuff
//	}
func (g *Graph) Iter() []struct{} {
	return make([]struct{}, len(g.nodes))
}

// Size returns the number of nodes in the graph
func (g *Graph) Size() int {
	return len(g.nodes)
}

// AddEdge returns true if the node was successfully added to the graph.
// returns false if the edge already exists in the graph
func (g *Graph) AddEdge(a, b int) bool {
	return (g.nodes[a].addEdge(b) && g.nodes[b].addEdge(a))
}

func (n *Node) addEdge(e int) bool {
	for v := range n.Edges {
		if v == e {
			return false
		}
	}
	n.Edges[e] = struct{}{}
	return true
}

// RemoveNode cuts a node out of the graph by removing all edges to/from it.
func (g *Graph) RemoveNode(n int) {
	if n < 0 || n > len(g.nodes) {
		return
	}
	for v := range g.nodes[n].Edges {
		g.removeEdge(n, v)
	}
}

func (g *Graph) removeEdge(a, b int) {
	g.nodes[a].removeEdge(b)
	g.nodes[b].removeEdge(a)
}

func (n *Node) removeEdge(e int) {
	for v := range n.Edges {
		if v == e {
			delete(n.Edges, v)
			return
		}
	}
}

// Degree provides the degree of the node, n.
func (n Node) Degree() int {
	return len(n.Edges)
}

func (g *Graph) String() string {
	var buf bytes.Buffer
	buf.WriteString("strict graph {\n")
	for _, n := range g.nodes {
		buf.WriteString(n.String())
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (n Node) String() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(n.Value) + "; ")
	for e := range n.Edges {
		if e > n.Value {
			buf.WriteString(strconv.Itoa(n.Value) + " -- " + strconv.Itoa(e) + "; ")
		}
	}
	buf.WriteString("\n")
	return buf.String()
}

// ByDegree implements sort.Interface for a Graph.
type ByDegree Graph

// Len is used for sorting by maximum degree
func (b ByDegree) Len() int { return len(b.nodes) }

// Less returns true if i has greater degree than j
func (b ByDegree) Less(i, j int) bool {
	iNode, jNode := b.sortOrder[i], b.sortOrder[j]
	return b.nodes[iNode].Degree() > b.nodes[jNode].Degree()
}

// Swap changes the rank indicated in the MaxDegree slice
func (b ByDegree) Swap(i, j int) {
	b.sortOrder[i], b.sortOrder[j] = b.sortOrder[j], b.sortOrder[i]
}
