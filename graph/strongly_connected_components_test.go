package graph

import (
	"os"
	"testing"
)

func TestDefaultGraph_Tarjan_14(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_14")
	if err != nil {
		t.Error(err)
	}
	rs := Tarjan(g)
	_ = rs
}
