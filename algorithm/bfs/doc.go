// Package bfs implements Breadth First Search algorithm (or BFS).
package bfs

// Pseudo Code CLRS page 595
//
// BFS(G, s)
// 	for each vertex u ∈ g.V - {s}
// 		u.color = WHITE
// 		u.d = ∞
// 		u.π = NIL
// 	// this is already done
// 	// when instantiating the graph
// 	// and instead of InVertices
// 	// we can just create another slice
// 	// inside Graph (Prev)
// 	// in order not to modify the original graph
//
// 	s.color = GRAY
// 	s.d = 0
// 	s.π = NIL
// 	Q = ∅
//
// 	ENQUEUE(Q, s)
//
// 	while Q ≠ ∅
// 		u = DEQUEUE(Q)
// 		for each v ∈ g.Adj[u]
// 			if v.color == WHITE
// 				v.color = GRAY
// 				v.d = u.d + 1
// 				v.π = u
// 				ENQUEUE(Q, v)
// 		u.color = BLACK
