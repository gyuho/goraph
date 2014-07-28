// Package spfw finds the all-pairs shortest paths using Floyd-Warshall algorithm.
package spfw

/*
http://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

- every combination of edges is tested
- incrementally improving an estimate on the shortest path
	between two vertices, until the estimate is optimal

Consider a graph G with vertices V numbered 1 through N.
Further consider a function shortestPath(i, j, k)
that returns the shortest possible path from i to j using vertices
only from the set {1,2,...,k} as intermediate points along the way.

Now, given this function, our goal is to find the shortest path
from each i to each j using only vertices 1 to k + 1.

For each of these pairs of vertices,
the true shortest path could be either
(1) a path that only uses vertices in the set {1, ..., k}
or
(2) a path that goes from i to k + 1 and then from k + 1 to j.

We know that the best path from i to j that only uses vertices 1
through k is defined by shortestPath(i, j, k), and it is clear that
if there were a better path from i to k + 1 to j,
then the length of this path would be the concatenation of the shortest path
from i to k + 1 (using vertices in {1, ..., k})
and
the shortest path from k + 1 to j (also using vertices in {1, ..., k}).


let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)

for each vertex v
	dist[v][v] ← 0

for each edge (u,v)
	dist[u][v] ← w(u,v)  // the weight of the edge (u,v)

for k from 1 to |V|
	for i from 1 to |V|
		for j from 1 to |V|
			if dist[i][j] > dist[i][k] + dist[k][j]
				dist[i][j] ← dist[i][k] + dist[k][j]
			end if
*/
