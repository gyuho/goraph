package graph

import "testing"

func TestfromJSON(t *testing.T) {
	rmap, err := fromJSON("../testdata/test_graph.json")
	if err != nil {
		t.Logf("%+v", rmap["graph_001"])
		t.Errorf("Error: %+v", err)
	}
	if rmap["graph_002"]["A"]["S"] != 15 || rmap["graph_002"]["S"]["A"] != 100 {
		t.Errorf("%+v", rmap)
	}
}

func TestToJSON(t *testing.T) {

}

func TestFromDOT(t *testing.T) {

}

func TestToDOT(t *testing.T) {

}
