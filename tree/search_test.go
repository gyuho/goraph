package tree

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	min := data.Min()
	if fmt.Sprintf("%v", min.Key) != "1" {
		t.Errorf("Unexpected %v", min.Key)
	}
}

func TestMax(t *testing.T) {
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	max := data.Max()
	if fmt.Sprintf("%v", max.Key) != "17" {
		t.Errorf("Unexpected %v", max.Key)
	}
}

func TestSearch(t *testing.T) {
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	nd1 := data.Search(Int(17))
	if fmt.Sprintf("%+v", Int(17)) != fmt.Sprintf("%+v", nd1.Key) {
		t.Errorf("Expected %v but %v", Int(17), nd1.Key)
	}
	nd2 := data.Search(Int(111))
	if nd2 != nil {
		t.Errorf("Expected nil but %v", nd2)
	}
}

func TestSearchChan(t *testing.T) {
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	ch1 := make(chan *Node)
	go data.SearchChan(Int(17), ch1)
	nd1 := <-ch1
	if fmt.Sprintf("%+v", Int(17)) != fmt.Sprintf("%+v", nd1.Key) {
		t.Errorf("Expected %v but %v", Int(17), nd1.Key)
	}
	ch2 := make(chan *Node)
	go data.SearchChan(Int(111), ch2)
	nd2 := <-ch2
	if nd2 != nil {
		t.Errorf("Expected nil but %v", nd2)
	}
}
