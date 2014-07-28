package kruskal

import (
	"github.com/gyuho/goraph/graph/gs"
	"github.com/gyuho/goraph/graph/gset"
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
func MST(g *gs.Graph) ([]*gs.Edge, float64) {
	mstedges := []*gs.Edge{}

	// for  each vertex  v ∈ G.V
	// 	Make-Set(v)
	dsets := gset.MakeGraphSet(g)

	// sort the edges of G.E  in increasing order of w
	edges := gset.SortEdges(g)

	// for each edge (u, v) ∈ G.E (increasing order)
	for _, edge := range *edges {
		// if Find-Set(u) ≠ Find-Set(v)
		ru := gset.FindSet(edge.(*gs.Edge).Src, dsets) // Find-Set(u)
		rv := gset.FindSet(edge.(*gs.Edge).Dst, dsets) // Find-Set(v)
		if !gs.SameVertex(ru, rv) {
			// A  =  A  U  {(u, v)}
			mstedges = append(mstedges, edge.(*gs.Edge))
			// gset.UnionByVtx(edge.(*gs.Edge).Src, edge.(*gs.Edge).Dst, &dsets)

			// Union(u, v)
			gset.UnionByRep(ru, rv, &dsets)
		}
	}
	var total float64
	for _, edg := range mstedges {
		total += edg.Weight
	}
	return mstedges, total
}
