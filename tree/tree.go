package tree

import "fmt"

// Data contains a Root node of a binary search tree.
type Data struct {
	Root *Node
}

// New returns a new Data with its root Node.
func New(root *Node) *Data {
	d := &Data{}
	d.Root = root
	return d
}

// Interface represents a single object in the tree.
type Interface interface {
	// Less returns true when the receiver item(key)
	// is less than the given(than) argument.
	Less(than Interface) bool
}

// Node is a Node and a Tree itself.
type Node struct {
	// Left is a left child Node.
	Left *Node

	Key Interface

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	return nd
}

func (d *Data) String() string {
	return d.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v", nd.Key)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

// Insert inserts a Node to a Data without replacement.
func (d *Data) Insert(nd *Node) {
	if d.Root == nd {
		return
	}
	d.Root = d.Root.insert(nd)
}

func (nd *Node) insert(node *Node) *Node {
	if nd == nil {
		return node
	}
	if nd.Key.Less(node.Key) {
		nd.Right = nd.Right.insert(node)
	} else {
		nd.Left = nd.Left.insert(node)
	}
	return nd
}

// Min returns the minimum key Node in the tree.
func (d Data) Min() *Node {
	nd := d.Root
	if nd == nil {
		return nil
	}
	for nd.Left != nil {
		nd = nd.Left
	}
	return nd
}

// Max returns the maximum key Node in the tree.
func (d Data) Max() *Node {
	nd := d.Root
	if nd == nil {
		return nil
	}
	for nd.Right != nil {
		nd = nd.Right
	}
	return nd
}

// Search does binary-search on a given key and returns the first Node with the key.
func (d Data) Search(key Interface) *Node {
	nd := d.Root
	// just updating the pointer value
	for nd != nil {
		switch {
		case nd.Key.Less(key):
			nd = nd.Right
		case key.Less(nd.Key):
			nd = nd.Left
		default:
			return nd
		}
	}
	return nil
}

// SearchChan does binary-search on a given key and return the first Node with the key.
func (d Data) SearchChan(key Interface, ch chan *Node) {
	searchChan(d.Root, key, ch)
	close(ch)
}

func searchChan(nd *Node, key Interface, ch chan *Node) {
	// leaf node
	if nd == nil {
		return
	}
	// when equal
	if !nd.Key.Less(key) && !key.Less(nd.Key) {
		ch <- nd
		return
	}
	searchChan(nd.Left, key, ch)  // left
	searchChan(nd.Right, key, ch) // right
}

// PreOrder traverses from Root, Left-SubTree, and Right-SubTree. (DFS)
func (d *Data) PreOrder(ch chan string) {
	preOrder(d.Root, ch)
	close(ch)
}

func preOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	ch <- fmt.Sprintf("%v", nd.Key) // root
	preOrder(nd.Left, ch)           // left
	preOrder(nd.Right, ch)          // right
}

// ComparePreOrder returns true if two Trees are same with PreOrder.
func ComparePreOrder(d1, d2 *Data) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go d1.PreOrder(ch1)
	go d2.PreOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// InOrder traverses from Left-SubTree, Root, and Right-SubTree. (DFS)
func (d *Data) InOrder(ch chan string) {
	inOrder(d.Root, ch)
	close(ch)
}

func inOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	inOrder(nd.Left, ch)            // left
	ch <- fmt.Sprintf("%v", nd.Key) // root
	inOrder(nd.Right, ch)           // right
}

// CompareInOrder returns true if two Trees are same with InOrder.
func CompareInOrder(d1, d2 *Data) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go d1.InOrder(ch1)
	go d2.InOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// PostOrder traverses from Left-SubTree, Right-SubTree, and Root.
func (d *Data) PostOrder(ch chan string) {
	postOrder(d.Root, ch)
	close(ch)
}

func postOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	postOrder(nd.Left, ch)          // left
	postOrder(nd.Right, ch)         // right
	ch <- fmt.Sprintf("%v", nd.Key) // root
}

// ComparePostOrder returns true if two Trees are same with PostOrder.
func ComparePostOrder(d1, d2 *Data) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go d1.PostOrder(ch1)
	go d2.PostOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// LevelOrder traverses the Tree with Breadth First Search.
// (http://en.wikipedia.org/wiki/Tree_traversal#Breadth-first_2)
//
//	levelorder(root)
//	  q = empty queue
//	  q.enqueue(root)
//	  while not q.empty do
//	    node := q.dequeue()
//	    visit(node)
//	    if node.left ≠ null then
//	      q.enqueue(node.left)
//	    if node.right ≠ null then
//	      q.enqueue(node.right)
//
func (d *Data) LevelOrder() []*Node {
	visited := []*Node{}
	queue := []*Node{d.Root}
	for len(queue) != 0 {
		nd := queue[0]
		queue = queue[1:len(queue):len(queue)]
		visited = append(visited, nd)
		if nd.Left != nil {
			queue = append(queue, nd.Left)
		}
		if nd.Right != nil {
			queue = append(queue, nd.Right)
		}
	}
	return visited
}
