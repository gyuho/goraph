package spd

import (
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestSPD(t *testing.T) {
	g4 := gs.FromJSON("../../files/testgraph.json", "testgraph.004")
	g4s := SPD(g4, g4.FindVertexByID("S"), g4.FindVertexByID("T"))
	g4c := "S(=0) → B(=14) → E(=32) → F(=38) → T(=44)"
	if g4s != g4c {
		t.Errorf("Should be same but\n%v\n%v", g4s, g4c)
	}

	g5 := gs.FromJSON("../../files/testgraph.json", "testgraph.005")
	g5s := SPD(g5, g5.FindVertexByID("A"), g5.FindVertexByID("E"))
	g5c := "A(=0) → C(=9) → F(=11) → E(=20)"
	if g5s != g5c {
		t.Errorf("testgraph5 Should be same but\n%v\n%v\n%v", g5s, g5c, g5.ShowPrev(g5.FindVertexByID("E")))
	}

	g10 := gs.FromJSON("../../files/testgraph.json", "testgraph.010")
	g10s := SPD(g10, g10.FindVertexByID("A"), g10.FindVertexByID("E"))
	g10c := "A(=0) → C(=9) → B(=19) → D(=34) → E(=36)"
	if g10s != g10c {
		t.Errorf("Should be same but\n%v\n%v", g10s, g10c)
	}

	g10o := gs.FromJSON("../../files/testgraph.json", "testgraph.010")
	g10so := SPD(g10o, g10o.FindVertexByID("E"), g10o.FindVertexByID("A"))
	g10co := "E(=0) → F(=9) → C(=11) → B(=21) → A(=22)"
	if g10so != g10co {
		t.Errorf("Should be same but\n%v\n%v", g10so, g10co)
	}

	g11 := gs.FromJSON("../../files/testgraph.json", "testgraph.011")
	g11s := SPD(g11, g11.FindVertexByID("S"), g11.FindVertexByID("T"))
	g11c := "S(=0) → A(=11) → B(=16) → D(=46) → E(=49) → T(=68)"
	if g11s != g11c {
		t.Errorf("Should be same but\n%v\n%v", g11s, g11c)
	}

	g11o := gs.FromJSON("../../files/testgraph.json", "testgraph.011")
	g11so := SPD(g11o, g11o.FindVertexByID("T"), g11o.FindVertexByID("S"))
	g11co := "T(=0) → D(=10) → E(=13) → B(=31) → S(=48)"
	if g11so != g11co {
		t.Errorf("Should be same but\n%v\n%v", g11so, g11co)
	}
}
