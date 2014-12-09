package gset

// Order changes everytime we run

import (
	"reflect"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func Test_NewSet(t *testing.T) {
	s := NewSet()
	if reflect.TypeOf(s) != reflect.TypeOf(Set{}) {
		t.Errorf("NewSet() should return Set type: %#v", s)
	}
}

func Test_Size(t *testing.T) {
	s := NewSet()
	if s.Size() != 0 {
		t.Errorf("NewSet() should return Set of size 0: %#v", s)
	}
}

func Test_IsEmpty(t *testing.T) {
	s := NewSet()
	if s.IsEmpty() != true {
		t.Errorf("IsEmpty() should return true: %#v", s)
	}
}

func Test_Insert(t *testing.T) {
	s := NewSet()
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	s.Insert(a, b)
	if s.IsEmpty() != false {
		t.Errorf("IsEmpty() should return false: %#v", s)
	}
	if s.Size() != 2 {
		t.Errorf("Size() should return 2: %#v", s)
	}

	value, exist := s[a]
	if !gs.SameVertex(value, a) {
		t.Errorf("s[a]'s value should be 1: %#v", value)
	}
	if exist != true {
		t.Errorf("s[a] should exist: %#v", value)
	}
}

func Test_InstantiateSet(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	s := InstantiateSet(a, b)
	if s.IsEmpty() != false {
		t.Errorf("IsEmpty() should return false: %#v", s)
	}
	if s.Size() != 2 {
		t.Errorf("Size() should return 5: %#v", s)
	}
	value, exist := s[a]
	if !gs.SameVertex(value, a) {
		t.Errorf("Should be same: %#v", s)
	}
	if exist != true {
		t.Errorf("s[2] should exist: %#v", s)
	}
}

/*
func Test_GetElements(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	s := InstantiateSet(a, b)
	sl := s.GetElements()
	if !gs.SameVertex(a, sl[0]) {
		t.Errorf("sl[0] should be gs.NewVertex(\"Google\"): %+v", sl[0])
	}
	if !gs.SameVertex(b, sl[1]) {
		t.Errorf("sl[1] should be gs.NewVertex(\"Apple\"): %+v", sl[1])
	}
}
*/

func Test_Contains(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	s := InstantiateSet(a, b)
	result := s.Contains(a)
	if result != true {
		t.Errorf("s.Contains(a) should return true: %#v", s)
	}
}

func Test_Delete(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	s := InstantiateSet(a, b)
	result1 := s.Delete(b)
	if result1 != true {
		t.Errorf("s.Delete(b) should return true: %#v", s)
	}
	if s.Size() != 1 {
		t.Errorf("s.Size() should return 4: %#v", s)
	}

	c := gs.NewVertex("Korea")
	result2 := s.Delete(c)
	if result2 != false {
		t.Errorf("s.Delete(100) should return false: %#v", s)
	}
}

func Test_Intersection(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")

	set1 := InstantiateSet(a, b, c)
	set2 := InstantiateSet(c, d, e)

	result := set1.Intersection(set2)
	if len(result) != 1 {
		t.Errorf("len(result) should return 1: %#v", len(result))
	}
	if !gs.SameVertex(result[0], c) {
		t.Errorf("Should return true: %+v", gs.SameVertex(result[0], c))
	}
}

func Test_Union(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")

	set1 := InstantiateSet(a, b, c)
	set2 := InstantiateSet(b, c, d, e)
	result := set1.Union(set2)
	if len(result) != 5 {
		t.Errorf("len(result) should return 5: %+v", result)
	}
	//if !gs.SameVertex(result[0], a) {
	//	t.Errorf("result[0] should return a: %+v", result)
	//}
	//if !gs.SameVertex(result[4], e) {
	//	t.Errorf("result[4] should return e: %+v", result)
	//}
}

func Test_Subtract(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")

	set1 := InstantiateSet(a, b, c, d, e)
	set2 := InstantiateSet(a, b)
	result := set1.Subtract(set2)
	if len(result) != 3 {
		t.Errorf("len(result) should return 3: %#v", len(result))
	}
	//if !gs.SameVertex(result[0], c) {
	//	t.Errorf("Should return true: %+v", gs.SameVertex(result[0], c))
	//}
	//if !gs.SameVertex(result[1], d) {
	//	t.Errorf("Should return true: %+v", gs.SameVertex(result[0], d))
	//}
}

func Test_IsEqual(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")

	set1 := InstantiateSet(a, b, c, d, e)
	set2 := InstantiateSet(a, b, c, d, e)
	set3 := InstantiateSet(a, e)

	if !set1.IsEqual(set2) {
		t.Errorf("Should return true: %+v", set1.IsEqual(set2))
	}

	if !set2.IsEqual(set1) {
		t.Errorf("Should return true: %+v", set2.IsEqual(set1))
	}

	if set1.IsEqual(set3) {
		t.Errorf("Should return false: %+v", set1.IsEqual(set3))
	}
}

func Test_Subset(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")

	set1 := InstantiateSet(a, b, c, d, e)
	set2 := InstantiateSet(b, d)

	if !set1.Subset(set2) {
		t.Errorf("Should be true: %+v", set1.Subset(set2))
	}

	if set2.Subset(set1) {
		t.Errorf("Should be false: %+v", set2.Subset(set1))
	}
}

func Test_Clone(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")

	set1 := InstantiateSet(a, b, c, d, e)
	set2 := set1.Clone()
	if !set1.IsEqual(set2) {
		t.Errorf("Should return true: %+v", set1.IsEqual(set2))
	}
}

/*
func Test_String(t *testing.T) {
	a := gs.NewVertex("Google")
	b := gs.NewVertex("Apple")
	c := gs.NewVertex("Annie")
	d := gs.NewVertex("Gyu-Ho")
	e := gs.NewVertex("Korea")
	set1 := InstantiateSet(a, b, c, d, e)
	result := set1.String()
	str := "Vertex Set: [ Google Apple Annie Gyu-Ho Korea ]"
	if result != str {
		t.Errorf("Does not match: %#v", result)
	}
	if result != fmt.Sprintf("%s", set1) {
		t.Errorf("Does not match: %#v", result)
	}
	if result != fmt.Sprintf("%v", set1) {
		t.Errorf("Does not match: %#v", result)
	}
}
*/
