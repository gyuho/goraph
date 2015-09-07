package graph

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDefaultGraph_Dijkstra_03(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_03")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, "S", "T")
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00)" {
		t.Errorf("Expected the shortest path S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00) but %s", strings.Join(ts, " → "))
	}
	if distance["T"] != 44.0 {
		t.Errorf("Expected 44.0 but %f", distance["T"])
	}
	fmt.Println("graph_03:", strings.Join(ts, " → "))
}
