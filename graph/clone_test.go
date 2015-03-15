package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestClone01(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_01")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph01[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph01[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph01 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph01[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph01[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph01 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone02(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph02[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph02[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph02 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph02[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph02[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph02 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone03(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_03")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph03[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph03[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph03 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph03[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph03[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph03 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone04(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_04")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph04[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph04[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph04 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph04[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph04[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph04 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone05(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_05")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph05[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph05[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph05 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph05[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph05[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph05 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone06(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_06")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph06[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph06[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph06 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph06[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph06[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph06 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone07(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_07")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph07[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph07[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph07 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph07[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph07[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph07 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone08(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_08")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph08[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph08[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph08 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph08[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph08[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph08 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone09(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_09")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph09[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph09[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph09 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph09[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph09[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph09 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone10(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_10")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph10[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph10[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph10 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph10[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph10[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph10 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone11(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_11")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph11[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph11[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph11 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph11[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph11[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph11 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone12(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_12")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph12[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph12[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph12 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph12[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph12[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph12 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone13(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_13")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph13[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph13[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph13 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph13[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph13[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph13 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone14(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_14")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph14[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph14[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph14 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph14[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph14[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph14 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone15(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_15")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph15[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph15[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph15 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph15[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph15[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph15 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone16(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_16")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph16[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph16[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph16 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph16[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph16[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph16 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}

func TestClone17(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_17")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	if data.GetNodeSize() != testdata.Graph17[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph17[0].TotalNodeCount, data.GetNodeSize())
	}
	for _, elem := range testdata.Graph17 {
		weight1 := data.GetEdgeWeight(data.GetNodeByID(elem.Nodes[0]), data.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}

	dataClone := data.Clone()
	if dataClone.GetNodeSize() != testdata.Graph17[0].TotalNodeCount {
		t.Errorf("Expected %d but %d", testdata.Graph17[0].TotalNodeCount, dataClone.GetNodeSize())
	}
	for _, elem := range testdata.Graph17 {
		weight1 := dataClone.GetEdgeWeight(dataClone.GetNodeByID(elem.Nodes[0]), dataClone.GetNodeByID(elem.Nodes[1]))
		weight2 := elem.Weight
		if weight1 != weight2 {
			t.Errorf("Expected %f but %f", weight2, weight1)
		}
	}
}
