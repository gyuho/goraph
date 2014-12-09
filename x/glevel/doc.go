// Package glevel is a graph data structure based on `LevelDB`.
// `LevelDB` is a key-value storage from Google.
// It simply maps key to value of string in order.
// https://github.com/google/leveldb
//
// On `LevelDB` we store vertices by ID and Value.
// And we store edges by Subject----->Object as key, and Weight as value.
// Which requires not to have duplicate edges from a Vertex.
//
// Additional APIs can be found at http://godoc.org/github.com/syndtr/goleveldb/leveldb.
package glevel
