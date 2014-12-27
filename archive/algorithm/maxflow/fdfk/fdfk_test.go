package fdfk

import (
	"testing"

	"github.com/gyuho/goraph/graph/gt"
)

func TestMaxFlow(t *testing.T) {
	g := gt.FromJSON("../../../files/testgraph.json", "testgraph.017")
	mf := MaxFlow(g, "S", "T")
	if mf != 28 {
		t.Errorf("expected 28 but %+v", mf)
	}
}
