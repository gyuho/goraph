// Package dfs implements Dapth First Search algorithm (or DFS).
package dfs

/*
Pseudo Code CLRS page 604

DFS(G)
	for each vertex u ∈ G.V
		u.color = WHITE
		u.π = NIL
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	time = 0

	for each vertex u ∈ G.V
		if u.color == WHITE
			DFS-Visit(G, u)

---

DFS-Visit(G, u)
	time = time + 1
	u.d = time
	u.color = GRAY

	for each v ∈ G.Adj[u]
		if v.color == WHITE
			v.π = u
			DFS-Visit(G, v)

	u.color = BLACK
	time = time + 1
	u.f = time
*/
