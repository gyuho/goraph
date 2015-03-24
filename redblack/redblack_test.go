package redblack

import (
	"bytes"
	"fmt"
	"testing"
)

type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

func TestTreeIntPreOrder(t *testing.T) {
	buf := new(bytes.Buffer)
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	ch := make(chan string)
	go data.PreOrder(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		buf.WriteString(v)
		buf.WriteString(" ")
	}
	if buf.String() != "5 (Black:true) 3 (Black:false) 1 (Black:false) 17 (Black:false) 7 (Black:false) " {
		t.Errorf("Unexpected %v", buf.String())
	}

	root2 := NewNode(Int(5))
	data2 := New(root2)
	data2.Insert(NewNode(Int(3)))
	data2.Insert(NewNode(Int(17)))
	data2.Insert(NewNode(Int(7)))
	data2.Insert(NewNode(Int(1)))
	if !ComparePreOrder(data, data2) {
		t.Errorf("Expected true but %v", data2)
	}
}

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

func TestFlipColor(t *testing.T) {
	root := NewNode(Int(3))
	data := New(root)
	data.Insert(NewNode(Int(1)))
	data.Insert(NewNode(Int(2)))
	data.Insert(NewNode(Int(5)))
	data.Insert(NewNode(Int(6)))
	data.Insert(NewNode(Int(4)))
	nd1 := data.Search(Int(5))
	if fmt.Sprintf("%v", nd1.Key) != "5" {
		t.Errorf("Unexpected %v", nd1.Key)
	}
	if fmt.Sprintf("%v", nd1.Left.Key) != "4" {
		t.Errorf("Unexpected %v", nd1.Left.Key)
	}
	if nd1.Left.Black {
		t.Errorf("Left should be red but %v", nd1.Left.Black)
	}
	if fmt.Sprintf("%v", nd1.Right.Key) != "6" {
		t.Errorf("Unexpected %v", nd1.Right.Key)
	}
	// FlipColor(nd1)
	// if !nd1.Left.Black {
	// 	t.Errorf("Left should be black but %v", nd1.Left.Black)
	// }
	// if !nd1.Right.Black {
	// 	t.Errorf("Right should be black but %v", nd1.Right.Black)
	// }
}
