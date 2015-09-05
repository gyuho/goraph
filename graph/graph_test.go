package graph

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	g := New()
	if reflect.TypeOf(g) != reflect.TypeOf(&Graph{}) {
		t.Errorf("Should be same but \n%+v\n%+v", g, &Graph{})
	}
}

func TestNewNode(t *testing.T) {
	nd := NewNode("A")
	if reflect.TypeOf(nd) != reflect.TypeOf(&Node{}) {
		t.Errorf("Should be same but \n%+v\n%+v", nd, &Node{})
	}
}

func TestAddNode(t *testing.T) {
	nodeToAdd1 := []*Node{
		NewNode("A"), NewNode("B"), NewNode("C"),
		NewNode("D"), NewNode("E"), NewNode("F"),
		NewNode("G"),
	}
	g := New()
	for _, nd := range nodeToAdd1 {
		exist := g.AddNode(nd)
		if !exist {
			t.Errorf("Shouldn't be false: %+v\n", g)
		}
	}
	nodeToAdd2 := []*Node{
		NewNode("A"), NewNode("B"), NewNode("C"),
	}
	for _, nd := range nodeToAdd2 {
		exist := g.AddNode(nd)
		if exist {
			t.Errorf("Shouldn't be false: %+v\n", g)
		}
	}
	if g.GetNodeSize() != 7 {
		t.Errorf("Expected 7 but %d", g.GetNodeSize())
	}
	g.AddNode(NewNode("X"))
	g.AddNode(NewNode("XX"))
	g.AddNode(NewNode("XXX"))
	if g.GetNodeSize() != 10 {
		t.Errorf("Expected 7 but %d", g.GetNodeSize())
	}
}

func TestConnect(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	if g.GetNodeSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", g)
	}
}

func TestInit(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	if g.GetNodeSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", g)
	}
	g.Init()
	if g.GetNodeSize() != 0 {
		t.Errorf("Expected 0 but %+v\n", g)
	}
}

func TestGetNodeByID(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	nd := g.GetNodeByID("B")
	if nd.ID != "B" {
		t.Errorf("Expected B but %+v", nd)
	}
}

func TestDeleteNode(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	nd := g.GetNodeByID("B")
	g.DeleteNode(nd)
	if g.GetNodeSize() != 2 {
		t.Errorf("Expected 2 but %+v", g)
	}

	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	g2.DeleteNode(g2.GetNodeByID("D"))
	if len(g2.GetNodeByID("C").WeightFrom) != 1 {
		t.Errorf("Expected 1 edge incoming to C but %s", g2)
	}
	if len(g2.GetNodeByID("C").WeightTo) != 2 {
		t.Errorf("Expected 2 edges outgoing from C but %s", g2)
	}
	if len(g2.GetNodeByID("F").WeightTo) != 2 {
		t.Errorf("Expected 2 edges outgoing from F but %s", g2)
	}
	if len(g2.GetNodeByID("F").WeightFrom) != 2 {
		t.Errorf("Expected 2 edges incoming to F but %s", g2)
	}
	if g2.GetNodeByID("D") != nil {
		t.Errorf("Expected nil but %s", g2)
	}
	if len(g2.GetNodeByID("B").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from B but %s", g2)
	}
	if len(g2.GetNodeByID("E").WeightFrom) != 4 {
		t.Errorf("Expected 4 edges incoming to E but %s", g2)
	}
	if len(g2.GetNodeByID("E").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from E but %s", g2)
	}
	if len(g2.GetNodeByID("T").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from T but %s", g2)
	}
	if len(g2.GetNodeByID("T").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from T but %s", g2)
	}
}

func TestDeleteEdge(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	nd1 := g.GetNodeByID("B")
	nd2 := g.GetNodeByID("C")
	g.DeleteEdge(nd1, nd2)
	if g.GetNodeSize() != 3 {
		t.Errorf("Expected 2 but %+v", g)
	}

	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	g2.DeleteEdge(g2.GetNodeByID("Z"), g2.GetNodeByID("Z"))

	g2.DeleteEdge(g2.GetNodeByID("B"), g2.GetNodeByID("D"))
	if len(g2.GetNodeByID("D").WeightFrom) != 4 {
		t.Errorf("Expected 4 edges incoming to D but %s", g2)
	}

	g2.DeleteEdge(g2.GetNodeByID("B"), g2.GetNodeByID("D"))
	if len(g2.GetNodeByID("D").WeightFrom) != 4 {
		t.Errorf("Expected 4 edges incoming to D but %s", g2)
	}

	g2.DeleteEdge(g2.GetNodeByID("B"), g2.GetNodeByID("C"))
	g2.DeleteEdge(g2.GetNodeByID("S"), g2.GetNodeByID("C"))
	if len(g2.GetNodeByID("S").WeightTo) != 2 {
		t.Errorf("Expected 2 edges outgoing from S but %s", g2)
	}

	g2.DeleteEdge(g2.GetNodeByID("C"), g2.GetNodeByID("E"))
	g2.DeleteEdge(g2.GetNodeByID("E"), g2.GetNodeByID("D"))
	if len(g2.GetNodeByID("E").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from E but %s", g2)
	}
	if len(g2.GetNodeByID("E").WeightFrom) != 3 {
		t.Errorf("Expected 3 edges incoming to E but %s", g2)
	}
	g2.DeleteEdge(g2.GetNodeByID("F"), g2.GetNodeByID("E"))
	if len(g2.GetNodeByID("E").WeightFrom) != 2 {
		t.Errorf("Expected 2 edges incoming to E but %s", g2)
	}
}

func TestGetUpdateEdgeWeight(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	g.Connect(NewNode("C"), NewNode("A"), 1.0)
	if g.GetEdgeWeight(g.GetNodeByID("C"), g.GetNodeByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", g)
	}
	g.UpdateEdgeWeight(g.GetNodeByID("C"), g.GetNodeByID("A"), 1.0)
	if g.GetEdgeWeight(g.GetNodeByID("C"), g.GetNodeByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", g)
	}
}

func TestString(t *testing.T) {
	g := New()
	g.Connect(NewNode("A"), NewNode("B"), 1.0)
	g.Connect(NewNode("B"), NewNode("C"), 10.0)
	g.Connect(NewNode("C"), NewNode("A"), 5.0)
	g.Connect(NewNode("C"), NewNode("A"), 15.0)
	str1 := fmt.Sprintf("%s", g)
	str2 := g.String()
	if !strings.Contains(str1, "Node: A | Outgoing Edge: [A] -- 1.000 --> [B]") ||
		!strings.Contains(str2, "Node: A | Outgoing Edge: [A] -- 1.000 --> [B]") {
		t.Error(str1, str2)
	}
	if "[A / 1 Outgoing / 1 Incoming Edges]" != fmt.Sprintf("%v", g.GetNodeByID("A")) {
		t.Errorf("Unexpected %v", g.GetNodeByID("A"))
	}
}
