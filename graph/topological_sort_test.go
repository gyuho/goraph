package graph

import (
	"fmt"
	"os"
	"testing"

	"github.com/gyuho/goraph/graph/testgraph"
)

func TestDefaultGraph_TopologicalSort_05(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_05")
	if err != nil {
		t.Error(err)
	}
	L, isDAG := TopologicalSort(g)
	if isDAG != true {
		t.Errorf("there is no directed cycle in the graph so isDAG should be true but %+v %+v", L, isDAG)
	}
	fmt.Println("graph_05:", L)
}

func TestDefaultGraph_TopologicalSort_06(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_06")
	if err != nil {
		t.Error(err)
	}
	L, isDAG := TopologicalSort(g)
	if isDAG != true {
		t.Errorf("there is no directed cycle in the graph so isDAG should be true but %+v %+v", L, isDAG)
	}
	fmt.Println("graph_06:", L)
}

func TestDefaultGraph_TopologicalSort_07(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_07")
	if err != nil {
		t.Error(err)
	}
	L, isDAG := TopologicalSort(g)
	if isDAG != false {
		t.Errorf("there is a directed cycle in the graph so isDAG should be false but %+v %+v", L, isDAG)
	}
	fmt.Println("graph_07:", L)
}

func TestDefaultGraph_TopologicalSort(t *testing.T) {
	for _, graph := range testgraph.GraphSlice {
		f, err := os.Open("testdata/graph.json")
		if err != nil {
			t.Error(err)
		}
		defer f.Close()
		g, err := newDefaultGraphFromJSON(f, graph.Name)
		if err != nil {
			t.Error(err)
		}
		L, isDAG := TopologicalSort(g)
		if isDAG != graph.IsDAG {
			t.Errorf("%s | IsDag are supposed to be %v but %+v %+v", graph.Name, graph.IsDAG, L, isDAG)
		}
	}
}
