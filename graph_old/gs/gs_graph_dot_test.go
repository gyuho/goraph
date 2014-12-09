package gs

import "testing"

func TestFromDOT(t *testing.T) {
	g := FromDOT("../../files/testgraph.003.dot")
	vs := g.GetVerticesSize()
	es := g.GetEdgesSize()

	if vs != 8 {
		t.Error("expected 8 but %v", vs)
	}

	if es != 30 {
		t.Error("expected 31 but %v", es)
	}
}
