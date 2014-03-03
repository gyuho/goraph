package tskahn

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
)

func Example_TSKahn_6() {
	testgraph6 := `
D|B,1|C,1
E|C,1|F,1
B|A,1
A|F,1
`
	g6 := gsd.ParseToGraph(testgraph6)
	g6s, ex6 := TSKahn(g6)
	fmt.Println(g6s, ex6)

	// Output:
	// D → E → B → C → A → F true
}

func Example_TSKahn_7() {
	testgraph7 := `
C|D,1|E,1
B|D,1
A|E,1|H,1
D|F,1|G,1|H,1
E|G,1
`
	g7 := gsd.ParseToGraph(testgraph7)
	g7s, ex7 := TSKahn(g7)
	fmt.Println(g7s, ex7)

	// Output:
	// C → B → A → D → E → F → H → G true
}

func Example_TSKahn_8() {
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
	fmt.Println(g8s, ex8)

	// Output:
	// No Topological Sort (Not a DAG, there is a cycle) false
}

func Example_TSKahn_9() {
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
	fmt.Println(g9s, ex9)

	// Output:
	// No Topological Sort (Not a DAG, there is a cycle) false
}
