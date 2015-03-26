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
	e1 := &Edge{Label: []byte("aaa")}
	e2 := &Edge{Label: []byte("bbb")}
	e3 := &Edge{Label: []byte("ccc")}
	e4 := &Edge{Label: []byte("golang")}
	edges := []*Edge{e1, e2, e3, e4}
	edge, edgeIdx, commonIdx, option := Edges(edges).searchPrefix([]byte("golangaaa"))
	if option != insertWithSplitValue {
		t.Error("Need Insert!")
	}
	if string(edge.Label) != "golang" {
		t.Errorf("Unexpected %s", edge.Label)
	}
	if edgeIdx != 3 {
		t.Errorf("Unexpected %d", edgeIdx)
	}
	if commonIdx != 5 {
		t.Errorf("Unexpected %d", commonIdx)
	}
}

func TestSplitEdge(t *testing.T) {
	edge := &Edge{Label: []byte("golangaaa")}
	edgeTop := splitEdge(edge, 5+1)
	if string(edgeTop.Label) != "golang" {
		t.Errorf("Unexpected %s", string(edgeTop.Label))
	}
	if string(edgeTop.ChildNode.Edges[0].Label) != "aaa" {
		t.Errorf("Unexpected %s", string(edgeTop.ChildNode.Edges[0].Label))
	}
	if edgeTop.ChildNode.Edges[0].ChildNode == nil {
		t.Errorf("Unexpected nil but %v", edgeTop.ChildNode.Edges[0].ChildNode)
	}
}

func TestInsert1(t *testing.T) {
	e1 := &Edge{Label: []byte("aaa")}
	e2 := &Edge{Label: []byte("bbb")}
	e3 := &Edge{Label: []byte("ccc")}
	e4 := &Edge{Label: []byte("golang")}
	edges := Edges([]*Edge{e1, e2, e3, e4})
	edges.insert([]byte("zzz"))
	if len(edges) != 5 {
		t.Errorf("Expected 6 but %d", len(edges))
	}
	edges.insert([]byte("golangxxx"))
	if len(edges) != 5 {
		t.Errorf("Expected 6 but %d", len(edges))
	}
	if string(edges[3].ChildNode.Edges[0].Label) != "xxx" {
		t.Errorf("Expected xxx but %s -> %s", edges[4].Label, edges[4].ChildNode.Edges[0].Label)
	}
}
