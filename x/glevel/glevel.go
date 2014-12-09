package glevel

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//
// TODO(gyuho)
// LevelDB stores data by key order.
// Would it be better if we tag vertex and edge
// by prefix, so that we can specify the retrieval range?
//
//

//
// TODO(gyuho)
// It might not be best to return 0 for non-existent keyword
// instead of return error, but it makes easier
// maybe harder to debug...
//
//

//
// TODO(gyuho)
// For now, we will assume only one outgoing edge from a vertex.
// GetOutEdges and GetOutVertices returns the slice of map
// with length 1.
// (I implemented this package specifically to use for hierarchy.
// Having more than two hypernym makes things too complicated. )
//
//

// OpenGraph returns the `LevelDB` database.
// Make sure to `defer db.Close()` to save.
func OpenGraph(filepath string) *leveldb.DB {
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// PutVertex adds a vertex to the graph database.
func PutVertex(db *leveldb.DB, id string, val float64) {
	err := db.Put([]byte(id), []byte(strconv.FormatFloat(val, 'f', 6, 64)), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetVertex returns the value of a vertex.
func GetVertex(db *leveldb.DB, id string) float64 {
	data, err := db.Get([]byte(id), nil)
	if err != nil {
		// log.Fatal(err)
		return 0.0
	}
	return StringToFloat64(string(data))
}

// DeleteVertex delete a vertex by its ID.
// Make sure to save with `db.Close()` at the end.
func DeleteVertex(db *leveldb.DB, id string) {
	err := db.Delete([]byte(id), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// PutEdge adds an edge to the graph database.
// The key will be a concatenated string of `Subject` + `----->` + `Object`.
func PutEdge(db *leveldb.DB, subject string, val float64, object string) {
	key := subject + "----->" + object
	err := db.Put([]byte(key), []byte(strconv.FormatFloat(val, 'f', 6, 64)), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetEdge (returns the edge weight.
func GetEdge(db *leveldb.DB, subject, object string) float64 {
	key := subject + "----->" + object
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		// log.Fatal(err)
		return 0.0
	}
	return StringToFloat64(string(data))
}

// DeleteEdge delete an edge by its ID.
// Make sure to save with `db.Close()` at the end.
func DeleteEdge(db *leveldb.DB, subject, object string) {
	key := subject + "----->" + object
	err := db.Delete([]byte(key), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetOutEdges returns all edges from a subject(source) vertex
// matching keys by its prefixes.
func GetOutEdges(db *leveldb.DB, subject string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(util.BytesPrefix([]byte(subject+"----->")), nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := string(iter.Key())
		value := StringToFloat64(string(iter.Value()))
		rs[key] = value
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// GetInEdges returns all edges from an object(destination) vertex
// matching keys by its suffixes.
// We cannot specify the iterator range since `LevelDB` stores by its key.
// This could be slower than prefix-matching.
func GetInEdges(db *leveldb.DB, object string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		if bytes.HasSuffix(key, []byte("----->"+object)) {
			value := StringToFloat64(string(iter.Value()))
			rs[string(key)] = value
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// GetOutVertices returns all vertices from a subject(source) vertex
// matching keys by its prefixes.
func GetOutVertices(db *leveldb.DB, subject string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(util.BytesPrefix([]byte(subject+"----->")), nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := string(iter.Key())
		value := StringToFloat64(string(iter.Value()))
		rs[key[len(subject+"----->"):]] = value
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// GetInVertices returns all vertices from an object(destination) vertex
// matching keys by its suffixes.
// We cannot specify the iterator range since `LevelDB` stores by its key.
// This could be slower than prefix-matching.
func GetInVertices(db *leveldb.DB, object string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		if strings.HasSuffix(key, "----->"+object) {
			value := StringToFloat64(string(iter.Value()))
			lm := len(key) - len("----->"+object)
			rs[key[:lm]] = value
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// StringVertex returns a string to describe a vertex.
func StringVertex(db *leveldb.DB, id string) string {
	rs := id + " has a value of " + fmt.Sprintf("%v", GetVertex(db, id)) + ".\n\n"

	m1 := GetOutVertices(db, id)
	rs = rs + "Outgoing Vertices:\n"
	if len(m1) == 0 {
		rs = rs + "No Outgoing Vertices...\n"
	} else {
		for object, value := range m1 {
			rs = rs + id + " -- " + fmt.Sprintf("%v", value) + " -> " + object + "\n"
		}
	}

	m2 := GetInVertices(db, id)
	rs = rs + "\nIncoming Vertices:\n"
	if len(m2) == 0 {
		rs = rs + "No Incoming Vertices...\n"
	} else {
		for subject, value := range m2 {
			rs = rs + subject + " -- " + fmt.Sprintf("%v", value) + " -> " + id + "\n"
		}
	}

	return rs
}

// SearchVertexByPrefix returns all vertices containing the prefix.
// It is relatively faster.
func SearchVertexByPrefix(db *leveldb.DB, prefix string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		key := iter.Key()
		// only when it's not an edge
		if !bytes.Contains(key, []byte("----->")) {
			value := StringToFloat64(string(iter.Value()))
			rs[string(key)] = value
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// SearchVertexBySuffix returns all vertices containing the suffix.
// It is relatively slower.
func SearchVertexBySuffix(db *leveldb.DB, suffix string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		// only when it's not an edge
		if !bytes.Contains(key, []byte("----->")) {
			if bytes.HasSuffix(key, []byte(suffix)) {
				value := StringToFloat64(string(iter.Value()))
				rs[string(key)] = value
			}
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// SearchVertexBySubstring returns all vertices containing the substring.
// Use only when you need. This can be slow.
func SearchVertexBySubstring(db *leveldb.DB, sub string) map[string]float64 {
	rs := make(map[string]float64)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		// only when it's not an edge
		if !bytes.Contains(key, []byte("----->")) {
			if bytes.Contains(key, []byte(sub)) {
				value := StringToFloat64(string(iter.Value()))
				rs[string(key)] = value
			}
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return rs
}

// StringToFloat64 converts string to float64.
func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
