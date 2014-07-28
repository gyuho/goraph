package prim

import (
	"strconv"
	"strings"

	"github.com/gyuho/goraph/graph/gs"
	"github.com/gyuho/goraph/graph/helper"
)

// VizDOT outputs the minimum spanning tree in DOT format.
func VizDOT(g *gs.Graph) string {
	lines := []string{}

	/*
		// map[string]map[string][]float64
		rm := g.ToMAP()
		tmpl := "\t%s -- %s [label=%v]"
		for srcNodeID, outMap := range rm {
			for outNodeID, fs := range outMap {
				lines = append(lines, fmt.Sprintf(tmpl, srcNodeID, outNodeID, fs[0]))
			}
		}
	*/

	line0 := "graph goraph {\n"

	_ = MST(g)
	for _, mvt := range *g.GetVertices() {
		if mvt.(*gs.Vertex).Prev.Len() != 0 {
			for _, vt := range *mvt.(*gs.Vertex).Prev {
				wt := g.GetEdgeWeight(mvt.(*gs.Vertex), vt.(*gs.Vertex))
				wtStr := strconv.FormatFloat(wt, 'f', -1, 64)
				lines = append(lines, "\t"+mvt.(*gs.Vertex).ID+" -- "+
					vt.(*gs.Vertex).ID+
					" [label="+wtStr+", color=blue]")
			}
		}
	}

	lineE := "\n}"

	return line0 + strings.Join(lines, "\n") + lineE
}

// VizDOTFile outputs the minimum spanning tree to a DOT format file.
func VizDOTFile(g *gs.Graph, fpath string) {
	sp := VizDOT(g)
	helper.WriteToFile(fpath, sp)
}
