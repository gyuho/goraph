// Pacakge avl implements an AVL tree.
package avl

import (
	"math"
	"sort"
)

// Tree is for binary tree.
type Tree struct {
	Left  *Tree
	Value int64
	Right *Tree
	Size  int64
}

// NewTree returns a new tree of input Value.
func NewTree(val int64) *Tree {
	return &Tree{
		Left:  nil, // (X) Left: new(Tree),
		Value: val,
		Right: nil, // (X) Right: new(Tree),
		Size:  1,
	}
}

// Inserts implements Insert with a variadic function.
func (T *Tree) Inserts(values ...int64) *Tree {
	for _, v := range values {
		T.Insert(v)
	}
	return T
}

// Insert inserts a new value(node) to the tree.
func (T *Tree) Insert(val int64) *Tree {
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	if T != nil && T.Value != val {
		T.Size += 1
	}
	if val < T.Value {
		T.Left = T.Left.Insert(val)
	} else if val > T.Value {
		T.Right = T.Right.Insert(val)
	}
	return T
}

// Find does Binary Search to find the value
// and returns Tree with the value as a root node.
func (T *Tree) Find(val int64) *Tree {
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	// To end recursion
	// set terminal node's left and right to nil
	if T.Value == val {
		return T
	}
	if val < T.Value {
		// Not working if we only have
		// T.Left.Find(val)
		return T.Left.Find(val)
	} else {
		return T.Right.Find(val)
	}
	return T
}

// SetValue updates the Value of the Tree.
func (T *Tree) SetValue(val int64) {
	T.Value = val
}

// Parent returns the parental Tree(node) of input value.
func (T *Tree) Parent(val int64) *Tree {
	// if the input value is root
	if val == T.Value {
		return T
	}
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	// we need to check if T.Left is nil or not
	// otherwise, it panics with the message:
	// panic: runtime error: invalid memory address
	// or nil pointer dereference [recovered]
	if T.Left != nil {
		if T.Left.Value == val {
			return T
		}
	}
	if T.Right != nil {
		if T.Right.Value == val {
			return T
		}
	}
	if val < T.Value {
		return T.Left.Parent(val)
	} else {
		return T.Right.Parent(val)
	}
	return T
}

// IsRoot returns true if the Node(tree) is a root of the tree.
func (T *Tree) IsRoot(val int64) bool {
	if T.Parent(val).Value == val {
		return true
	} else {
		return false
	}
}

// IsLeaf returns true if the Node(tree) is a leaf.
func (T *Tree) IsLeaf(val int64) bool {
	nd := T.Find(val)
	if nd == nil {
		return false
	}
	if nd.Left == nil && nd.Right == nil {
		return true
	} else {
		return false
	}
}

// Copy returns a copy of the tree.
func (T *Tree) Copy() *Tree {
	t := NewTree(T.Value)
	t.Left = T.Left
	t.Right = T.Right
	return t
}

// FindMin returns the minimum(left-most) value of the tree.
func (T *Tree) FindMin() int64 {
	curT := T.Copy()
	for curT.Left != nil {
		curT = curT.Left
	}
	return curT.Value
}

// FindMax returns the maximum(right-most) value of the tree.
func (T *Tree) FindMax() int64 {
	curT := T.Copy()
	for curT.Right != nil {
		curT = curT.Right
	}
	return curT.Value
}

// GetSize returns the size of the node with the input value
// which is the number of the children node + 1.
func (T *Tree) GetSize(val int64) int64 {
	if T.Value == val {
		return T.Size
	}
	node := T.Find(val)
	return node.Size
}

// GetHeight returns the height of the node with the input value.
func (T *Tree) GetHeight(val int64) int64 {
	// Height h = ⌊log_2 n⌋
	// n = 15 (15 nodes), then the height of the root node
	// is = ⌊log_2 15⌋ = 3

	// in order to truncate
	// need to make a copy by passing as a parameter
	float64ToInt64 := func(num float64) int64 {
		return int64(num)
	}
	tree := T.Find(val)
	return float64ToInt64(math.Floor(math.Log2(float64(tree.GetSize(val)))))
}

// GetHeightLeft returns the height of the left sub-tree
// of the input value(node).
func (T *Tree) GetHeightLeft(val int64) int64 {
	tree := T.Find(val)
	if tree.Left == nil {
		return 0
	}
	h := T.GetHeight(tree.Left.Value)
	return h + 1
}

// GetHeightRight returns the height of the right sub-tree
// of the input value(node).
func (T *Tree) GetHeightRight(val int64) int64 {
	tree := T.Find(val)
	if tree.Right == nil {
		return 0
	}
	// (X) T.GetHeight(tree.Value)
	h := T.GetHeight(tree.Right.Value)
	return h + 1
}

// Height returns the difference between
// GetHeightLeft and GetHeightRight.
func (T *Tree) Height(val int64) int64 {
	return T.GetHeightLeft(val) - T.GetHeightRight(val)
}

