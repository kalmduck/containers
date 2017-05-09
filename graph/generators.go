package graph

import (
	"math/rand"
	"time"
)

// NewCompleteGraph creates a complete graph with n nodes.
func NewCompleteGraph(n int) *Graph {
	g := New(n)
	for i := range g.Nodes {
		for j := 0; j < i; j++ {
			g.AddEdge(j, i)
		}
	}
	return g
}

// NewRandomGraph generates a graph with n nodes and density d.
func NewRandomGraph(n, d int) *Graph {
	g := New(n)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range g.Nodes {
		for j := 0; j < n; j++ {
			if i != j && rand.Intn(100) <= d {
				g.AddEdge(i, j)
			}
		}
	}
	return g
}
