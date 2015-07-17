// Package redblack implements a Left-Leaning Red-Black tree
// , which is a 2-3 balanced binary search tree:
//	* Each edge(link between nodes) is either black or red.
//	* Every path from root to null edge has the same number of black edges.
//	* Red edges lean left.
//	* Two red edges in a row are not allowed.
//	* Only rotate red edges.
//
// Reference
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/08Penn.pdf
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//	* https://www.youtube.com/watch?v=lKmLBOJXZHI
//	* https://github.com/petar/GoLLRB
//
package redblack // import "github.com/gyuho/goraph/redblack"
