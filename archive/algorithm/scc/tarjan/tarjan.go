package tarjan

import (
	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// SCC returns the Strongly Connected Components using Tarjan's algorithm.
// (Wikipedia/Tarjan's_strongly_connected_components_algorithm)
func SCC(g *gs.Graph) [][]string {
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
	var idx int64
	Vertices := g.GetVertices()

	stack := slice.NewSequence()
	result := [][]string{}

	for _, vtx := range *Vertices {
		// if (v.index is undefined)
		if vtx.(*gs.Vertex).StampD > 9999999998 {
			runSCC(&idx, vtx.(*gs.Vertex), &result, stack)
		}
	}
	return result
}

func runSCC(idx *int64, vtx *gs.Vertex, result *[][]string, stack *slice.Sequence) {
	vtx.StampD = *idx
	vtx.StampF = *idx
	*idx = *idx + 1
	stack.PushBack(vtx)

	ovs := vtx.GetOutVertices()
	for _, w := range *ovs {
		// if the node was not discovered yet
		if w.(*gs.Vertex).StampD > 9999999998 {
			runSCC(idx, w.(*gs.Vertex), result, stack)
			vtx.StampF = getMinimum64(vtx.StampF, w.(*gs.Vertex).StampF)
		} else if contains(w.(*gs.Vertex), stack) {
			vtx.StampF = getMinimum64(vtx.StampF, w.(*gs.Vertex).StampD)
		}
	}

	sccSli := []string{}
	// if (v.lowlink = v.index)
	if vtx.StampF == vtx.StampD {
		// w := stack.PopBack()
		w := &gs.Vertex{}
		for !gs.SameVertex(w, vtx) {
			w = stack.PopBack().(*gs.Vertex)
			sccSli = append(sccSli, w.ID)
		}
	}
	if len(sccSli) != 0 {
		*result = append(*result, sccSli)
	}
}

// getMinimum64 returns the smaller element
// between v1 and v2.
func getMinimum64(v1, v2 int64) int64 {
	if v1 > v2 {
		return v2
	}
	return v1
}

// contains returns true if vtx exists in the slice sl.
func contains(vtx *gs.Vertex, sl *slice.Sequence) bool {
	for _, val := range *sl {
		if val.(*gs.Vertex).ID == vtx.ID {
			return true
		}
	}
	return false
}
