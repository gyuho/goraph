package gl

import (
	"os"
	"testing"
)

func Test_JSON_ToMAP(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.003")
	rm := g.ToMAP()
	if len(rm) != 8 {
		t.Errorf("In testgraph1, expected 8 vertices but %+v", rm)
	}
	testCases3 := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 20.0}, // Updated (Added 6)
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
		wgt := rm[testCase.vertices[0]][testCase.vertices[1]][0]
		if wgt != testCase.weight {
			t.Errorf("In testgraph3, Expected '%#v'. But %#v.",
				testCase.weight, wgt,
				rm[testCase.vertices[0]][testCase.vertices[1]])
		}
	}
}

func Test_JSON_ToJSONFile(t *testing.T) {
	g3 := FromJSON("../../files/testgraph.json", "testgraph.003")
	g3.ToJSONFile("./tmp.json")
	g3j := FromJSON("./tmp.json", "goraph")
	l3j := g3j.GetVerticesSize()
	if l3j != 8 {
		t.Errorf("In testgraph3, expected 8 vertices but %v", l3j)
	}
	testCases3 := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 20.0}, // Updated (Added 6)
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
		wgt := g3j.GetEdgeWeight(g3j.FindVertexByID(testCase.vertices[0]), g3j.FindVertexByID(testCase.vertices[1]))
		if wgt != testCase.weight {
			t.Errorf("In testgraph3, Expected '%#v'. But %#v. %#v, %#v, %#v",
				testCase.weight, wgt,
				g3j.GetEdgeWeight(
					g3j.FindVertexByID("S"),
					g3j.FindVertexByID("A")),
				testCase.vertices[0], testCase.vertices[1])
		}
	}
	for _, testCase := range testCases3 {
		wgt := g3j.GetEdge(g3j.FindVertexByID(testCase.vertices[0]), g3j.FindVertexByID(testCase.vertices[1])).Weight
		if wgt != testCase.weight {
			t.Errorf("In testgraph3, Expected '%#v'. But %#v. %#v, %#v, %#v",
				testCase.weight, wgt,
				g3j.GetEdgeWeight(
					g3j.FindVertexByID("S"),
					g3j.FindVertexByID("A")),
				testCase.vertices[0], testCase.vertices[1])
		}
	}
	os.RemoveAll("./tmp.json")
}

func Test_JSON_GetVertices(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	l := g.GetVerticesSize()
	if l != 8 {
		t.Errorf("In testgraph1, expected 8 vertices but %v", l)
	}
}

func Test_JSON_GetVerticesSize(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	r := g.GetVerticesSize()
	if r != 8 {
		t.Errorf("In testgraph1, expected 8 vertices but %v", r)
	}
}

func Test_JSON_GetEdges(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	l := g.GetEdgesSize()
	// since it's bidirectional
	if l != 30 {
		t.Errorf("In testgraph1, expected 30 edges but %v", l)
	}
}

func Test_JSON_GetEdgesSize(t *testing.T) {
	g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
	r1 := g1.GetEdgesSize()
	if r1 != 30 {
		t.Errorf("In testgraph1, expected 30 edges but %v", r1)
	}
	g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
	r2 := g2.GetEdgesSize()
	if r2 != 24 {
		t.Errorf("In testgraph2, expected 24 edges but %v", r2)
	}
}

func Test_JSON_FindVertexByID(t *testing.T) {
	g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
	vs := g1.FindVertexByID("S")
	if vs.GetOutVerticesSize() != 3 {
		t.Errorf("In testgraph1, expected 3 but %v", vs.GetOutVerticesSize())
	}

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
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		r := g1.FindVertexByID(testCase.vtx)
		if r == nil {
			t.Errorf("In testgraph1, %+v should exist. But %#v", testCase.vtx, r)
		}
	}
	for _, testCase := range testCases {
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		r := g2.FindVertexByID(testCase.vtx)
		if r == nil {
			t.Errorf("In testgraph2, %+v should exist. But %#v", testCase.vtx, r)
		}
	}
}

func Test_JSON_CreateAndAddToGraph(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	_ = g.CreateAndAddToGraph("X")
	s := g.GetVerticesSize()
	if s != 9 {
		t.Errorf("In testgraph1, Created X vertex so it should now contain 9 vertices but %v", s)
	}
}

func Test_JSON_ImmediateDominate(t *testing.T) {
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
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		r := g2.ImmediateDominate(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		if r != testCase.imd {
			t.Errorf("testgraph2: %+v does not go to %#v with one edge", testCase.vts[0], testCase.vts[1])
		}
	}
}

