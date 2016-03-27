package goraph

import (
	"fmt"
	"os"
	"testing"
)

func TestDefaultGraph_Tarjan_14(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_14")
	if err != nil {
		t.Error(err)
	}
	scc := Tarjan(g)
	if len(scc) != 4 {
		t.Fatalf("Expected 4 but %v", scc)
	}
	fmt.Println("Tarjan graph_14:", scc)
}

func TestDefaultGraph_Tarjan_15(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_15")
	if err != nil {
		t.Error(err)
	}
	scc := Tarjan(g)
	if len(scc) != 4 {
		t.Fatalf("Expected 4 but %v", scc)
	}
	fmt.Println("Tarjan graph_15:", scc)
}