// IsBalanced returns true if the Height of the Tree
// with the input value is balanced.
func (T *Tree) IsBalanced(val int64) bool {
	Tree := T.Find(val)
	return -1 <= Tree.Height(val) && Tree.Height(val) <= 1
}

// CheckTreeBalance returns true if all nodes in the Tree
// is balanced. Otherwise, it returns false and the unbalanced
// nodes.
func (T *Tree) CheckTreeBalance() (bool, []int64) {
	ch := make(chan int64)
	go WalkInOrder(T, ch)
	result := []int64{}
	for v := range ch {
		result = append(result, v)
	}
	nbs := []int64{}
	for _, v := range result {
		if !T.IsBalanced(v) {
			nbs = append(nbs, v)
		}
	}
	if len(nbs) != 0 {
		return false, nbs
	} else {
		return true, nbs
	}
}

// BalanceInsert inserts one value to a Tree
// and tells if the Tree is LL, RR, LR, RL.
func (T *Tree) BalanceInsert(val int64) (*Tree, string) {
	T.Insert(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	if T.IsBalanced(Parent.Value) {
		return T, "Balanced"
	} else {
		switch T.Height(Parent.Value) {
		case 2: // LL or LR
			if Parent.Left.Left != nil && Parent.Left.Right == nil {
				return T, "LL"
			}
			if Parent.Left.Left == nil && Parent.Left.Right != nil {
				return T, "LR"
			}
		case -2: // RR or RL
			if Parent.Right.Right != nil && Parent.Right.Left == nil {
				return T, "RR"
			}
			if Parent.Right.Right == nil && Parent.Right.Left != nil {
				return T, "RL"
			}
		}
	}
	return T, "None"
}

type Int64Slice []int64

func (p Int64Slice) Len() int {
	return len(p)
}
func (p Int64Slice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// int64Sort sorts 3 integers and return the slice in order.
func int64Sort(v1, v2, v3 int64) []int64 {
	slice := []int64{v1, v2, v3}
	sort.Sort(Int64Slice(slice))
	return slice
}

// BalanceLL balances a LL tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceLL(val int64) *Tree {
	//
	//     Parent
	//      /
	//     pt
	//    /
	//  node
	//
	// to
	//
	//     Parent
	//      /  \
	//    pt   node
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Right = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Left = nil

	// 2. Delete the moved node itself
	// (X) T.Size -= 1
	// we just move the node not deleting from the whole tree
	node = nil

	return T
}

// BalanceLR balances a LR tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceLR(val int64) *Tree {
	//
	//     Parent
	//      /
	//     pt
	//       \
	//       node
	//
	// to
	//
	//     Parent
	//      /  \
	//     pt  node
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Right = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Right = nil

	// 2. Delete the moved node itself
	node = nil

	return T
}

// BalanceRR balances a RR tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceRR(val int64) *Tree {
	//
	//     Parent
	//        \
	//         pt
	//          \
	//          node
	//
	// to
	//
	//     Parent
	//      /  \
	//   node   pt
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Left = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Right = nil

	// 2. Delete the moved node itself
	// (X) T.Size -= 1
	// we just move the node not deleting from the whole tree
	node = nil

	return T
}

// BalanceRL balances a RL tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceRL(val int64) *Tree {
	//
	//     Parent
	//         \
	//          pt
	//          /
	//        node
	//
	// to
	//
	//     Parent
	//      /  \
	//   node   pt
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Left = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Left = nil

	// 2. Delete the moved node itself
	node = nil

	return T
}

// BalanceInserts implements Insert with a variadic function.
func (T *Tree) BalanceInserts(values ...int64) *Tree {
	for _, val := range values {
		_, rs := T.BalanceInsert(val)

		switch rs {

		case "LL":
			T.BalanceLL(val)

		case "LR":
			T.BalanceLR(val)

		case "RR":
			T.BalanceRR(val)

		case "RL":
			T.BalanceRL(val)

		}
	}
	return T
}

// RotateRight does right rotation on the Tree
// rooted with the value val.
func (T *Tree) RotateRight(val int64) *Tree {
	/*
	       y                               x
	      / \     Right Rotation          /  \
	     x   T3   – - – - – - – >        T1   y
	    / \       < - - - - - - -            / \
	   T1  T2     Left Rotation            T2  T3
	*/

	y := T.Find(val)
	yVal := y.Value

	x := y.Left
	xVal := x.Value

	T3 := y.Right
	T3Val := T3.Value

	T1 := x.Left
	T1Val := T1.Value

	T2 := x.Right
	T2Val := T2.Value

	// Update Tree
	T3.Right = &Tree{nil, 0, nil, 1}
	T3.Left = T2
	T3.Size = T3.Size + 2
	x.Right = nil
	//x.Left = nil
	x.Size = x.Size - 2

	// Update Values
	y.SetValue(xVal)
	x.SetValue(T1Val)
	T3.SetValue(yVal)
	T3.Left.SetValue(T2Val)
	T3.Right.SetValue(T3Val)

	return T
}

