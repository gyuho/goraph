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
	if len(s) != 1 || !gsd.SameVertex(s[a], a) {
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
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	// fmt.Printf("%+v", s[0])
	// DisJoint Set (Rep: S) / Vertex Set: [ S ]
	if len(s) != 8 {
		test.Errorf("Should return 8 but: %+v", s)
	}
}

func Test_SetContains(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	vertices := g.GetVertices()
	if !s[0].SetContains((*vertices)[0].(*gsd.Vertex)) {
		test.Errorf("Should return\n%+v", s[0].SetContains((*vertices)[0].(*gsd.Vertex)))
	}
}

func Test_GetSet(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	vertices := g.GetVertices()
	ds := GetSet((*vertices)[0].(*gsd.Vertex), s)
	result := fmt.Sprintf("%+v", ds)
	rc := "DisJoint Set (Rep: S) / Vertex Set: [ S ]"
	if result != rc {
		test.Errorf("Should return\n%+v\nbut\n%+v", rc, result)
	}
}

func Test_FindSet(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	vertices := g.GetVertices()
	fs := FindSet((*vertices)[0].(*gsd.Vertex), s)
	result := fmt.Sprintf("%+v", fs.ID)
	rc := "S"
	if result != rc {
		test.Errorf("Should return\n%+v\nbut\n%+v", rc, result)
	}
}

func Test_UnionDisJointSet(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	us := UnionDisJointSet(s[0], s[1])
	result := fmt.Sprintf("%+v", us)
	rc := "DisJoint Set (Rep: S) / Vertex Set: [ S A ]"
	if result != rc {
		test.Errorf("Should return\n%+v\nbut\n%+v", rc, result)
	}
}

func Test_UnionByRep(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	vertices := g.GetVertices()
	fs := UnionByRep((*vertices)[0].(*gsd.Vertex), (*vertices)[1].(*gsd.Vertex), &s)
	// fmt.Println((*vertices)[0].(*gsd.Vertex).ID) // S
	// fmt.Println((*vertices)[1].(*gsd.Vertex).ID) // A
	// fmt.Printf("%q", *fs)
	// ["DisJoint Set (Rep: S) / Vertex Set: [ S A ]" "DisJoint Set (Rep: B) / Vertex Set: [ B ]" "DisJoint Set (Rep: C) / Vertex Set: [ C ]" "DisJoint Set (Rep: D) / Vertex Set: [ D ]" "DisJoint Set (Rep: T) / Vertex Set: [ T ]" "DisJoint Set (Rep: E) / Vertex Set: [ E ]" "DisJoint Set (Rep: F) / Vertex Set: [ F ]"]

	if len(*fs) != 7 {
		test.Errorf("Should return 7 but\n%+v", *fs)
	}
}

func Test_UnionByVtx(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	s := MakeGraphSet(g)
	vertices := g.GetVertices()
	fs := UnionByVtx((*vertices)[0].(*gsd.Vertex), (*vertices)[1].(*gsd.Vertex), &s)
	// fmt.Println((*vertices)[0].(*gsd.Vertex).ID) // S
	// fmt.Println((*vertices)[1].(*gsd.Vertex).ID) // A
	// fmt.Printf("%q", *fs)
	// ["DisJoint Set (Rep: S) / Vertex Set: [ S A ]" "DisJoint Set (Rep: B) / Vertex Set: [ B ]" "DisJoint Set (Rep: C) / Vertex Set: [ C ]" "DisJoint Set (Rep: D) / Vertex Set: [ D ]" "DisJoint Set (Rep: T) / Vertex Set: [ T ]" "DisJoint Set (Rep: E) / Vertex Set: [ E ]" "DisJoint Set (Rep: F) / Vertex Set: [ F ]"]

	if len(*fs) != 7 {
		test.Errorf("Should return 7 but\n%+v", *fs)
	}
}

func Test_SortEdges(test *testing.T) {
	g := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.003")
	edges := g.GetEdges()
	fmt.Println("Before Sorting:")
	for _, edge := range *edges {
		fmt.Printf("%v ", edge.(*gsd.Edge).Weight)
	}
	// 100 6 14 200 15 5 20 44 14 5 30 18 9 24 20 30 2 11 16 18 24 2 6 19 11 6 6 44 16 6 19
	println()

	SortEdges(g)
	fmt.Println("After Sorting:")
	for _, edge := range *edges {
		fmt.Printf("%v ", edge.(*gsd.Edge).Weight)
	}
	// 2 2 5 5 6 6 6 6 6 9 11 11 14 14 15 16 16 18 18 19 19 20 20 24 24 30 30 44 44 100 200
}
