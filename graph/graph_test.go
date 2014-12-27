package graph

import (
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
	if data.GetVertexSize() != 3 {
		t.Errorf("Expected 3 but %+v\n", data)
	}
}

func TestInit(t *testing.T) {

}

func TestGetVertexSize(t *testing.T) {

}

func TestString(t *testing.T) {

}

func TestFindVertexByID(t *testing.T) {

}

func TestDeleteVertex(t *testing.T) {

}

func TestDeleteEdge(t *testing.T) {

}

func TestClone(t *testing.T) {

}
