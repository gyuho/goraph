package tskahn

import (
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestTSKahn(t *testing.T) {
	/*
		g6 := gs.FromJSON("../../files/testgraph.json", "testgraph.006")
		g6s, ex6 := TSKahn(g6)
		g6c := "D → E → B → C → A → F"
		if ex6 != true || g6s != g6c {
			t.Errorf("Should exist with %v and should be same but\n%v\n%v\n%v", ex6, g6s, g6c, g6.GetEdgesSize())
		}

		g7 := gs.FromJSON("../../files/testgraph.json", "testgraph.007")
		g7s, ex7 := TSKahn(g7)
		g7c := "A → B → C → D → E → F → H → G"
		if ex7 != true || g7s != g7c {
			t.Errorf("Should exist with %v and should be same but\n%v\n%v\n%v", ex7, g7s, g7c, g7.GetEdgesSize())
		}
	*/

	g8 := gs.FromJSON("../../files/testgraph.json", "testgraph.008")
	g8s, ex8 := TSKahn(g8)
	g8c := "No Topological Sort (Not a DAG, there is a cycle)"
	if ex8 != false || g8s != g8c {
		t.Errorf("Should't exist with %v and should be same but\n%v\n%v", ex8, g8s, g8c)
	}

	g9 := gs.FromJSON("../../files/testgraph.json", "testgraph.009")
	g9s, ex9 := TSKahn(g9)
	g9c := "No Topological Sort (Not a DAG, there is a cycle)"
	if ex9 != false || g9s != g9c {
		t.Errorf("Should't exist with %v and should be same but\n%v\n%v", ex9, g9s, g9c)
	}
}
