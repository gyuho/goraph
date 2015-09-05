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
	g, err := FromJSON(file, "test_graph_04")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := g.Dijkstra(g.GetNodeByID("S"), g.GetNodeByID("T"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00)" {
		t.Errorf("Expected the shortest path S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[g.GetNodeByID("T")] != 44.0 {
		t.Errorf("Expected 44.0 but %f", mapToDistance[g.GetNodeByID("T")])
	}
}

func TestDijkstra05(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_05")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := g.Dijkstra(g.GetNodeByID("A"), g.GetNodeByID("E"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "A(0.00) → C(9.00) → F(11.00) → E(20.00)" {
		t.Errorf("Expected the shortest path A(0.00) → C(9.00) → F(11.00) → E(20.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[g.GetNodeByID("E")] != 20.0 {
		t.Errorf("Expected 20.0 but %f", mapToDistance[g.GetNodeByID("E")])
	}
}

func TestDijkstra10(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_10")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := g.Dijkstra(g.GetNodeByID("A"), g.GetNodeByID("E"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "A(0.00) → C(9.00) → B(19.00) → D(34.00) → E(36.00)" {
		t.Errorf("Expected the shortest path A(0.00) → C(9.00) → B(19.00) → D(34.00) → E(36.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[g.GetNodeByID("E")] != 36.0 {
		t.Errorf("Expected 36.0 but %f", mapToDistance[g.GetNodeByID("E")])
	}
}

func TestDijkstra10Reverse(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_10")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := g.Dijkstra(g.GetNodeByID("E"), g.GetNodeByID("A"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "E(0.00) → F(9.00) → C(11.00) → B(21.00) → A(22.00)" {
		t.Errorf("Expected the shortest path E(0.00) → F(9.00) → C(11.00) → B(21.00) → A(22.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[g.GetNodeByID("A")] != 22.0 {
		t.Errorf("Expected 22.0 but %f", mapToDistance[g.GetNodeByID("A")])
	}
}

func TestDijkstra11(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_11")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := g.Dijkstra(g.GetNodeByID("S"), g.GetNodeByID("T"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → A(11.00) → B(16.00) → D(46.00) → E(49.00) → T(68.00)" {
		t.Errorf("Expected the shortest path S(0.00) → A(11.00) → B(16.00) → D(46.00) → E(49.00) → T(68.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[g.GetNodeByID("T")] != 68.0 {
		t.Errorf("Expected 68.0 but %f", mapToDistance[g.GetNodeByID("T")])
	}
}

func TestDijkstra11Reverse(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	g, err := FromJSON(file, "test_graph_11")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	pathSlice, mapToDistance := g.Dijkstra(g.GetNodeByID("T"), g.GetNodeByID("S"))
	ts := []string{}
	for _, v := range pathSlice {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v.ID, mapToDistance[v]))
	}
	if strings.Join(ts, " → ") != "T(0.00) → D(10.00) → E(13.00) → B(31.00) → S(48.00)" {
		t.Errorf("Expected the shortest path T(0.00) → D(10.00) → E(13.00) → B(31.00) → S(48.00) but %s", strings.Join(ts, " → "))
	}
	if mapToDistance[g.GetNodeByID("S")] != 48.0 {
		t.Errorf("Expected 48.0 but %f", mapToDistance[g.GetNodeByID("S")])
	}
}
