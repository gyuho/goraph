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

func TestFlipColor(t *testing.T) {
	root := NewNode(Int(3))
	data := New(root)
	data.Insert(NewNode(Int(1)))
	data.Insert(NewNode(Int(2)))
	data.Insert(NewNode(Int(5)))
	data.Insert(NewNode(Int(6)))
	data.Insert(NewNode(Int(4)))
	nd := data.Search(Int(5))
	if fmt.Sprintf("%v", nd.Key) != "5" {
		t.Errorf("Unexpected %v", nd.Key)
	}
	if fmt.Sprintf("%v", nd.Left.Key) != "4" {
		t.Errorf("Unexpected %v", nd.Left.Key)
	}
	if nd.Left.Black {
		t.Errorf("Left should be red but %v", nd.Left.Black)
	}
	if nd.Right.Black {
		t.Errorf("Right should be red but %v", nd.Right.Black)
	}
	if fmt.Sprintf("%v", nd.Right.Key) != "6" {
		t.Errorf("Unexpected %v", nd.Right.Key)
	}
	flipColor(nd)
	if !nd.Left.Black {
		t.Errorf("Left should be black but %v", nd.Left.Black)
	}
	if !nd.Right.Black {
		t.Errorf("Right should be black but %v", nd.Right.Black)
	}
}
