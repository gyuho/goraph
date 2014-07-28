package kosaraju

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

// Order is a bit different than ParseToGraph.
func TestDFS(test *testing.T) {
	g4 := gs.FromJSON("../../../files/testgraph.json", "testgraph.004")
	DFS(g4)
	/*
		g4s := DFS(g4)
		g4c := "S → B → A → D → E → F → T → C"
		if g4s != g4c {
			test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
		}
	*/
	allvisited4 := true
	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited4 = false
		}
	}
	if !allvisited4 {
		test.Errorf("All vertices should be marked black")
	}

	g5 := gs.FromJSON("../../../files/testgraph.json", "testgraph.005")
	DFS(g5)
	/*
		g5s := DFS(g5)
		g5c := "A → B → C → D → E → F"
		if g5s != g5c {
			test.Errorf("Should be same but\n%v\n%v", g5s, g5c)
		}
	*/
	allvisited5 := true
	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited5 = false
		}
	}
	if !allvisited5 {
		test.Errorf("All vertices should be marked black")
	}
}

func TestDFSStart(test *testing.T) {
	g4 := gs.FromJSON("../../../files/testgraph.json", "testgraph.004")
	DFSStart(g4, g4.FindVertexByID("S"))
	/*
		g4s := DFSStart(g4, g4.FindVertexByID("S"))
		g4c := "S → B → A → D → E → F → T → C"
		if g4s != g4c {
			test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
		}
	*/
	allvisited4 := true
	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited4 = false
		}
	}
	if !allvisited4 {
		test.Errorf("All vertices should be marked black")
	}

	g5 := gs.FromJSON("../../../files/testgraph.json", "testgraph.005")
	DFSStart(g5, g5.FindVertexByID("A"))
	/*
		g5s := DFSStart(g5, g5.FindVertexByID("A"))
		g5c := "A → B → C → D → E → F"
		if g5s != g5c {
			test.Errorf("Should be same but\n%v\n%v", g5s, g5c)
		}
	*/
	allvisited5 := true
	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited5 = false
		}
	}
	if !allvisited5 {
		test.Errorf("All vertices should be marked black")
	}
}

func TestDFSStart_1(test *testing.T) {
	g4 := gs.FromJSON("../../../files/testgraph.json", "testgraph.004")
	DFSStart(g4, g4.FindVertexByID("A"))
	/*
		g4s := DFSStart(g4, g4.FindVertexByID("A"))
		g4c := "A → S → B → D → E → F → T → C"
		if g4s != g4c {
			test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
		}
	*/
	allvisited4 := true
	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited4 = false
		}
	}
	if !allvisited4 {
		test.Errorf("All vertices should be marked black")
	}

	g5 := gs.FromJSON("../../../files/testgraph.json", "testgraph.005")

	DFSStart(g5, g5.FindVertexByID("E"))
	/*
		g5s := DFSStart(g5, g5.FindVertexByID("E"))
		g5c := "E → F → A → B → C → D"
		if g5s != g5c {
			test.Errorf("Should be same but\n%v\n%v", g5s, g5c)
		}
	*/
	allvisited5 := true
	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited5 = false
		}
	}
	if !allvisited5 {
		test.Errorf("All vertices should be marked black")
	}
}

func TestDFSStart_2(test *testing.T) {
	g4 := gs.FromJSON("../../../files/testgraph.json", "testgraph.004")
	DFSStart(g4, g4.FindVertexByID("D"))
	/*
		g4s := DFSStart(g4, g4.FindVertexByID("D"))
		g4c := "D → A → S → B → E → F → T → C"
		if g4s != g4c {
			test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
		}
	*/
	allvisited4 := true
	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited4 = false
		}
	}
	if !allvisited4 {
		test.Errorf("All vertices should be marked black")
	}

	g5 := gs.FromJSON("../../../files/testgraph.json", "testgraph.005")
	DFSStart(g5, g5.FindVertexByID("B"))
	/*
		g5s := DFSStart(g5, g5.FindVertexByID("B"))
		g5c := "B → A → C → D → E → F"
		if g5s != g5c {
			test.Errorf("Should be same but\n%v\n%v", g5s, g5c)
		}
	*/
	allvisited5 := true
	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gs.Vertex).Color) {
			allvisited5 = false
		}
	}
	if !allvisited5 {
		test.Errorf("All vertices should be marked black")
	}
}

func TestDFSandSCC(t *testing.T) {
	g15 := gs.FromJSON("../../../files/testgraph.json", "testgraph.015")
	rs15 := DFSandSCC(g15, g15.FindVertexByID("B"))
	fmt.Println(rs15)
	// [A H G F E D C B]
	if len(rs15) != 8 {
		t.Errorf("expected 8 but %v", rs15)
	}

	g16 := gs.FromJSON("../../../files/testgraph.json", "testgraph.016")
	rs16 := DFSandSCC(g16, g16.FindVertexByID("C"))
	fmt.Println(rs16)
	// [H E J I D C]

	if len(rs16) != 6 {
		t.Errorf("expected 6 but %v", rs16)
	}
}
