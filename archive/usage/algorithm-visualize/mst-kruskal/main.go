// +build
package main

import (
	"github.com/gyuho/goraph/algorithm/mst/kruskal"
	"github.com/gyuho/goraph/graph/gs"
)

func main() {
	g14 := gs.FromJSON("../../../files/testgraph.json", "testgraph.014")
	kruskal.VizDOTFile(g14, "test.dot")
}
