package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testgraph"
)

func TestfromJSON(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
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
}

func TestFromJSON01(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_01")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph01 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON02(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph02 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON03(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_03")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph03 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON04(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_04")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph04 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON05(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_05")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 6 {
		t.Errorf("Expected 6 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph05 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON06(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_06")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 6 {
		t.Errorf("Expected 6 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph06 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON07(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_07")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph07 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON08(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_08")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 6 {
		t.Errorf("Expected 6 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph08 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON09(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_09")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph09 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON10(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_10")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 6 {
		t.Errorf("Expected 6 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph10 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON11(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_11")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph11 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON12(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_12")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 5 {
		t.Errorf("Expected 5 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph12 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON13(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_13")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 5 {
		t.Errorf("Expected 5 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph13 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON14(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_14")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 9 {
		t.Errorf("Expected 9 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph14 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON15(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_15")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph15 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON16(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_16")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 10 {
		t.Errorf("Expected 10 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph16 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestFromJSON17(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_17")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != 8 {
		t.Errorf("Expected 8 but %v", data.GetNodeSize())
	}
	for _, elem := range testgraph.Graph17 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Vertices[0]), data.GetNodeByID(elem.Vertices[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}
