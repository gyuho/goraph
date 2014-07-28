package gs

import "testing"

func TestPath(t *testing.T) {
	g4 := FromJSON("../../files/testgraph.json", "testgraph.004")
	rb4 := g4.Path(g4.FindVertexByID("S"), g4.FindVertexByID("F"))
	if !rb4 {
		t.Errorf("expected true but %+v\n", rb4)
	}

	g5 := FromJSON("../../files/testgraph.json", "testgraph.005")
	rb5 := g5.Path(g5.FindVertexByID("F"), g5.FindVertexByID("B"))
	if !rb5 {
		t.Errorf("expected true but %+v\n", rb5)
	}

	g16 := FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16 := g16.Path(g16.FindVertexByID("C"), g16.FindVertexByID("B"))
	if rb16 {
		t.Errorf("expected false but %+v\n", rb16)
	}

	g16i := FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16i := g16i.Path(g16i.FindVertexByID("I"), g16i.FindVertexByID("J"))
	if !rb16i {
		t.Errorf("expected true but %+v\n", rb16i)
	}

	g16j := FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16j := g16j.Path(g16j.FindVertexByID("J"), g16j.FindVertexByID("I"))
	if rb16j {
		t.Errorf("expected false but %+v\n", rb16j)
	}

	g16d := FromJSON("../../files/testgraph.json", "testgraph.016")
	rb16d := g16d.Path(g16d.FindVertexByID("D"), g16d.FindVertexByID("E"))
	if !rb16d {
		t.Errorf("expected true but %+v\n", rb16d)
	}
}
