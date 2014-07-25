package gl

import "testing"

func TestGetEdgeWeightDupl(t *testing.T) {
	g := FromJSONDupl("../../files/testgraph.json", "testgraph.001")
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
		weights := g.GetEdgeWeightDupl(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			t.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, weights)
		}
	}

	g3 := FromJSONDupl("../../files/testgraph.json", "testgraph.003")
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
		weights := g3.GetEdgeWeightDupl(g3.FindVertexByID(testCase.vertices[0]), g3.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			t.Errorf("In testgraph3, Expected '%#v'. But %#v. %#v, %#v, %#v", testCase.weight, weights, g3.GetEdgeWeightDupl(g3.FindVertexByID("S"), g3.FindVertexByID("A")), testCase.vertices[0], testCase.vertices[1])
		}
	}
}

func Test_JSON_DeleteEdge_Dupl(t *testing.T) {
	g := NewGraph()
	a := g.CreateAndAddToGraph("Google")
	b := g.CreateAndAddToGraph("Apple")
	c := g.CreateAndAddToGraph("Samsung")
	g.ConnectDupl(a, b, 0.0)
	g.ConnectDupl(a, b, 1.0)
	g.ConnectDupl(a, b, 2.0)
	g.ConnectDupl(a, b, 3.0)
	g.ConnectDupl(a, c, 4.0)
	g.ConnectDupl(a, c, 5.0)
	g.ConnectDupl(a, c, 6.0)
	g.ConnectDupl(b, c, 7.0)

	// Connect would only create 3 edges
	if g.GetEdgesSize() != 8 {
		t.Errorf("expected 8 edges but %v", g.GetEdgesSize())
	}

	g.DeleteEdge(a, c)
	if g.GetEdgesSize() != 5 {
		t.Errorf("expected 5 edges but %v", g.GetEdgesSize())
	}

	g.DeleteEdge(b, c)
	if g.GetEdgesSize() != 4 {
		t.Errorf("expected 4 edges but %v", g.GetEdgesSize())
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
		g1 := FromJSONDupl("../../files/testgraph.json", "testgraph.001")
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
		g2 := FromJSONDupl("../../files/testgraph.json", "testgraph.002")
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
		g1 := FromJSONDupl("../../files/testgraph.json", "testgraph.001")
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
		g2 := FromJSONDupl("../../files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[0]), g2.FindVertexByID(testCase.vts[1]))
		g2.DeleteEdge(g2.FindVertexByID(testCase.vts[1]), g2.FindVertexByID(testCase.vts[0]))
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			t.Errorf("(Bi-direction) In testgraph2, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}
}

func Test_DeleteVertex(t *testing.T) {
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
		g1 := FromJSONDupl("../../files/testgraph.json", "testgraph.001")
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
		g2 := FromJSONDupl("../../files/testgraph.json", "testgraph.002")
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
