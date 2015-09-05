package graph

import "container/heap"

// Prim finds the minimum spanning tree with min-heap (priority queue).
// Start a free from an arbitrary root Node r and grow the tree until
// it spans all the Nodes in the graph. Maintain the heap with the minimum
// weight value of edges.
// (http://en.wikipedia.org/wiki/Prim%27s_algorithm)
//
//	for  each vertex  u ∈ G.V
//		u.key = ∞
//		u.π = NIL
//
//	r.key = 0
//	Q = G.V
//
//	while  Q ≠ ø
//		u = Extract-Min(Q)
//		for each v ∈ G.Adj[u]
//			if v ∈ Q  and  v.key > w(u, v)
//				v.π = u
//				v.key = w(u, v)
//
func (g *Graph) Prim() map[Edge]bool {

	var src *Node
	for nd := range g.NodeMap {
		src = nd
		break
	}

	mapToDistance := make(map[*Node]float32)
	mapToDistance[src] = 0.0

	minHeap := &nodeDistanceHeap{}

	// initialize mapToDistance
	for nd := range g.NodeMap {
		if nd != src {
			mapToDistance[nd] = 2147483646.0
		}
		ndd := nodeDistance{}
		ndd.node = nd
		ndd.distance = mapToDistance[nd]
		heap.Push(minHeap, ndd)
	}

	mapToPrevID := make(map[string]string)
	heap.Init(minHeap)

	for minHeap.Len() != 0 {

		elem := heap.Pop(minHeap)

		for ov, weight := range elem.(nodeDistance).node.WeightTo {
			isExist := false
			for _, one := range *minHeap {
				if ov == one.node {
					isExist = true
					break
				}
			}
			if isExist && mapToDistance[ov] > weight {
				mapToDistance[ov] = weight
				minHeap.updateDistance(ov, weight)
				heap.Init(minHeap)

				mapToPrevID[ov.ID] = elem.(nodeDistance).node.ID
			}
		}
		for iv, weight := range elem.(nodeDistance).node.WeightTo {
			isExist := false
			for _, one := range *minHeap {
				if iv == one.node {
					isExist = true
					break
				}
			}
			if isExist && mapToDistance[iv] > weight {
				mapToDistance[iv] = weight
				minHeap.updateDistance(iv, weight)
				heap.Init(minHeap)

				mapToPrevID[iv.ID] = elem.(nodeDistance).node.ID
			}
		}
	}

	rmap := make(map[Edge]bool)
	for k, v := range mapToPrevID {
		one := Edge{}
		one.Src = g.GetNodeByID(v)
		one.Dst = g.GetNodeByID(k)
		one.Weight = g.GetEdgeWeight(one.Src, one.Dst)
		rmap[one] = true
	}

	return rmap
}
