package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/graph/testdata"
)

func TestNewDefaultGraphFromJSON(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(g.String())
	if g.VertexToChildren["C"]["S"] != 9.0 {
		t.Errorf("weight from C to S must be 9.0 but %f", g.VertexToChildren["C"]["S"])
	}
	for _, graph := range testdata.GraphSlice {
		f, err := os.Open("testdata/graph.json")
		if err != nil {
			t.Error(err)
		}
		g, err := NewDefaultGraphFromJSON(f, graph.Name)
		if err != nil {
			t.Error(err)
		}
		if len(g.Vertices) != graph.TotalNodeCount {
			t.Errorf("%s | Expected %d but %d", graph.Name, graph.TotalNodeCount, len(g.Vertices))
		}
		for _, elem := range graph.EdgeToWeight {
			weight1, err := g.GetWeight(elem.Nodes[0], elem.Nodes[1])
			if err != nil {
				t.Error(err)
			}
			weight2 := elem.Weight
			if weight1 != weight2 {
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
		f.Close()
	}
}
