package gsd

import "testing"

func Test_JSONGraphT(test *testing.T) {
	g7 := JSONGraphT("../../example_files/testgraph.json", "testgraph.007")
	if g7.FindVertexByID("B").GetInVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g7.FindVertexByID("B").GetInVerticesSize())
	}
	if g7.FindVertexByID("B").GetOutVerticesSize() != 0 {
		test.Errorf("Should be 0 but %+v", g7.FindVertexByID("B").GetOutVerticesSize())
	}
	if g7.FindVertexByID("D").GetInVerticesSize() != 3 {
		test.Errorf("Should be 3 but %+v", g7.FindVertexByID("D").GetInVerticesSize())
	}
	if g7.FindVertexByID("D").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g7.FindVertexByID("D").GetOutVerticesSize())
	}
	if g7.FindVertexByID("E").GetInVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g7.FindVertexByID("E").GetInVerticesSize())
	}
	if g7.FindVertexByID("E").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g7.FindVertexByID("E").GetOutVerticesSize())
	}
	if g7.FindVertexByID("H").GetInVerticesSize() != 0 {
		test.Errorf("Should be 0 but %+v", g7.FindVertexByID("H").GetInVerticesSize())
	}
	if g7.FindVertexByID("H").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g7.FindVertexByID("H").GetOutVerticesSize())
	}

	g10 := JSONGraphT("../../example_files/testgraph.json", "testgraph.010")
	if g10.FindVertexByID("A").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("A").GetInVerticesSize())
	}
	if g10.FindVertexByID("A").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("A").GetOutVerticesSize())
	}
	if g10.FindVertexByID("B").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("B").GetInVerticesSize())
	}
	if g10.FindVertexByID("B").GetOutVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g10.FindVertexByID("B").GetOutVerticesSize())
	}
	if g10.FindVertexByID("C").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("C").GetInVerticesSize())
	}
	if g10.FindVertexByID("C").GetOutVerticesSize() != 4 {
		test.Errorf("Should be 4 but %+v", g10.FindVertexByID("C").GetOutVerticesSize())
	}
	if g10.FindVertexByID("D").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("D").GetInVerticesSize())
	}
	if g10.FindVertexByID("D").GetOutVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g10.FindVertexByID("D").GetOutVerticesSize())
	}
	if g10.FindVertexByID("E").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("E").GetInVerticesSize())
	}
	if g10.FindVertexByID("E").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("E").GetOutVerticesSize())
	}
}

/*
func Test_Transpose(test *testing.T) {
	g7 := JSONGraph("../../example_files/testgraph.json", "testgraph.007")
	g7.Transpose()
	if g7.FindVertexByID("B").GetInVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g7.FindVertexByID("B").GetInVerticesSize())
	}
	if g7.FindVertexByID("B").GetOutVerticesSize() != 0 {
		test.Errorf("Should be 0 but %+v", g7.FindVertexByID("B").GetOutVerticesSize())
	}
	if g7.FindVertexByID("D").GetInVerticesSize() != 3 {
		test.Errorf("Should be 3 but %+v", g7.FindVertexByID("D").GetInVerticesSize())
	}
	if g7.FindVertexByID("D").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g7.FindVertexByID("D").GetOutVerticesSize())
	}
	if g7.FindVertexByID("E").GetInVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g7.FindVertexByID("E").GetInVerticesSize())
	}
	if g7.FindVertexByID("E").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g7.FindVertexByID("E").GetOutVerticesSize())
	}
	if g7.FindVertexByID("H").GetInVerticesSize() != 0 {
		test.Errorf("Should be 0 but %+v", g7.FindVertexByID("H").GetInVerticesSize())
	}
	if g7.FindVertexByID("H").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g7.FindVertexByID("H").GetOutVerticesSize())
	}

	g10 := JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10.Transpose()
	if g10.FindVertexByID("A").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("A").GetInVerticesSize())
	}
	if g10.FindVertexByID("A").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("A").GetOutVerticesSize())
	}
	if g10.FindVertexByID("B").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("B").GetInVerticesSize())
	}
	if g10.FindVertexByID("B").GetOutVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g10.FindVertexByID("B").GetOutVerticesSize())
	}
	if g10.FindVertexByID("C").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("C").GetInVerticesSize())
	}
	if g10.FindVertexByID("C").GetOutVerticesSize() != 4 {
		test.Errorf("Should be 4 but %+v", g10.FindVertexByID("C").GetOutVerticesSize())
	}
	if g10.FindVertexByID("D").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("D").GetInVerticesSize())
	}
	if g10.FindVertexByID("D").GetOutVerticesSize() != 1 {
		test.Errorf("Should be 1 but %+v", g10.FindVertexByID("D").GetOutVerticesSize())
	}
	if g10.FindVertexByID("E").GetInVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("E").GetInVerticesSize())
	}
	if g10.FindVertexByID("E").GetOutVerticesSize() != 2 {
		test.Errorf("Should be 2 but %+v", g10.FindVertexByID("E").GetOutVerticesSize())
	}
}
*/
