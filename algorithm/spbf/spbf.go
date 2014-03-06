// Package spd finds the shortest path using Bellman-Ford algorithm.
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

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

// SPBF returns false if there is a negatively-weighted
// cycle that is reachable from the source vertex
// , which means that there is no shortest path
// since the negatively-weighted cycle adds up the
// infinite negative. Otherwise it returns the shortest
// path in the graph that can contain negative edges.
func SPBF(g *gsd.Graph, src, dst string) (string, bool) {
	source := g.FindVertexByID(src)
	destin := g.FindVertexByID(dst)

	source.StampD = 0

	// for each vertex u ∈ g.V
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}

		// relax
		edges := g.GetEdges()
		for _, edge := range *edges {
			if edge == nil {
				continue
			}
			a := edge.(*gsd.Edge).Dst.StampD
			b := edge.(*gsd.Edge).Src.StampD
			c := edge.(*gsd.Edge).Weight

			if a > b+int64(c) {
				edge.(*gsd.Edge).Dst.StampD = b + int64(c)
			}

			// Update Prev
			if edge.(*gsd.Edge).Dst.Prev.Len() == 0 {
				edge.(*gsd.Edge).Dst.Prev.PushBack(edge.(*gsd.Edge).Src)
			} else {
				ex := false
				pvs := edge.(*gsd.Edge).Dst.Prev
				for _, v := range *pvs {
					if v == nil {
						continue
					}
					// if fmt.Sprintf("%v", v.(*gsd.Vertex).ID) == fmt.Sprintf("%v", edge.(*gsd.Edge).Src.ID) {
					if v.(*gsd.Vertex) == edge.(*gsd.Edge).Src {
						ex = true
					}
				}

				if ex == false {
					edge.(*gsd.Edge).Dst.Prev.PushBack(edge.(*gsd.Edge).Src)
				}
			}
		}
	}

	edges := g.GetEdges()
	for _, edge := range *edges {
		if edge == nil {
			continue
		}
		a := edge.(*gsd.Edge).Dst.StampD
		b := edge.(*gsd.Edge).Src.StampD
		c := edge.(*gsd.Edge).Weight

		if a > b+int64(c) {
			return "There is negative weighted cycle (No Shortest Path)", false
		}
	}
	result := slice.NewSequence()
	TrackSPBF(g, source, destin, destin, result)

	s := ""
	for _, v := range *result {
		if v == nil {
			continue
		}
		s += fmt.Sprintf("%v(=%v) → ", v.(*gsd.Vertex).ID, v.(*gsd.Vertex).StampD)
	}

	return s[:len(s)-5], true
}

func TrackSPBF(g *gsd.Graph, source, target, end *gsd.Vertex, result *slice.Sequence) {
	// Add target first
	if result.Len() == 0 {
		result.PushFront(end)
	}

	// End recursion when we have NON-connected graph (len = 0)
	// End recursion when we get to Source that has no Prev
	if target.Prev.Len() == 0 {
		return
	}

	// when there is only start source and another vertex
	// that is already visited
	// which means that there is only source vertex left to visit
	tps := target.Prev
	if tps.Len() == 2 {
		for _, vtx := range *tps {
			// if fmt.Sprintf("%v", source.ID) != fmt.Sprintf("%v", vtx.(*gsd.Vertex).ID) {
			if source != vtx.(*gsd.Vertex) {
				for _, v := range *result {
					// if fmt.Sprintf("%v", vtx.(*gsd.Vertex).ID) == fmt.Sprintf("%v", v.(*gsd.Vertex).ID) {
					if vtx.(*gsd.Vertex) == v.(*gsd.Vertex) {
						result.PushFront(source)
						return
					}
				}
			}
		}
	}

	// Add the smallest vertex from back
	// that is not source, destination
	// that is not in the result
	ps := target.Prev
	ts := []string{}
	for _, vtx1 := range *ps {
		if vtx1 == nil {
			continue
		}
		// b1 := fmt.Sprintf("%v", vtx1.(*gsd.Vertex).ID) == fmt.Sprintf("%v", source.ID)
		// b2 := fmt.Sprintf("%v", vtx1.(*gsd.Vertex).ID) == fmt.Sprintf("%v", end.ID)
		b1 := vtx1.(*gsd.Vertex) == source
		b2 := vtx1.(*gsd.Vertex) == end

		b3 := false
		if result.Len() != 0 {
			for _, vtx2 := range *result {
				// if fmt.Sprintf("%v", vtx2.(*gsd.Vertex).ID) == fmt.Sprintf("%v", vtx1.(*gsd.Vertex).ID) {
				if vtx2.(*gsd.Vertex) == vtx1.(*gsd.Vertex) {
					b3 = true
				}
			}
		}

		if !b1 && !b2 && !b3 {
			ts = append(ts, fmt.Sprintf("%v", vtx1.(*gsd.Vertex).ID))
		}
	}

	if len(ts) == 0 {
		return
	}

	// find the Prev vertex with the smallest StampD
	sm := g.FindVertexByID(ts[0])
	for _, id := range ts {
		vtx := g.FindVertexByID(id)
		if sm.StampD > vtx.StampD {
			sm = vtx
		}
	}
	// now we know what is the smallest one

	result.PushFront(sm)
	TrackSPBF(g, source, sm, end, result)
}
