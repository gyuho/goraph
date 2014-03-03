package bfs

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_BFS(test *testing.T) {
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
	fmt.Println("Breadth First Search on testgraph4:")
	g4 := gsd.ParseToGraph(testgraph4)
	g4s := BFS(g4, g4.FindVertexByID("S"))
	g4c := "S → A → B → C → D → T → E → F"
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

	println()
	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	fmt.Println("Breadth First Search on testgraph5:")
	g5 := gsd.ParseToGraph(testgraph5)
	g5s := BFS(g5, g5.FindVertexByID("A"))
	g5c := "A → B → C → F → D → E"
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
	fmt.Println("Breadth First Search on testgraph4:")
	g4 := gsd.ParseToGraph(testgraph4)
	fmt.Println(BFS(g4, g4.FindVertexByID("S")))
	// S → A → B → C → D → T → E → F
	// S(timestamp: 0) → A(timestamp: 1) → B(timestamp: 1) → C(timestamp: 1) → D(timestamp: 2) → T(timestamp: 2) → E(timestamp: 2) → F(timestamp: 3)

	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		fmt.Printf(" %v(Color: %v) /", vtx.(*gsd.Vertex).ID, vtx.(*gsd.Vertex).Color)
	}
	// S(Color: black) / A(Color: black) / B(Color: black) / C(Color: black) / D(Color: black) / T(Color: black) / E(Color: black) / F(Color: black) /

	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	fmt.Println("Breadth First Search on testgraph5:")
	g5 := gsd.ParseToGraph(testgraph5)
	fmt.Println(BFS(g5, g5.FindVertexByID("A")))
	// A → B → C → F → D → E
	// A(timestamp: 0) → B(timestamp: 1) → C(timestamp: 1) → F(timestamp: 1) → D(timestamp: 2) → E(timestamp: 2)

	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		fmt.Printf(" %v(Color: %v) /", vtx.(*gsd.Vertex).ID, vtx.(*gsd.Vertex).Color)
	}
	// A(Color: black) / B(Color: black) / C(Color: black) / F(Color: black) / D(Color: black) / E(Color: black) /

	// Output:
}

/*
got:
Breadth First Search on testgraph4:
S → A → B → C → D → T → E → F
 S(Color: black) / A(Color: black) / B(Color: black) / C(Color: black) / D(Color: black) / T(Color: black) / E(Color: black) / F(Color: black) /
Breadth First Search on testgraph5:
A → B → C → F → D → E
 A(Color: black) / B(Color: black) / C(Color: black) / F(Color: black) / D(Color: black) / E(Color: black) /
want:
*/
