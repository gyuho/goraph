// Package tskahn finds topological sort
// based on algorithm by Arthur Kahn(1962).
// (Reference: http://dl.acm.org/citation.cfm?doid=368996.369025)
package tskahn

/*
DAG_Kahn(G)
S as a set of nodes with no incoming edges
while  S.length  !=  0
	remove a node n from S
	add n to the tail of L
	for each vertex m that comes out of n
		remove the edge from n to m
		if m has no other incoming edges
			add m to S

if graph still has edges
	panic to return error
	(Graph has at least one cycle)
else
	return L (topological sort)
*/
