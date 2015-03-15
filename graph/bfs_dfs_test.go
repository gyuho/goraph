package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestBfs(t *testing.T) {
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
		oneNode := &Node{}
		for elem := range data.NodeMap {
			oneNode = elem
			break
		}
		rs := data.Bfs(oneNode)
		if len(rs) != data.GetNodeSize() {
			t.Errorf("Not traversed all: %s", data)
		}
		traversedNodeID := make(map[string]bool)
		for _, nd := range rs {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
			if _, ok := traversedNodeID[nd.ID]; !ok {
				traversedNodeID[nd.ID] = true
			}
		}
		if len(traversedNodeID) != data.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range data.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}

func TestDfsStack(t *testing.T) {
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
		oneNode := &Node{}
		for elem := range data.NodeMap {
			oneNode = elem
			break
		}
		rs := data.DfsStack(oneNode)
		if len(rs) != data.GetNodeSize() {
			t.Errorf("Not traversed all: %s", data)
		}
		traversedNodeID := make(map[string]bool)
		for _, nd := range rs {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
			if _, ok := traversedNodeID[nd.ID]; !ok {
				traversedNodeID[nd.ID] = true
			}
		}
		if len(traversedNodeID) != data.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range data.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}

func TestDfs(t *testing.T) {
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
		oneNode := &Node{}
		for elem := range data.NodeMap {
			oneNode = elem
			break
		}
		rs := []*Node{}
		data.Dfs(oneNode, &rs)
		if len(rs) != data.GetNodeSize() {
			t.Errorf("Not traversed all: %s", data)
		}
		traversedNodeID := make(map[string]bool)
		for _, nd := range rs {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
			if _, ok := traversedNodeID[nd.ID]; !ok {
				traversedNodeID[nd.ID] = true
			}
		}
		if len(traversedNodeID) != data.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range data.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}
