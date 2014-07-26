package gs

import "testing"

func Test_CopyVertex(test *testing.T) {
	a := NewVertex("Google")
	b := CopyVertex(a)
	if !SameVertex(a, b) {
		test.Error("Should be same but %v", SameVertex(a, b))
	}

	g1 := FromJSON("../../files/testgraph.json", "testgraph.003")
	s1 := g1.FindVertexByID("S")
	s2 := CopyVertex(s1)
	if !SameVertex(s1, s2) {
		test.Error("Should be same but %v", SameVertex(s1, s2))
	}
}

func Test_CopyEdge(test *testing.T) {
	g1 := FromJSON("../../files/testgraph.json", "testgraph.003")
	s := g1.FindVertexByID("S")
	b := g1.FindVertexByID("B")
	edge1 := g1.GetEdge(s, b)
	edge2 := CopyEdge(edge1)

	if !SameEdge(edge1, edge2) {
		test.Error("Should be same but %v", SameEdge(edge1, edge2))
	}
}
