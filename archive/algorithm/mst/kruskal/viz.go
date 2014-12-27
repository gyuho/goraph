package kruskal

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

	mstedges, _ := MST(g)
	for _, edge := range mstedges {
		wt := g.GetEdgeWeight(edge.Src, edge.Dst)
		wtStr := strconv.FormatFloat(wt, 'f', -1, 64)
		lines = append(lines, "\t"+edge.Src.ID+" -- "+edge.Dst.ID+" [label="+wtStr+", color=blue]")
	}

	lineE := "}"

	return line0 + strings.Join(lines, "\n") + "\n" + lineE
}

// VizDOTFile outputs the minimum spanning tree to a DOT format file.
func VizDOTFile(g *gs.Graph, fpath string) {
	sp := VizDOT(g)
	helper.WriteToFile(fpath, sp)
}
