package goraph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testgraph"
)

func TestNewDefaultGraphInterface(t *testing.T) {
	g := NewDefaultGraph()
	t.Log(g.String())
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	_, err = NewDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
}

func TestNewDefaultGraphFromJSON(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(g.String())
	if g.VertexToChildren["C"]["S"] != 9.0 {
		t.Errorf("weight from C to S must be 9.0 but %f", g.VertexToChildren["C"]["S"])
	}
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
		if len(g.Vertices) != graph.TotalVertexCount {
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalVertexCount, len(g.Vertices))
		}
		for _, elem := range graph.EdgeToWeight {
			weight1, err := g.GetWeight(elem.Nodes[0], elem.Nodes[1])
			if err != nil {
				t.Error(err)
			}
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}

func TestDefaultGraph_GetVertices(t *testing.T) {
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
		if len(g.GetVertices()) != graph.TotalVertexCount {
			t.Errorf("wrong number of vertices: %s", g)
		}
	}
}

func TestDefaultGraph_Init(t *testing.T) {
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
		g.Init()
		if len(g.Vertices) != 0 {
			t.Errorf("not initialized: %s", g)
		}
	}
}

func TestDefaultGraph_DeleteVertex(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_01")
	if err != nil {
		t.Error(err)
	}
	if !g.DeleteVertex("D") {
		t.Error("D does not exist in the graph")
	}
	if g.FindVertex("D") {
		t.Errorf("Expected false but %s", g)
	}
	if v, err := g.GetParents("C"); err != nil || len(v) != 1 {
		t.Fatalf("Expected 1 edge incoming to C but %v\n\n%s", err, g)
	}
	if v, err := g.GetChildren("C"); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges outgoing from C but %v\n\n%s", err, g)
	}
	if v, err := g.GetChildren("F"); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges outgoing from F but %v\n\n%s", err, g)
	}
	if v, err := g.GetParents("F"); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges incoming to F but %v\n\n%s", err, g)
	}
	if v, err := g.GetChildren("B"); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from B but %v\n\n%s", err, g)
	}
	if v, err := g.GetParents("E"); err != nil || len(v) != 4 {
		t.Fatalf("Expected 4 edges incoming to E but %v\n\n%s", err, g)
	}
	if v, err := g.GetChildren("E"); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from E but %v\n\n%s", err, g)
	}
	if v, err := g.GetChildren("T"); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from T but %v\n\n%s", err, g)
	}
}

func TestDefaultGraph_DeleteEdge(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	g, err := newDefaultGraphFromJSON(f, "graph_01")
	if err != nil {
		t.Error(err)
	}

	if err := g.DeleteEdge("B", "D"); err != nil {
		t.Error(err)
	}
	if v, err := g.GetParents("D"); err != nil || len(v) != 4 {
		t.Errorf("Expected 4 edges incoming to D but %v\n\n%s", err, g)
	}

	if err := g.DeleteEdge("B", "C"); err != nil {
		t.Error(err)
	}
	if err := g.DeleteEdge("S", "C"); err != nil {
		t.Error(err)
	}
	if v, err := g.GetChildren("S"); err != nil || len(v) != 2 {
		t.Errorf("Expected 2 edges outgoing from S but %v\n\n%s", err, g)
	}

	if err := g.DeleteEdge("C", "E"); err != nil {
		t.Error(err)
	}
	if err := g.DeleteEdge("E", "D"); err != nil {
		t.Error(err)
	}
	if v, err := g.GetChildren("E"); err != nil || len(v) != 3 {
		t.Errorf("Expected 3 edges outgoing from E but %v\n\n%s", err, g)
	}
	if v, err := g.GetParents("E"); err != nil || len(v) != 3 {
		t.Errorf("Expected 3 edges incoming to E but %v\n\n%s", err, g)
	}

	if err := g.DeleteEdge("F", "E"); err != nil {
		t.Error(err)
	}
	if v, err := g.GetParents("E"); err != nil || len(v) != 2 {
		t.Errorf("Expected 2 edges incoming to E but %v\n\n%s", err, g)
	}
}

func TestDefaultGraph_ReplaceEdge(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := newDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	if err := g.ReplaceEdge("C", "S", 1.0); err != nil {
		t.Error(err)
	}
	if v, err := g.GetWeight("C", "S"); err != nil || v != 1.0 {
		t.Errorf("weight from C to S must be 1.0 but %v\n\n%v", err, g)
	}
}
