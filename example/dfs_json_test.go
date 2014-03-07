package example

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/dfs"
	"github.com/gyuho/goraph/graph/gsd"
	// go test -v github.com/gyuho/goraph/example
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example/dfs_json_test.go
)

func Test_JSON_DFS(test *testing.T) {
	fmt.Println("Depth First Search on testgraph4:")
	g4 := gsd.JSONGraph("../testgraph/testgraph.json", "testgraph.004")
	fmt.Println(dfs.DFS(g4))
	// S → B → A → T → F → E → C → D

	println()
	fmt.Println("Depth First Search on testgraph5:")
	g5 := gsd.JSONGraph("../testgraph/testgraph.json", "testgraph.005")
	fmt.Println(dfs.DFS(g5))
	// A → B → C → D → E → F
}
