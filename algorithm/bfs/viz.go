package bfs

import (
	"fmt"

	"strings"

	"github.com/gyuho/goraph/graph/gs"
	"github.com/gyuho/goraph/graph/helper"
)

// VizDOT outputs the shortest path in DOT format.
func VizDOT(g *gs.Graph, src *gs.Vertex) string {
	// map[string]map[string][]float64
	rm := g.ToMAP()
	tmpl := "\t%s -> %s [label=%v]"
	lines := []string{}
	for srcNodeID, outMap := range rm {
		for outNodeID, fs := range outMap {
			lines = append(lines, fmt.Sprintf(tmpl, srcNodeID, outNodeID, fs[0]))
		}
	}

	line0 := "digraph goraph {\n"

	path := BFS(g, src)
	// path = strings.Replace(path, "→", "--", -1)
	path += " [label=BFS, color=blue]\n"

	lineE := "}"

	return line0 + strings.Join(lines, "\n") + "\n" + path + lineE
}

// VizDOTFile outputs the shortest path to a DOT format file.
func VizDOTFile(g *gs.Graph, src *gs.Vertex, fpath string) {
	sp := VizDOT(g, src)
	helper.WriteToFile(fpath, sp)
}
