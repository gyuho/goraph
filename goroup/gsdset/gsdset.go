// Package gsdset implements set operations with the package graph/gsd.
package gsdset

import "github.com/gyuho/goraph/graph/gsd"

// Set maps each Vertex of graph to its Representative of DisJointSet.
type Set map[*gsd.Vertex]*gsd.Vertex

// NewSet returns a new Set.
// Map supports the built-in function "make"
// so we do not have to use "new" and
// "make" does not return pointer.
func NewSet() Set {
	return make(Set)
}

// Size returns the size of set.
func (s Set) Size() int {
	return len(s)
}

// IsEmpty returns true if the set is empty.
func (s Set) IsEmpty() bool {
	return s.Size() == 0
}

// Insert insert values to the set.
func (s Set) Insert(vertices ...*gsd.Vertex) {
	for _, value := range vertices {
		if _, exist := s[value]; exist {
			s[value] = value
		} else {
			s[value] = value
		}
	}
}

// InstantiateSet instantiates a set object
// with initial elements.
func InstantiateSet(vertices ...*gsd.Vertex) Set {
	n := NewSet()
	for _, value := range vertices {
		n.Insert(value)
	}
	return n
}

// GetElements returns the set elements.
// It makes another slice to return keys in map.
// The keys that occur more than one are
// considered to occur once in this case.
func (s Set) GetElements() []*gsd.Vertex {
	slice := []*gsd.Vertex{}
	for key := range s {
		slice = append(slice, key)
	}
	return slice
}

// Contains returns true if the value exists in the Set.
func (s Set) Contains(vtx *gsd.Vertex) bool {
	if _, exist := s[vtx]; exist {
		return true
	} else {
		return false
	}
}

// Delete deletes the value, or return false
// if the value does not exist in the Set.
func (s Set) Delete(vtx *gsd.Vertex) bool {
	if _, exist := s[vtx]; exist {
		delete(s, vtx)
		return true
	} else {
		return false
	}
}

// Intersection returns values common in both sets.
func (s Set) Intersection(a Set) []*gsd.Vertex {
	slice := []*gsd.Vertex{}
	for key := range a {
		if _, exist := s[key]; exist {
			slice = append(slice, key)
		}
	}
	return slice
}

// Union returns the union of two sets.
func (s Set) Union(a Set) []*gsd.Vertex {
	slice := []*gsd.Vertex{}
	for key := range s {
		slice = append(slice, key)
	}
	for key := range a {
		if _, exist := s[key]; !exist {
			slice = append(slice, key)
		}
	}
	return slice
}

// Subtract returns the set "s" - "a".
func (s Set) Subtract(a Set) []*gsd.Vertex {
	slice := []*gsd.Vertex{}
	for key := range s {
		if _, exist := a[key]; !exist {
			slice = append(slice, key)
		}
	}
	return slice
}

// IsEqual returns true if the two sets are same,
// regardless of its frequency.
func (s Set) IsEqual(a Set) bool {
	if s.Size() != a.Size() {
		return false
	}
	// for every element of s
	for key, _ := range s {
		// check if it exists in the Set "a"
		if _, exist := a[key]; !exist {
			return false
		}
		// if we consider its frequency
		// include
		// else if value1 != value2 {
		//	return false
	}
	return true
}

// Subset returns true if "a" is a subset of "s".
func (s Set) Subset(a Set) bool {
	if s.Size() < a.Size() {
		return false
	}
	for key := range a {
		if _, exist := s[key]; !exist {
			return false
		}
	}
	return true
}

// Clone returns a cloned set
// but does not clone its frequency.
func (s Set) Clone() Set {
	sl := s.GetElements()
	return InstantiateSet(sl...)
}

// String prints out the Set information.
func (s Set) String() string {
	result := "Vertex Set: [ "
	for _, vtx := range s.GetElements() {
		result = result + vtx.ID + " "
	}
	return result + "]"
}
