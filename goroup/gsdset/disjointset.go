package gsdset

import (
	"sort"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

type DisJointSet struct {
	Rep  *gsd.Vertex // Representitive Vertex
	Mems Set         // Set Members
	// Each Set member is mapped to Rep
}

// NewDisJointSet returns a new DisJointSet.
func NewDisJointSet() *DisJointSet {
	return &DisJointSet{
		Rep:  nil,
		Mems: nil,
	}
}

// String prints out the DisJointSet information.
func (s DisJointSet) String() string {
	return "DisJoint Set (Rep: " + s.Rep.ID + ") / " + s.Mems.String()
}

// makeSet creates a Set of one input Vertex.
func makeSet(vtx *gsd.Vertex) Set {
	set := make(Set)
	set[vtx] = vtx
	return set
}

// MakeSet creates a Set of one input Vertex
// with the vertex as Rep.
func MakeSet(vtx *gsd.Vertex) *DisJointSet {
	return &DisJointSet{
		Rep:  vtx,
		Mems: makeSet(vtx),
	}
}

// MakeGraphSet creates a Set of Vertices
// from the input Graph.
func MakeGraphSet(graph *gsd.Graph) []*DisJointSet {
	vertices := graph.GetVertices()
	result := []*DisJointSet{}
	for _, vtx := range *vertices {
		result = append(result, MakeSet(vtx.(*gsd.Vertex)))
	}
	return result
}

// SetContains returns true if the value exists in the set.
func (s DisJointSet) SetContains(vtx *gsd.Vertex) bool {
	if _, exist := s.Mems[vtx]; exist {
		return true
	} else {
		return false
	}
}

// GetSet returns the DisJointSet that vtx belongs to.
func GetSet(vtx *gsd.Vertex, ds []*DisJointSet) *DisJointSet {
	for _, set := range ds {
		if set.SetContains(vtx) {
			return set
		}
	}
	return nil
}

// FindSet returns the Rep of the DisJointSet that vtx belongs to.
func FindSet(vtx *gsd.Vertex, ds []*DisJointSet) *gsd.Vertex {
	set := GetSet(vtx, ds)
	if set != nil {
		return set.Rep
	}
	return nil
}

// UnionDisJointSet creates a Set of Vertices
// from the input Graph, setting s1's Rep as a new Rep.
func UnionDisJointSet(s1, s2 *DisJointSet) *DisJointSet {
	rs := NewDisJointSet()
	rs.Rep = s1.Rep
	union := s1.Mems.Union(s2.Mems)

	// Convert DisJointSet to Set
	set := make(Set)
	for _, vtx := range union {
		set[vtx] = vtx
	}

	// Mems Set
	rs.Mems = set
	return rs
}

// UnionByRep joins two subsets each of which is
// represented by its Rep x and y.
// We pass pointer to update the input slice.
func UnionByRep(rep1, rep2 *gsd.Vertex, ds *[]*DisJointSet) *[]*DisJointSet {
	var PI int // index that does merge
	var CI int // index to be merged

	// need to know which element to update
	for k, v := range *ds {
		if gsd.SameVertex(rep1, v.Rep) {
			PI = k
		}
		if gsd.SameVertex(rep2, v.Rep) {
			CI = k
		}
	}

	// set1 := GetSet(rep1, *ds)
	// set2 := GetSet(rep2, *ds)

	// merge into PI the DisJointSet in index CI
	// (X) UnionDisJointSet(*ds[PI], *ds[CI])
	ns := UnionDisJointSet((*ds)[PI], (*ds)[CI])
	(*ds)[PI] = ns

	// and delete the DisJointSet in CI
	copy((*ds)[CI:], (*ds)[CI+1:])

	// resize
	*ds = (*ds)[:len(*ds)-1 : len(*ds)-1]

	return ds
}

// UnionByVtx joins two subsets that contain x and y.
// It joins two subsets into a single subset and update
// the slice of DisJointSet.
func UnionByVtx(vtx1, vtx2 *gsd.Vertex, ds *[]*DisJointSet) *[]*DisJointSet {
	rep1 := FindSet(vtx1, *ds)
	rep2 := FindSet(vtx1, *ds)
	return UnionByRep(rep1, rep2, ds)
}

// (X) type EdgeSlice *slice.Sequence
type EdgeSlice slice.Sequence

func (s EdgeSlice) Len() int {
	return len(s)
}
func (s EdgeSlice) Less(i, j int) bool {
	return s[i].(*gsd.Edge).Weight < s[j].(*gsd.Edge).Weight
}
func (s EdgeSlice) Swap(i, j int) {
	// (X) s[i].(*gsd.Edge)
	// we only need to swap in the Sequence
	s[i], s[j] = s[j], s[i]
}

// SortEdges sorts Edges of graph in an increasing order
// by its weight and return the sorted Edges in slice.
func SortEdges(graph *gsd.Graph) *slice.Sequence {
	// func (g Graph) GetEdges() *slice.Sequence {
	edges := graph.GetEdges()
	sort.Sort(EdgeSlice(*edges))
	return edges
}
