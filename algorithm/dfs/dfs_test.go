package dfs

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_DFS(test *testing.T) {
	testgraph4 := `
S|A,15|B,14|C,9
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`
	g4 := gsd.ParseToGraph(testgraph4)
	g4s := DFS(g4)
	g4c := "S → A → B → D → E → F → T → C"
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

	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	g5 := gsd.ParseToGraph(testgraph5)
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
