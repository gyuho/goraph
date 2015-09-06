package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDefaultGraphFromJSON(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_00", true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(111)
	fmt.Println(g.String())
}
