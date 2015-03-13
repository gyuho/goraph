package graph

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewData(t *testing.T) {
	data := NewData()
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
	data := NewData()
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
	data := NewData()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	t.Logf("%+v\n", data)
	if data.GetNodeSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", data)
	}
}

func TestInit(t *testing.T) {
	data := NewData()
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

func TestString(t *testing.T) {
	data := NewData()
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
}

func TestGetNodeByID(t *testing.T) {
	data := NewData()
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
	data := NewData()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	t.Logf("\n%+v", data)
	nd := data.GetNodeByID("B")
	data.DeleteNode(nd)
	if data.GetNodeSize() != 2 {
		t.Errorf("Expected 2 but %+v", data)
	}
	t.Logf("\n%+v", data)
}

func TestDeleteEdge(t *testing.T) {
	data := NewData()
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
}

func TestGetUpdateEdgeWeight(t *testing.T) {
	data := NewData()
	data.Connect(NewNode("A"), NewNode("B"), 1.0)
	data.Connect(NewNode("B"), NewNode("C"), 10.0)
	data.Connect(NewNode("C"), NewNode("A"), 5.0)
	data.Connect(NewNode("C"), NewNode("A"), 15.0)
	data.Connect(NewNode("C"), NewNode("A"), 1.0)
	if data.GetEdgeWeight(data.GetNodeByID("C"), data.GetNodeByID("A")) != 21.000 {
		t.Errorf("Expected 21 but\n%+v", data)
	}
	data.UpdateEdgeWeight(data.GetNodeByID("C"), data.GetNodeByID("A"), 1.0)
	if data.GetEdgeWeight(data.GetNodeByID("C"), data.GetNodeByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", data)
	}
}
