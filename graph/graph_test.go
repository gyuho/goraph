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

}

func TestConnect(t *testing.T) {

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
