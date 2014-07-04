package sp

import (
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SP(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	g4s, okg4s := SP(g4, "S", "T")
	g4c := "S(=0) → B(=14) → E(=32) → F(=38) → T(=44)"
	okg4c := true
	if g4s != g4c || okg4s != okg4c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g4s, g4c, okg4s)
	}

	g5 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.005")
	g5s, okg5s := SP(g5, "A", "E")
	g5c := "A(=0) → C(=9) → F(=11) → E(=20)"
	okg5c := true
	if g5s != g5c || okg5s != okg5c {
		test.Errorf("testgraph5 Should be same but\n%v\n%v\n%v", g5s, g5c, okg5s)
	}

	g10 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10s, okg10s := SP(g10, "A", "E")
	g10c := "A(=0) → C(=9) → B(=19) → D(=34) → E(=36)"
	okg10c := true
	if g10s != g10c || okg10s != okg10c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10s, g10c, okg10s)
	}

	g10o := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10so, okg10so := SP(g10o, "E", "A")
	g10co := "E(=0) → F(=9) → C(=11) → B(=21) → A(=22)"
	okg10co := true
	if g10so != g10co || okg10so != okg10co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10so, g10co, okg10so)
	}

	g11 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.011")
	g11s, okg11s := SP(g11, "S", "T")
	g11c := "S(=0) → A(=11) → B(=16) → D(=46) → E(=49) → T(=68)"
	okg11c := true
	if g11s != g11c || okg11s != okg11c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11s, g11c, okg11s)
	}

	g11o := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.011")
	g11so, okg11so := SP(g11o, "T", "S")
	g11co := "T(=0) → D(=10) → E(=13) → B(=31) → S(=48)"
	okg11co := true
	if g11so != g11co || okg11so != okg11co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11so, g11co, okg11so)
	}

	g12 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.012")
	s12, ok12 := SP(g12, "S", "T")
	s12c := "S(=0) → A(=7) → C(=4) → B(=2) → T(=-2)"
	ok12c := true
	if s12 != s12c || ok12 != ok12c {
		test.Errorf("Should be same but %v", s12)
	}

	g13 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.013")
	s13, ok13 := SP(g13, "S", "T")
	s13c := "There is negative weighted cycle (No Shortest Path)"
	ok13c := false
	if s13 != s13c || ok13 != ok13c {
		test.Errorf("Should be same but %v", s13)
	}
}
