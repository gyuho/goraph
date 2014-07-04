package fdfk

import (
	"testing"

	"github.com/gyuho/goraph/graph/gt"
)

func Test_JSON_MaxFlow(test *testing.T) {
	g := gt.JSONGraph("../../../example_files/testgraph.json", "testgraph.017")
	mf := MaxFlow(g, "S", "T")
	if mf != 28 {
		test.Errorf("Should return 28 but %+v", mf)
	}
}
