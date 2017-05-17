package graph

// PairwiseConnectivity computes the average number connected nodes in the graph.
func (g *Graph) PairwiseConnectivity() int {
	var connectivity int
	c := make(chan int)
	for i := range g.Nodes {
		go pCon(g, i, c)
	}
	for _ = range g.Nodes {
		connectivity += <-c
	}
	return connectivity / len(g.Nodes)
}

func pCon(g *Graph, n int, c chan int) {
	visited := make([]bool, len(g.Nodes))
	c <- pConSub(g, n, n, visited)
}

func pConSub(g *Graph, n, p int, visited []bool) int {
	visited[n] = true
	var connections int
	node := g.Nodes[n]
	for con := range node.Edges {
		if !visited[con] {
			connections++
			connections += pConSub(g, con, n, visited)
		}
	}
	return connections
}
