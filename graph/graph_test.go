package graph

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	g := New(5)
	if len(g.nodes) != 5 {
		t.Error("Length is off in New().\n")
	}
}

func TestAddEdge(t *testing.T) {
	g := New(5)
	g.addEdge(1, 2)
	if g.nodes[1].Edges[0] != 2 {
		t.Error("Edge a->b not created.")
	}
	if g.nodes[2].Edges[0] != 1 {
		t.Error("Edge b->a not created.")
	}
	g.addEdge(2, 1)
	if len(g.nodes[1].Edges) != 1 {
		t.Error("Edge duplicated.")
	}
	if len(g.nodes[2].Edges) != 1 {
		t.Error("Edge duplicated.")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := New(2)
	g.addEdge(0, 1)
	g.removeEdge(0, 1)
	if len(g.nodes[0].Edges) != 0 || len(g.nodes[1].Edges) != 0 {
		t.Error("Edge not removed.")
	}
}

func TestRemoveNode(t *testing.T) {
	g := New(3)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.RemoveNode(0)
	if len(g.nodes[0].Edges) != 0 || len(g.nodes[1].Edges) != 1 {
		t.Error("Node not removed.")
	}
	g.RemoveNode(4)
}

func TestCompleteGraph(t *testing.T) {
	g := NewCompleteGraph(20)
	fmt.Println(g)
}

func TestRandomeGraph(t *testing.T) {
	g := NewRandomGraph(10, 100)
	fmt.Println(g)
}
