package graph

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	data := New()
	if reflect.TypeOf(data) != reflect.TypeOf(&Data{}) {
		t.Errorf("Should be same but \n%+v\n%+v", data, &Data{})
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
	data := New()
	for _, nd := range nodeToAdd1 {
		exist := data.AddNode(nd)
		if !exist {
			t.Errorf("Shouldn't be false: %+v\n", data)
		}
	}
	nodeToAdd2 := []*Node{
		NewNode("A"), NewNode("B"), NewNode("C"),
	}
	for _, nd := range nodeToAdd2 {
		exist := data.AddNode(nd)
		if exist {
			t.Errorf("Shouldn't be false: %+v\n", data)
		}
	}
	if data.GetNodeSize() != 7 {
		t.Errorf("Expected 7 but %d", data.GetNodeSize())
	}
	data.AddNode(NewNode("X"))
	data.AddNode(NewNode("XX"))
	data.AddNode(NewNode("XXX"))
	if data.GetNodeSize() != 10 {
		t.Errorf("Expected 7 but %d", data.GetNodeSize())
	}
}

func TestConnect(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	if data.GetNodeSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", data)
	}
}

func TestInit(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	if data.GetNodeSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", data)
	}
	data.Init()
	if data.GetNodeSize() != 0 {
		t.Errorf("Expected 0 but %+v\n", data)
	}
}

func TestGetNodeByID(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	nd := data.GetNodeByID("B")
	if nd.ID != "B" {
		t.Errorf("Expected B but %+v", nd)
	}
}

func TestDeleteNode(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	nd := data.GetNodeByID("B")
	data.DeleteNode(nd)
	if data.GetNodeSize() != 2 {
		t.Errorf("Expected 2 but %+v", data)
	}

	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	data2.DeleteNode(data2.GetNodeByID("D"))
	if len(data2.GetNodeByID("C").WeightFrom) != 1 {
		t.Errorf("Expected 1 edge incoming to C but %s", data2)
	}
	if len(data2.GetNodeByID("C").WeightTo) != 2 {
		t.Errorf("Expected 2 edges outgoing from C but %s", data2)
	}
	if len(data2.GetNodeByID("F").WeightTo) != 2 {
		t.Errorf("Expected 2 edges outgoing from F but %s", data2)
	}
	if len(data2.GetNodeByID("F").WeightFrom) != 2 {
		t.Errorf("Expected 2 edges incoming to F but %s", data2)
	}
	if data2.GetNodeByID("D") != nil {
		t.Errorf("Expected nil but %s", data2)
	}
	if len(data2.GetNodeByID("B").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from B but %s", data2)
	}
	if len(data2.GetNodeByID("E").WeightFrom) != 4 {
		t.Errorf("Expected 4 edges incoming to E but %s", data2)
	}
	if len(data2.GetNodeByID("E").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from E but %s", data2)
	}
	if len(data2.GetNodeByID("T").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from T but %s", data2)
	}
	if len(data2.GetNodeByID("T").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from T but %s", data2)
	}
}

func TestDeleteEdge(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	nd1 := data.GetNodeByID("B")
	nd2 := data.GetNodeByID("C")
	data.DeleteEdge(nd1, nd2)
	if data.GetNodeSize() != 3 {
		t.Errorf("Expected 2 but %+v", data)
	}

	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	data2.DeleteEdge(data2.GetNodeByID("Z"), data2.GetNodeByID("Z"))

	data2.DeleteEdge(data2.GetNodeByID("B"), data2.GetNodeByID("D"))
	if len(data2.GetNodeByID("D").WeightFrom) != 4 {
		t.Errorf("Expected 4 edges incoming to D but %s", data2)
	}

	data2.DeleteEdge(data2.GetNodeByID("B"), data2.GetNodeByID("D"))
	if len(data2.GetNodeByID("D").WeightFrom) != 4 {
		t.Errorf("Expected 4 edges incoming to D but %s", data2)
	}

	data2.DeleteEdge(data2.GetNodeByID("B"), data2.GetNodeByID("C"))
	data2.DeleteEdge(data2.GetNodeByID("S"), data2.GetNodeByID("C"))
	if len(data2.GetNodeByID("S").WeightTo) != 2 {
		t.Errorf("Expected 2 edges outgoing from S but %s", data2)
	}

	data2.DeleteEdge(data2.GetNodeByID("C"), data2.GetNodeByID("E"))
	data2.DeleteEdge(data2.GetNodeByID("E"), data2.GetNodeByID("D"))
	if len(data2.GetNodeByID("E").WeightTo) != 3 {
		t.Errorf("Expected 3 edges outgoing from E but %s", data2)
	}
	if len(data2.GetNodeByID("E").WeightFrom) != 3 {
		t.Errorf("Expected 3 edges incoming to E but %s", data2)
	}
	data2.DeleteEdge(data2.GetNodeByID("F"), data2.GetNodeByID("E"))
	if len(data2.GetNodeByID("E").WeightFrom) != 2 {
		t.Errorf("Expected 2 edges incoming to E but %s", data2)
	}
}

func TestGetUpdateEdgeWeight(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	data.Connect(NewNode("C"), NewNode("A"), 1.0)
	if data.GetEdgeWeight(data.GetNodeByID("C"), data.GetNodeByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", data)
	}
	data.UpdateEdgeWeight(data.GetNodeByID("C"), data.GetNodeByID("A"), 1.0)
	if data.GetEdgeWeight(data.GetNodeByID("C"), data.GetNodeByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", data)
	}
}

func TestString(t *testing.T) {
	data := New()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	str1 := fmt.Sprintf("%s", data)
	str2 := data.String()
	if !strings.Contains(str1, "Node: A | Outgoing Edge: [A] -- 1.000 --> [B]") ||
		!strings.Contains(str2, "Node: A | Outgoing Edge: [A] -- 1.000 --> [B]") {
		t.Error(str1, str2)
	}
	if "[A / 1 Outgoing / 1 Incoming Edges]" != fmt.Sprintf("%v", data.GetNodeByID("A")) {
		t.Errorf("Unexpected %v", data.GetNodeByID("A"))
	}
}
