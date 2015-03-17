package graph

import (
	"os"
	"testing"
)

func TestPrim14(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_14")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	rmap := data.Prim()
	var total float32
	for k := range rmap {
		total += k.Weight
	}
	if total != 37.0 {
		t.Errorf("Expected total 37.0 but %.2f", total)
	}
}
