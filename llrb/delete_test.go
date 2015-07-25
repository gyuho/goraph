package llrb

import (
	"fmt"
	"testing"
)

func TestDeleteMin1(t *testing.T) {
	root := NewNode(String("S"))
	data := New(root)
	data.Insert(NewNode(String("E")))
	data.Insert(NewNode(String("A")))
	key := data.DeleteMin()
	if fmt.Sprintf("%v", key) != "A" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", data) != "[[E(false)] S(true)]" {
		t.Errorf("Unexpected %v", data)
	}
}

func TestDeleteMax1(t *testing.T) {
	root := NewNode(String("S"))
	data := New(root)
	data.Insert(NewNode(String("E")))
	data.Insert(NewNode(String("A")))
	key := data.DeleteMax()
	if fmt.Sprintf("%v", key) != "S" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", data) != "[[A(false)] E(true)]" {
		t.Errorf("Unexpected %v", data)
	}
}

func TestDelete1(t *testing.T) {
	root := NewNode(String("S"))
	data := New(root)
	data.Insert(NewNode(String("E")))
	data.Insert(NewNode(String("A")))
	key := data.Delete(String("E"))
	if fmt.Sprintf("%v", key) != "E" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", data) != "[[A(false)] S(true)]" {
		t.Errorf("Unexpected %v", data)
	}
}
