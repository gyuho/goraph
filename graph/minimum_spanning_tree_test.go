package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestDefaultGraph_Kruskal_13(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_13")
	if err != nil {
		t.Error(err)
	}
	A := Kruskal(g)
	total := 0.0
	for edge := range A {
		total += edge.Weight
	}
	if total != 37.0 {
		t.Errorf("Expected total 37.0 but %.2f", total)
	}
	fmt.Println("Kruskal from graph_13:", A)
}

func TestDefaultGraph_Prim_13(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_13")
	if err != nil {
		t.Error(err)
	}
	for v := range g.GetVertices() {
		A := Prim(g, v)
		total := 0.0
		for edge := range A {
			total += edge.Weight
		}
		if total != 37.0 {
			t.Errorf("Expected total 37.0 but %.2f", total)
		}
		fmt.Println("Prim from graph_13:", A, "with", v)
	}
}
