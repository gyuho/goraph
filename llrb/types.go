package llrb

import "fmt"

type nodeStruct struct {
	ID    string
	Value int
}

func (n nodeStruct) Less(b Interface) bool {
	return n.Value < b.(nodeStruct).Value
}

func (n nodeStruct) String() string {
	return fmt.Sprintf("%s(%d)", n.ID, n.Value)
}

type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
}

type String string

// Less returns true if string(a) < string(b).
func (a String) Less(b Interface) bool {
	return a < b.(String)
}
