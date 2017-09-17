package graph

// PairwiseConnectivity computes the average number connected nodes in the graph.
func (g *Graph) PairwiseConnectivity() int {
	var connectivity int
	c := make(chan int)
	for i := range g.nodes {
		go pCon(g, i, c)
	}
	for _ = range g.nodes {
		connectivity += <-c
	}
	return connectivity / len(g.nodes)
}

func pCon(g *Graph, n int, c chan int) {
	visited := make([]bool, len(g.nodes))
	c <- pConSub(g, n, n, visited)
}

func pConSub(g *Graph, n, p int, visited []bool) int {
	visited[n] = true
	var connections int
	node := g.nodes[n]
	for con := range node.Edges {
		if !visited[con] {
			connections++
			connections += pConSub(g, con, n, visited)
		}
	}
	return connections
}
