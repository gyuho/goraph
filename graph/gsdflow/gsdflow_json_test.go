package gsdflow

import "testing"

func Test_JSONGraph(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	vs := g.GetVerticesSize()
	es := g.GetEdgesSize()

	if vs != 8 {
		test.Error("Should return 8 but: %v", vs)
	}

	if es != 31 {
		test.Error("Should return 31 but: %v", es)
	}
}

func Test_JSON_GetVertices(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	l := g.GetVerticesSize()
	if l != 8 {
		test.Error("In testgraph1, it should have 8 vertices but", l)
	}
}

func Test_JSON_GetVerticesSize(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	r := g.GetVerticesSize()
	if r != 8 {
		test.Error("In testgraph1, it should have 8 vertices but", r)
	}
}

func Test_JSON_GetEdges(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	l := g.GetEdgesSize()
	// since it's bidirectional
	if l != 30 {
		test.Error("In testgraph1, it should have 30 edges but", l)
	}
}

func Test_JSON_GetEdgesSize(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	r1 := g1.GetEdgesSize()
	if r1 != 30 {
		test.Error("In testgraph1, it should have 30 edges but", r1)
	}
	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	r2 := g2.GetEdgesSize()
	if r2 != 24 {
		test.Error("In testgraph2, It should have 24 edges but", r2)
	}
}

func Test_JSON_CreateAndAddToGraph(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	_ = g.CreateAndAddToGraph("X")
	s := g.GetVerticesSize()
	if s != 9 {
		test.Error("In testgraph1, Created X vertex so it should now contain 9 vertices but", s)
	}
}

func Test_JSON_GetOutVertices(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	d := g.FindVertexByID("D")
	l := d.GetOutVerticesSize()
	if l != 5 {
		test.Error("In testgraph1, D should have 5 outgoing vertices but", l)
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
		test.Error("In testgraph1, F and B should exist as outgoing vertices of D but", existF, existB)
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
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
			test.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.outedges, n)
		}
	}
}

func Test_JSON_GetInVertices(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	d := g.FindVertexByID("D")
	l := d.GetInVerticesSize()
	if l != 5 {
		test.Error("In testgraph1, D should have 5 outgoing edges but", l)
	}

	s := g.FindVertexByID("S")
	se := s.GetInVerticesSize()
	if se != 3 {
		test.Error("In testgraph1, S only have 3 incoming vertices but", se)
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
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
			test.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.inedges, n)
		}
	}
}

