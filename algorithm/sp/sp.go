// Package sp returns the shortest path in the graph.
// If the graph contains the negative edges, it runs
// the Bellman-Ford algorithms. Otherwise, it runs
// Dijkstra algorithm.
package sp

import (
	"github.com/gyuho/goraph/algorithm/spbf"
	"github.com/gyuho/goraph/algorithm/spd"
	"github.com/gyuho/goraph/graph/gsd"
)

func SP(g *gsd.Graph, src, dst string) (string, bool) {
	nc := false
	edges := g.GetEdges()
	for _, edge := range *edges {
		if edge.(*gsd.Edge).Weight < 0 {
			nc = true
		}
	}

	// if there is a negative edge
	// run Bellman-Ford
	if nc == true {
		a, b := spbf.SPBF(g, src, dst)
		return a, b
	} else {
		return spd.SPD(g, src, dst), true
	}
}
