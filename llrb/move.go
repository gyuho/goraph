package llrb

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
