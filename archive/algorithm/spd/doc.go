// Package spd finds the shortest path using Dijkstra algorithm.
// It does not work with negative edges.
package spd

/*
Dijkstra(G, source, target)
	for each vertex v ∈ G.V
		v.d = ∞
		v.π = nil
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	source.d = 0

	// Min-Priority queue Q, keyed by their d values
	Q = G.V

	while Q ≠ ∅
		Min-Heapify(Q)
		u = Extract-Min(Q)
		if u.d = ∞
			break
		for each vertex v ∈ G.Adj[u]
			if v.d > u.d + w(u,v)
				v.d = u.d + w(u,v)
				v.π = u
*/
