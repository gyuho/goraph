package llrb

// Delete deletes a Node from a tree
// Just like deletion in binary search tree,
// this maintains the balance and Left-leaning tree properties by:
//	1. Root node is always black.
//	2. Rebalance the tree as of its parental node.
// (Another implementation can be found https://github.com/petar/GoLLRB).
func (tr *Tree) Delete(nd *Node) {
	if nd == nil {
		return
	}
	parent := tr.SearchParent(nd.Key)

	// you need to dereference the pointer
	// and update with a value
	// in order to change the original struct

	if nd.Left != nil && nd.Right != nil {
		// if two children

		// #1. Find the node to substitute
		// the to-be-deleted node
		//
		// either get the biggest of left sub-tree
		tmp := new(Tree)
		tmp.Root = nd.Left
		tmpNode := tmp.Max()
		//
		// OR
		//
		// get the smallest of right sub-tree
		// tmp := new(Data)
		// tmp.Root = nd.Right
		// tmpNode := nd.Right.Min()
		//
		replacingNode := tr.Search(tmpNode.Key)
		parentOfReplacingNode := tr.SearchParent(replacingNode.Key)

		// order matters!
		if replacingNode.Key.Less(nd.Key) {
			// replacing with the left child
			replacingNode.Right = nd.Right

			// inherit the sub-tree
			if nd.Left.Key.Less(replacingNode.Key) ||
				replacingNode.Key.Less(nd.Left.Key) {
				// if different
				replacingNode.Left = nd.Left

				// destroy the old pointer in sub-tree
				if parentOfReplacingNode.Key.Less(replacingNode.Key) {
					// deleting right child of parentOfReplacingNode
					parentOfReplacingNode.Right = nil
				} else {
					// deleting left child of parentOfReplacingNode
					parentOfReplacingNode.Left = nil
				}
			}

		} else {
			// replacing with the right child
			replacingNode.Left = nd.Left

			// inherit the sub-tree
			if nd.Right.Key.Less(replacingNode.Key) ||
				replacingNode.Key.Less(nd.Right.Key) {

				// destroy the old pointer in sub-tree
				if parentOfReplacingNode.Key.Less(replacingNode.Key) {
					// deleting right child of parentOfReplacingNode
					parentOfReplacingNode.Right = nil
				} else {
					// deleting left child of parentOfReplacingNode
					parentOfReplacingNode.Left = nil
				}
			}
		}

		// #2. Update the parent, child node
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = replacingNode
		} else {
			if parent.Key.Less(nd.Key) {
				// deleting right child of parent
				parent.Right = replacingNode
			} else {
				// deleting left child of parent
				parent.Left = replacingNode
			}
		}

	} else if nd.Left != nil && nd.Right == nil {
		// only left child
		// #1. Update the parent node
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = nd.Left
		} else {
			if parent.Key.Less(nd.Key) {
				// right child of parent
				parent.Right = nd.Left
			} else {
				// left child of parent
				parent.Left = nd.Left
			}
		}

	} else if nd.Left == nil && nd.Right != nil {
		// only right child
		// #1. Update the parent node
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = nd.Right
		} else {
			if parent.Key.Less(nd.Key) {
				// right child of parent
				parent.Right = nd.Right
			} else {
				// left child of parent
				parent.Left = nd.Right
			}
		}
	} else {
		// no child
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = nil
		} else {
			if parent.Key.Less(nd.Key) {
				// right child of parent
				parent.Right = nil
			} else {
				// left child of parent
				parent.Left = nil
			}
		}
	}

	// At the end, delete the node
	// this is not necessary because it will be
	// garbage collected
	*nd = Node{}

	// because this is inside function
	// this won't change the actual node
	//
	// nd = new(Node)
	// nd = nil

	// Root node must be always black.
	tr.Root.Black = true
	if parent != nil {
		parent = balance(parent)
	} else {
		tr.Root = balance(tr.Root)
	}
}
