package graph

// BellmanFord returns the shortest path using Bellman-Ford algorithm
// This algorithm works with negative weight edges.
// (http://courses.csail.mit.edu/6.006/spring11/lectures/lec15.pdf).
//
//	for v in V:
//		v.d = ∞
//		v.π = nil
//
//	source.d = 0
//
//	for 1 to |V|-1:
//		for v.d > u.d + w(u,v):
//			relax(u,v)
//
//	relax(u,v):
//		if v.d > u.d + w(u,v)
//			v.d = u.d + w(u,v)
//			v.π = u
//
// O (V E)
func (d *Data) BellmanFord(src, target *Node) {

}