func Test_JSON_ImmediateDominate(test *testing.T) {
	testCases2 := []struct {
		vts []string
		imd bool
	}{
		{[]string{"S", "A"}, true},
		{[]string{"A", "S"}, true},
		{[]string{"A", "T"}, true},
		{[]string{"D", "E"}, false},
		{[]string{"E", "D"}, true},
		{[]string{"F", "S"}, false},
		{[]string{"B", "C"}, false},
		{[]string{"D", "A"}, false},
		{[]string{"E", "F"}, true},
		{[]string{"D", "A"}, false},
		{[]string{"T", "C"}, false},
	}
	for _, testCase := range testCases2 {
		g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
		r := g2.ImmediateDominate(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		if r != testCase.imd {
			test.Errorf("testgraph2: %+v does not go to %#v with one edge", testCase.vts[0], testCase.vts[1])
		}
	}
}

func Test_JSON_Connect(test *testing.T) {
	g := NewGraph()
	a := NewVertex("Google")
	b := NewVertex("Apple")
	c := NewVertex("Samsung")
	g.Connect(a, b, 0.0)
	g.Connect(a, c, 0.0)
	l := a.GetOutVerticesSize()
	if l != 2 {
		test.Error("Google should have 2 outgoing edges but", l)
	}
	le := g.GetEdgesSize()
	if le != 2 {
		test.Error("This graph should have 2 edges but", le)
	}
}

func Test_JSON_FindVertexByID(test *testing.T) {
	testCases := []struct {
		vtx   string
		exist bool
	}{
		{"S", true},
		{"A", true},
		{"B", true},
		{"C", true},
		{"D", true},
		{"E", true},
		{"F", true},
		{"T", true},
	}
	for _, testCase := range testCases {
		g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
		r := g1.FindVertexByID(testCase.vtx)
		if r == nil {
			test.Errorf("In testgraph1, %+v should exist. But %#v", testCase.vtx, r)
		}
	}
	for _, testCase := range testCases {
		g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
		r := g2.FindVertexByID(testCase.vtx)
		if r == nil {
			test.Errorf("In testgraph2, %+v should exist. But %#v", testCase.vtx, r)
		}
	}
}

func Test_JSON_DeleteInVertex(test *testing.T) {
	g := NewGraph()
	a := g.CreateAndAddToGraph("Google")
	b := g.CreateAndAddToGraph("Apple")
	c := g.CreateAndAddToGraph("Samsung")
	g.Connect(b, a, 0.0)
	g.Connect(c, a, 0.0)
	a.DeleteInVertex(b)
	n := a.GetInVerticesSize()
	if n != 1 {
		test.Error("Should only have 1 incoming vertex:", n)
	}
}

func Test_JSON_DeleteOutVertex(test *testing.T) {
	g := NewGraph()
	a := g.CreateAndAddToGraph("Google")
	b := g.CreateAndAddToGraph("Apple")
	c := g.CreateAndAddToGraph("Samsung")
	g.Connect(a, b, 0.0)
	g.Connect(a, c, 0.0)
	a.DeleteOutVertex(b)
	n := a.GetOutVerticesSize()
	if n != 1 {
		test.Error("Should only have 1 outgoing vertex:", n)
	}
}

func Test_JSON_DeleteEdge(test *testing.T) {
	g := NewGraph()
	a := g.CreateAndAddToGraph("Google")
	b := g.CreateAndAddToGraph("Apple")
	c := g.CreateAndAddToGraph("Samsung")
	g.Connect(a, b, 0.0)
	g.Connect(a, b, 1.0)
	g.Connect(a, b, 2.0)
	g.Connect(a, b, 3.0)
	g.Connect(a, c, 4.0)
	g.Connect(a, c, 5.0)
	g.Connect(a, c, 6.0)
	g.Connect(b, c, 7.0)

	g.DeleteEdge(a, c)
	if g.GetEdgesSize() != 5 {
		test.Error("Should only have 5 edges:", g.GetEdgesSize())
	}

	g.DeleteEdge(b, c)
	if g.GetEdgesSize() != 4 {
		test.Error("Should only have 4 edges:", g.GetEdgesSize())
	}

	g.DeleteEdge(a, b)
	if g.GetEdgesSize() != 0 {
		test.Error("Should only have 0 edge:", g.GetEdgesSize())
	}

	testCases1 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 29},
		{[]string{"S", "A"}, 29},
		{[]string{"S", "C"}, 29},
		{[]string{"A", "B"}, 29},
		{[]string{"A", "D"}, 29},
		{[]string{"A", "S"}, 29},
		{[]string{"A", "T"}, 29},
		{[]string{"T", "A"}, 29},
		{[]string{"D", "E"}, 29},
		{[]string{"E", "D"}, 29},
		{[]string{"C", "E"}, 29},
		{[]string{"B", "E"}, 29},
		{[]string{"D", "T"}, 29},
		{[]string{"T", "D"}, 29},
		{[]string{"F", "E"}, 29},
		{[]string{"E", "F"}, 29},
		{[]string{"E", "T"}, 29},
	}
	for _, testCase := range testCases1 {
		g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
		o := g1.GetEdgesSize()
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[0]), g1.FindVertexByID(testCase.vts[1]))
		// g1.DeleteEdge(g1.FindVertexByID(testCase.vts[1]), g1.FindVertexByID(testCase.vts[0]))
		n := g1.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("In testgraph1, %+v is deleted. Expected '%#v' edges left. But %#v, originally %#v", testCase.vts, testCase.edgnum, n, o)
		}
	}

	testCases2 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 23},
		{[]string{"S", "A"}, 23},
		{[]string{"S", "C"}, 23},
		{[]string{"A", "B"}, 23},
		{[]string{"A", "D"}, 23},
		{[]string{"A", "S"}, 23},
		{[]string{"A", "T"}, 23},
		{[]string{"T", "A"}, 23},
		{[]string{"D", "E"}, 24},
		{[]string{"E", "D"}, 23},
		{[]string{"C", "E"}, 23},
		{[]string{"B", "E"}, 23},
		{[]string{"D", "T"}, 24},
		{[]string{"T", "D"}, 23},
		{[]string{"F", "E"}, 23},
		{[]string{"E", "F"}, 23},
		{[]string{"E", "T"}, 23},
	}
	for _, testCase := range testCases2 {
		g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		// g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("In testgraph2, %+v is deleted. Expected '%#v' edges left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}

	testCases1b := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 28},
		{[]string{"S", "A"}, 28},
		{[]string{"S", "C"}, 28},
		{[]string{"A", "B"}, 28},
		{[]string{"A", "D"}, 28},
		{[]string{"A", "S"}, 28},
		{[]string{"A", "T"}, 28},
		{[]string{"T", "A"}, 28},
		{[]string{"D", "E"}, 28},
		{[]string{"E", "D"}, 28},
		{[]string{"C", "E"}, 28},
		{[]string{"B", "E"}, 28},
		{[]string{"D", "T"}, 28},
		{[]string{"T", "D"}, 28},
		{[]string{"F", "E"}, 28},
		{[]string{"E", "F"}, 28},
		{[]string{"E", "T"}, 28},
	}
	for _, testCase := range testCases1b {
		g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[0]), g1.FindVertexByID(testCase.vts[1]))
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[1]), g1.FindVertexByID(testCase.vts[0]))
		n := g1.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("(Bi-direction) In testgraph1, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}

	testCases2b := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 22},
		{[]string{"S", "A"}, 22},
		{[]string{"S", "C"}, 22},
		{[]string{"A", "B"}, 22},
		{[]string{"A", "D"}, 23},
		{[]string{"A", "S"}, 22},
		{[]string{"F", "D"}, 23},
		{[]string{"T", "A"}, 22},
		{[]string{"D", "E"}, 23},
		{[]string{"E", "D"}, 23},
		{[]string{"C", "E"}, 23},
		{[]string{"B", "E"}, 22},
		{[]string{"D", "T"}, 23},
		{[]string{"T", "D"}, 23},
		{[]string{"F", "E"}, 22},
		{[]string{"E", "F"}, 22},
		{[]string{"E", "T"}, 22},
	}
	for _, testCase := range testCases2b {
		g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("(Bi-direction) In testgraph2, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}
}

func Test_JSON_DeleteVertex(test *testing.T) {
	testCases1 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "A", "D"}, 10},
		{[]string{"A", "T", "C"}, 12},
		{[]string{"D", "E", "S"}, 6},
		{[]string{"F", "S", "T"}, 12},
		{[]string{"B", "C", "D"}, 10},
		{[]string{"D", "A", "E"}, 6},
		{[]string{"E", "F", "A"}, 8},
		{[]string{"D", "A", "T"}, 10},
	}
	for _, testCase := range testCases1 {
		g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
		for _, v := range testCase.vts {
			g1.DeleteVertex(g1.FindVertexByID(v))
		}

		n := g1.GetEdgesSize()
		v := g1.GetVerticesSize()
		if n != testCase.edgnum {
			test.Errorf("testgraph1: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left. %+v", testCase.vts, testCase.edgnum, n, v, g1.FindVertexByID(testCase.vts[0]))
		}
	}

	testCases2 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "A", "D"}, 9},
		{[]string{"A", "T", "C"}, 9},
		{[]string{"D", "E", "S"}, 6},
		{[]string{"F", "S", "T"}, 8},
		{[]string{"B", "C", "D"}, 10},
		{[]string{"D", "A", "E"}, 6},
		{[]string{"E", "F", "A"}, 6},
		{[]string{"D", "A", "T"}, 9},
	}
	for _, testCase := range testCases2 {
		g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
		for _, v := range testCase.vts {
			g2.DeleteVertex(g2.FindVertexByID(v))
		}
		n := g2.GetEdgesSize()
		v := g2.GetVerticesSize()
		if n != testCase.edgnum {
			test.Errorf("testgraph2: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left.", testCase.vts, testCase.edgnum, n, v)
		}
	}
}
