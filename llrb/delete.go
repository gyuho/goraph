package llrb

import "fmt"

// Left and Right children must be present
func moveRedFromRightToLeft(nd *Node) *Node {
	fmt.Println("moveRedFromRightToLeft:", nd.Key)
	flipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = rotateToRight(nd.Right)
		nd = rotateToLeft(nd)
		flipColor(nd)
	}
	return nd
}

// Left and Right children must be present
func moveRedFromLeftToRight(nd *Node) *Node {
	fmt.Println("moveRedFromLeftToRight:", nd.Key)
	flipColor(nd)
	if isRed(nd.Left.Left) {
		fmt.Println("moveRedFromLeftToRight isRed(nd.Left.Left):", nd.Key)
		nd = rotateToRight(nd)
		flipColor(nd)
	}
	return nd
}

func fixUp(nd *Node) *Node {
	fmt.Println("fixUp", nd.Key)
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
	fmt.Println("deleteMin", nd.Key)
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}
	if !isRed(nd.Left) && !isRed(nd.Left.Left) {
		nd = moveRedFromRightToLeft(nd)
	}
	var deleted Interface
	nd.Left, deleted = deleteMin(nd.Left)
	return fixUp(nd), deleted
}

// Delete deletes the node with a given key and returns the key.
// It returns nil if it does not exist in the tree.
//
//	Delete Algorithm:
//	1. Start 'delete' from tree Root.
//
//	2. Call 'delete' method recursively on each Node from binary search path.
//		- e.g. If the key to delete is greater than Root's key
//			, call 'delete' on Right Node.
//
//
//	# start
//
//	3. Recursive 'tree.delete(nd, key)'
//
//		if key < nd.Key:
//
//			if nd.Left is empty:
//				print "then nothing to delete, so return nil"
//				return nd, nil
//
//			if (nd.Left is Black) and (nd.Left.Left is Black):
//				print "then move Red from Right to Left to update nd"
//				nd = moveRedFromRightToLeft(nd)
//
//			print "recursively call 'delete' to update nd.Left"
//			nd.Left, deleted = tr.delete(nd.Left, key)
//
//		else if key >= nd.Key:
//
//			if nd.Left is Red:
//				print "then rotateToRight(nd) to update nd"
//				nd = rotateToRight(nd)
//
//			if (key == nd.Key) and nd.Right is empty:
//				print "then return nil, nd.Key to recursively update nd"
//				return nil, nd.Key
//
//			if (nd.Right is not empty)
//				and (nd.Right is Black)
//				and (nd.Right.Left is Black):
//				print "then move Red from Left to Right to update nd"
//				nd = moveRedFromLeftToRight(nd)
//
//			# nd gets updated by this step
//
//			if (key == nd.Key):
//				print "then deleteMin of nd.Right to update nd.Right"
//				nd.Right, subDeleted = deleteMin(nd.Right)
//
//				print "and then update nd.Key with deleteMin(nd.Right)"
//				deleted, nd.Key = nd.Key, subDeleted
//
//			else if key != nd.Key:
//				print "recursively call 'delete' to update nd.Right"
//				nd.Right, deleted = tr.delete(nd.Right, key)
//
//	# end
//
//
//	4. If the tree's Root is not nil, set Root Black.
//
//	5. Return the Interface(nil if the key does not exist.)
//
func (tr *Tree) Delete(key Interface) Interface {
	var deleted Interface
	tr.Root, deleted = tr.delete(tr.Root, key)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

func (tr *Tree) delete(nd *Node, key Interface) (*Node, Interface) {
	fmt.Println("calling delete on", nd.Key, "for the key", key)
	if nd == nil {
		return nil, nil
	}

	var deleted Interface

	// if key is Less than nd.Key
	if key.Less(nd.Key) {

		// if key is Less than nd.Key
		// and nd.Left is empty
		if nd.Left == nil {

			// then nothing to delete
			// so return the nil
			return nd, nil
		}

		// if key is Less than nd.Key
		// and nd.Left is Black
		// and nd.Left.Left is Black
		if !isRed(nd.Left) && !isRed(nd.Left.Left) {

			// then moveRedFromRightToLeft(nd)
			nd = moveRedFromRightToLeft(nd)
		}

		// and recursively call tr.delete(nd.Left, key)
		nd.Left, deleted = tr.delete(nd.Left, key)

	} else {
		// if key is not Less than nd.Key
		//(or key is greater than or equal to nd.Key)
		//(or key >= nd.Key)

		// and nd.Left is Red
		if isRed(nd.Left) {

			// then rotateToRight(nd)
			nd = rotateToRight(nd)
		}

		// and nd.Key is not Less than key
		// (or nd.Key >= key)
		// (or key == nd.Key)
		// and nd.Right is empty
		if !nd.Key.Less(key) && nd.Right == nil {
			fmt.Println("!nd.Key.Less(key) && nd.Right == nil when", nd.Key)
			// then return nil to delete the key
			return nil, nd.Key
		}

		// and nd.Right is not empty
		// and nd.Right is Black
		// and nd.Right.Left is Black
		if nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) {
			fmt.Println("nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when", nd.Key)

			// then moveRedFromLeftToRight(nd)
			nd = moveRedFromLeftToRight(nd)
		}

		// and key == nd.Key
		if !nd.Key.Less(key) {

			var subDeleted Interface

			// then deleteMin(nd.Right)
			nd.Right, subDeleted = deleteMin(nd.Right)
			if subDeleted == nil {
				panic("Unexpected nil value")
			}

			// and update nd.Key with deleteMin(nd.Right)
			deleted, nd.Key = nd.Key, subDeleted

		} else {
			// if updated nd.Key is Less than key (nd.Key < key) to update nd.Right
			fmt.Println("nd.Right, deleted = tr.delete(nd.Right, key) at", nd.Key)
			nd.Right, deleted = tr.delete(nd.Right, key)
		}
	}

	return fixUp(nd), deleted
}
