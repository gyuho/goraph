package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestDefaultGraph_DFS(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := DFS(g, "S")
	fmt.Println(rs) // [S C E B A D T F]
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}
