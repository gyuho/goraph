package spbf

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_SPBF(test *testing.T) {
	testgraph12 := `
S|A,7|B,6
A|C,-3|T,9
B|A,8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	fmt.Println("Bellman-Ford Shortest Path on testgraph12:")
	g12 := gsd.ParseToGraph(testgraph12)
	s12, ok12 := SPBF(g12, "S", "T")
	s12c := "S(=0) → A(=7) → C(=4) → B(=2) → T(=-2)"
	ok12c := true
	if s12 != s12c || ok12 != ok12c {
		test.Errorf("Should be same but %v", s12)
	}

	testgraph13 := `
S|A,7|B,6
A|C,-3|T,9
B|A,-8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	fmt.Println("Bellman-Ford Shortest Path on testgraph13:")
	g13 := gsd.ParseToGraph(testgraph13)
	s13, ok13 := SPBF(g13, "S", "T")
	s13c := "There is negative weighted cycle (No Shortest Path)"
	ok13c := false
	if s13 != s13c || ok13 != ok13c {
		test.Errorf("Should be same but %v", s13)
	}
}
