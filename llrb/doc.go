// Package llrb implements a Left-Leaning Red-Black tree
// , which is a 2-3 balanced binary search tree:
//	* Every node(or edge) is either black or red.
//	* Every path from root to null Node has the same number of black nodes.
//	* Red nodes lean left.
//	* Two red nodes in a row are not allowed.
//
// Reference
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/08Penn.pdf
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//	* https://www.cs.princeton.edu/~rs/talks/LLRB/RedBlack.pdf
//	* https://www.youtube.com/watch?v=lKmLBOJXZHI
//	* https://github.com/petar/GoLLRB
//
package llrb // import "github.com/gyuho/goraph/llrb"
