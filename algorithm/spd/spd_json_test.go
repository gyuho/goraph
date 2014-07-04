package spd

import (
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SPD(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	g4s := SPD(g4, "S", "T")
	g4c := "S(=0) → B(=14) → E(=32) → F(=38) → T(=44)"
	if g4s != g4c {
		test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
	}

	g5 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.005")
	g5s := SPD(g5, "A", "E")
	g5c := "A(=0) → C(=9) → F(=11) → E(=20)"
	if g5s != g5c {
		test.Errorf("testgraph5 Should be same but\n%v\n%v\n%v", g5s, g5c, g5.ShowPrev("E"))
	}

	g10 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10s := SPD(g10, "A", "E")
	g10c := "A(=0) → C(=9) → B(=19) → D(=34) → E(=36)"
	if g10s != g10c {
		test.Errorf("Should be same but\n%v\n%v", g10s, g10c)
	}

	g10o := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10so := SPD(g10o, "E", "A")
	g10co := "E(=0) → F(=9) → C(=11) → B(=21) → A(=22)"
	if g10so != g10co {
		test.Errorf("Should be same but\n%v\n%v", g10so, g10co)
	}

	g11 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.011")
	g11s := SPD(g11, "S", "T")
	g11c := "S(=0) → A(=11) → B(=16) → D(=46) → E(=49) → T(=68)"
	if g11s != g11c {
		test.Errorf("Should be same but\n%v\n%v", g11s, g11c)
	}

	g11o := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.011")
	g11so := SPD(g11o, "T", "S")
	g11co := "T(=0) → D(=10) → E(=13) → B(=31) → S(=48)"
	if g11so != g11co {
		test.Errorf("Should be same but\n%v\n%v", g11so, g11co)
	}
}
