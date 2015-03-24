package redblack

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

// Left and Right children must be present
func moveRedToLeft(nd *Node) *Node {
	flipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = rotateToRight(nd.Right)
		nd = rotateToLeft(nd)
		flipColor(nd)
	}
	return nd
}

// Left and Right children must be present
func moveRedToRight(nd *Node) *Node {
	flipColor(nd)
	if isRed(nd.Left.Left) {
		nd = rotateToRight(nd)
		flipColor(nd)
	}
	return nd
}

// DeleteMin deletes the minimum element in the tree and returns the
// deleted item or nil otherwise.
func (d *Data) DeleteMin() Interface {
	var deleted Interface
	d.Root, deleted = deleteMin(d.Root)
	if d.Root != nil {
		d.Root.Black = true
	}
	return deleted
}

// deleteMin code for LLRB 2-3 trees
func deleteMin(nd *Node) (*Node, Interface) {
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}

	if !isRed(nd.Left) && !isRed(nd.Left.Left) {
		nd = moveRedToLeft(nd)
	}

	var deleted Interface
	nd.Left, deleted = deleteMin(nd.Left)

	return fixUp(nd), deleted
}
