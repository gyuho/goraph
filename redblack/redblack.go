package redblack

import "fmt"

// Data contains a Root node of a binary search tree.
type Data struct {
	Root *Node
}

// New returns a new Data with its root Node.
func New(root *Node) *Data {
	d := &Data{}
	root.Black = true
	d.Root = root
	// d.Root.Black = true
	return d
}

// Interface represents a single object in the tree.
type Interface interface {
	// Less returns true when the receiver item(key)
	// is less than the given(than) argument.
	Less(than Interface) bool
}

// Node is a Node and a Tree itself.
type Node struct {
	// Left is a left child Node.
	Left *Node

	Key   Interface
	Black bool

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
	return nd
}

// Insert inserts a Node to a Data.
func (d *Data) Insert(nd *Node) {
	if d.Root == nd {
		return
	}
	d.Root = d.Root.insert(nd)
}

func (nd *Node) insert(node *Node) *Node {
	if nd == nil {
		return node
	}
	if nd.Key.Less(node.Key) {
		nd.Right = nd.Right.insert(node)
		return nd
	}
	nd.Left = nd.Left.insert(node)
	return nd
}

// Min returns the minimum key Node in the tree.
func (d Data) Min() *Node {
	nd := d.Root
	if nd == nil {
		return nil
	}
	for nd.Left != nil {
		nd = nd.Left
	}
	return nd
}

// Max returns the maximum key Node in the tree.
func (d Data) Max() *Node {
	nd := d.Root
	if nd == nil {
		return nil
	}
	for nd.Right != nil {
		nd = nd.Right
	}
	return nd
}

// Search does binary-search on a given key and returns the first Node with the key.
func (d Data) Search(key Interface) *Node {
	nd := d.Root
	for nd != nil {
		switch {
		case nd.Key.Less(key):
			nd = nd.Right
		case key.Less(nd.Key):
			nd = nd.Left
		default:
			return nd
		}
	}
	return nil
}

// PreOrder traverses from Root, Left-SubTree, and Right-SubTree. (DFS)
func (d *Data) PreOrder(ch chan string) {
	preOrder(d.Root, ch)
	close(ch)
}

func preOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	ch <- fmt.Sprintf("%v (Black:%v)", nd.Key, nd.Black) // root
	preOrder(nd.Left, ch)                                // left
	preOrder(nd.Right, ch)                               // right
}

// ComparePreOrder returns true if two Trees are same with PreOrder.
func ComparePreOrder(d1, d2 *Data) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go d1.PreOrder(ch1)
	go d2.PreOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}
