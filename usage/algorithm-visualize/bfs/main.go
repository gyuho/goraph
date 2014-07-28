package main

import (
	"github.com/gyuho/goraph/algorithm/bfs"
	"github.com/gyuho/goraph/graph/gs"
)

func main() {
	g4 := gs.FromJSON("../../../files/testgraph.json", "testgraph.004")
	bfs.VizDOTFile(g4, g4.FindVertexByID("S"), "test.dot")
}
