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
		g, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		if g.GetNodeSize() != graph.TotalNodeCount {
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalNodeCount, g.GetNodeSize())
		}
		for _, elem := range graph.EdgeToWeight {
			weight1 := g.GetEdgeWeight(g.GetNodeByID(elem.Nodes[0]), g.GetNodeByID(elem.Nodes[1]))
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("%s | Expected %f but %f", graph.Name, weight2, weight1)
			}
		}
		gClone := g.Clone()
		if gClone.GetNodeSize() != graph.TotalNodeCount {
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalNodeCount, gClone.GetNodeSize())
		}
		for _, elem := range graph.EdgeToWeight {
			weight1 := gClone.GetEdgeWeight(gClone.GetNodeByID(elem.Nodes[0]), gClone.GetNodeByID(elem.Nodes[1]))
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}
