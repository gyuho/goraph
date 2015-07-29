package llrb

import "fmt"

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	fmt.Println("RotateToLeft", nd.Key)
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	fmt.Println("RotateToRight", nd.Key)
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	fmt.Println("FlipColor", nd.Key)
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// MoveRedFromRightToLeft moves Red Node
// from Right sub-Tree to Left sub-Tree.
// Left and Right children must be present
func MoveRedFromRightToLeft(nd *Node) *Node {
	fmt.Println("MoveRedFromRightToLeft", nd.Key)
	FlipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = RotateToRight(nd.Right)
		nd = RotateToLeft(nd)
		FlipColor(nd)
	}
	fmt.Println("returning", nd.Key)
	return nd
}

// MoveRedFromLeftToRight moves Red Node
// from Left sub-Tree to Right sub-Tree.
// Left and Right children must be present
func MoveRedFromLeftToRight(nd *Node) *Node {
	fmt.Println("MoveRedFromLeftToRight", nd.Key)
	FlipColor(nd)
	if isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
		FlipColor(nd)
	}
	return nd
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	fmt.Println("Balance", nd.Key)
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

// FixUp fixes the balances of the Node.
func FixUp(nd *Node) *Node {
	fmt.Println("FixUp", nd.Key)
	if isRed(nd.Right) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}
