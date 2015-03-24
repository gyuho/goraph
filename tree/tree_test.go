package tree

import (
	"bytes"
	"fmt"
	"testing"
)

type nodeStruct struct {
	ID    string
	Value int
}

func (n nodeStruct) Less(b Interface) bool {
	return n.Value < b.(nodeStruct).Value
}

func (n nodeStruct) String() string {
	return fmt.Sprintf("%s(%d)", n.ID, n.Value)
}

func TestTreeNodeStruct(t *testing.T) {
	buf1 := new(bytes.Buffer)
	root1 := NewNode(nodeStruct{"A", 5})
	data1 := New(root1)
	data1.Insert(NewNode(nodeStruct{"B", 3}))
	data1.Insert(NewNode(nodeStruct{"C", 17}))
	data1.Insert(NewNode(nodeStruct{"D", 7}))
	data1.Insert(NewNode(nodeStruct{"E", 1}))
	ch1 := make(chan string)
	go data1.PreOrder(ch1)
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		buf1.WriteString(v)
		buf1.WriteString(" ")
	}

	buf2 := new(bytes.Buffer)
	root2 := NewNode(nodeStruct{"A", 5})
	data2 := New(root2)
	data2.Insert(NewNode(nodeStruct{"B", 3}))
	data2.Insert(NewNode(nodeStruct{"C", 17}))
	data2.Insert(NewNode(nodeStruct{"D", 7}))
	data2.Insert(NewNode(nodeStruct{"E", 1}))
	ch2 := make(chan string)
	go data2.PreOrder(ch2)
	for {
		v, ok := <-ch2
		if !ok {
			break
		}
		buf2.WriteString(v)
		buf2.WriteString(" ")
	}
	if buf1.String() != buf2.String() {
		t.Errorf("Expected the same but %s | %s", buf1.String(), buf2.String())
	}
	if !ComparePreOrder(data1, data2) {
		t.Error("Expected the same but %v | %v", data1, data2)
	}

	buf3 := new(bytes.Buffer)
	for _, elem := range data2.LevelOrder() {
		buf3.WriteString(fmt.Sprintf("%v", elem.Key))
		buf3.WriteString(" ")
	}
	if buf3.String() != "A(5) B(3) C(17) E(1) D(7) " {
		t.Errorf("Unexpected %v", buf3.String())
	}
}

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

func TestSearchChannel(t *testing.T) {
	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(3)))
	data.Insert(NewNode(Int(17)))
	data.Insert(NewNode(Int(7)))
	data.Insert(NewNode(Int(1)))
	ch1 := make(chan *Node)
	go data.SearchChannel(Int(17), ch1)
	nd1 := <-ch1
	if fmt.Sprintf("%+v", Int(17)) != fmt.Sprintf("%+v", nd1.Key) {
		t.Errorf("Expected %v but %v", Int(17), nd1.Key)
	}
	ch2 := make(chan *Node)
	go data.SearchChannel(Int(111), ch2)
	nd2 := <-ch2
	if nd2 != nil {
		t.Errorf("Expected nil but %v", nd2)
	}
}
