package bfs

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

// Order changes different every time it runs.

func TestBFS(t *testing.T) {
	g4 := gs.FromJSON("../../files/testgraph.json", "testgraph.004")
	_ = BFS(g4, g4.FindVertexByID("S"))
	/*
		g4c1 := "S → B → A → D → E → T → F → C"
		g4c2 := "S → B → D → E → A → T → F → C"
		g4c3 := "S → B → E → A → D → C → F → T"
		if g4s != g4c1 && g4s != g4c2 && g4s != g4c3 {
			t.Errorf("Should be same but\n%v\n%v\n%v", g4s, g4c1, g4c2)
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
		t.Errorf("All vertices should be marked black")
	}

	g5 := gs.FromJSON("../../files/testgraph.json", "testgraph.005")
	_ = BFS(g5, g5.FindVertexByID("A"))
	/*
		g5c1 := "A → C → F → B → D → E"
		g5c2 := "A → B → C → F → D → E"
		if g5s != g5c1 && g5s != g5c2 {
			t.Errorf("Should be same but\n%v\n%v\n%v", g5s, g5c1, g5c2)
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
		t.Errorf("All vertices should be marked black")
	}
}

func TestPath(t *testing.T) {
	g4 := gs.FromJSON("../../files/testgraph.json", "testgraph.004")
	rb4 := Path(g4, g4.FindVertexByID("S"), g4.FindVertexByID("F"))
	if !rb4 {
		t.Errorf("Should return true but %+v\n", rb4)
	}

	g16 := gs.FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16 := Path(g16, g16.FindVertexByID("C"), g16.FindVertexByID("B"))
	if rb16 {
		t.Errorf("Should return false but %+v\n", rb16)
	}

	g16i := gs.FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16i := Path(g16i, g16i.FindVertexByID("I"), g16i.FindVertexByID("J"))
	if !rb16i {
		t.Errorf("Should return true but %+v\n", rb16i)
	}

	g16j := gs.FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16j := Path(g16j, g16j.FindVertexByID("J"), g16j.FindVertexByID("I"))
	if rb16j {
		t.Errorf("Should return false but %+v\n", rb16j)
	}

	g16d := gs.FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16d := Path(g16d, g16d.FindVertexByID("D"), g16d.FindVertexByID("E"))
	if !rb16d {
		t.Errorf("Should return true but %+v\n", rb16d)
	}
}
