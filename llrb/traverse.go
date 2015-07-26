package llrb

import "fmt"

// PreOrder traverses from Root, Left-SubTree, and Right-SubTree. (DFS)
func (tr *Tree) PreOrder(ch chan string) {
	preOrder(tr.Root, ch)
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
func ComparePreOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.PreOrder(ch1)
	go t2.PreOrder(ch2)
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
func (tr *Tree) InOrder(ch chan string) {
	inOrder(tr.Root, ch)
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
func CompareInOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.InOrder(ch1)
	go t2.InOrder(ch2)
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
func (tr *Tree) PostOrder(ch chan string) {
	postOrder(tr.Root, ch)
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
func ComparePostOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.PostOrder(ch1)
	go t2.PostOrder(ch2)
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
func (tr *Tree) LevelOrder() []*Node {
	visited := []*Node{}
	queue := []*Node{tr.Root}
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
