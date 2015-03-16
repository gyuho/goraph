package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestDijkstra(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_04")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	mapToDistance, mapToPrev := data.Dijkstra(data.GetNodeByID("S"))
	fmt.Println(mapToDistance)
	fmt.Println(mapToPrev)
}
