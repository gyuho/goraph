package graph

import (
	"os"
	"testing"

	"github.com/gyuho/goraph/testdata"
)

func TestTopologicalDag(t *testing.T) {
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
		rs, isDag := data.TopologicalDag()
		if isDag != graph.IsDag {
			t.Errorf("%s | IsDag are supposed to be %v but %+v %+v", graph.Name, graph.IsDag, rs, isDag)
		}
	}
}
