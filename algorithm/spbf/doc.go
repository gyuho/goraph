// Package spbf finds the shortest path using Bellman-Ford algorithm.
// It works with negative edges.
package spbf

/*
	SPBF(G, source, target)
		// Initialize-Single-Source(G,s)
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

		// for each vertex
		for  i = 1  to  |G.V| - 1
			for  each edge (u, v) ∈ G.E
				Relax(u, v, w)

		for  each edge (u, v) ∈ G.E
			if  v.d > u.d + w(u, v)
				return FALSE

		return TRUE
*/
