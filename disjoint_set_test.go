package goraph

import "testing"

func TestDisjointSet(t *testing.T) {
	forests := NewForests()
	for _, name := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"} {
		MakeDisjointSet(forests, name)
	}
	for _, u := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"} {
		if FindSet(forests, u) == nil {
			t.Errorf("%s must be found", u)
		}
		if FindSet(forests, u).represent != u {
			t.Errorf("%s's represent must be %s but %s", u, u, FindSet(forests, u).represent)
		}
	}
	Union(forests, FindSet(forests, "A"), FindSet(forests, "B"))
	if FindSet(forests, "A").represent != "A" {
		t.Errorf("A's represent must be A but %s", FindSet(forests, "A").represent)
	}
	if FindSet(forests, "B").represent != "A" {
		t.Errorf("B's represent must be A but %s", FindSet(forests, "B").represent)
	}
}
