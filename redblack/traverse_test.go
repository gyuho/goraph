package redblack

import (
	"bytes"
	"fmt"
	"testing"
)

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
	if buf.String() != "5 3 1 17 7 " {
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

func TestTreeIntInOrder(t *testing.T) {
	buf := new(bytes.Buffer)
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	ch := make(chan string)
	go data.InOrder(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		buf.WriteString(v)
		buf.WriteString(" ")
	}
	if buf.String() != "1 3 5 7 17 " {
		t.Errorf("Unexpected %v", buf.String())
	}

	root2 := NewNode(Int(5))
	data2 := New(root2)
	data2.Insert(NewNode(Int(3)))
	data2.Insert(NewNode(Int(17)))
	data2.Insert(NewNode(Int(7)))
	data2.Insert(NewNode(Int(1)))
	if !CompareInOrder(data, data2) {
		t.Errorf("Expected true but %v", data2)
	}
}

func TestTreeIntPostOrder(t *testing.T) {
	buf := new(bytes.Buffer)
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	ch := make(chan string)
	go data.PostOrder(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		buf.WriteString(v)
		buf.WriteString(" ")
	}
	if buf.String() != "1 3 7 17 5 " {
		t.Errorf("Unexpected %v", buf.String())
	}

	root2 := NewNode(Int(5))
	data2 := New(root2)
	data2.Insert(NewNode(Int(3)))
	data2.Insert(NewNode(Int(17)))
	data2.Insert(NewNode(Int(7)))
	data2.Insert(NewNode(Int(1)))
	if !ComparePostOrder(data, data2) {
		t.Errorf("Expected true but %v", data2)
	}
}

func TestTreeIntLevelOrder(t *testing.T) {
	buf := new(bytes.Buffer)
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	nodes := data.LevelOrder()
	for _, v := range nodes {
		buf.WriteString(fmt.Sprintf("%v ", v.Key))
	}
	if buf.String() != "5 3 17 1 7 " {
		t.Errorf("Unexpected %v", buf.String())
	}
}
