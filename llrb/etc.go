package llrb

// DeleteMin deletes the minimum element in the tree and returns the
// deleted item or nil otherwise.
func (tr *Tree) DeleteMin() Interface {
	var deleted Interface
	tr.Root, deleted = DeleteMin(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

// DeleteMax code for LLRB 2-3 trees
func DeleteMax(nd *Node) (*Node, Interface) {
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}
	if !isRed(nd.Right) && !isRed(nd.Right.Left) {
		nd = MoveRedFromRightToLeft(nd)
	}
	var deleted Interface
	nd.Right, deleted = DeleteMax(nd.Right)
	return FixUp(nd), deleted
}

// DeleteMax deletes the maximum element in the tree and returns the
// deleted item or nil otherwise.
func (tr *Tree) DeleteMax() Interface {
	var deleted Interface
	tr.Root, deleted = DeleteMax(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}
