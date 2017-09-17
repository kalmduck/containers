package graph

import (
	"fmt"
	"sort"
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
	g.AddEdge(1, 2)
	if _, ok := g.nodes[1].Edges[2]; !ok {
		t.Error("Edge a->b not created.")
	}
	if _, ok := g.nodes[2].Edges[1]; !ok {
		t.Error("Edge b->a not created.")
	}
	g.AddEdge(2, 1)
	if len(g.nodes[1].Edges) != 1 {
		t.Error("Edge duplicated.")
	}
	if len(g.nodes[2].Edges) != 1 {
		t.Error("Edge duplicated.")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := New(2)
	g.AddEdge(0, 1)
	g.removeEdge(0, 1)
	if len(g.nodes[0].Edges) != 0 || len(g.nodes[1].Edges) != 0 {
		t.Error("Edge not removed.")
	}
}

func TestRemoveNode(t *testing.T) {
	g := New(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
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

func TestRandomGraph(t *testing.T) {
	g := NewRandomDensityGraph(10, 100)
	fmt.Println(g)
}

func TestRandomEdgeGraph(t *testing.T) {
	g := NewRandomEdgeGraph(10, 15)
	fmt.Println(g)
}

func TestConnectivity(t *testing.T) {
	g := New(2)
	if g.PairwiseConnectivity() != 0 {
		t.Error("Non-zero connectivity for edgeless graph")
	}
	g.AddEdge(0, 1)
	if pc := g.PairwiseConnectivity(); pc != 1 {
		t.Errorf("Two-node, one edge should get pc of 2\nactual %d", pc)
	}
	g = New(3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	if pc := g.PairwiseConnectivity(); pc != 2 {
		t.Errorf("three-node, two edge should get pc of 6\nactual %d", pc)
	}
	g.AddEdge(0, 2)
	if pc := g.PairwiseConnectivity(); pc != 2 {
		t.Errorf("three-node, three-edge should get pc of 6\nactual %d", pc)
	}
}

func TestDegreeSort(t *testing.T) {
	g := New(4)
	g.AddEdge(3, 2)
	g.AddEdge(3, 1)
	sort.Sort(ByDegree(*g))
	if g.nodes[g.sortOrder[0]].Value != 3 {
		t.Errorf("Node 3 should have highest degree.\n"+
			"Sort resulted with %d having highest degree.\n", g.nodes[0].Value)
	}
	g = NewRandomEdgeGraph(20, 25)
	sort.Sort(ByDegree(*g))
	pc := g.PairwiseConnectivity()
	fmt.Printf("Pairwise Connectivity = %d\n", pc)
	fmt.Println(g)
	fmt.Println(g.sortOrder)
}
