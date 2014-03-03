package spbf

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
)

func Example_SPBF() {
	testgraph12 := `
S|A,7|B,6
A|C,-3|T,9
B|A,8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	g12 := gsd.ParseToGraph(testgraph12)
	fmt.Println(SPBF(g12, "S", "T"))
	// S(=0) → A(=7) → C(=4) → B(=2) → T(=-2) true
	fmt.Println(g12.ShowPrev("T"), "/ T's StampD:", g12.FindVertexByID("T").StampD)
	fmt.Println(g12.ShowPrev("B"), "/ B's StampD:", g12.FindVertexByID("B").StampD)
	fmt.Println(g12.ShowPrev("C"), "/ C's StampD:", g12.FindVertexByID("C").StampD)
	fmt.Println(g12.ShowPrev("A"), "/ A's StampD:", g12.FindVertexByID("A").StampD)
	fmt.Println(g12.ShowPrev("S"), "/ S's StampD:", g12.FindVertexByID("S").StampD)
	/*
	   Prev of T:  A B / T's StampD: -2
	   Prev of B:  S C / B's StampD: 2
	   Prev of C:  A B T / C's StampD: 4
	   Prev of A:  S B / A's StampD: 7
	   Prev of S:  T / S's StampD: 0

	   Add the smallest from back
	   	that is not source, destination
	   	that is not in the result
	*/

	testgraph13 := `
S|A,7|B,6
A|C,-3|T,9
B|A,-8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	g13 := gsd.ParseToGraph(testgraph13)
	fmt.Println(SPBF(g13, "S", "T"))
	// There is negative weighted cycle (No Shortest Path) false

	// Output:
	// S(=0) → A(=7) → C(=4) → B(=2) → T(=-2) true
	// Prev of T:  A B / T's StampD: -2
	// Prev of B:  S C / B's StampD: 2
	// Prev of C:  A B T / C's StampD: 4
	// Prev of A:  S B / A's StampD: 7
	// Prev of S:  T / S's StampD: 0
	// There is negative weighted cycle (No Shortest Path) false
}
