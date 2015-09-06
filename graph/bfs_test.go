package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestDefaultGraph_BFS(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := BFS(g, "S")
	fmt.Println("BFS:", rs) // [S A B C D T E F]
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}
