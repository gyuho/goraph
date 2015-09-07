package graph

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDefaultGraph_BellmanFord_11(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_11")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := BellmanFord(g, "S", "T")
	if err != nil {
		t.Errorf("There should be no negative-weight cycle but found one with %v", err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → A(7.00) → C(4.00) → B(2.00) → T(-2.00)" {
		t.Errorf("Expected the shortest path S(0.00) → A(7.00) → C(4.00) → B(2.00) → T(-2.00) but %s", strings.Join(ts, " → "))
	}
	if distance["T"] != -2.0 {
		t.Errorf("Expected -2.0 but %f", distance["T"])
	}
	fmt.Println("graph_11:", strings.Join(ts, " → "))
}
