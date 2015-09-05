package graph

import (
	"os"
	"strings"
	"testing"
)

func TestBellmanFord12(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_12")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance, noNegCycle := g.BellmanFord(g.GetNodeByID("S"), g.GetNodeByID("T"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, v.ID)
	}
	if strings.Join(ts, "-") != "S-A-C-B-T" {
		t.Errorf("Expected the shortest path S-A-C-B-T but %s", strings.Join(ts, "-"))
	}
	if mapToDistance[g.GetNodeByID("T")] != -2.0 {
		t.Errorf("Expected the shortest distance of -2.0 (from S to T) but %.2f", mapToDistance[g.GetNodeByID("T")])
	}
	if !noNegCycle {
		t.Errorf("There should be no negative-weight cycle but found one with %v", noNegCycle)
	}
}

func TestBellmanFord13(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_13")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance, noNegCycle := g.BellmanFord(g.GetNodeByID("S"), g.GetNodeByID("T"))
	if pathSlice != nil || mapToDistance != nil {
		t.Errorf("Expected nil, nil but %v, %v", pathSlice, mapToDistance)
	}
	if noNegCycle {
		t.Errorf("There should be a negative-weight cycle but found one with %v", noNegCycle)
	}
}
