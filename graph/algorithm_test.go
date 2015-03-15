package graph

import (
	"os"
	"testing"
)

func TestBfs(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs2 := data2.Bfs(data2.GetNodeByID("S"))
	if len(rs2) != data2.GetNodeSize() {
		t.Errorf("Not traversed all: %s", data2)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs2 {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data2.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data2.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestDfsStack(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs2 := data2.DfsStack(data2.GetNodeByID("S"))
	if len(rs2) != data2.GetNodeSize() {
		t.Errorf("Not traversed all: %s", data2)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs2 {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data2.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data2.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestDfs(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs2 := []*Node{}
	data2.Dfs(data2.GetNodeByID("S"), &rs2)
	if len(rs2) != data2.GetNodeSize() {
		t.Errorf("Not traversed all: %s", data2)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs2 {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data2.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data2.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestTopologicalDagGraph2(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data2, err := FromJSON(file, "test_graph_02")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs2, isDag2 := data2.TopologicalDag()
	if rs2 != nil || isDag2 {
		t.Errorf("test_graph_02 has a cycle (not a DAG). Expected nil and false but %+v %v", rs2, isDag2)
	}
}

func TestTopologicalDagGraph6(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data6, err := FromJSON(file, "test_graph_06")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs6, isDag6 := data6.TopologicalDag()
	if rs6 == nil || !isDag6 {
		t.Errorf("test_graph_06 has no cycle (DAG). Not expected nil or false but %+v %v", rs6, isDag6)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs6 {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data6.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data6.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestTopologicalDagGraph7(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data7, err := FromJSON(file, "test_graph_07")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs7, isDag7 := data7.TopologicalDag()
	if rs7 == nil || !isDag7 {
		t.Errorf("test_graph_07 has no cycle (DAG). Not expected nil or false but %+v %v", rs7, isDag7)
	}
	traversedNodeID := make(map[string]bool)
	for _, nd := range rs7 {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
		if _, ok := traversedNodeID[nd.ID]; !ok {
			traversedNodeID[nd.ID] = true
		}
	}
	if len(traversedNodeID) != data7.GetNodeSize() {
		t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
	}
	for nd := range data7.NodeMap {
		if nd.Color == "white" {
			t.Errorf("%v", nd)
		}
	}
}

func TestTopologicalDagGraph8(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data8, err := FromJSON(file, "test_graph_08")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs8, isDag8 := data8.TopologicalDag()
	if rs8 != nil || isDag8 {
		t.Errorf("test_graph_08 has a cycle (not a DAG). Expected nil and false but %+v %v", rs8, isDag8)
	}
}

func TestTopologicalDagGraph9(t *testing.T) {
	file, err := os.Open("../testgraph/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()

	data9, err := FromJSON(file, "test_graph_09")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	rs9, isDag9 := data9.TopologicalDag()
	if rs9 != nil || isDag9 {
		t.Errorf("test_graph_09 has a cycle (not a DAG). Expected nil and false but %+v %v", rs9, isDag9)
	}
}
