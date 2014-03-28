package kosaraju

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

// Order is a bit different than ParseToGraph.
func Test_JSON_DFS(test *testing.T) {
	g4 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.004")
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

	g5 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.005")
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

func Test_JSON_DFSs(test *testing.T) {
	g4 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.004")
	g4s := DFSs(g4, g4.FindVertexByID("S"))
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

	g5 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.005")
	g5s := DFSs(g5, g5.FindVertexByID("A"))
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

func Test_JSON_DFSs_1(test *testing.T) {
	g4 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.004")
	g4s := DFSs(g4, g4.FindVertexByID("A"))
	g4c := "A → S → B → D → E → F → T → C"
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

	g5 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.005")
	g5s := DFSs(g5, g5.FindVertexByID("E"))
	g5c := "E → F → A → B → C → D"
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

func Test_JSON_DFSs_2(test *testing.T) {
	g4 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.004")
	g4s := DFSs(g4, g4.FindVertexByID("D"))
	g4c := "D → A → S → B → E → F → T → C"
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

	g5 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.005")
	g5s := DFSs(g5, g5.FindVertexByID("B"))
	g5c := "B → A → C → D → E → F"
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

func Test_JSON_DFS_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.015")
	fmt.Println(DFS_SCC(g15, g15.FindVertexByID("B")))

	g16 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.016")
	fmt.Println(DFS_SCC(g16, g16.FindVertexByID("C")))
}
