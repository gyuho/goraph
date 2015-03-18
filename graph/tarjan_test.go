package graph

import (
	"os"
	"testing"
)

func TestTarjan15(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_15")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	scc := data.Tarjan()
	if len(scc) != 4 {
		t.Errorf("Expected 4 but %v", scc)
	}
	// for _, row := range scc {
	// 	for _, elem := range row {
	// 		fmt.Println(elem)
	// 	}
	// 	println()
	// }
	/*
	   [H / 1 Outgoing / 3 Incoming Edges]

	   [G / 2 Outgoing / 2 Incoming Edges]
	   [F / 1 Outgoing / 3 Incoming Edges]

	   [D / 2 Outgoing / 1 Incoming Edges]
	   [C / 2 Outgoing / 2 Incoming Edges]

	   [A / 1 Outgoing / 1 Incoming Edges]
	   [E / 2 Outgoing / 1 Incoming Edges]
	   [B / 3 Outgoing / 1 Incoming Edges]
	*/
}

func TestTarjan16(t *testing.T) {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	defer file.Close()
	data, err := FromJSON(file, "test_graph_16")
	if err != nil {
		t.Errorf("Error: %+v", err)
	}
	scc := data.Tarjan()
	if len(scc) != 4 {
		t.Errorf("Expected 4 but %v", scc)
	}
	// for _, row := range scc {
	// 	for _, elem := range row {
	// 		fmt.Println(elem)
	// 	}
	// 	println()
	// }
	/*
	   [J / 1 Outgoing / 3 Incoming Edges]
	   [E / 1 Outgoing / 2 Incoming Edges]

	   [I / 1 Outgoing / 1 Incoming Edges]

	   [H / 1 Outgoing / 3 Incoming Edges]
	   [D / 4 Outgoing / 1 Incoming Edges]
	   [C / 2 Outgoing / 2 Incoming Edges]

	   [F / 1 Outgoing / 1 Incoming Edges]
	   [B / 2 Outgoing / 1 Incoming Edges]
	   [A / 2 Outgoing / 1 Incoming Edges]
	   [G / 2 Outgoing / 2 Incoming Edges]
	*/
}
