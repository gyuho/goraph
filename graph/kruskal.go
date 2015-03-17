package graph

import "sort"

// Kruskal finds the minimum spanning tree with disjoint-set data structure.
// (http://en.wikipedia.org/wiki/Kruskal%27s_algorithm)
//
//	KRUSKAL(G):
//	1 A = ∅
//	2 foreach v ∈ G.V:
//	3   MAKE-SET(v)
//	4 foreach (u, v) ordered by weight(u, v), increasing:
//	5    if FIND-SET(u) ≠ FIND-SET(v):
//	6       A = A ∪ {(u, v)}
//	7       UNION(u, v)
//	8 return A
//
func (d *Data) Kruskal() map[Edge]bool {
	dsMap := make(map[*disjointSet]bool)
	for nd := range d.NodeMap {

		one := &disjointSet{}
		one.rep = nd
		one.member = make(map[*Node]*Node)
		one.member[nd] = nd

		dsMap[one] = true
	}

	edges := d.GetEdges()
	sort.Sort(edgeSlice(edges))
	rmap := make(map[Edge]bool)

	for _, edge := range edges {

		ds1 := &disjointSet{}
		for k := range dsMap {
			if _, ok := k.member[edge.Src]; ok {
				ds1 = k
				break
			}
		}

		ds2 := &disjointSet{}
		for k := range dsMap {
			if _, ok := k.member[edge.Dst]; ok {
				ds2 = k
				break
			}
		}

		if ds1.rep != ds2.rep {
			rmap[edge] = true
			delete(dsMap, ds1)
			delete(dsMap, ds2)
			dsMap[ds1.union(ds2)] = true
		}
	}

	// rs := []Edge{}
	// for k := range rmap {
	// 	rs = append(rs, k)
	// }

	return rmap
}

// disjointSet implements disjoint set
// mapping each Node to the representative Node.
type disjointSet struct {
	rep    *Node
	member map[*Node]*Node
}

// union unions d1 with d2 with d1's rep as a new rep.
func (d1 *disjointSet) union(d2 *disjointSet) *disjointSet {
	newSet := &disjointSet{}
	newSet.rep = d1.rep
	newSet.member = d1.member
	for k := range d2.member {
		if _, ok := newSet.member[k]; !ok {
			newSet.member[k] = newSet.rep
		}
	}
	return newSet
}

type edgeSlice []Edge

func (e edgeSlice) Len() int           { return len(e) }
func (e edgeSlice) Less(i, j int) bool { return e[i].Weight < e[j].Weight }
func (e edgeSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
