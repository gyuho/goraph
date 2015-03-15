package graph

// Bfs does Breadth First Search and return the result in visited order.
// Bfs traverses graphs in an arbitrary order. The time complexity is O(|V| + |E|).
// Bfs uses queue. Dfs uses recursion or stack.
// (http://en.wikipedia.org/wiki/Breadth-first_search)
//
//	1  procedure Bfs(G,v) is
//	2      let Q be a queue
//	3      Q.push(v)
//	4      label v as discovered
//	5      while Q is not empty
//	6         v ← Q.pop()
//	7         for all edges from v to w in G.adjacentEdges(v) do
//	8             if w is not labeled as discovered
//	9                 Q.push(w)
//	10                label w as discovered
//
func (d *Data) Bfs(src *Node) []*Node {

	result := []*Node{}

	src.Color = "black"
	queue := []*Node{src}

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:len(queue):len(queue)]

		for ov := range front.WeightTo {

			if ov == nil {
				continue
			}

			if ov.Color == "white" {
				ov.Color = "black"
				queue = append(queue, ov)
			}

		}

		front.Color = "black"
		result = append(result, front)
	}

	return result
}

// DfsStack searches a graph with depth-first.
// (http://en.wikipedia.org/wiki/Depth-first_search)
//
//	1  procedure DFS-iterative(G,v):
//	2      let S be a stack
//	3      S.push(v)
//	4      while S is not empty
//	5            v = S.pop()
//	6            if v is not labeled as discovered:
//	7                label v as discovered
//	8                for all edges from v to w in G.adjacentEdges(v) do
//	9                    S.push(w)
//
func (d *Data) DfsStack(src *Node) []*Node {

	result := []*Node{}
	stack := []*Node{src}

	for len(stack) != 0 {

		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1 : len(stack)-1]

		if back.Color == "white" {

			back.Color = "black"
			result = append(result, back)

			for ov := range back.WeightTo {
				stack = append(stack, ov)
			}
		}

	}

	return result
}

// Dfs recursively traverses a graph.
// (http://en.wikipedia.org/wiki/Depth-first_search)
//
//	1  procedure DFS(G,v):
//	2      label v as discovered
//	3      for all edges from v to w in G.adjacentEdges(v) do
//	4          if vertex w is not labeled as discovered then
//	5              recursively call DFS(G,w)
//
func (d *Data) Dfs(src *Node, result *[]*Node) {
	if src.Color == "black" {
		return
	}

	src.Color = "black"
	*result = append(*result, src)

	for ov := range src.WeightTo {
		if ov.Color == "white" {
			d.Dfs(ov, result)
		}
	}
}

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
func (d Data) TopologicalDag() ([]*Node, bool) {
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
