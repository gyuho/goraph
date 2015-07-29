package llrb

import "fmt"

// Left and Right children must be present
func moveRedToLeft(nd *Node) *Node {
	fmt.Println("moveRedToLeft:", nd.Key)
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
	fmt.Println("moveRedToRight:", nd.Key)
	flipColor(nd)
	if isRed(nd.Left.Left) {
		fmt.Println("moveRedToRight isRed(nd.Left.Left):", nd.Key)
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
		nd = moveRedToLeft(nd)
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
//	//
//	2. Call 'delete' method recursively on each Node from binary search path.
//		- e.g. If the key to delete is greater than Root's key
//			, call 'delete' on Right Node.
//	//
//	// start
//	3. 'tree.delete(Node, Key)'
//
//	// end
//	//
//	4. If the tree's Root is not nil, set Root Black.
//	//
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

			// then moveRedToLeft(nd)
			nd = moveRedToLeft(nd)
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

			// then moveRedToRight(nd)
			nd = moveRedToRight(nd)
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
			// if updated nd.Key is Less than key (nd.Key < key)
			fmt.Println("nd.Right, deleted = tr.delete(nd.Right, key) at", nd.Key)
			nd.Right, deleted = tr.delete(nd.Right, key)
		}
	}

	return fixUp(nd), deleted
}
