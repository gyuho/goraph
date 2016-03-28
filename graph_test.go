package goraph

import (
	"fmt"
	"os"
	"testing"

	"github.com/gyuho/goraph/testgraph"
)

func TestNewGraph(t *testing.T) {
	g1 := NewGraph()
	fmt.Println("g1:", g1.String())

	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g2, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("g2:", g2.String())
}

func TestNewGraphFromJSON_graph(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	jg, err := NewGraphFromJSON(f, "graph_00")
	g, ok := jg.(*graph)
	if err != nil || !ok {
		t.Fatalf("nil graph %v", err)
	}
	if g.nodeToTargets[g.GetID("C")][g.GetID("S")] != 9.0 {
		t.Fatalf("weight from C to S must be 9.0 but %f", g.nodeToTargets[g.GetID("C")][g.GetID("S")])
	}
	for _, tg := range testgraph.GraphSlice {
		f, err := os.Open("testdata/graph.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		jg, err := NewGraphFromJSON(f, tg.Name)
		g, ok := jg.(*graph)
		if err != nil || !ok {
			t.Fatalf("nil graph %v", err)
		}
		if g.GetNodeCount() != tg.TotalVertexCount {
			t.Fatalf("%s | Expected %d but %d", tg.Name, tg.TotalVertexCount, g.GetNodeCount())
		}
		for _, elem := range tg.EdgeToWeight {
			weight1, err := g.GetWeight(g.GetID(elem.Nodes[0]), g.GetID(elem.Nodes[1]))
			if err != nil {
				t.Fatal(err)
			}
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Fatalf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}

func TestGraph_GetVertices(t *testing.T) {
	for _, tg := range testgraph.GraphSlice {
		f, err := os.Open("testdata/graph.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		g, err := NewGraphFromJSON(f, tg.Name)
		if err != nil {
			t.Fatal(err)
		}
		if g.GetNodeCount() != tg.TotalVertexCount {
			t.Fatalf("wrong number of vertices: %s", g)
		}
	}
}

func TestGraph_Init(t *testing.T) {
	for _, tg := range testgraph.GraphSlice {
		f, err := os.Open("testdata/graph.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		g, err := NewGraphFromJSON(f, tg.Name)
		if err != nil {
			t.Fatal(err)
		}
		g.Init()
		if g.GetNodeCount() != 0 {
			t.Fatalf("not initialized: %s", g)
		}
	}
}

func TestGraph_DeleteNode(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_01")
	if err != nil {
		t.Fatal(err)
	}
	if !g.DeleteNode(g.GetID("D")) {
		t.Fatal("D does not exist in the graph")
	}
	if g.GetNode("D") != nil {
		t.Fatalf("Expected Node but %s", g.GetNode("D"))
	}
	if v, err := g.GetSources(g.GetID("C")); err != nil || len(v) != 1 {
		t.Fatalf("Expected 1 edge incoming to C but %v\n\n%s", err, g)
	}
	if v, err := g.GetTargets(g.GetID("C")); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges outgoing from C but %v\n\n%s", err, g)
	}
	if v, err := g.GetTargets(g.GetID("F")); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges outgoing from F but %v\n\n%s", err, g)
	}
	if v, err := g.GetSources(g.GetID("F")); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges incoming to F but %v\n\n%s", err, g)
	}
	if v, err := g.GetTargets(g.GetID("B")); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from B but %v\n\n%s", err, g)
	}
	if v, err := g.GetSources(g.GetID("E")); err != nil || len(v) != 4 {
		t.Fatalf("Expected 4 edges incoming to E but %v\n\n%s", err, g)
	}
	if v, err := g.GetTargets(g.GetID("E")); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from E but %v\n\n%s", err, g)
	}
	if v, err := g.GetTargets(g.GetID("T")); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from T but %v\n\n%s", err, g)
	}
}

func TestGraph_DeleteEdge(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	g, err := NewGraphFromJSON(f, "graph_01")
	if err != nil {
		t.Fatal(err)
	}

	if err := g.DeleteEdge(g.GetID("B"), g.GetID("D")); err != nil {
		t.Fatal(err)
	}
	if v, err := g.GetSources(g.GetID("D")); err != nil || len(v) != 4 {
		t.Fatalf("Expected 4 edges incoming to D but %v\n\n%s", err, g)
	}

	if err := g.DeleteEdge(g.GetID("B"), g.GetID("C")); err != nil {
		t.Fatal(err)
	}
	if err := g.DeleteEdge(g.GetID("S"), g.GetID("C")); err != nil {
		t.Fatal(err)
	}
	if v, err := g.GetTargets(g.GetID("S")); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges outgoing from S but %v\n\n%s", err, g)
	}

	if err := g.DeleteEdge(g.GetID("C"), g.GetID("E")); err != nil {
		t.Fatal(err)
	}
	if err := g.DeleteEdge(g.GetID("E"), g.GetID("D")); err != nil {
		t.Fatal(err)
	}
	if v, err := g.GetTargets(g.GetID("E")); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges outgoing from E but %v\n\n%s", err, g)
	}
	if v, err := g.GetSources(g.GetID("E")); err != nil || len(v) != 3 {
		t.Fatalf("Expected 3 edges incoming to E but %v\n\n%s", err, g)
	}

	if err := g.DeleteEdge(g.GetID("F"), g.GetID("E")); err != nil {
		t.Fatal(err)
	}
	if v, err := g.GetSources(g.GetID("E")); err != nil || len(v) != 2 {
		t.Fatalf("Expected 2 edges incoming to E but %v\n\n%s", err, g)
	}
}

func TestGraph_ReplaceEdge(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Fatal(err)
	}
	if err := g.ReplaceEdge(g.GetID("C"), g.GetID("S"), 1.0); err != nil {
		t.Fatal(err)
	}
	if v, err := g.GetWeight(g.GetID("C"), g.GetID("S")); err != nil || v != 1.0 {
		t.Fatalf("weight from C to S must be 1.0 but %v\n\n%v", err, g)
	}
}
