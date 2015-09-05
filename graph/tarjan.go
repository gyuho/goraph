package graph

// Tarjan finds the strongly connected components.
// In the mathematics, a directed graph is "strongly connected"
// if every Node(vertex) is reachable from every other node.
// A graph is strongly connected if there is a path in each
// direction between each pair of node of a graph. Then a pair
// of nodes u and v is strongly connected to each other
// because there is a path in each direction.
// "Strongly connected components" of an arbitrary graph
// partition into sub-graphs that are themselves strongly connected.
// That is, "strongly connected component" of a directed graph
// is a sub-graph that is strongly connected.
// Formally, "Strongly connected components" of a graph is a maximal
// set of nodes C in G.V such that for all u, v ∈ C, there is a path
// both from u to v, and from v to u.
// (http://en.wikipedia.org/wiki/Strongly_connected_component, http://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm)
//
//	// v.index = index of a node to record the order of being discovered
//	// v.lowlink = the smallest index of any index reachable from v, including v itself
//	index := 0
//	S := empty
//	for each v in V do
//		if (v.index is undefined) then
//			strongconnect(v)
//		end if
//	end for
//
//	function strongconnect(v)
//		// Set the depth index for v to the smallest unused index
//		v.index := index
//		v.lowlink := index
//		index := index + 1
//		S.push(v)
//		v.onStack := true
//
//		// Consider successors of v
//		for each (v, w) in E do
//			if (w.index is undefined) then
//				// Successor w has not yet been visited; recurse on it
//				strongconnect(w)
//				v.lowlink  := min(v.lowlink, w.lowlink)
//			else if (w.onStack) then
//				// Successor w is in stack S and hence in the current SCC
//				v.lowlink  := min(v.lowlink, w.index)
//			end if
//		end for
//
//		// If v is a root node, pop the stack and generate an SCC
//		if (v.lowlink = v.index) then
//			start a new strongly connected component
//			repeat
//				w := S.pop()
//				w.onStack := false
//				add w to current strongly connected component
//			until (w = v)
//			output the current strongly connected component
//		end if
//	end function
//
func (g *Graph) Tarjan() [][]*Node {

	var globalIdx int
	mapNodeToIndex := make(map[*Node]int)
	mapNodeToLowLink := make(map[*Node]int)
	stack := []*Node{}
	stackMap := make(map[*Node]bool)
	result := [][]*Node{}

	for nd := range g.NodeMap {
		if _, ok := mapNodeToIndex[nd]; !ok {
			g.tarjan(nd, &globalIdx, mapNodeToIndex, mapNodeToLowLink, &stack, stackMap, &result)
		}
	}
	return result
}

func (g *Graph) tarjan(
	nd *Node,
	globalIdx *int,
	mapNodeToIndex map[*Node]int,
	mapNodeToLowLink map[*Node]int,
	stack *[]*Node,
	stackMap map[*Node]bool,
	result *[][]*Node,
) {
	mapNodeToIndex[nd] = *globalIdx
	mapNodeToLowLink[nd] = *globalIdx
	*globalIdx = *globalIdx + 1

	*stack = append(*stack, nd)
	stackMap[nd] = true

	for ov := range nd.WeightTo {
		if _, ok := mapNodeToIndex[ov]; !ok {
			// successor ov has not yet been visited; recurse on it
			g.tarjan(ov, globalIdx, mapNodeToIndex, mapNodeToLowLink, stack, stackMap, result)
			mapNodeToLowLink[nd] = min(mapNodeToLowLink[nd], mapNodeToLowLink[ov])
		} else if _, ok := stackMap[ov]; ok {
			// successor ov is in stack and hence in the current SCC
			mapNodeToLowLink[nd] = min(mapNodeToLowLink[nd], mapNodeToIndex[ov])
		}
	}
	// for iv := range nd.WeightFrom {
	// 	if _, ok := mapNodeToIndex[iv]; !ok {
	// 		// successor iv has not yet been visited; recurse on it
	// 		g.tarjan(iv, globalIdx, mapNodeToIndex, mapNodeToLowLink, stack, stackMap, result)
	// 		mapNodeToLowLink[nd] = min(mapNodeToLowLink[nd], mapNodeToLowLink[iv])
	// 	} else if _, ok := stackMap[iv]; ok {
	// 		// successor iv is in stack and hence in the current SCC
	// 		mapNodeToLowLink[nd] = min(mapNodeToLowLink[nd], mapNodeToIndex[iv])
	// 	}
	// }

	// If v is a root node, pop the stack and generate an SCC
	if mapNodeToLowLink[nd] == mapNodeToIndex[nd] {
		// start a new strongly connected component
		slice := []*Node{}
		for {
			w := (*stack)[len(*stack)-1]
			// PopBack
			*stack = (*stack)[:len(*stack)-1 : len(*stack)-1]
			delete(stackMap, w)
			slice = append(slice, w)
			if w == nd {
				*result = append(*result, slice)
				// slice = []*Node{}
				break
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
