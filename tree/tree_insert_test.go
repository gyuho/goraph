package tree

import (
	"bytes"
	"fmt"
	"testing"
)

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
