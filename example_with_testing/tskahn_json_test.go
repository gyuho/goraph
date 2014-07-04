package example_with_testing

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/tskahn"
	"github.com/gyuho/goraph/graph/gsd"
	// go test -v github.com/gyuho/goraph/example_with_testing
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example_with_testing/tskahn_json_test.go
)

func Test_JSON_TSKahn(test *testing.T) {
	fmt.Println("Topological Sort on testgraph6:")
	g6 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.006")
	g6s, ex6 := tskahn.TSKahn(g6)
	fmt.Println(g6s, ex6)
	// D → E → B → C → A → F true

	fmt.Println("Topological Sort on testgraph7:")
	g7 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.007")
	g7s, ex7 := tskahn.TSKahn(g7)
	fmt.Println(g7s, ex7)
	// A → B → C → D → E → F → H → G true

	// let's create a cyclic graph that is NOT a DAG
	// this has a cycle of A→E→D→B→A
	fmt.Println("Topological Sort on testgraph8:")
	g8 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.008")
	g8s, ex8 := tskahn.TSKahn(g8)
	fmt.Println(g8s, ex8)
	// No Topological Sort (Not a DAG, there is a cycle) false

	// this is NOT a DAG
	// this has a cycle of C→D→G→H→F→E→A→B→C
	fmt.Println("Topological Sort on testgraph9:")
	g9 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.009")
	g9s, ex9 := tskahn.TSKahn(g9)
	fmt.Println(g9s, ex9)
	// No Topological Sort (Not a DAG, there is a cycle) false
}
