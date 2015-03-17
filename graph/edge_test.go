package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestGetEdges(t *testing.T) {
	for _, graph := range testdata.GraphSlice {
		file, err := os.Open("../testdata/data.json")
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		defer file.Close()
		data, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		if len(data.GetEdges()) != graph.TotalEdgeCount {
			t.Logf("%+v", data)
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalEdgeCount, len(data.GetEdges()))
		}
		for _, elem := range graph.EdgeToWeight {
			weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}
