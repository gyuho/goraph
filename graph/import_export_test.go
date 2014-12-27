package graph

import "testing"

func TestFromJSON(t *testing.T) {
	rmap, err := FromJSON("../testdata/test_graph.json")
	if err != nil {
		t.Logf("%+v", rmap["graph_001"])
		t.Errorf("Error: %+v", err)
	}
}

func TestToJSON(t *testing.T) {

}

func TestFromDOT(t *testing.T) {

}

func TestToDOT(t *testing.T) {

}
