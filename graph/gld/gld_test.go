package gld

import (
	"reflect"
	"testing"
)

/*
A|B,7|C,9|F,20
Connect A with B with an edge from A "to" B of weight value 7
Connect A with C with an edge from A "to" C of weight value 9
Connect A with F with an edge from A "to" F of weight value 20

F|A,20|C,2|E,9
Connect F with A with an edge from F "to" A of weight value 20
Connect F with C with an edge from F "to" C of weight value 2
Connect F with E with an edge from F "to" E of weight value 9
*/

var testgraph1 string = `
S|A,100|C,200|B,14
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`

var testgraph2 string = `
S|A,100|C,200|B,14
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
E|B,18|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`

func Test_NewGraph(test *testing.T) {
	a := NewGraph()
	r := reflect.TypeOf(a)
	c := &Graph{}
	if r != reflect.TypeOf(c) {
		test.Error("Type Error")
	}
}

func Test_GetVertices(test *testing.T) {
	g := ParseToGraph(testgraph1)
	r := g.GetVertices()
	l := r.Len()
	if l != 8 {
		test.Error("In testgraph1, It should have 8 vertices but", l)
	}
}

func Test_GetEdges(test *testing.T) {
	g := ParseToGraph(testgraph1)
	edges := g.GetEdges()
	check := false
	for e := edges.Front(); e != nil; e = e.Next() {
		if e.Value.(*Edge).Weight == 200 {
			check = true
		}
	}
	if !check {
		test.Error("In testgraph1, there must be an edge of weight 200")
	}
}

func Test_GetVerticesSize(test *testing.T) {
	g := ParseToGraph(testgraph1)
	r := g.GetVerticesSize()
	if r != 8 {
		test.Error("In testgraph1, It should have 8 vertices but", r)
	}
}

func Test_GetEdgesSize(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	r1 := g1.GetEdgesSize()
	if r1 != 30 {
		test.Error("In testgraph1, It should have 30 edges but", r1)
	}
	g2 := ParseToGraph(testgraph2)
	r2 := g2.GetEdgesSize()
	if r2 != 24 {
		test.Error("In testgraph2, It should have 24 edges but", r2)
	}
}

func Test_NewVertex(test *testing.T) {
	v := NewVertex("Google")
	if v.ID != "Google" {
		test.Error("ID should be Google but", v.ID)
	}

	r := v.OutVertices.Len()
	if r != 0 {
		test.Error("OutVertices should be empty but", r)
	}

	c := v.StampD
	if c != 9999999999 {
		test.Error("StampD should be 9999999999 but", c)
	}
}

func Test_NewEdge(test *testing.T) {
	a := NewVertex("Google")
	b := NewVertex("Apple")
	e := NewEdge(a, b, 123.45)
	if e.Src.ID != a.ID {
		test.Error("Source vertex should be Google but", e.Src.ID)
	}
	if e.Dst.ID != b.ID {
		test.Error("Destination vertex should be Apple but", e.Dst.ID)
	}
}

func Test_ImmediateDominate(test *testing.T) {
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
		g2 := ParseToGraph(testgraph2)
		r := g2.ImmediateDominate(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		if r != testCase.imd {
			test.Errorf("testgraph2: %+v does not go to %#v with one edge", testCase.vts[0], testCase.vts[1])
		}
	}
}

func Test_GetEdgeWeight(test *testing.T) {
	g := ParseToGraph(testgraph1)
	testCases := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 14.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},
		{[]string{"T", "A"}, 44.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"C", "E"}, 24.0},
		{[]string{"B", "E"}, 18.0},
		{[]string{"D", "T"}, 16.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},
		{[]string{"S", "C"}, 200.0},
		{[]string{"S", "A"}, 100.0},
	}
	for _, testCase := range testCases {
		weights := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, weights)
		}
	}

	var testgraph3 string = `
S|A,100|C,200|B,14|B,6
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`

	g3 := ParseToGraph(testgraph3)
	testCases3 := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 14.0}, // duplicate edges
		{[]string{"S", "B"}, 6.0},  // duplicate edges
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},
		{[]string{"T", "A"}, 44.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"C", "E"}, 24.0},
		{[]string{"B", "E"}, 18.0},
		{[]string{"D", "T"}, 16.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},
		{[]string{"S", "C"}, 200.0},
		{[]string{"S", "A"}, 100.0},
	}
	for _, testCase := range testCases3 {
		weights := g3.GetEdgeWeight(g3.FindVertexByID(testCase.vertices[0]), g3.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			test.Errorf("In testgraph3, Expected '%#v'. But %#v. %#v, %#v, %#v", testCase.weight, weights, g3.GetEdgeWeight(g3.FindVertexByID("S"), g3.FindVertexByID("A")), testCase.vertices[0], testCase.vertices[1])
		}
	}
}

