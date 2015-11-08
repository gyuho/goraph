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
	mu   sync.Mutex // guards the following
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
	forests.mu.Lock()
	forests.data[newDS] = struct{}{}
	forests.mu.Unlock()
}

// FindSet returns the DisjointSet with the represent u.
func FindSet(forests *Forests, u string) *DisjointSet {
	forests.mu.Lock()
	defer forests.mu.Unlock()
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
	forests.mu.Lock()
	forests.data[newDS] = struct{}{}
	delete(forests.data, ds1)
	delete(forests.data, ds2)
	forests.mu.Unlock()
}
