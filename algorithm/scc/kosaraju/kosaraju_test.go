package kosaraju

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.015")
	gr15 := gsd.JSONGraphT("../../../example_files/testgraph.json", "testgraph.015")
	fmt.Println(SCC(g15, gr15))
	// [[B E A] [D C] [G F] [H]]

	// TODO
	// g16 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.016")
	// gr16 := gsd.JSONGraphT("../../../example_files/testgraph.json", "testgraph.016")
	// fmt.Println(SCC(g16, gr16))
	// [[B F G A] [D H C] [I] [E J]]
}

func Test_JSON_DFS_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.015")
	fmt.Println(DFS_SCC(g15, g15.FindVertexByID("B")))
	// [A H G F E D C B]

	g16 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.016")
	fmt.Println(DFS_SCC(g16, g16.FindVertexByID("C")))
	// [H E J I D C]
}
