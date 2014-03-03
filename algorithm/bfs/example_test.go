package bfs

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
)

func Example_BFS() {
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
	fmt.Println(BFS(g4, g4.FindVertexByID("S")))
	// S → A → B → C → D → T → E → F
	// S(timestamp: 0) → A(timestamp: 1) → B(timestamp: 1) → C(timestamp: 1) → D(timestamp: 2) → T(timestamp: 2) → E(timestamp: 2) → F(timestamp: 3)

	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		fmt.Printf("%v(Color: %v) / ", vtx.(*gsd.Vertex).ID, vtx.(*gsd.Vertex).Color)
	}
	fmt.Printf("\n")
	// S(Color: black) / A(Color: black) / B(Color: black) / C(Color: black) / D(Color: black) / T(Color: black) / E(Color: black) / F(Color: black) /

	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	g5 := gsd.ParseToGraph(testgraph5)
	fmt.Println(BFS(g5, g5.FindVertexByID("A")))
	// A → B → C → F → D → E
	// A(timestamp: 0) → B(timestamp: 1) → C(timestamp: 1) → F(timestamp: 1) → D(timestamp: 2) → E(timestamp: 2)

	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		fmt.Printf("%v(Color: %v) / ", vtx.(*gsd.Vertex).ID, vtx.(*gsd.Vertex).Color)
	}
	// A(Color: black) / B(Color: black) / C(Color: black) / F(Color: black) / D(Color: black) / E(Color: black) /

	// Output:
	// S → A → B → C → D → T → E → F
	// S(Color: black) / A(Color: black) / B(Color: black) / C(Color: black) / D(Color: black) / T(Color: black) / E(Color: black) / F(Color: black) /

	// A → B → C → F → D → E
	// A(Color: black) / B(Color: black) / C(Color: black) / F(Color: black) / D(Color: black) / E(Color: black) /
}
