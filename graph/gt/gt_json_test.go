package gt

import "testing"

func Test_JSON_ParseToGraph(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	if g1.GetVerticesSize() != 8 {
		test.Error("The graph should be of vertex size 8 but:", g1.GetVerticesSize())
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	if g2.GetVerticesSize() != 8 {
		test.Error("The graph should be of vertex size 8 but:", g2.GetVerticesSize())
	}
}

func Test_JSON_GetVertices(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	vs1 := g1.GetVertices()
	if len(vs1) != 8 {
		test.Errorf("The graph should be of 8 vertices but: %v", vs1)
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	vs2 := g2.GetVertices()
	if len(vs2) != 8 {
		test.Errorf("The graph should be of 8 vertices but: %v", vs2)
	}
}

func Test_JSON_GetEdgesSize(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	if g1.GetEdgesSize() != 30 {
		test.Error("The graph should be of edge size 30 but:", g1.GetEdgesSize())
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	if g2.GetEdgesSize() != 24 {
		test.Error("The graph should be of edge size 24 but:", g1.GetEdgesSize())
	}
}

func Test_JSON_GetEdgeWeight(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	if g1.GetEdgeWeight("S", "C") != 200 {
		test.Error("Should return 200 but:", g1.GetEdgeWeight("S", "C"))
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	if g2.GetEdgeWeight("D", "B") != 0.0 {
		test.Error("Should return 0.0 but:", g2.GetEdgeWeight("D", "B"))
	}
	if g2.GetEdgeWeight("E", "D") != 2.0 {
		test.Error("Should return 2.0 but:", g2.GetEdgeWeight("E", "D"))
	}
}

func Test_JSON_RemoveEdge(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	g1.RemoveEdge("S", "C")
	if g1.GetEdgesSize() != 29 {
		test.Error("Should return 29 but:", g1.GetEdgesSize())
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	g2.RemoveEdge("D", "B")
	if g2.GetEdgesSize() != 24 {
		test.Error("Should return 24 but:", g1.GetEdgesSize())
	}

	g2.RemoveEdge("E", "D")
	if g2.GetEdgesSize() != 23 {
		test.Error("Should return 23 but:", g1.GetEdgesSize())
	}
}

func Test_JSON_HasEdge(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	if g1.HasEdge("S", "C") != true {
		test.Error("Should return true but:", g1.HasEdge("S", "C"))
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	if g2.HasEdge("D", "B") != false {
		test.Error("Should return false but:", g2.HasEdge("D", "B"))
	}

	if g2.HasEdge("E", "D") != true {
		test.Error("Should return true but:", g2.HasEdge("E", "D"))
	}
}

func Test_JSON_GetInVertices(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	if len(g1.GetInVertices("S")) != 3 {
		test.Error("Should return 3 but:", len(g1.GetInVertices("S")))
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	if len(g2.GetInVertices("B")) != 3 {
		test.Error("Should return 3 but:", len(g2.GetInVertices("B")))
	}

	if len(g2.GetInVertices("D")) != 5 {
		test.Error("Should return 5 but:", len(g2.GetInVertices("D")))
	}
}

func Test_JSON_GetOutVertices(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	if len(g1.GetOutVertices("S")) != 3 {
		test.Error("Should return 3 but:", len(g1.GetOutVertices("S")))
	}

	g2 := JSONGraph("../../example_files/testgraph.json", "testgraph.002")
	if len(g2.GetOutVertices("B")) != 4 {
		test.Error("Should return 4 but:", len(g2.GetOutVertices("B")))
	}

	if len(g2.GetOutVertices("D")) != 0 {
		test.Error("Should return 0 but:", len(g2.GetOutVertices("D")))
	}

	if len(g2.GetOutVertices("C")) != 2 {
		test.Error("Should return 2 but:", len(g2.GetOutVertices("C")))
	}
}

func Test_JSON_Initialize(test *testing.T) {
	g1 := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	g1.Initialize(99999.999)

	cnum := true
	mat := g1.Matrix
	for r := range mat {
		for c := range mat[r] {
			if mat[r][c] != 99999.999 {
				cnum = false
			}
		}
	}

	if !cnum {
		test.Error("Should return true but %v", mat)
	}
}
