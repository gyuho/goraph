package llrb

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
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

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

func (d *Data) String() string {
	return d.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// rotateToLeft when there is a right-leaning link.
func rotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// rotateToRight when there are two left red links in a row.
// Then flip color.
func rotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// flipColor flips the color.
// Left and Right children must be present
func flipColor(nd *Node) {
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

func balance(nd *Node) *Node {
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = rotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = rotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		flipColor(nd)
	}
	return nd
}

func fixUp(nd *Node) *Node {
	if isRed(nd.Right) {
		nd = rotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = rotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		flipColor(nd)
	}
	return nd
}

// Insert inserts a Node to a Data without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
func (d *Data) Insert(nd *Node) {
	if d.Root == nd {
		return
	}
	d.Root = d.Root.insert(nd)
	d.Root.Black = true
}

func (nd *Node) insert(node *Node) *Node {
	if nd == nil {
		return node
	}
	if nd.Key.Less(node.Key) {
		nd.Right = nd.Right.insert(node)
	} else {
		nd.Left = nd.Left.insert(node)
	}
	return balance(nd)
}
