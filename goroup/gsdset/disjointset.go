package gsdset

import "github.com/gyuho/goraph/graph/gsd"

type DisJointSet struct {
	Rep  *gsd.Vertex // Representitive Vertex
	Mems Set         // Set Members
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
	set[vtx] = 1
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
func MakeGraphSet(gp *gsd.Graph) []*DisJointSet {
	vertices := gp.GetVertices()
	result := []*DisJointSet{}
	for _, vtx := range *vertices {
		result = append(result, MakeSet(vtx.(*gsd.Vertex)))
	}
	return result
}

// UnionDisJointSet creates a Set of Vertices
// from the input Graph, setting s1's Rep as a new Rep.
func UnionDisJointSet(s1, s2 *DisJointSet) *DisJointSet {
	rs := NewDisJointSet()
	rs.Rep = s1.Rep
	union := s1.Mems.Union(s2.Mems)
	set := make(Set)
	for _, vtx := range union {
		set[vtx] = 1
	}
	rs.Mems = set
	return rs
}

// FindSet returns the Rep of the DisJointSet
// that vtx belongs to.
func FindSet(vtx *gsd.Vertex, ds []*DisJointSet) *DisJointSet {
	for _, set := range ds {
		if gsd.SameVertex(vtx, set.Rep) {
			return set
		}
	}
	return nil
}
