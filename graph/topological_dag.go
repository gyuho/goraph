package graph

// TopologicalDag does topological sort(ordering) with DFS.
// It returns true if the Graph is a DAG. (no cycle, have a topological sort)
// It returns false if the Graph is not a DAG. (cycle, have no topological sort)
// (http://en.wikipedia.org/wiki/Topological_sorting)
//
//	L ← Empty list that will contain the sorted nodes
//	while there are unmarked nodes do
//	    select an unmarked node n
//	    visit(n)
//
//	function visit(node n)
//	    if n has a temporary mark then stop (not a DAG)
//	    if n is not marked (i.e. has not been visited yet) then
//	        mark n temporarily
//	        for each node m with an edge from n to m do
//	            visit(m)
//	        mark n permanently
//	        unmark n temporarily
//	        add n to head of L
//
func (d *Data) TopologicalDag() ([]*Node, bool) {
	result := []*Node{}
	isDag := true
	for nd := range d.NodeMap {
		if nd.Color != "white" {
			continue
		}
		d.topologicalDag(nd, &result, &isDag)
	}

	if !isDag {
		return nil, false
	}

	return result, true
}

// topologicalDag recursively traverses the Graph with DFS.
func (d *Data) topologicalDag(src *Node, result *[]*Node, isDag *bool) {
	if src == nil {
		return
	}
	if src.Color == "gray" {
		*isDag = false
		return
	}
	if src.Color == "white" {
		src.Color = "gray"
		for ov := range src.WeightTo {
			d.topologicalDag(ov, result, isDag)
		}
		src.Color = "black"
		// PushFront
		copied := make([]*Node, len(*result)+1)
		copied[0] = src
		copy(copied[1:], *result)
		*result = copied
	}
}
