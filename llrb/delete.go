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

// DeleteMin deletes the minimum element in the tree and returns the
// deleted item or nil otherwise.
func (tr *Tree) DeleteMin() Interface {
	var deleted Interface
	tr.Root, deleted = deleteMin(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

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

// Delete deletes the node with a given key and returns the key.
// It returns nil if it does not exist in the tree.
func (tr *Tree) Delete(key Interface) Interface {
	var deleted Interface
	tr.Root, deleted = tr.delete(tr.Root, key)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

func (tr *Tree) delete(nd *Node, key Interface) (*Node, Interface) {
	var deleted Interface
	if nd == nil {
		return nil, nil
	}
	if key.Less(nd.Key) {
		if nd.Left == nil {
			// nothing to delete
			return nd, nil
		}
		if !isRed(nd.Left) && !isRed(nd.Left.Left) {
			nd = moveRedToLeft(nd)
		}
		nd.Left, deleted = tr.delete(nd.Left, key)
	} else {
		if isRed(nd.Left) {
			nd = rotateToRight(nd)
		}
		if !nd.Key.Less(key) && nd.Right == nil {
			return nil, nd.Key
		}
		if nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) {
			nd = moveRedToRight(nd)
		}
		if !nd.Key.Less(key) {
			var subDeleted Interface
			nd.Right, subDeleted = deleteMin(nd.Right)
			if subDeleted == nil {
				panic("Unexpected nil value")
			}
			deleted, nd.Key = nd.Key, subDeleted
		} else {
			nd.Right, deleted = tr.delete(nd.Right, key)
		}
	}

	return fixUp(nd), deleted
}
