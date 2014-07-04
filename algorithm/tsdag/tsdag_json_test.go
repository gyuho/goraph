package tsdag

import (
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_TSDAG(test *testing.T) {
	g6 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.006")
	g6s, ex6 := TSDAG(g6)
	g6c := "E → D → C → B → A → F"
	if ex6 != true || g6s != g6c {
		test.Errorf("Should exist with %v and should be same but\n%v\n%v", ex6, g6s, g6c)
	}

	g7 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.007")
	g7s, ex7 := TSDAG(g7)
	g7c := "C → B → D → F → A → H → E → G"
	if ex7 != true || g7s != g7c {
		test.Errorf("Should exist with %v and should be same but\n%v\n%v", ex7, g7s, g7c)
	}

	g8 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.008")
	g8s, ex8 := TSDAG(g8)
	g8c := "No Topological Sort (Not a DAG, there is a cycle)"
	if ex8 != false || g8s != g8c {
		test.Errorf("Should't exist with %v and should be same but\n%v\n%v", ex8, g8s, g8c)
	}

	g9 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.009")
	g9s, ex9 := TSDAG(g9)
	g9c := "No Topological Sort (Not a DAG, there is a cycle)"
	if ex9 != false || g9s != g9c {
		test.Errorf("Should't exist with %v and should be same but\n%v\n%v", ex9, g9s, g9c)
	}
}
