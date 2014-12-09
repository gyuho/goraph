// Package graph implements Graph data structure.
// It is `goraph`'s default graph data structure.
// It uses adjacency list and slice.
// Be aware that this does not allow duplicate(parallel) edges.
// Duplicating edges will overwrite the edge weights, adding weight values.
// (http://www.cs.cmu.edu/afs/cs/academic/class/15210-s12/www/lectures/lecture18.pdf)
package graph
