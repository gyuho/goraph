package kosaraju

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.015")
	gr15 := gsd.JSONGraphT("../../../testgraph/testgraph.json", "testgraph.015")
	fmt.Println(SCC(g15, gr15))
	// [[B E A] [D C] [G F] [H]]

	g16 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.016")
	gr16 := gsd.JSONGraphT("../../../testgraph/testgraph.json", "testgraph.016")
	fmt.Println(SCC(g16, gr16))
	// [[B F G A] [D H C] [I] [E J]]
}

func Test_JSON_DFS_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.015")
	fmt.Println(DFS_SCC(g15, g15.FindVertexByID("B")))

	g16 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.016")
	fmt.Println(DFS_SCC(g16, g16.FindVertexByID("C")))
}

func Test_JSON_Contains(test *testing.T) {
	g15 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.015")
	a := g15.FindVertexByID("A")
	b := g15.FindVertexByID("B")
	ovs := a.GetOutVertices()
	if !Contains(b, ovs) {
		test.Errorf("Should contain B in OutVertices but %+v", Contains(b, ovs))
	}
}
