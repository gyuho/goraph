package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/graph/testdata"
)

func TestDefaultGraph_TopologicalSort(t *testing.T) {
	for _, graph := range testdata.GraphSlice {
		f, err := os.Open("testdata/graph.json")
		if err != nil {
			t.Error(err)
		}
		defer f.Close()
		g, err := NewDefaultGraphFromJSON(f, graph.Name)
		if err != nil {
			t.Error(err)
		}
		L, isDAG := TopologicalSort(g)
		if isDAG != graph.IsDAG {
			t.Errorf("%s | IsDag are supposed to be %v but %+v %+v", graph.Name, graph.IsDAG, L, isDAG)
		}
	}
}
