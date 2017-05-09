package graph

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	g := New(5)
	if len(g.Nodes) != 5 {
		t.Error("Length is off in New().\n")
	}
}

func TestAddEdge(t *testing.T) {
	g := New(5)
	g.AddEdge(1, 2)
	if g.Nodes[1].Edges[0] != 2 {
		t.Error("Edge a->b not created.")
	}
	if g.Nodes[2].Edges[0] != 1 {
		t.Error("Edge b->a not created.")
	}
	g.AddEdge(2, 1)
	if len(g.Nodes[1].Edges) != 1 {
		t.Error("Edge duplicated.")
	}
	if len(g.Nodes[2].Edges) != 1 {
		t.Error("Edge duplicated.")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := New(2)
	g.AddEdge(0, 1)
	g.removeEdge(0, 1)
	if len(g.Nodes[0].Edges) != 0 || len(g.Nodes[1].Edges) != 0 {
		t.Error("Edge not removed.")
	}
}

func TestRemoveNode(t *testing.T) {
	g := New(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.RemoveNode(0)
	if len(g.Nodes[0].Edges) != 0 || len(g.Nodes[1].Edges) != 1 {
		t.Error("Node not removed.")
	}
	g.RemoveNode(4)
}

func TestCompleteGraph(t *testing.T) {
	g := NewCompleteGraph(20)
	fmt.Println(g)
}

func TestRandomGraph(t *testing.T) {
	g := NewRandomDensityGraph(10, 100)
	fmt.Println(g)
}

func TestRandomEdgeGraph(t *testing.T) {
	g := NewRandomEdgeGraph(10, 15)
	fmt.Println(g)
}