func Test_UpdateWeight(test *testing.T) {
	g := ParseToGraph(testgraph1)
	testCases := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 914.0},
		{[]string{"A", "B"}, 95.0},
		{[]string{"A", "D"}, 920.0},
		{[]string{"A", "T"}, 944.0},
		{[]string{"T", "A"}, 944.0},
		{[]string{"D", "E"}, 92.0},
		{[]string{"E", "D"}, 92.0},
		{[]string{"C", "E"}, 924.0},
		{[]string{"B", "E"}, 918.0},
		{[]string{"D", "T"}, 916.0},
		{[]string{"T", "D"}, 916.0},
		{[]string{"F", "E"}, 96.0},
		{[]string{"E", "F"}, 96.0},
		{[]string{"E", "T"}, 919.0},
		{[]string{"S", "C"}, 9200.0},
		{[]string{"S", "A"}, 9100.0},
	}
	for _, testCase := range testCases {
		g.UpdateWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), testCase.weight)
		weights := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, weights)
		}
	}
}

func Test_Connect(test *testing.T) {
	g := NewGraph()
	a := NewVertex("Google")
	b := NewVertex("Apple")
	c := NewVertex("Samsung")
	g.Connect(a, b, 0.0)
	g.Connect(a, c, 0.0)
	l := a.OutVertices.Len()
	if l != 2 {
		test.Error("Google should have 2 outgoing edges but", l)
	}
	le := g.GetEdgesSize()
	if le != 2 {
		test.Error("This graph should have 2 edges but", le)
	}
}

func Test_GetOutVertices(test *testing.T) {
	g := ParseToGraph(testgraph1)
	d := g.FindVertexByID("D")
	l := d.GetOutVertices().Len()
	if l != 5 {
		test.Error("In testgraph1, D should have 5 outgoing vertices but", l)
	}

	list := d.GetOutVertices()
	existF := false
	for vtx := list.Front(); vtx != nil; vtx = vtx.Next() {
		if "F" == vtx.Value.(*Vertex).ID {
			existF = true
		}
	}

	existB := false
	for vtx := list.Front(); vtx != nil; vtx = vtx.Next() {
		if "B" == vtx.Value.(*Vertex).ID {
			existB = true
		}
	}

	if !existF || !existB {
		test.Error("In testgraph1, F and B should exist as outgoing vertices of D but", existF, existB)
	}

	g2 := ParseToGraph(testgraph2)
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
		n := v.GetOutVertices().Len()
		if n != testCase.outedges {
			test.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.outedges, n)
		}
	}
}

func Test_GetInVertices(test *testing.T) {
	g := ParseToGraph(testgraph1)
	d := g.FindVertexByID("D")
	l := d.GetInVertices().Len()
	if l != 5 {
		test.Error("In testgraph1, D should have 5 outgoing edges but", l)
	}

	s := g.FindVertexByID("S")
	se := s.GetInVertices().Len()
	if se != 3 {
		test.Error("In testgraph1, S only have 3 incoming vertices but", se)
	}

	g2 := ParseToGraph(testgraph2)
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
		n := v.GetInVertices().Len()
		if n != testCase.inedges {
			test.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.inedges, n)
		}
	}
}

func Test_FindVertexByID(test *testing.T) {
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
		g1 := ParseToGraph(testgraph1)
		r := g1.FindVertexByID(testCase.vtx)
		if r == nil {
			test.Errorf("In testgraph1, %+v should exist. But %#v", testCase.vtx, r)
		}
	}
	for _, testCase := range testCases {
		g2 := ParseToGraph(testgraph2)
		r := g2.FindVertexByID(testCase.vtx)
		if r == nil {
			test.Errorf("In testgraph2, %+v should exist. But %#v", testCase.vtx, r)
		}
	}
}

func Test_CreateAndAddToGraph(test *testing.T) {
	g := ParseToGraph(testgraph1)
	_ = g.CreateAndAddToGraph("X")
	s := g.GetVerticesSize()
	if s != 9 {
		test.Error("In testgraph1, Created X vertex so it should now contain 9 vertices but", s)
	}
}

func Test_DeleteEdge(test *testing.T) {
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
		g1 := ParseToGraph(testgraph1)
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[0]), g1.FindVertexByID(testCase.vts[1]))
		// g1.DeleteEdge(g1.FindVertexByID(testCase.vts[1]), g1.FindVertexByID(testCase.vts[0]))
		n := g1.GetEdges().Len()
		if n != testCase.edgnum {
			test.Errorf("In testgraph1, %+v is deleted. Expected '%#v' edges left. But %#v", testCase.vts, testCase.edgnum, n)
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
		g2 := ParseToGraph(testgraph2)
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		// g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdges().Len()
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
		g1 := ParseToGraph(testgraph1)
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[0]), g1.FindVertexByID(testCase.vts[1]))
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[1]), g1.FindVertexByID(testCase.vts[0]))
		n := g1.GetEdges().Len()
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
		g2 := ParseToGraph(testgraph2)
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdges().Len()
		if n != testCase.edgnum {
			test.Errorf("(Bi-direction) In testgraph2, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}
}

func Test_DeleteVertex(test *testing.T) {
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
		g1 := ParseToGraph(testgraph1)
		for _, v := range testCase.vts {
			g1.DeleteVertex(g1.FindVertexByID(v))
		}
		n := g1.GetEdgesSize()
		v := g1.GetVerticesSize()
		if n != testCase.edgnum {
			test.Errorf("testgraph1: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left.", testCase.vts, testCase.edgnum, n, v)
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
		g2 := ParseToGraph(testgraph2)
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
