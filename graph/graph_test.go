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

func TestNewVertex(t *testing.T) {
	vtx := NewVertex("A")
	if reflect.TypeOf(vtx) != reflect.TypeOf(&Vertex{}) {
		t.Errorf("Should be same but \n%+v\n%+v", vtx, &Vertex{})
	}
}

func TestAddVertex(t *testing.T) {
	vertexToAdd1 := []*Vertex{
		NewVertex("A"), NewVertex("B"), NewVertex("C"),
		NewVertex("D"), NewVertex("E"), NewVertex("F"),
		NewVertex("G"),
	}
	data := NewData()
	for _, vtx := range vertexToAdd1 {
		exist := data.AddVertex(vtx)
		if !exist {
			t.Errorf("Shouldn't be false: %+v\n", data)
		}
	}
	vertexToAdd2 := []*Vertex{
		NewVertex("A"), NewVertex("B"), NewVertex("C"),
	}
	for _, vtx := range vertexToAdd2 {
		exist := data.AddVertex(vtx)
		if exist {
			t.Errorf("Shouldn't be false: %+v\n", data)
		}
	}
	if data.GetVertexSize() != 7 {
		t.Errorf("Expected 7 but %d", data.GetVertexSize())
	}
	data.AddVertex(NewVertex("X"))
	data.AddVertex(NewVertex("XX"))
	data.AddVertex(NewVertex("XXX"))
	if data.GetVertexSize() != 10 {
		t.Errorf("Expected 7 but %d", data.GetVertexSize())
	}
}

func TestConnect(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	t.Logf("%+v\n", data)
	if data.GetVertexSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", data)
	}
}

func TestInit(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	if data.GetVertexSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", data)
	}
	data.Init()
	if data.GetVertexSize() != 0 {
		t.Errorf("Expected 0 but %+v\n", data)
	}
}

func TestString(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	str1 := fmt.Sprintf("%s", data)
	str2 := data.String()
	if !strings.Contains(str1, "Vertex: A | Outgoing Edge: [A] -- 1.000 --> [B]") ||
		!strings.Contains(str2, "Vertex: A | Outgoing Edge: [A] -- 1.000 --> [B]") {
		t.Error(str1, str2)
	}
}

func TestFindVertexByID(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	vtx := data.FindVertexByID("B")
	if vtx.ID != "B" {
		t.Errorf("Expected B but %+v", vtx)
	}
}

func TestDeleteVertex(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	t.Logf("\n%+v", data)
	vtx := data.FindVertexByID("B")
	data.DeleteVertex(vtx)
	if data.GetVertexSize() != 2 {
		t.Errorf("Expected 2 but %+v", data)
	}
	t.Logf("\n%+v", data)
}

func TestDeleteEdge(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	vtx1 := data.FindVertexByID("B")
	vtx2 := data.FindVertexByID("C")
	data.DeleteEdge(vtx1, vtx2)
	if data.GetVertexSize() != 3 {
		t.Errorf("Expected 2 but %+v", data)
	}
}

func TestGetUpdateEdgeWeight(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 1.0)
	if data.GetEdgeWeight(data.FindVertexByID("C"), data.FindVertexByID("A")) != 21.000 {
		t.Errorf("Expected 21 but\n%+v", data)
	}
	data.UpdateEdgeWeight(data.FindVertexByID("C"), data.FindVertexByID("A"), 1.0)
	if data.GetEdgeWeight(data.FindVertexByID("C"), data.FindVertexByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", data)
	}
}
