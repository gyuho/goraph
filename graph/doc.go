// Package graph implements graph data structure with Go structs, pointers, and map.
// (http://en.wikipedia.org/wiki/Adjacency_list)
// This package does not allow duplicate(parallel) edges, referred to as a multigraph.
// Connecting two nodes with duplicate edges overwrites the current weight value (not adding-up!).
// http://www.cs.cmu.edu/afs/cs/academic/class/15210-s12/www/lectures/lecture18.pdf also replaces
// parallel edges with a single edge, while having redundant(parallel) edges can sometimes be convenient
package graph // import "github.com/gyuho/goraph/graph"
