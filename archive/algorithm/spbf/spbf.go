package spbf

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// SPBF returns false if there is a negatively-weighted
// cycle that is reachable from the source vertex
// , which means that there is no shortest path
// since the negatively-weighted cycle adds up the
// infinite negative. Otherwise it returns the shortest
// path in the graph that can contain negative edges.
func SPBF(g *gs.Graph, src, dst *gs.Vertex) (string, bool) {
	src.StampD = 0

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
			a := edge.(*gs.Edge).Dst.StampD
			b := edge.(*gs.Edge).Src.StampD
			c := edge.(*gs.Edge).Weight

			if a > b+int64(c) {
				edge.(*gs.Edge).Dst.StampD = b + int64(c)
			}

			// Update Prev
			if edge.(*gs.Edge).Dst.Prev.Len() == 0 {
				edge.(*gs.Edge).Dst.Prev.PushBack(edge.(*gs.Edge).Src)
			} else {
				ex := false
				pvs := edge.(*gs.Edge).Dst.Prev
				for _, v := range *pvs {
					if v == nil {
						continue
					}
					// if fmt.Sprintf("%v", v.(*gs.Vertex).ID) == fmt.Sprintf("%v", edge.(*gs.Edge).Src.ID) {
					if v.(*gs.Vertex) == edge.(*gs.Edge).Src {
						ex = true
					}
				}

				if ex == false {
					edge.(*gs.Edge).Dst.Prev.PushBack(edge.(*gs.Edge).Src)
				}
			}
		}
	}

	edges := g.GetEdges()
	for _, edge := range *edges {
		if edge == nil {
			continue
		}
		a := edge.(*gs.Edge).Dst.StampD
		b := edge.(*gs.Edge).Src.StampD
		c := edge.(*gs.Edge).Weight

		if a > b+int64(c) {
			return "There is negative weighted cycle (No Shortest Path)", false
		}
	}
	result := slice.NewSequence()
	TrackSPBF(g, src, dst, dst, result)

	var rs string
	for _, v := range *result {
		if v == nil {
			continue
		}
		rs += fmt.Sprintf("%v(=%v) → ", v.(*gs.Vertex).ID, v.(*gs.Vertex).StampD)
	}

	return rs[:len(rs)-5], true
}

func TrackSPBF(g *gs.Graph, src, dst, end *gs.Vertex, result *slice.Sequence) {
	// Add target first
	if result.Len() == 0 {
		result.PushFront(end)
	}

	// End recursion when we have NON-connected graph (len = 0)
	// End recursion when we get to Source that has no Prev
	if dst.Prev.Len() == 0 {
		return
	}

	// when there is only start source and another vertex
	// that is already visited
	// which means that there is only source vertex left to visit
	tps := dst.Prev
	if tps.Len() == 2 {
		for _, vtx := range *tps {
			// if fmt.Sprintf("%v", src.ID) != fmt.Sprintf("%v", vtx.(*gs.Vertex).ID) {
			if src != vtx.(*gs.Vertex) {
				for _, v := range *result {
					// if fmt.Sprintf("%v", vtx.(*gs.Vertex).ID) == fmt.Sprintf("%v", v.(*gs.Vertex).ID) {
					if vtx.(*gs.Vertex) == v.(*gs.Vertex) {
						result.PushFront(src)
						return
					}
				}
			}
		}
	}

	// Add the smallest vertex from back
	// that is not source, destination
	// that is not in the result
	ps := dst.Prev
	ts := []string{}
	for _, vtx1 := range *ps {
		if vtx1 == nil {
			continue
		}
		// b1 := fmt.Sprintf("%v", vtx1.(*gs.Vertex).ID) == fmt.Sprintf("%v", src.ID)
		// b2 := fmt.Sprintf("%v", vtx1.(*gs.Vertex).ID) == fmt.Sprintf("%v", end.ID)
		b1 := vtx1.(*gs.Vertex) == src
		b2 := vtx1.(*gs.Vertex) == end

		b3 := false
		if result.Len() != 0 {
			for _, vtx2 := range *result {
				// if fmt.Sprintf("%v", vtx2.(*gs.Vertex).ID) == fmt.Sprintf("%v", vtx1.(*gs.Vertex).ID) {
				if vtx2.(*gs.Vertex) == vtx1.(*gs.Vertex) {
					b3 = true
				}
			}
		}

		if !b1 && !b2 && !b3 {
			ts = append(ts, fmt.Sprintf("%v", vtx1.(*gs.Vertex).ID))
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
	TrackSPBF(g, src, sm, end, result)
}
