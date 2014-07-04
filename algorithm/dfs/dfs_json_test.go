package dfs

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

// Order is a bit different than ParseToGraph.
func Test_JSON_DFS(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	g4s := DFS(g4)
	g4c := "S → B → A → D → E → F → T → C"
	if g4s != g4c {
		test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
	}

	allvisited4 := true
	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gsd.Vertex).Color) {
			allvisited4 = false
		}
	}
	if !allvisited4 {
		test.Errorf("All vertices should be marked black")
	}

	g5 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.005")
	g5s := DFS(g5)
	g5c := "A → B → C → D → E → F"
	if g5s != g5c {
		test.Errorf("Should be same but\n%v\n%v", g5s, g5c)
	}

	allvisited5 := true
	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gsd.Vertex).Color) {
			allvisited5 = false
		}
	}
	if !allvisited5 {
		test.Errorf("All vertices should be marked black")
	}
}

func Test_JSON_Path(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	rb4 := Path(g4, g4.FindVertexByID("S"), g4.FindVertexByID("F"))
	if !rb4 {
		test.Errorf("Should return true but %+v\n", rb4)
	}

	g16 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.016")
	rb16 := Path(g16, g16.FindVertexByID("C"), g16.FindVertexByID("B"))
	if rb16 {
		test.Errorf("Should return false but %+v\n", rb16)
	}

	g16i := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.016")
	rb16i := Path(g16i, g16i.FindVertexByID("I"), g16i.FindVertexByID("J"))
	if !rb16i {
		test.Errorf("Should return true but %+v\n", rb16i)
	}

	g16j := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.016")
	rb16j := Path(g16j, g16j.FindVertexByID("J"), g16j.FindVertexByID("I"))
	if rb16j {
		test.Errorf("Should return false but %+v\n", rb16j)
	}

	g16d := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.016")
	rb16d := Path(g16d, g16d.FindVertexByID("D"), g16d.FindVertexByID("E"))
	if !rb16d {
		test.Errorf("Should return true but %+v\n", rb16d)
	}
}
