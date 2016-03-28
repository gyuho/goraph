package goraph

import (
	"fmt"
	"os"
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := BFS(g, StringID("S"))
	fmt.Println("BFS:", rs) // [S A B C D T E F]
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}

func TestGraph_DFS(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := DFS(g, StringID("S"))
	fmt.Println("DFS:", rs) // [S C E B A D T F]
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}

func TestGraph_DFSRecursion(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := DFSRecursion(g, StringID("S"))
	fmt.Println("DFSRecursion:", rs) // [S C E T A B D F]
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}
