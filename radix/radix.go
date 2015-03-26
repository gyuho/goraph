package radix

import "sort"

// Data contains a Root node of a radix trie.
type Data struct {
	Root *Node
}

// New returns a new Data with its root Node.
func New(root *Node) *Data {
	d := &Data{}
	d.Root = root
	return d
}

// Node contains edges and possible leaf node.
type Node struct {
	Edges Edges
}

// NewNode returns a new Node.
func NewNode() *Node {
	nd := &Node{}
	nd.Edges = []*Edge{}
	return nd
}

// Edge connects nodes.
type Edge struct {
	Label     []byte
	ChildNode *Node
}

// Edges is a slice of Edge.
// type Keys []Interface would update the receiver
// but not type Edges []Edge.
type Edges []*Edge

func (s Edges) Len() int {
	return len(s)
}

func (s Edges) Less(i, j int) bool {
	return string(s[i].Label) < string(s[j].Label)
}

func (s Edges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort sorts the Edges in an ascending order of Label.
func (s Edges) sort() {
	sort.Sort(s)
}

// searchPrefix returns the new Edge with common prefix, edge index and index of character for the longest commond prefix.
// For the value 'hello', it returns the Edge with label 'he' and index 2 as a char index.
// It returns true either when we need to insert a whole new edge
// or to split current Edge.

const (
	insertWithNoCommonPrifix = iota
	skip
	insertWithSplitValue
	insertWithSplitBoth
)

// searchPrefix returns the Edge with common prefix and its index in Edges
// , and the index of last common prefix character.
// It returns the option needed for insertion.
func (s Edges) searchPrefix(val []byte) (*Edge, int, int, int) {
	// sort before search
	s.sort()

	// Search uses binary search to find
	// and return the smallest index i in [0, n) at which f(i) is true
	// edgeIdx := sort.Search(len(s), func(i int) bool {
	// 	return s[i].Label[0] == val[0]
	// })
	var edgeIdx int
	for k, elem := range s {
		if elem.Label[0] == val[0] {
			edgeIdx = k
			break
		}
	}

	// #1. no common prefix found
	if edgeIdx == 0 {
		return nil, 0, 0, insertWithNoCommonPrifix
	}

	// common prefix, then must there be only one edge
	edge := s[edgeIdx]

	// find the common prefix index
	commonIdx := 0
	for idx, char := range val {
		if len(edge.Label) > idx {
			if edge.Label[idx] == char {
				commonIdx = idx
			} else {
				break
			}
		}
	}

	// #2. value is a substring of label
	if commonIdx == len(val)-1 {
		return nil, 0, 0, skip
	}

	// #3. exists common prefix but value needs split
	if commonIdx == len(edge.Label)-1 {
		return edge, edgeIdx, commonIdx, insertWithSplitValue
	}

	// #4. exists common prefix but both need split
	return edge, edgeIdx, commonIdx, insertWithSplitBoth
}

// without receiving pointer, it won't update the Edges.
func (s *Edges) insert(val []byte) {

	edge, edgeIdx, commonIdx, option := s.searchPrefix(val)

	switch option {

	case insertWithNoCommonPrifix:
		// insert a whole new edge (no common prefix)
		newEdge := &Edge{}
		newEdge.Label = val
		*s = append(*s, newEdge)

	case skip:

	case insertWithSplitValue:
		newNode := NewNode()
		newEdge := &Edge{}
		newEdge.Label = val[commonIdx+1:]
		newNode.Edges = append(newNode.Edges, newEdge)
		edge.ChildNode = newNode

	case insertWithSplitBoth:
		edgeTop := splitEdge(edge, commonIdx+1)
		newEdge := &Edge{}
		newEdge.Label = val[commonIdx+1:]
		edgeTop.ChildNode.Edges = append(edgeTop.ChildNode.Edges, newEdge)
		(*s)[edgeIdx] = edgeTop
		*s = append(*s, newEdge)
	}

	return
}

// Insert inserts a value to a radix trie.
func (d *Data) Insert(val []byte) {

}

func splitEdge(e *Edge, idx int) *Edge {
	edgeBottom := &Edge{}
	edgeBottom.Label = e.Label[idx:]
	edgeBottom.ChildNode = &Node{}
	if e.ChildNode != nil {
		edgeBottom.ChildNode.Edges = append(edgeBottom.ChildNode.Edges, e.ChildNode.Edges...)
	}

	edgeTop := &Edge{}
	edgeTop.Label = e.Label[:idx]
	edgeTop.ChildNode = NewNode()
	edgeTop.ChildNode.Edges = append(edgeTop.ChildNode.Edges, edgeBottom)

	return edgeTop
}
