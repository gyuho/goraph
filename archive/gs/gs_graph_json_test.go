package gs

import "testing"

func TestFromJSON(t *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.003")
	vs := g.GetVerticesSize()
	es := g.GetEdgesSize()

	if vs != 8 {
		t.Error("expected 8 but %v", vs)
	}

	if es != 30 {
		t.Error("expected 30 but %v", es)
	}
}
