// Package redblack implements a Left-Leaning Red-Black tree, which is a 2-3 balanced binary search tree.
//	* No node has two red links connected to it.
//	* Every path from root to null link has the same number of black links.
//	* Red links lean left.
//
// Binary-Search does not examine the color and in the same way as a regular BST.
// It performs better with better balance. Just make sure during insertion, maintain the balance property:
//	* If the red links are leaning in the right, we need to rotate-left.
//	* Sometimes we need to temporarily rotate to the right.
//
// Reference
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/08Penn.pdf
//	* http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//	* https://github.com/petar/GoLLRB
//
package redblack // import "github.com/gyuho/goraph/redblack"
