// Package kosaraju implements Kosaraju's Strongly Connected Components algorithm.
package kosaraju

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

/*
http://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
It makes use of the fact that the transpose graph
(the same graph with the direction of every edge reversed)
has exactly the same strongly connected components as the original graph.


Let G be a directed graph
Let S be an empty stack

While S does not contain all vertices:
	Choose an arbitrary vertex v not in S.
		Perform a DFS starting at v.
			Each time that DFS finishes
			expanding a vertex u, push u onto S.

Reverse the directions of all arcs to obtain the transpose graph.

While S is nonempty:
	Pop the top vertex v from S.
		Perform a DFS starting at v in the transpose graph.

		The set of visited vertices will give
		the strongly connected component containing v

		record this and
		remove all these vertices from the graph G and the stack S.
*/
