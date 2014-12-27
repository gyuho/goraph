// +build
package main

import (
	"github.com/gyuho/goraph/algorithm/dfs"
	"github.com/gyuho/goraph/graph/gs"
)

func main() {
	g4 := gs.FromJSON("../../../files/testgraph.json", "testgraph.004")
	dfs.VizDOTFile(g4, "test.dot")
}
