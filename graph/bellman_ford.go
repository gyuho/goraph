package graph

// BellmanFord returns the shortest path using Bellman-Ford algorithm
// This algorithm works with negative weight edges. Time complexity is O (V E).
// (http://courses.csail.mit.edu/6.006/spring11/lectures/lec15.pdf).
//
//	1.  for v in V:
//	2.  	v.d = ∞
//	3.  	v.π = nil
//	4.
//	5.  source.d = 0
//	6.
//	7.  for 1 to |V|-1:
//	8.  	for v.d > u.d + w(u,v):
//	9.  		relax(u,v)
//	10.
//	11.  relax(u,v):
//	12.  	if v.d > u.d + w(u,v)
//	13.  		v.d = u.d + w(u,v)
//	14.  		v.π = u
//
func (d *Data) BellmanFord(src, target *Node) {

}
