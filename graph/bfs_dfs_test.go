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
		g, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		oneNode := &Node{}
		for elem := range g.NodeMap {
			oneNode = elem
			break
		}
		rs := g.Bfs(oneNode)
		if len(rs) != g.GetNodeSize() {
			t.Errorf("Not traversed all: %s", g)
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
		if len(traversedNodeID) != g.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range g.NodeMap {
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
		g, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		oneNode := &Node{}
		for elem := range g.NodeMap {
			oneNode = elem
			break
		}
		rs := g.DfsStack(oneNode)
		if len(rs) != g.GetNodeSize() {
			t.Errorf("Not traversed all: %s", g)
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
		if len(traversedNodeID) != g.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range g.NodeMap {
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
		g, err := FromJSON(file, graph.Name)
		if err != nil {
			t.Errorf("Error: %+v", err)
		}
		oneNode := &Node{}
		for elem := range g.NodeMap {
			oneNode = elem
			break
		}
		rs := []*Node{}
		g.Dfs(oneNode, &rs)
		if len(rs) != g.GetNodeSize() {
			t.Errorf("Not traversed all: %s", g)
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
		if len(traversedNodeID) != g.GetNodeSize() {
			t.Errorf("Found duplicate Node ID: %+v", traversedNodeID)
		}
		for nd := range g.NodeMap {
			if nd.Color == "white" {
				t.Errorf("%v", nd)
			}
		}
	}
}
