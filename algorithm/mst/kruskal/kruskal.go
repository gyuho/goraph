// Package kruskal implements Kruskal's Minimum Spanning Tree algorithm.
package kruskal

import (
	"github.com/gyuho/goraph/goroup/gsdset"
	"github.com/gyuho/goraph/graph/gsd"
)

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

// MST implements Kruskal's Minimum Spanning Tree algorithm.
// It returns the edges and total weight of Minimum Spanning Tree.
func MST(g *gsd.Graph) ([]*gsd.Edge, float64) {
	mstedges := []*gsd.Edge{}

	// for  each vertex  v ∈ G.V
	// 	Make-Set(v)
	dsets := gsdset.MakeGraphSet(g)

	// sort the edges of G.E  in increasing order of w
	edges := gsdset.SortEdges(g)

	// for each edge (u, v) ∈ G.E (increasing order)
	for _, edge := range *edges {
		// if Find-Set(u) ≠ Find-Set(v)
		ru := gsdset.FindSet(edge.(*gsd.Edge).Src, dsets) // Find-Set(u)
		rv := gsdset.FindSet(edge.(*gsd.Edge).Dst, dsets) // Find-Set(v)
		if !gsd.SameVertex(ru, rv) {
			// A  =  A  U  {(u, v)}
			mstedges = append(mstedges, edge.(*gsd.Edge))
			// gsdset.UnionByVtx(edge.(*gsd.Edge).Src, edge.(*gsd.Edge).Dst, &dsets)

			// Union(u, v)
			gsdset.UnionByRep(ru, rv, &dsets)
		}
	}
	var total float64
	for _, edg := range mstedges {
		total += edg.Weight
	}
	return mstedges, total
}
