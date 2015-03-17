package graph

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
func (d *Data) Prim() map[Edge]bool {

	return nil
}
