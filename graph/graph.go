package graph

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

// A Graph is a simple undirected graph
type Graph struct {
	nodes []*Node
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
	for i := range g.nodes {
		g.nodes[i] = &Node{Value: i}
	}
	return g
}

// NewCompleteGraph creates a complete graph with n nodes.
func NewCompleteGraph(n int) *Graph {
	g := New(n)
	for i := range g.nodes {
		for j := 0; j < i; j++ {
			g.addEdge(j, i)
		}
	}
	return g
}

// NewRandomGraph generates a graph with n nodes and density d.
func NewRandomGraph(n, d int) *Graph {
	g := New(n)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range g.nodes {
		for j := 0; j < n; j++ {
			if i != j && rand.Intn(100) <= d {
				g.addEdge(i, j)
			}
		}
	}
	return g
}

// addEdge returns true if the node was successfully added to the graph.
// returns false if the edge already exists in the graph
func (g *Graph) addEdge(a, b int) bool {
	return (g.nodes[a].addEdge(b) && g.nodes[b].addEdge(a))
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
	for _, e := range n.Edges {
		buf.WriteString(strconv.Itoa(n.Value) + " -- " + strconv.Itoa(e) + "; ")
	}
	buf.WriteString("\n")
	return buf.String()
}
