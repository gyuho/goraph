package llrb

import "fmt"

// RotateToLeft when there is a right-leaning link.
func RotateToLeft(nd *Node) *Node {
	fmt.Println("RotateToLeft:", nd.Key)
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

// RotateToRight when there are two left red links in a row.
func RotateToRight(nd *Node) *Node {
	fmt.Println("RotateToRight:", nd.Key)
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
	fmt.Println("FlipColor:", nd.Key)
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
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
