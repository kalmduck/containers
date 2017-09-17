package graph

import (
	"math/rand"
	"time"
)

// NewCompleteGraph creates a complete graph with n nodes.
func NewCompleteGraph(n int) *Graph {
	g := New(n)
	for i := range g.nodes {
		for j := 0; j < i; j++ {
			g.AddEdge(j, i)
		}
	}
	return g
}

// NewRandomDensityGraph generates a graph with n nodes and density d.
func NewRandomDensityGraph(n, d int) *Graph {
	g := New(n)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range g.nodes {
		for j := 0; j < n; j++ {
			if i != j && rand.Intn(100) <= d {
				g.AddEdge(i, j)
			}
		}
	}
	return g
}

// NewRandomEdgeGraph generates a graph with n nodes and e edges
func NewRandomEdgeGraph(n, e int) *Graph {
	g := New(n)
	rand.Seed(time.Now().UTC().UnixNano())
	for e > 0 {
		a := rand.Intn(n)
		b := rand.Intn(n)
		if g.AddEdge(a, b) {
			e--
		}
	}
	return g
}
