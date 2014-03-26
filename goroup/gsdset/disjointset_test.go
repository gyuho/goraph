package gsdset

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_makeSet(test *testing.T) {
	a := gsd.NewVertex("Google")
	s := makeSet(a)
	if len(s) != 1 || s[a] != 1 {
		test.Errorf("makeSet should have one member: %+v", s)
	}
}

func Test_MakeSet(test *testing.T) {
	a := gsd.NewVertex("Google")
	s := MakeSet(a)
	if reflect.TypeOf(s) != reflect.TypeOf(&DisJointSet{}) {
		test.Errorf("Should be same but: %+v", s)
	}
}

func Test_MakeGraphSet(test *testing.T) {
	g := gsd.JSONGraph("../../testgraph/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	// fmt.Printf("%+v", s[0])
	// DisJoint Set (Rep: S) / Vertex Set: [ S ]
	if len(s) != 8 {
		test.Errorf("Should return 8 but: %+v", s)
	}
}

func Test_UnionDisJointSet(test *testing.T) {
	g := gsd.JSONGraph("../../testgraph/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	us := UnionDisJointSet(s[0], s[1])
	result := fmt.Sprintf("%+v", us)
	rc := "DisJoint Set (Rep: S) / Vertex Set: [ S A ]"
	if result != rc {
		test.Errorf("Should return\n%+v\nbut\n%+v", rc, result)
	}
}

func Test_FindSet(test *testing.T) {
	g := gsd.JSONGraph("../../testgraph/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	vertices := g.GetVertices()
	ds := FindSet((*vertices)[0].(*gsd.Vertex), s)
	result := fmt.Sprintf("%+v", ds)
	rc := "DisJoint Set (Rep: S) / Vertex Set: [ S ]"
	if result != rc {
		test.Errorf("Should return\n%+v\nbut\n%+v", rc, result)
	}
}
