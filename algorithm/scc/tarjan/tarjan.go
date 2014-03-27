// Package tarjan implements Tarjan's Strongly Connected Components algorithm.
package tarjan

import (
	"strings"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

/*
http://en.wikipedia.org/wiki/Strongly_connected_component

In the mathematical theory of directed graphs,
a graph is said to be "strongly connected"
if every vertex is reachable from every other vertex.

A directed graph is called "strongly connected"
if there is a "path" in each direction
between each pair of vertices of the graph.

a pair of vertices u and v are said to be
strongly connected to each other
if there is a path in each direction between them.

binary relation of being strongly connected is an equivalence relation


"strongly connected components" of an arbitrary directed graph
form a partition into subgraphs that are themselves strongly connected.

"strongly connected component" of a directed graph G
is a subgraph that is strongly connected
, and is maximal with this property

SCC of G is a maximal set of vertices C in V
such that for all u, v ∈ C,
there is a path both from u to v and from v to u.

http://en.wikipedia.org/wiki/Tarjan's_strongly_connected_components_algorithm

*/

// SCC returns the Strongly Connected Components using Tarjan's algorithm.
// (Wikipedia/Tarjan's_strongly_connected_components_algorithm)
func SCC(g *gsd.Graph) string {
	//
	// v.index
	//	numbers the nodes consecutively in the order
	//	in which they are discovered
	//
	// v.lowlink
	//	represents (roughly speaking) the smallest index
	//	of any node known to be reachable from v, including v itself
	//
	// if v.lowlink < v.index
	// 	v must be left on the stack
	//
	// if v.lowlink == v.index
	//	whereas v must be removed as the root
	//	of a strongly connected component
	//
	//
	//	StampD:      9999999999,  <--- use as index
	//	StampF:      9999999999,  <--- use as lowlink
	//
	result := []string{}
	var idx int64
	idx = 0
	Vertices := g.GetVertices()
	stack := slice.NewSequence()
	for _, vtx := range *Vertices {
		// if (v.index is undefined)
		if vtx.(*gsd.Vertex).StampD > 9999999998 {
			result = append(result, scc(g, &idx, vtx.(*gsd.Vertex), stack))
		}
	}
	return strings.Join(result, "\n")
}

// GetMinimum64 returns the smaller element
// between v1 and v2.
func GetMinimum64(v1, v2 int64) int64 {
	if v1 > v2 {
		return v2
	} else {
		return v1
	}
}

// Contains returns true if vtx exists in the slice sl.
func Contains(vtx *gsd.Vertex, sl *slice.Sequence) bool {
	for _, val := range *sl {
		if val.(*gsd.Vertex).ID == vtx.ID {
			return true
		}
	}
	return false
}

func scc(g *gsd.Graph, idx *int64, vtx *gsd.Vertex, stack *slice.Sequence) string {
	vtx.StampD = *idx
	vtx.StampF = *idx
	*idx = *idx + 1
	stack.PushBack(vtx)

	ovs := vtx.GetOutVertices()
	for _, w := range *ovs {
		if w.(*gsd.Vertex).StampD > 9999999998 {
			scc(g, idx, w.(*gsd.Vertex), stack)
			vtx.StampF = GetMinimum64(vtx.StampF, w.(*gsd.Vertex).StampF)
		} else if Contains(w.(*gsd.Vertex), stack) {
			vtx.StampF = GetMinimum64(vtx.StampF, w.(*gsd.Vertex).StampD)
		}
	}

	result := []string{}

	// if (v.lowlink = v.index)
	if vtx.StampF == vtx.StampD {
		w := stack.PopBack()
		for !gsd.SameVertex(w.(*gsd.Vertex), vtx) {
			w = stack.PopBack()
			result = append(result, w.(*gsd.Vertex).ID)
		}
	}
	return strings.Join(result, "\n")
}
