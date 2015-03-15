package graph

// Bfs does Breadth First Search and return the result in visited order.
// Bfs traverses graphs in an arbitrary order. Time complexity is O(|V| + |E|).
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

	if src == nil {
		return nil
	}

	result := []*Node{}

	src.Color = "black"
	queue := []*Node{src}

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:len(queue):len(queue)]

		for ov := range front.WeightTo {
			if ov.Color == "white" {
				ov.Color = "black"
				queue = append(queue, ov)
			}
		}
		for iv := range front.WeightFrom {
			if iv.Color == "white" {
				iv.Color = "black"
				queue = append(queue, iv)
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

	if src == nil {
		return nil
	}

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
			for iv := range back.WeightFrom {
				stack = append(stack, iv)
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

	if src == nil {
		return
	}

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
	for iv := range src.WeightFrom {
		if iv.Color == "white" {
			d.Dfs(iv, result)
		}
	}
}
