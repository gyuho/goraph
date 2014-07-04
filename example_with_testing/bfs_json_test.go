// Package example_with_testing is to show code usage.
// It is to be run with the command go test -v.
package example_with_testing

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/bfs"
	"github.com/gyuho/goraph/graph/gsd"
	// go test -v github.com/gyuho/goraph/example_with_testing
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example_with_testing/bfs_json_test.go
)

func Test_JSON_BFS(test *testing.T) {
	fmt.Println("Breadth First Search on testgraph4:")
	g4 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.004")
	fmt.Println(bfs.BFS(g4, g4.FindVertexByID("S")))
	// S → B → A → D → E → T → C → F

	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		fmt.Printf(" %v(Color: %v) /", vtx.(*gsd.Vertex).ID, vtx.(*gsd.Vertex).Color)
	}
	// S(Color: black) / B(Color: black) / A(Color: black) / D(Color: black) / T(Color: black) / E(Color: black) / C(Color: black) / F(Color: black) /

	println()
	fmt.Println("Breadth First Search on testgraph5:")
	g5 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.005")
	fmt.Println(bfs.BFS(g5, g5.FindVertexByID("A")))
	// A → B → C → F → D → E

	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		fmt.Printf(" %v(Color: %v) /", vtx.(*gsd.Vertex).ID, vtx.(*gsd.Vertex).Color)
	}
	// A(Color: black) / B(Color: black) / C(Color: black) / F(Color: black) / D(Color: black) / E(Color: black) /
}
