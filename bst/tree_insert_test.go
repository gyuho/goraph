package bst

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTreeNodeStruct(t *testing.T) {
	buf1 := new(bytes.Buffer)
	root1 := NewNode(nodeStruct{"A", 5})
	tr1 := New(root1)
	tr1.Insert(NewNode(nodeStruct{"B", 3}))
	tr1.Insert(NewNode(nodeStruct{"C", 17}))
	tr1.Insert(NewNode(nodeStruct{"D", 7}))
	tr1.Insert(NewNode(nodeStruct{"E", 1}))
	ch1 := make(chan string)
	go tr1.PreOrder(ch1)
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
	tr2 := New(root2)
	tr2.Insert(NewNode(nodeStruct{"B", 3}))
	tr2.Insert(NewNode(nodeStruct{"C", 17}))
	tr2.Insert(NewNode(nodeStruct{"D", 7}))
	tr2.Insert(NewNode(nodeStruct{"E", 1}))
	ch2 := make(chan string)
	go tr2.PreOrder(ch2)
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
	if !ComparePreOrder(tr1, tr2) {
		t.Error("Expected the same but %v | %v", tr1, tr2)
	}

	buf3 := new(bytes.Buffer)
	for _, elem := range tr2.LevelOrder() {
		buf3.WriteString(fmt.Sprintf("%v", elem.Key))
		buf3.WriteString(" ")
	}
	if buf3.String() != "A(5) B(3) C(17) E(1) D(7) " {
		t.Errorf("Unexpected %v", buf3.String())
	}
}
