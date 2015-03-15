package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestFromJSON(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	// jsonStream, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return nil, err
	// }
	rmap, err := fromJSON(file)
	if err != nil {
		t.Logf("%+v", rmap["test_graph_01"])
		t.Errorf("Error: %+v", err)
	}
	if rmap["test_graph_02"]["A"]["S"] != 15 || rmap["test_graph_02"]["S"]["A"] != 100 {
		t.Errorf("%+v", rmap)
	}
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
				t.Errorf("Expected %f but %f", weight2, weight1)
			}
		}
	}
}
