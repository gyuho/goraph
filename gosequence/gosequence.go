// Package gosequence is a small package to handle sequence in Go.
// It provides slice operations: insert and delete(pop) that can
// be used for stack and queue data structure. The present methods
// are already supported by the standard package 'container/list'
// (http://golang.org/pkg/container/list/) that uses doubly linked list.
// You may use slice for the sake of performance and type safety.
// Refer to the link for more detail about whether to use list or slice.
// (https://groups.google.com/d/msg/golang-nuts/mPKCoYNwsoU/tLefhE7tQjMJ)
// It also provides map operations, trying to achieve the orderness in
// unordered data structure map. Note that the biggest difference between
// map and sequence is that map is unordered and sequence is ordered.
package gosequence

import (
	"fmt"
	"sort"
)

// Sequence can contain any type of values
// , because its data is a slice of interface{} type.
// It is an empty interface, which means that it can
// be satisfied by any type of value.
// We deal with sequence of order, so it is natural
// to use slice as our basic data structure.
type Sequence []interface{}

// NewSequence returns a new object of Sequence.
func NewSequence() *Sequence {
	return new(Sequence)
}

// Len returns the length of sequence.
// If the method needs to mutate the receiver
// , the receiver must be a pointer.
// For more detail, refer to this link.
// (http://golang.org/doc/faq#methods_on_values_or_pointers)
func (s Sequence) Len() int {
	// if (s Sequence)
	// len(*s)
	return len(s)
}

// IsEmpty returns true if the sequence is empty.
func (s Sequence) IsEmpty() bool {
	return len(s) == 0
}

// Init initializes the sequence.
func (s *Sequence) Init() {
	s = NewSequence()
}

// Copy returns a copy of the sequence.
// This is useful because ":=" operator
// does deep copy and when we manipulate
// either one, then the other one also changes.
func (s Sequence) Copy() Sequence {
	ts := NewSequence()
	for _, v := range s {
		ts.PushBack(v)
	}
	return *ts
}

// Find returns the index of input value
// , with true if the value exists in the sequence.
func (s Sequence) Find(val interface{}) (int, bool) {
	for key, value := range s {
		if fmt.Sprintf("%v", value) == fmt.Sprintf("%v", val) {
			return key, true
		}
	}
	return 0, false
}

// PushFront adds an element to the front of sequence.
func (s *Sequence) PushFront(val interface{}) {
	temp := make([]interface{}, len(*s)+1)
	temp[0] = val
	copy(temp[1:], *s)
	*s = temp
}

// PushBack adds an element to the back of sequence.
func (s *Sequence) PushBack(val interface{}) {
	*s = append(*s, val)
}

// DeepSliceDelete deletes the element in the index n.
func (s *Sequence) DeepSliceDelete(n int) {
	copy((*s)[n:], (*s)[n+1:])
	(*s)[len(*s)-1] = nil
	// zero value of type or nil

	(*s) = (*s)[:len(*s)-1 : len(*s)-1]
}

// FindAndDelete finds the element and delete it.
func (s *Sequence) FindAndDelete(val interface{}) bool {
	idx, ok := s.Find(val)
	if !ok {
		return false
	}
	s.DeepSliceDelete(idx)
	return true
}

// DeepSliceCut deletes the elements from indices a to b.
func (s *Sequence) DeepSliceCut(a, b int) {
	if b > len(*s)-1 || a < 0 || a > b {
		panic("Index out of range! You can cut only inside slice.")
	}
	diff := b - a + 1
	idx := a
	i := 0
	for i < diff {
		s.DeepSliceDelete(idx)
		i++
	}
}

// Front returns the first(front) element of sequence.
func (s Sequence) Front() interface{} {
	return s[0]
}

// Back returns the last(back) element of sequence.
func (s Sequence) Back() interface{} {
	return s[s.Len()-1]
}

// PopFront removes the front(first) element of sequence
// and return it at the same time.
func (s *Sequence) PopFront() interface{} {
	tm := (*s).Front()
	(*s).DeepSliceDelete(0)
	return tm
}

// PopBack removes the back(last) element of sequence
// and return it at the same time.
func (s *Sequence) PopBack() interface{} {
	tm := (*s).Back()
	(*s).DeepSliceDelete((*s).Len() - 1)
	return tm
}

// GetElements returns a slice of all values.
func (s Sequence) GetElements() []interface{} {
	tm := s.Copy()
	slice := []interface{}{}
	for tm.Len() != 0 {
		slice = append(slice, tm.PopFront())
	}
	return slice
}

// CommonPrefix returns the longest common leading components
// among all Sequence. Python commonPrefix compares the maximal
// Sequence with the minimal Sequence, which only takes linear time
// , whereas this compares every possible pair among all Sequence
// , which makes it slower, but still quadratic, than Python.
// This is to find the common prefix among all Sequence
// , not just between maximal and minimal Sequence.
func CommonPrefix(more ...Sequence) []interface{} {
	minl := more[0]
	min := more[0].Len()
	// to get the Sequence of the shortest length
	for _, value := range more {
		if value.Len() < min {
			minl = value
			min = value.Len()
		}
	}
	// traverse the minimal Sequence
	// and compare with other Sequence
	// elements in the same index
	for key, value := range minl {
		// if any value in other Sequence
		// is different than the one
		// in minimal Sequence
		for _, other := range more {
			if value != other[key] {
				return minl[:key]
			}
		}
	}
	return minl
}

// MapSF is a key/value pair of map[string]float64
type MapSF struct {
	key   string
	value float64
}

// MapSFSlice is a slice of MapSF
// to be used in sort.Interface.
type MapSFSlice []MapSF

// sort.Interface for MapSFSlice.
func (p MapSFSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p MapSFSlice) Len() int      { return len(p) }
func (p MapSFSlice) Less(i, j int) bool {
	return p[i].value < p[j].value
}

// SortMapSFByValue sorts the map by its value,
// and return a sorted MapSFSlice.
// This can handle the keys with duplicate values.
// The original map does not change
// because the map itself is unordered.
func SortMapSFByValue(m map[string]float64) MapSFSlice {
	p := make(MapSFSlice, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapSF{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

// MapSI is a key/value pair of map[string]int64
type MapSI struct {
	key   string
	value int64
}

// MapSISlice is a slice of MapSI
// to be used in sort.Interface.
type MapSISlice []MapSI

// sort.Interface for MapSISlice.
func (p MapSISlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p MapSISlice) Len() int      { return len(p) }
func (p MapSISlice) Less(i, j int) bool {
	return p[i].value < p[j].value
}

// SortMapSIByValue sorts the map by its value,
// and return a sorted MapSISlice.
// This can handle the keys with duplicate values.
// The original map does not change
// because the map itself is unordered.
func SortMapSIByValue(m map[string]int64) MapSISlice {
	p := make(MapSISlice, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapSI{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

// SwitchMap converts key and value of map data structure.
// Map does not allow to have duplicate keys, so we delete
// one of the elements with the same value, before switching.
// We convert the arguments to map[interface{}]interface{} type.
func SwitchMap(m map[interface{}]interface{}) map[interface{}]interface{} {
	nm := make(map[interface{}]interface{})
	for key, value := range m {
		if _, exist := nm[value]; exist {
			delete(m, key)
		} else {
			nm[value] = key
		}
	}
	return nm
}
