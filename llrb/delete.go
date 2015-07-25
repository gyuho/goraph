package llrb

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
func (d *Data) DeleteMin() Interface {
	var deleted Interface
	d.Root, deleted = deleteMin(d.Root)
	if d.Root != nil {
		d.Root.Black = true
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
func (d *Data) DeleteMax() Interface {
	var deleted Interface
	d.Root, deleted = deleteMax(d.Root)
	if d.Root != nil {
		d.Root.Black = true
	}
	return deleted
}

// Delete deletes the node with a given key and returns the key.
// It returns nil if it does not exist in the tree.
func (d *Data) Delete(key Interface) Interface {
	var deleted Interface
	d.Root, deleted = d.delete(d.Root, key)
	if d.Root != nil {
		d.Root.Black = true
	}
	return deleted
}

func (d *Data) delete(nd *Node, key Interface) (*Node, Interface) {
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
		nd.Left, deleted = d.delete(nd.Left, key)
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
			nd.Right, deleted = d.delete(nd.Right, key)
		}
	}

	return fixUp(nd), deleted
}