func Test_JSON_Prev(t *testing.T) {
	g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
	pr := g2.ShowPrev(g2.FindVertexByID("S"))
	if pr != "Prev of S: " {
		t.Errorf("testgraph2: expected ~ but\n%v", pr)
	}

	g2.FindVertexByID("S").AddPrevVertex(NewVertex("X"))
	g2.FindVertexByID("S").AddPrevVertex(NewVertex("Y"))
	pr = g2.ShowPrev(g2.FindVertexByID("S"))
	if pr != "Prev of S: X, Y" {
		t.Errorf("testgraph2: expected ~ but\n%v", pr)
	}

	if prs := g2.FindVertexByID("S").GetPrevSize(); prs != 2 {
		t.Errorf("testgraph2: expected ~ but\n%v", prs)
	}
}

func Test_JSON_GetEdgeWeight(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
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
		weight := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		if weight != testCase.weight {
			t.Errorf("In testgraph1, Expected '%#v'. But %#v", weight, testCase.weight)
		}
	}

	g3 := FromJSON("../../files/testgraph.json", "testgraph.003")
	testCases3 := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 20.0}, // Updated (Added 6)
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
		wgt := g3.GetEdgeWeight(g3.FindVertexByID(testCase.vertices[0]), g3.FindVertexByID(testCase.vertices[1]))
		if wgt != testCase.weight {
			t.Errorf("In testgraph3, Expected '%#v'. But %#v. %#v, %#v, %#v", testCase.weight, wgt, g3.GetEdgeWeight(g3.FindVertexByID("S"), g3.FindVertexByID("A")), testCase.vertices[0], testCase.vertices[1])
		}
	}
}

func Test_JSON_UpdateWeight(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
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
		wgt := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		if wgt != testCase.weight {
			t.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, wgt)
		}
	}
}

func Test_JSON_DeleteVertex(t *testing.T) {
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
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		for _, v := range testCase.vts {
			g1.DeleteVertex(g1.FindVertexByID(v))
		}

		n := g1.GetEdgesSize()
		v := g1.GetVerticesSize()
		if n != testCase.edgnum {
			t.Errorf("testgraph1: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left. %+v", testCase.vts, testCase.edgnum, n, v, g1.FindVertexByID(testCase.vts[0]))
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
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		for _, v := range testCase.vts {
			g2.DeleteVertex(g2.FindVertexByID(v))
		}
		n := g2.GetEdgesSize()
		v := g2.GetVerticesSize()
		if n != testCase.edgnum {
			t.Errorf("testgraph2: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left.", testCase.vts, testCase.edgnum, n, v)
		}
	}
}

func Test_JSON_DeleteEdge(t *testing.T) {
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

	if g.GetEdgesSize() != 3 {
		t.Errorf("expected 3 edges but %v", g.GetEdgesSize())
	}

	g.DeleteEdge(a, c)
	if g.GetEdgesSize() != 2 {
		t.Errorf("expected 2 edges but %v", g.GetEdgesSize())
	}

	g.DeleteEdge(b, c)
	if g.GetEdgesSize() != 1 {
		t.Errorf("expected 1 edges but %v", g.GetEdgesSize())
	}

	g.DeleteEdge(a, b)
	if g.GetEdgesSize() != 0 {
		t.Errorf("expected 0 edge: but %v", g.GetEdgesSize())
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
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		o := g1.GetEdgesSize()
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[0]), g1.FindVertexByID(testCase.vts[1]))
		// g1.DeleteEdge(g1.FindVertexByID(testCase.vts[1]), g1.FindVertexByID(testCase.vts[0]))
		n := g1.GetEdgesSize()
		if n != testCase.edgnum {
			t.Errorf("In testgraph1, %+v is deleted. Expected '%#v' edges left. But %#v, originally %#v", testCase.vts, testCase.edgnum, n, o)
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
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		// g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			t.Errorf("In testgraph2, %+v is deleted. Expected '%#v' edges left. But %#v", testCase.vts, testCase.edgnum, n)
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
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[0]), g1.FindVertexByID(testCase.vts[1]))
		g1.DeleteEdge(g1.FindVertexByID(testCase.vts[1]), g1.FindVertexByID(testCase.vts[0]))
		n := g1.GetEdgesSize()
		if n != testCase.edgnum {
			t.Errorf("(Bi-direction) In testgraph1, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
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
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			t.Errorf("(Bi-direction) In testgraph2, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}
}
