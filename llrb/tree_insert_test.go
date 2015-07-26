package llrb

import (
	"fmt"
	"testing"
)

func TestIsRed(t *testing.T) {
	if !isRed(NewNode(Int(5))) {
		t.Error("Expected Red")
	}
}

func TestInsert1(t *testing.T) {
	root := NewNode(String("S"))
	data := New(root)
	data.Insert(NewNode(String("E")))
	data.Insert(NewNode(String("A")))
	nd := data.Search(String("E"))
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
	root := NewNode(String("S"))
	data := New(root)
	data.Insert(NewNode(String("E")))
	data.Insert(NewNode(String("A")))
	data.Insert(NewNode(String("R")))
	data.Insert(NewNode(String("C")))
	data.Insert(NewNode(String("H")))
	data.Insert(NewNode(String("X")))
	data.Insert(NewNode(String("M")))
	data.Insert(NewNode(String("P")))
	data.Insert(NewNode(String("L")))
	nd := data.Search(String("E"))
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

func TestBalanceInsert1(t *testing.T) {
	root := NewNode(Float64(1))
	data := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		data.Insert(NewNode(Float64(num)))
	}
}

func TestBalanceInsert2(t *testing.T) {
	root := NewNode(Float64(1))
	data := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		data.Insert(NewNode(Float64(num)))
	}
}
