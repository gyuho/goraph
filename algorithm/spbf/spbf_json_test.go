package spbf

import (
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SPBF(test *testing.T) {
	g12 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.012")
	s12, ok12 := SPBF(g12, "S", "T")
	s12c := "S(=0) → A(=7) → C(=4) → B(=2) → T(=-2)"
	ok12c := true
	if s12 != s12c || ok12 != ok12c {
		test.Errorf("Should be same but %v", s12)
	}

	g13 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.013")
	s13, ok13 := SPBF(g13, "S", "T")
	s13c := "There is negative weighted cycle (No Shortest Path)"
	ok13c := false
	if s13 != s13c || ok13 != ok13c {
		test.Errorf("Should be same but %v", s13)
	}
}
