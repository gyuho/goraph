package goraph

import (
	"fmt"
	"os"
	"testing"
)

func TestGraph_Kruskal_13(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_13")
	if err != nil {
		t.Error(err)
	}
	A, err := Kruskal(g)
	if err != nil {
		t.Error(err)
	}
	total := 0.0
	for edge := range A {
		total += edge.Weight()
	}
	if total != 37.0 {
		t.Errorf("Expected total 37.0 but %.2f", total)
	}
	fmt.Println("Kruskal from graph_13:", A)
}

func TestGraph_Prim_13(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_13")
	if err != nil {
		t.Error(err)
	}
	for v := range g.GetNodes() {
		A, err := Prim(g, v)
		if err != nil {
			t.Error(err)
		}
		total := 0.0
		for edge := range A {
			total += edge.Weight()
		}
		if total != 37.0 {
			t.Errorf("Expected total 37.0 but %.2f", total)
		}
		fmt.Println("Prim from graph_13:", A, "with", v)
	}
}
