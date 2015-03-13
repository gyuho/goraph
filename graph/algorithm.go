package graph

// Breadth First Search: Queue
// Depth First Search: Stack / Recursion

// Bfs does Breadth First Search and return the result in visited order.
// Bfs traverses graphs in an arbitrary order.
//
// 1  procedure Bfs(G,v) is
// 2      let Q be a queue
// 3      Q.push(v)
// 4      label v as discovered
// 5      while Q is not empty
// 6         v ← Q.pop()
// 7         for all edges from v to w in G.adjacentEdges(v) do
// 8             if w is not labeled as discovered
// 9                 Q.push(w)
// 10                label w as discovered
//
// O(|V| + |E|)
func (d Data) Bfs(src *Vertex) []*Vertex {

	result := []*Vertex{}

	src.Color = "black"
	queue := []*Vertex{src}

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
//
// 1  procedure DFS-iterative(G,v):
// 2      let S be a stack
// 3      S.push(v)
// 4      while S is not empty
// 5            v = S.pop()
// 6            if v is not labeled as discovered:
// 7                label v as discovered
// 8                for all edges from v to w in G.adjacentEdges(v) do
// 9                    S.push(w)
//
func (d Data) DfsStack(src *Vertex) []*Vertex {

	result := []*Vertex{}
	stack := []*Vertex{src}

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

// DfsRecursive recursively traverse a graph.
//
// 1  procedure DFS(G,v):
// 2      label v as discovered
// 3      for all edges from v to w in G.adjacentEdges(v) do
// 4          if vertex w is not labeled as discovered then
// 5              recursively call DFS(G,w)
//
func (d Data) DfsRecursive(src *Vertex, result *[]*Vertex) {

	if src.Color == "black" {
		return
	}

	src.Color = "black"
	*result = append(*result, src)

	for ov := range src.WeightTo {
		if ov.Color == "white" {
			d.DfsRecursive(ov, result)
		}
	}
}

// TopologicalDfs does topological sorting with DFS.
//
// L ← Empty list that will contain the sorted nodes
// while there are unmarked nodes do
//     select an unmarked node n
//     visit(n)
// function visit(node n)
//     if n has a temporary mark then stop (not a DAG)
//     if n is not marked (i.e. has not been visited yet) then
//         mark n temporarily
//         for each node m with an edge from n to m do
//             visit(m)
//         mark n permanently
//         unmark n temporarily
//         add n to head of L
//
func (d Data) TopologicalDfs(src *Vertex, result *[]*Vertex) {

}
