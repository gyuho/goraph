package sp

import (
	"github.com/gyuho/goraph/algorithm/spbf"
	"github.com/gyuho/goraph/algorithm/spd"
	"github.com/gyuho/goraph/graph/gs"
)

// SP returns the shortest path from `src` to `dst` vertex.
func SP(g *gs.Graph, src, dst *gs.Vertex) (string, bool) {
	nc := false
	edges := g.GetEdges()
	for _, edge := range *edges {
		if edge.(*gs.Edge).Weight < 0 {
			nc = true
		}
	}

	// if there is a negative edge
	// run Bellman-Ford
	if nc == true {
		a, b := spbf.SPBF(g, src, dst)
		return a, b
	}
	return spd.SPD(g, src, dst), true
}
