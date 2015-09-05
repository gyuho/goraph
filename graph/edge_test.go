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
		g, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		if len(g.GetEdges()) != graph.TotalEdgeCount {
			t.Logf("%+v", g)
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalEdgeCount, len(g.GetEdges()))
		}
		for _, elem := range graph.EdgeToWeight {
			weight1 := g.GetEdgeWeight(g.GetNodeByID(elem.Nodes[0]), g.GetNodeByID(elem.Nodes[1]))
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}
