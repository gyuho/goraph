package llrb

import (
	"fmt"
	"testing"
)

func TestDeleteMin1(t *testing.T) {
	root := NewNode(String("S"))
	tr := New(root)
	tr.Insert(NewNode(String("E")))
	tr.Insert(NewNode(String("A")))
	key := tr.DeleteMin()
	if fmt.Sprintf("%v", key) != "A" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", tr) != "[[E(false)] S(true)]" {
		t.Errorf("Unexpected %v", tr)
	}
}

func TestDeleteMax1(t *testing.T) {
	root := NewNode(String("S"))
	tr := New(root)
	tr.Insert(NewNode(String("E")))
	tr.Insert(NewNode(String("A")))
	key := tr.DeleteMax()
	if fmt.Sprintf("%v", key) != "S" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", tr) != "[[A(false)] E(true)]" {
		t.Errorf("Unexpected %v", tr)
	}
}

func TestDelete1(t *testing.T) {
	root := NewNode(String("S"))
	tr := New(root)
	tr.Insert(NewNode(String("E")))
	tr.Insert(NewNode(String("A")))
	key := tr.Delete(String("E"))
	if fmt.Sprintf("%v", key) != "E" {
		t.Errorf("Expected A but %v", key)
	}
	if fmt.Sprintf("%v", tr) != "[[A(false)] S(true)]" {
		t.Errorf("Unexpected %v", tr)
	}
}
