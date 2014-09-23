package glevel

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

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
	err := db.Put([]byte(id), []byte(fmt.Sprintf("%v", val)), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetVertex returns the value of a vertex.
func GetVertex(db *leveldb.DB, id string) float64 {
	data, err := db.Get([]byte(id), nil)
	if err != nil {
		log.Fatal(err)
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
	err := db.Put([]byte(key), []byte(fmt.Sprintf("%v", val)), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GetEdge (returns the edge weight.
func GetEdge(db *leveldb.DB, subject, object string) float64 {
	key := subject + "----->" + object
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal(err)
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
		key := string(iter.Key())
		if strings.HasSuffix(key, "----->"+object) {
			value := StringToFloat64(string(iter.Value()))
			rs[key] = value
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

// StringToFloat64 converts string to float64.
func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
