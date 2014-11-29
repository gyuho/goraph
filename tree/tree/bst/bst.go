// Package bst implements binary search tree.
package bst

import (
	"container/list"
	"fmt"
	"math"
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
	// To end recursion
	// set terminal node's left and right to nil
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	if T != nil && T.Value != val {
		T.Size += 1
	}
	if val < T.Value {
		// Insert into the left tree
		T.Left = T.Left.Insert(val)
		// We don't need to do this
		// (X) to increase the size of the left sub-tree
		// (X) T.Left.Size += 1
	} else if val > T.Value {
		// Insert into the right tree
		T.Right = T.Right.Insert(val)
		// (X) to increase the size of the right sub-tree
		// (X) T.Right.Size += 1
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
	return float64ToInt64(math.Floor(math.Log2(float64(T.GetSize(val)))))
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

// Delete deletes the node of input value.
func (T *Tree) Delete(val int64) *Tree {
	// ** Deleting a leaf (node with no children)
	if T.IsLeaf(val) {
		// First delete as a child
		Parent := T.Parent(val)
		if val < Parent.Value { // Left Node
			Parent.Left = nil
		} else if val > Parent.Value { // Right Node
			Parent.Right = nil
		}
		// Then delete the node itself
		// Be careful with order
		// Do nil at the end
		T.Size -= 1
		TT := T.Find(val)
		_ = TT
		TT = nil
		return T
	}

	// we need to access from the node
	// not from the root of the Tree
	VT := T.Find(val)

	// ** Deleting a node with one child:
	// Remove the node and replace it with its child.
	// (1) only Left child
	// if T.Left != nil && T.Right == nil {
	if VT.Left != nil && VT.Right == nil {
		// First delete as a child
		// and replace it with the child of to-be-deleted node
		Parent := T.Parent(val)
		if val < Parent.Value { // if it's Left Node
			Parent.Left = VT.Left
		} else if val > Parent.Value { // if it's Right Node
			Parent.Right = VT.Left
		}
		// Then delete the node itself
		VT = nil
		T.Size -= 1
		return T
	}

	// ** Deleting a node with one child:
	// Remove the node and replace it with its child.
	// (2) only Right child
	if VT.Left == nil && VT.Right != nil {
		// First delete as a child
		// and replace it with the child of to-be-deleted node
		Parent := T.Parent(val)
		if val < Parent.Value { // if it's Left Node
			Parent.Left = VT.Right
		} else if val > Parent.Value { // if it's Right Node
			Parent.Right = VT.Right
		}
		// Then delete the node itself
		VT = nil
		T.Size -= 1
		return T
	}

	// TODO: Inefficient...
	// ** Deleting a node with two children
	// Move up the Left Child
	// rightmost node in the left subtree,
	// the inorder predecessor 6, is identified.
	// Its value is copied into the node being deleted.
	if VT.Left != nil && VT.Right != nil {
		ch := make(chan int64)
		is := ValuePreOrder(T, ch)
		return Construct(val, is)
	}

	return T
}

// TreePrint prints out the tree.
func (T *Tree) TreePrint() string {
	if T == nil {
		return "()"
	}
	s := ""
	if T.Left != nil {
		s += T.Left.TreePrint() + " "
	}
	s += fmt.Sprintf("%v", T.Value)
	if T.Right != nil {
		s += " " + T.Right.TreePrint()
	}
	return "(" + s + ")"
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

// StringInOrder returns the traversed string
// of the tree in the order of Left, Root, Right.
func StringInOrder(T *Tree, ch chan int64) string {
	go func() {
		defer close(ch)
		walkInOrder(T, ch)
	}()
	s := ""
	for v := range ch {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

// StringPreOrder returns the traversed string
// of the tree in the order of Root, Left, Right.
func StringPreOrder(T *Tree, ch chan int64) string {
	go func() {
		defer close(ch)
		walkPreOrder(T, ch)
	}()
	s := ""
	for v := range ch {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

// StringPreOrder returns the traversed string
// of the tree in the order of Left, Right, Root.
func StringPostOrder(T *Tree, ch chan int64) string {
	go func() {
		defer close(ch)
		walkPostOrder(T, ch)
	}()
	s := ""
	for v := range ch {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

// Copy returns a copy of the tree.
func (T *Tree) Copy() *Tree {
	t := NewTree(T.Value)
	t.Left = T.Left
	t.Right = T.Right
	return t
}

// WalkLevelOrder traverses the tree from the top.
func WalkLevelOrder(T *Tree) *list.List {
	result := list.New()
	queue := list.New()
	result.PushBack(T)
	queue.PushBack(T)
	for queue.Len() > 0 {
		elem := queue.Front()
		tn := elem.Value.(*Tree).Copy()
		queue.Remove(elem)
		if tn.Left != nil {
			result.PushBack(tn.Left)
			queue.PushBack(tn.Left)
		}
		if tn.Right != nil {
			result.PushBack(tn.Right)
			queue.PushBack(tn.Right)
		}
	}
	return result
}

// StringLevelOrder traverses the tree from the top.
func StringLevelOrder(T *Tree) string {
	list := WalkLevelOrder(T)
	// fmt.Println(list.Len())
	s := ""
	for elem := list.Front(); elem != nil; elem = elem.Next() {
		s += fmt.Sprintf("%v ", elem.Value.(*Tree).Value)
	}
	return s
}

// ValuePreOrder returns the traversed integer slice
// of the tree in the order of Root, Left, Right.
func ValuePreOrder(T *Tree, ch chan int64) []int64 {
	go func() {
		defer close(ch)
		walkPreOrder(T, ch)
	}()
	slice := []int64{}
	for v := range ch {
		slice = append(slice, v)
	}
	return slice
}

// Construct creates the Tree out of input slice,
// except the input value.
func Construct(val int64, slice []int64) *Tree {
	// delete the value from slice
	findIdx := func(val int64, sl []int64) int {
		for k, v := range sl {
			if val == v {
				return k
			}
		}
		return 0
	}
	idx := findIdx(val, slice)
	copy(slice[idx:], slice[idx+1:])
	slice = slice[:len(slice)-1 : len(slice)-1] // resize
	Tr := NewTree(slice[0])
	Tr.Inserts(slice[1:]...)
	return Tr
}

// SameSlice returns true if the two int64 slices
// are equal.
func SameSlice(s1, s2 []int64) bool {
	if len(s1) != len(s2) {
		return false
	}
	result := true
	for k, v := range s2 {
		if s1[k] != v {
			result = false
		}
	}
	return result
}
