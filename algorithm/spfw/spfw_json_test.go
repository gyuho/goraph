package spfw

import (
	"testing"

	"github.com/gyuho/goraph/graph/gt"
)

func Test_JSON_SPFW(test *testing.T) {
	g4 := gt.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	g4s, g4m := SPFW(g4, "S", "T")
	g4c := 44.0
	if g4s != g4c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g4s, g4c, g4m)
	}

	g5 := gt.JSONGraph("../../example_files/testgraph.json", "testgraph.005")
	g5s, g5m := SPFW(g5, "A", "E")
	g5c := 20.0
	if g5s != g5c {
		test.Errorf("testgraph5 Should be same but\n%v\n%v\n%v", g5s, g5c, g5m)
	}

	g10 := gt.JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10s, g10m := SPFW(g10, "A", "E")
	g10c := 36.0
	if g10s != g10c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10s, g10c, g10m)
	}

	g10o := gt.JSONGraph("../../example_files/testgraph.json", "testgraph.010")
	g10so, g10m := SPFW(g10o, "E", "A")
	g10co := 22.0
	if g10so != g10co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10so, g10co, g10m)
	}

	g11 := gt.JSONGraph("../../example_files/testgraph.json", "testgraph.011")
	g11s, g11m := SPFW(g11, "S", "T")
	g11c := 68.0
	if g11s != g11c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11s, g11c, g11m)
	}

	g11o := gt.JSONGraph("../../example_files/testgraph.json", "testgraph.011")
	g11so, g11om := SPFW(g11o, "T", "S")
	g11co := 48.0
	if g11so != g11co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11so, g11co, g11om)
	}
}
