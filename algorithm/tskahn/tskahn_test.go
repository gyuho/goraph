package tskahn

import (
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_TSKahn(test *testing.T) {
	testgraph6 := `
D|B,1|C,1
E|C,1|F,1
B|A,1
A|F,1
`
	g6 := gsd.ParseToGraph(testgraph6)
	g6s, ex6 := TSKahn(g6)
	g6c := "D → E → B → C → A → F"
	if ex6 != true || g6s != g6c {
		test.Errorf("Should exist with %v and should be same but\n%v\n%v\n%v", ex6, g6s, g6c, g6.GetEdgesSize())
	}

	testgraph7 := `
C|D,1|E,1
B|D,1
A|E,1|H,1
D|F,1|G,1|H,1
E|G,1
`
	g7 := gsd.ParseToGraph(testgraph7)
	g7s, ex7 := TSKahn(g7)
	g7c := "C → B → A → D → E → F → H → G"
	if ex7 != true || g7s != g7c {
		test.Errorf("Should exist with %v and should be same but\n%v\n%v\n%v", ex7, g7s, g7c, g7.GetEdgesSize())
	}

	// let's create a cyclic graph that is NOT a DAG
	// this has a cycle of A→E→D→B→A
	testgraph8 := `
D|B,1|C,1
E|C,1|F,1|D,1
B|A,1
A|F,1|E,1
`
	g8 := gsd.ParseToGraph(testgraph8)
	g8s, ex8 := TSKahn(g8)
	g8c := "No Topological Sort (Not a DAG, there is a cycle)"
	if ex8 != false || g8s != g8c {
		test.Errorf("Should't exist with %v and should be same but\n%v\n%v", ex8, g8s, g8c)
	}

	// this is NOT a DAG
	// this has a cycle of C→D→G→H→F→E→A→B→C
	testgraph9 := `
C|D,1|E,1
B|D,1|C,1
A|E,1|H,1|B,1
D|F,1|G,1|H,1
E|G,1|A,1
G|H,1
F|E,1
`
	g9 := gsd.ParseToGraph(testgraph9)
	g9s, ex9 := TSKahn(g9)
	g9c := "No Topological Sort (Not a DAG, there is a cycle)"
	if ex9 != false || g9s != g9c {
		test.Errorf("Should't exist with %v and should be same but\n%v\n%v", ex9, g9s, g9c)
	}
}
