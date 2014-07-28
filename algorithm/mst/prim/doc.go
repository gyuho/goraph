// Package prim implements Prim's Minimum Spanning Tree algorithm.
package prim

/*
Prim Algorithm uses Min-Priority Queue
like Dijkstra algorithm shortest path algorithm.
(https://github.com/gyuho/goraph/tree/master/algorithm/spd)

Tree starts from an arbitrary root vertex  r
The tree grows until it spans all the vertices in V

All vertices that are not in the tree reside
in a min-priority queue Q based on the key attribute

v.key  is the minimum weight of any edge
from v to a vertex in the tree
v.key  is ∞  if there is no such edge

CLRS p.634

for  each vertex  u ∈ G.V
	u.key = ∞
	u.π = NIL

r.key = 0
Q = G.V

while  Q ≠ ø
	u = Extract-Min(Q)
	for each v ∈ G.Adj[u]
		if v ∈ Q  and  v.key > w(u, v)
			v.π = u
			v.key = w(u, v)
*/
