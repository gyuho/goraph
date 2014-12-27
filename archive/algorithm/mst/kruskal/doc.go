// Package kruskal implements Kruskal's Minimum Spanning Tree algorithm.
package kruskal

/*
A = ø
for  each vertex  v ∈ G.V
	Make-Set(v)

sort the edges of G.E  in increasing order of w

for each edge (u, v) ∈ G.E (increasing order)
	if Find-Set(u) ≠ Find-Set(v)
		A  =  A  U  {(u, v)}
		Union(u, v)

return A
*/
