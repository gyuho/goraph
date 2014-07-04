package example_with_testing

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/tskahn"
	"github.com/gyuho/goraph/graph/gsd"
	// go test -v github.com/gyuho/goraph/example_with_testing
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example_with_testing/tskahn_test.go
)

func Test_TSKahn(test *testing.T) {
	testgraph6 := `
D|B,1|C,1
E|C,1|F,1
B|A,1
A|F,1
`
	fmt.Println("Topological Sort on testgraph6:")
	g6 := gsd.ParseToGraph(testgraph6)
	g6s, ex6 := tskahn.TSKahn(g6)
	fmt.Println(g6s, ex6)
	// D → E → B → C → A → F true

	testgraph7 := `
C|D,1|E,1
B|D,1
A|E,1|H,1
D|F,1|G,1|H,1
E|G,1
`
	fmt.Println("Topological Sort on testgraph7:")
	g7 := gsd.ParseToGraph(testgraph7)
	g7s, ex7 := tskahn.TSKahn(g7)
	fmt.Println(g7s, ex7)
	// C → B → A → D → E → F → H → G true

	// let's create a cyclic graph that is NOT a DAG
	// this has a cycle of A→E→D→B→A
	testgraph8 := `
D|B,1|C,1
E|C,1|F,1|D,1
B|A,1
A|F,1|E,1
`
	fmt.Println("Topological Sort on testgraph8:")
	g8 := gsd.ParseToGraph(testgraph8)
	g8s, ex8 := tskahn.TSKahn(g8)
	fmt.Println(g8s, ex8)
	// No Topological Sort (Not a DAG, there is a cycle) false

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
	fmt.Println("Topological Sort on testgraph9:")
	g9 := gsd.ParseToGraph(testgraph9)
	g9s, ex9 := tskahn.TSKahn(g9)
	fmt.Println(g9s, ex9)
	// No Topological Sort (Not a DAG, there is a cycle) false
}
