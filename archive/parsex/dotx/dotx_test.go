package dotx

import (
	"testing"

	"github.com/gyuho/goraph/parsex"
)

func TestGetNodes(t *testing.T) {
	ns := GetNodes("../../files/testgraph.003.dot")
	sl := []string{"S", "A", "B", "C", "D", "E", "F", "T"}
	if !parsex.EqualSliceElem(ns, sl) {
		t.Errorf("expected true but: %v", ns)
	}
}

func TestGetGraphMap(t *testing.T) {
	gname, ns := GetGraphMap("../../files/testgraph.003.dot")
	sl := []string{"S", "A", "B", "C", "D", "E", "F", "T"}
	if gname != "testgraph003" {
		t.Errorf("expected testgraph003 but %v", gname)
	}
	if len(ns) != len(sl) {
		t.Errorf("expected 8 but: %v", ns)
	}
	if v, ok := ns["S"]["B"]; v[0] != 20 || !ok {
		t.Errorf("expected [20], true but %v %v", v, ok)
	}
}
