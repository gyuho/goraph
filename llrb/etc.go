package llrb

// deleteMax code for LLRB 2-3 trees
func deleteMax(nd *Node) (*Node, Interface) {
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}
	if !isRed(nd.Right) && !isRed(nd.Right.Left) {
		nd = moveRedToLeft(nd)
	}
	var deleted Interface
	nd.Right, deleted = deleteMax(nd.Right)
	return fixUp(nd), deleted
}

// DeleteMax deletes the maximum element in the tree and returns the
// deleted item or nil otherwise.
func (tr *Tree) DeleteMax() Interface {
	var deleted Interface
	tr.Root, deleted = deleteMax(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}
