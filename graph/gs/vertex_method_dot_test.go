package gs

import "testing"

func Test_DOT_GetOutVertices(t *testing.T) {
	g := FromDOT("../../files/testgraph.001.dot")
	d := g.FindVertexByID("D")
	l := d.GetOutVerticesSize()
	if l != 5 {
		t.Errorf("In testgraph1, D should have 5 outgoing vertices but %v", l)
	}

	slice := d.GetOutVertices()
	existF := false
	for _, v := range *slice {
		if "F" == v.(*Vertex).ID {
			existF = true
		}
	}

	existB := false
	for _, v := range *slice {
		if "B" == v.(*Vertex).ID {
			existB = true
		}
	}

	if !existF || !existB {
		t.Errorf("In testgraph1, F and B should exist as outgoing vertices of D but %v", existF, existB)
	}

	g2 := FromDOT("../../files/testgraph.002.dot")
	testCases := []struct {
		vtx      string
		outedges int
	}{
		{"S", 3},
		{"A", 4},
		{"B", 4},
		{"C", 2},
		{"D", 0},
		{"E", 4},
		{"F", 3},
		{"T", 4},
	}

	for _, testCase := range testCases {
		v := g2.FindVertexByID(testCase.vtx)
		n := v.GetOutVerticesSize()
		if n != testCase.outedges {
			t.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.outedges, n)
		}
	}
}

func Test_DOT_GetInVertices(t *testing.T) {
	g := FromDOT("../../files/testgraph.001.dot")
	d := g.FindVertexByID("D")
	l := d.GetInVerticesSize()
	if l != 5 {
		t.Errorf("In testgraph1, D should have 5 outgoing edges but %v", l)
	}

	s := g.FindVertexByID("S")
	se := s.GetInVerticesSize()
	if se != 3 {
		t.Errorf("In testgraph1, S only have 3 incoming vertices but %v", se)
	}

	g2 := FromDOT("../../files/testgraph.002.dot")
	testCases := []struct {
		vtx     string
		inedges int
	}{
		{"S", 3},
		{"A", 3},
		{"B", 3},
		{"C", 1},
		{"D", 5},
		{"E", 4},
		{"F", 2},
		{"T", 3},
	}

	for _, testCase := range testCases {
		v := g2.FindVertexByID(testCase.vtx)
		n := v.GetInVerticesSize()
		if n != testCase.inedges {
			t.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.inedges, n)
		}
	}
}
