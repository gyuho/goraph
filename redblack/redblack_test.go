package redblack

import (
	"fmt"
	"testing"
)

type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

func TestIsRed(t *testing.T) {
	if !isRed(NewNode(Int(5))) {
		t.Error("Expected Red")
	}
}

type Str string

// Less returns true if string(a) < string(b).
func (a Str) Less(b Interface) bool {
	return a < b.(Str)
}

func TestInsert1(t *testing.T) {
	root := NewNode(Str("S"))
	data := New(root)
	data.Insert(NewNode(Str("E")))
	data.Insert(NewNode(Str("A")))
	nd := data.Search(Str("E"))
	if fmt.Sprintf("%v", nd.Key) != "E" {
		t.Errorf("Unexpected %v", nd.Key)
	}
	if fmt.Sprintf("%v", nd.Left.Key) != "A" {
		t.Errorf("Unexpected %v", nd.Left.Key)
	}
	if fmt.Sprintf("%v", nd.Right.Key) != "S" {
		t.Errorf("Unexpected %v", nd.Right.Key)
	}
	if !nd.Left.Black {
		t.Errorf("Left should be black but %v", nd.Left.Black)
	}
	if !nd.Right.Black {
		t.Errorf("Right should be black but %v", nd.Right.Black)
	}
}

func TestInsert2(t *testing.T) {
	root := NewNode(Str("S"))
	data := New(root)
	data.Insert(NewNode(Str("E")))
	data.Insert(NewNode(Str("A")))
	data.Insert(NewNode(Str("R")))
	data.Insert(NewNode(Str("C")))
	data.Insert(NewNode(Str("H")))
	data.Insert(NewNode(Str("X")))
	data.Insert(NewNode(Str("M")))
	data.Insert(NewNode(Str("P")))
	data.Insert(NewNode(Str("L")))
	nd := data.Search(Str("E"))
	if fmt.Sprintf("%v", nd.Key) != "E" {
		t.Errorf("Unexpected %v", nd.Key)
	}
	if fmt.Sprintf("%v", nd.Left.Key) != "C" {
		t.Errorf("Unexpected %v", nd.Left.Key)
	}
	if fmt.Sprintf("%v", nd.Right.Key) != "L" {
		t.Errorf("Unexpected %v", nd.Right.Key)
	}
	if !nd.Left.Black {
		t.Errorf("Left should be black but %v", nd.Left.Black)
	}
	if nd.Left.Left.Black {
		t.Errorf("Left, Left should be red but %v", nd.Left.Left.Black)
	}
	if !nd.Right.Black {
		t.Errorf("Right should be black but %v", nd.Right.Black)
	}
	if nd.Right.Left.Black {
		t.Errorf("Right, Left should be red but %v", nd.Right.Black)
	}
	if fmt.Sprintf("%v", nd.Right.Left.Key) != "H" {
		t.Errorf("Unexpected %v", nd.Right.Left.Key)
	}
	if fmt.Sprintf("%v", data.Root.Key) != "M" {
		t.Errorf("Unexpected %v", data.Root.Key)
	}
	if fmt.Sprintf("%v", data.Root.Right.Right.Key) != "X" {
		t.Errorf("Unexpected %v", data.Root.Right.Right.Key)
	}
	if fmt.Sprintf("%v", data.Root.Right.Right.Left.Key) != "S" {
		t.Errorf("Unexpected %v", data.Root.Right.Right.Left.Key)
	}
	if fmt.Sprintf("%v", data.Root.Left.Key) != "E" {
		t.Errorf("Unexpected %v", data.Root.Left.Key)
	}
	if fmt.Sprintf("%v", data.Root.Right.Left.Key) != "P" {
		t.Errorf("Unexpected %v", data.Root.Right.Left.Key)
	}
}

func TestDeleteMin1(t *testing.T) {
	root := NewNode(Str("S"))
	data := New(root)
	data.Insert(NewNode(Str("E")))
	data.Insert(NewNode(Str("A")))
	key := data.DeleteMin()
	if fmt.Sprintf("%v", key) != "A" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", data) != "[[E(false)] S(true)]" {
		t.Errorf("Unexpected %v", data)
	}
}

func TestDeleteMax1(t *testing.T) {
	root := NewNode(Str("S"))
	data := New(root)
	data.Insert(NewNode(Str("E")))
	data.Insert(NewNode(Str("A")))
	key := data.DeleteMax()
	if fmt.Sprintf("%v", key) != "S" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", data) != "[[A(false)] E(true)]" {
		t.Errorf("Unexpected %v", data)
	}
}

func TestDelete1(t *testing.T) {
	root := NewNode(Str("S"))
	data := New(root)
	data.Insert(NewNode(Str("E")))
	data.Insert(NewNode(Str("A")))
	key := data.Delete(Str("E"))
	if fmt.Sprintf("%v", key) != "E" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", data) != "[[A(false)] S(true)]" {
		t.Errorf("Unexpected %v", data)
	}
}
