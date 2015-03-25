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
	nd.Edges = []Edge{}
	return nd
}

// Edge connects nodes.
type Edge struct {
	Label     []byte
	ChildNode *Node
}

// Edges is a slice of Edge.
type Edges []Edge

func (s Edges) Len() int {
	return len(s)
}

func (s Edges) Less(i, j int) bool {
	return string(s[i].Label) < string(s[j].Label)
}

func (s Edges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sort sorts the Edges in an ascending order of Label.
func (s Edges) Sort() {
	sort.Sort(s)
}

// SearchPrefix returns the Edge with common prefix and the index of longest commond prefix.
// For the value 'hello', it returns the Edge with label 'he' and index 2.
func (s Edges) SearchPrefix(val []byte) (Edge, int) {
	// must sort before search
	s.Sort()

	// Search uses binary search to find
	// and return the smallest index i in [0, n) at which f(i) is true
	idx := sort.Search(len(s), func(i int) bool {
		return s[i].Label[0] == val[0]
	})

	// not found
	if idx == len(s) {
		return Edge{}, 0
	}

	// common prefix, then must be only one edg
	edg := s[idx]

	// when label is longer, no need to insert
	if len(edg.Label) >= len(val) {
		return Edge{}, 0
	}

	// when val is longer, find the common prefix
	loc := 0
	for idx, char := range val {
		if len(edg.Label) > idx {
			if edg.Label[idx] == char {
				loc = idx
			}
		} else {
			break
		}
	}
	return edg, loc
}
