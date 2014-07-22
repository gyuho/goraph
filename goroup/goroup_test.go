package goroup

import (
	"reflect"
	"testing"
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
	s.Insert(1, 2, -.9, "A", 0, 2, 2, 2)
	if s.IsEmpty() != false {
		t.Errorf("IsEmpty() should return false: %#v", s)
	}
	if s.Size() != 5 {
		t.Errorf("Size() should return 5: %#v", s)
	}

	value, exist := s[2]
	if value != 4 {
		t.Errorf("s[2]'s value should be 4: %#v", value)
	}
	if exist != true {
		t.Errorf("s[2] should exist: %#v", value)
	}
}

func Test_InstantiateSet(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	if s.IsEmpty() != false {
		t.Errorf("IsEmpty() should return false: %#v", s)
	}
	if s.Size() != 5 {
		t.Errorf("Size() should return 5: %#v", s)
	}
	value, exist := s[2]
	if value != 1 {
		t.Errorf("value should be 1: %#v", s)
	}
	if exist != true {
		t.Errorf("s[2] should exist: %#v", s)
	}
}

func Test_GetElements(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0, 10, 20)
	sl := s.GetElements()
	if len(sl) != 7 {
		t.Errorf("len(sl) should be 7: %#v", s)
	}
}

func Test_Contains(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	result := s.Contains(-0.9)
	if result != true {
		t.Errorf("s.Contains(-0.9) should return true: %#v", s)
	}
}

func Test_Delete(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	result1 := s.Delete(-0.9)
	if result1 != true {
		t.Errorf("s.Delete(-0.9) should return true: %#v", s)
	}
	result2 := s.Delete("A")
	if result2 != true {
		t.Errorf("s.Delete('A') should return true: %#v", s)
	}
	if s.Size() != 3 {
		t.Errorf("s.Size() should return 3: %#v", s)
	}
	result3 := s.Delete(100)
	if result3 != false {
		t.Errorf("s.Delete(100) should return false: %#v", s)
	}
}

func Test_Intersection(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	a := InstantiateSet(1, 2)
	result := s.Intersection(a)
	if len(result) != 2 {
		t.Errorf("len(result) should return 2: %#v", s)
	}
	if result[0] != []interface{}{1, 2}[0] {
		t.Errorf("result[0] should return 1: %#v", s)
	}
	if result[1] != []interface{}{1, 2}[1] {
		t.Errorf("result[1] should return 2: %#v", s)
	}
}

func Test_Union(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	a := InstantiateSet(100, 200)
	result := s.Union(a)
	if len(result) != 7 {
		t.Errorf("len(result) should return 7: %#v", s)
	}
}

func Test_Subtract(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	a := InstantiateSet(1, 2)
	result := s.Subtract(a)
	if len(result) != 3 {
		t.Errorf("len(result) should return 3: %#v", s)
	}
}

func Test_IsEqual(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)

	a := InstantiateSet(1, 2)
	result1 := s.IsEqual(a)
	if result1 != false {
		t.Errorf("result1 should be false: %#v", s)
	}

	b := InstantiateSet(1, 2, -.9, "A", 0)
	result2 := s.IsEqual(b)
	if result2 != true {
		t.Errorf("result2 should be true: %#v", s)
	}
}

func Test_Subset(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	a := InstantiateSet(1, 2)
	result1 := s.Subset(a)
	if result1 != true {
		t.Errorf("result1 should be true: %#v", s)
	}

	b := InstantiateSet(1, 2, -.9, "A", 0, 100)
	result2 := s.Subset(b)
	if result2 != false {
		t.Errorf("result2 should be false: %#v", s)
	}
}

func Test_Clone(t *testing.T) {
	s := InstantiateSet(1, 2, -.9, "A", 0)
	a := s.Clone()
	result := s.IsEqual(a)
	if result != true {
		t.Errorf("result should be true: %#v", s)
	}
}
