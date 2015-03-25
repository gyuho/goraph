package radix

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	nd := NewNode()
	if reflect.TypeOf(nd) != reflect.TypeOf(&Node{}) {
		t.Error(nd)
	}
}

func TestSearchPrefix(t *testing.T) {
	e1 := Edge{Label: []byte("aaa")}
	e2 := Edge{Label: []byte("bbb")}
	e3 := Edge{Label: []byte("ccc")}
	e4 := Edge{Label: []byte("golang")}
	e5 := Edge{Label: []byte("clang")}
	edges := []Edge{e1, e2, e3, e4, e5}
	edg, idx := Edges(edges).SearchPrefix([]byte("golangaaa"))
	if string(edg.Label) != "golang" {
		t.Errorf("Unexpected %s", edg.Label)
	}
	if idx != 5 {
		t.Errorf("Unexpected %d", idx)
	}
}
