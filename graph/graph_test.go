package graph

import (
	"fmt"
	"reflect"
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
		exist, err := data.AddVertex(vtx)
		if !exist || err != nil {
			t.Errorf("Expected no error: %+v\n", data)
		}
	}
	vertexToAdd2 := []*Vertex{
		NewVertex("A"), NewVertex("B"), NewVertex("C"),
	}
	for _, vtx := range vertexToAdd2 {
		exist, err := data.AddVertex(vtx)
		if exist || err == nil {
			t.Errorf("Expected error: %+v\n", data)
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
	str := `Vertex: A | Outgoing Edges: [A] -- 1.000 --> [B]
Vertex: A | Incoming Edges: [C] -- 20.000 --> [A]
Vertex: B | Outgoing Edges: [B] -- 10.000 --> [C]
Vertex: B | Incoming Edges: [A] -- 1.000 --> [B]
Vertex: C | Outgoing Edges: [C] -- 20.000 --> [A]
Vertex: C | Incoming Edges: [B] -- 10.000 --> [C]`
	str1 := fmt.Sprintf("%+v", data)
	str2 := fmt.Sprintf("%s", data)
	str3 := data.String()
	if str != str1 || str1 != str2 || str2 != str3 || str3 != str1 {
		t.Errorf("Expected the same:\n%s\n\n%s\n\n%s\n\n%s",
			str, str1, str2, str3)
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
	if _, ok := data.OutEdges[vtx]; ok {
		t.Errorf("Expected false but %+v", data)
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
	if len(data.OutEdges[vtx1]) != 0 {
		t.Errorf("Expected false but\n%+v", data)
	}
	str1 := `Vertex: A | Outgoing Edges: [A] -- 1.000 --> [B]
Vertex: A | Incoming Edges: [C] -- 20.000 --> [A]
Vertex: B | Incoming Edges: [A] -- 1.000 --> [B]
Vertex: C | Outgoing Edges: [C] -- 20.000 --> [A]`
	str2 := data.String()
	if str1 != str2 {
		t.Errorf("Expected the same but\n%s\n\n%s", str1, str2)
	}
}

func TestGetUpdateEdgeWeight(t *testing.T) {
	data := NewData()
	data.Connect(NewVertex("A"), NewVertex("B"), 1.0)
	data.Connect(NewVertex("B"), NewVertex("C"), 10.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 5.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 15.0)
	data.Connect(NewVertex("C"), NewVertex("A"), 1.0)
	str1 := `Vertex: A | Outgoing Edges: [A] -- 1.000 --> [B]
Vertex: A | Incoming Edges: [C] -- 21.000 --> [A]
Vertex: B | Outgoing Edges: [B] -- 10.000 --> [C]
Vertex: B | Incoming Edges: [A] -- 1.000 --> [B]
Vertex: C | Outgoing Edges: [C] -- 21.000 --> [A]
Vertex: C | Incoming Edges: [B] -- 10.000 --> [C]`
	str2 := data.String()
	if str1 != str2 {
		t.Errorf("Expected the same but\n%s\n\n%s", str1, str2)
	}
	if data.GetEdgeWeight(data.FindVertexByID("C"), data.FindVertexByID("A")) != 21.000 {
		t.Errorf("Expected 21 but\n%+v", data)
	}

	data.UpdateEdgeWeight(data.FindVertexByID("C"), data.FindVertexByID("A"), 1.0)
	str3 := `Vertex: A | Outgoing Edges: [A] -- 1.000 --> [B]
Vertex: A | Incoming Edges: [C] -- 21.000 --> [A]
Vertex: B | Outgoing Edges: [B] -- 10.000 --> [C]
Vertex: B | Incoming Edges: [A] -- 1.000 --> [B]
Vertex: C | Outgoing Edges: [C] -- 1.000 --> [A]
Vertex: C | Incoming Edges: [B] -- 10.000 --> [C]`
	str4 := data.String()
	if str3 != str4 {
		t.Errorf("Expected the same but\n%s\n\n%s", str3, str4)
	}
	if data.GetEdgeWeight(data.FindVertexByID("C"), data.FindVertexByID("A")) != 1.000 {
		t.Errorf("Expected 1 but\n%+v", data)
	}
}
