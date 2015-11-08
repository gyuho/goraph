package graph

import "sync"

// DisjointSet implements disjoint set.
// (https://en.wikipedia.org/wiki/Disjoint-set_data_structure)
type DisjointSet struct {
	represent string
	member    map[string]struct{}
}

// Forests is a set of DisjointSet.
type Forests struct {
	sync.Mutex
	data map[*DisjointSet]struct{}
}

// NewForests creates a new Forests.
func NewForests() *Forests {
	set := &Forests{}
	set.data = make(map[*DisjointSet]struct{})
	return set
}

// MakeDisjointSet creates a DisjointSet.
func MakeDisjointSet(forests *Forests, vtx string) {
	newDS := &DisjointSet{}
	newDS.represent = vtx
	member := make(map[string]struct{})
	member[vtx] = struct{}{}
	newDS.member = member
	forests.Lock()
	forests.data[newDS] = struct{}{}
	forests.Unlock()
}

// FindSet returns the DisjointSet with the represent u.
func FindSet(forests *Forests, u string) *DisjointSet {
	forests.Lock()
	defer forests.Unlock()
	for data := range forests.data {
		if data.represent == u {
			return data
		}
		for k := range data.member {
			if k == u {
				return data
			}
		}
	}
	return nil
}

// Union unions two DisjointSet, with ds1's represent.
func Union(forests *Forests, ds1, ds2 *DisjointSet) {
	newDS := &DisjointSet{}
	newDS.represent = ds1.represent
	newDS.member = ds1.member
	for k := range ds2.member {
		newDS.member[k] = struct{}{}
	}
	forests.Lock()
	forests.data[newDS] = struct{}{}
	delete(forests.data, ds1)
	delete(forests.data, ds2)
	forests.Unlock()
}