// RotateLeft does left rotation on the Tree
// rooted with the value val.
func (T *Tree) RotateLeft(val int64) *Tree {
	/*
	       y                               x
	      / \     Right Rotation          /  \
	     x   T3   – - – - – - – >        T1   y
	    / \       < - - - - - - -            / \
	   T1  T2     Left Rotation            T2  T3
	*/

	x := T.Find(val)
	xVal := x.Value

	T1 := x.Left
	T1Val := T1.Value

	y := x.Right
	yVal := y.Value

	T2 := y.Left
	T2Val := T2.Value

	T3 := y.Right
	T3Val := T3.Value

	// Update Tree
	T1.Right = T2
	T1.Left = &Tree{nil, 0, nil, 1}
	T1.Size = T1.Size + 2
	// y.Right = nil
	y.Left = nil
	y.Size = y.Size - 2

	// Update Values
	x.SetValue(yVal)
	y.SetValue(T3Val)
	T1.SetValue(xVal)
	T1.Left.SetValue(T1Val)
	T1.Right.SetValue(T2Val)

	return T
}

// walkPreOrder traverses the tree in the order of
// Root, Left, Right.
func walkPreOrder(T *Tree, ch chan int64) {
	ch <- T.Value

	if T.Left != nil {
		walkPreOrder(T.Left, ch)
	}

	if T.Right != nil {
		walkPreOrder(T.Right, ch)
	}
}

// WalkPreOrder traverses the tree in the order of
// Root, Left, Right.
func WalkPreOrder(T *Tree, ch chan int64) {
	walkPreOrder(T, ch)
	close(ch)
}

// walkInOrder traverses the tree in the order of
// Left, Root, Right.
func walkInOrder(T *Tree, ch chan int64) {
	// if left sub-tree does exist
	// recursively traverse the left sub-tree
	if T.Left != nil {
		walkInOrder(T.Left, ch)
	}

	// send the value of the present root node
	ch <- T.Value

	if T.Right != nil {
		walkInOrder(T.Right, ch)
	}
}

// WalkInOrder traverses the tree in the order of
// Left, Root, Right.
func WalkInOrder(T *Tree, ch chan int64) {
	walkInOrder(T, ch)
	close(ch)
}

// walkPostOrder traverses the tree in the order of
// Left, Right, Root.
func walkPostOrder(T *Tree, ch chan int64) {
	if T.Left != nil {
		walkPostOrder(T.Left, ch)
	}

	if T.Right != nil {
		walkPostOrder(T.Right, ch)
	}

	ch <- T.Value
}

// WalkPostOrder traverses the tree in the order of
// Left, Right, Root.
func WalkPostOrder(T *Tree, ch chan int64) {
	walkPostOrder(T, ch)
	close(ch)
}

// Same returns true if the two trees are same.
func Same(T1, T2 *Tree) bool {
	ch1, ch2 := make(chan int64), make(chan int64)
	go WalkInOrder(T1, ch1)
	go WalkInOrder(T2, ch2)

	for {
		// if the two trees are the same
		// all values that are sent to channel
		// 				should be equal
		// and the time that the channel gets closed
		// 				should also be equal
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// TRUE when TWO trees are different
		if v1 != v2 || ok1 != ok2 {
			return false
		}

		// TRUE here only when TWO trees are SAME
		// if one tree gets closed first
		// break and return true
		if !ok1 {
			break
		}
	}
	return true
}

func (T *Tree) TreeInserts(values ...int64) *Tree {
	// First insert with BalanceInserts
	T.BalanceInserts(values...)

	// And check the balancing status
	rb, sl := T.CheckTreeBalance()
	if rb && len(sl) != 1 {
		return T
	}

	// Rebalance the unbalanced nodes
	for _, val := range sl {
		switch T.Height(val) {
		case 2: // Left Left or Left Right case
			NT := T.Find(val)
			NTLC := NT.Left // Left Child
			// Left Left case
			if NT.Right.Right == nil &&
				NT.Right.Left == nil &&
				NTLC.Right.Left == nil &&
				NTLC.Right.Right == nil &&
				NTLC.Left.Left != nil &&
				NTLC.Left.Right != nil {
				T.RotateRight(NT.Value)
			} else {
				// Left Right case
				// Rotate Left on Left Child
				T.RotateLeft(NTLC.Value)
				// Rotate Right on Root
				T.RotateRight(NT.Value)
			}
		case -2: // Right Right or Right Left
			NT := T.Find(val)
			NTRC := NT.Right // Right Child
			// Right Right case
			if NT.Left.Left == nil &&
				NT.Left.Right == nil &&
				NTRC.Left.Left == nil &&
				NTRC.Left.Right == nil &&
				NTRC.Right.Left != nil &&
				NTRC.Right.Right != nil {
				T.RotateLeft(NT.Value)
			} else {
				// Left Right case
				// Rotate Left on Left Child
				T.RotateRight(NTRC.Value)
				// Rotate Right on Root
				T.RotateLeft(NT.Value)
			}
		}
	}
	return T
}
