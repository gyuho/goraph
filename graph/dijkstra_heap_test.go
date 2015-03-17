package graph

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDijkstra04(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_04")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := data.Dijkstra(data.GetNodeByID("S"), data.GetNodeByID("T"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00)" {
		t.Errorf("Expected the shortest path S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[data.GetNodeByID("T")] != 44.0 {
		t.Errorf("Expected 44.0 but %f", mapToDistance[data.GetNodeByID("T")])
	}
}

func TestDijkstra05(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_05")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := data.Dijkstra(data.GetNodeByID("A"), data.GetNodeByID("E"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "A(0.00) → C(9.00) → F(11.00) → E(20.00)" {
		t.Errorf("Expected the shortest path A(0.00) → C(9.00) → F(11.00) → E(20.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[data.GetNodeByID("E")] != 20.0 {
		t.Errorf("Expected 20.0 but %f", mapToDistance[data.GetNodeByID("T")])
	}
}
