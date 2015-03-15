package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestClone(t *testing.T) {
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
		if data.GetNodeSize() != graph.TotalNodeCount {
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalNodeCount, data.GetNodeSize())
		}
		for _, elem := range graph.EdgeToWeight {
			weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("%s | Expected %f but %f", graph.Name, weight2, weight1)
			}
		}
		dataClone := data.Clone()
		if dataClone.GetNodeSize() != graph.TotalNodeCount {
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalNodeCount, dataClone.GetNodeSize())
		}
		for _, elem := range graph.EdgeToWeight {
			weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}
