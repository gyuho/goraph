// Package llrb implements a Left-Leaning Red-Black tree
// , which is a 2-3 balanced binary search tree:
//	* Each edge(link between nodes) is either black or red.
//	* Or every node is either black or red — easier to code.
//	* Every path from root to null edge has the same number of black edges.
//	* Red nodes(edges) lean left.
//	* Two red nodes in a row are not allowed.
//	* We only rotate red ones.
//
// Reference
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/08Penn.pdf
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//	* https://www.youtube.com/watch?v=lKmLBOJXZHI
//	* https://github.com/petar/GoLLRB
//
package llrb // import "github.com/gyuho/goraph/llrb"
