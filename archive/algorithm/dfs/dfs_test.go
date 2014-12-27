package dfs

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

// Order changes different every time it runs.

func TestDFS(t *testing.T) {
	g4 := gs.FromJSON("../../files/testgraph.json", "testgraph.004")
	_ = DFS(g4)
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
	_ = DFS(g5)
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
