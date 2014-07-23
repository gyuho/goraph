package jsonx

import (
	"io/ioutil"
	"log"

	"github.com/gyuho/goraph/gson"
	"github.com/gyuho/goraph/parsex"
)

// JG returns the JSON data in map[string]interface{} format.
func JG(fpath string) map[string]interface{} {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	js, err := gson.NewJSON([]byte(file))
	if err != nil {
		log.Fatal(err)
	}
	m, err := js.Map()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

// GetGraph returns the data of specified graph from JSON file.
func GetGraph(fpath, gname string) map[string]interface{} {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	js, err := gson.NewJSON([]byte(file))
	if err != nil {
		log.Fatal(err)
	}
	m, err := js.Get(gname).Map()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

// GetNodes returns the list of nodes in the specified graph.
func GetNodes(fpath, gname string) []string {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	js, err := gson.NewJSON([]byte(file))
	if err != nil {
		log.Fatal(err)
	}
	gm, err := js.Get(gname).Map()
	if err != nil {
		log.Fatal(err)
	}

	// keyslice contains the keys of the map
	keyslice := []string{}
	for k := range gm {
		keyslice = append(keyslice, k)
	}
	result := []string{}
	result = append(result, keyslice...)
	for _, kv := range keyslice {
		mn, _ := js.Get(gname).Get(kv).Map()
		for n := range mn {
			result = append(result, n)
		}
	}
	return parsex.UniqElemStr(result)
}

// GetGraphMap returns the graph data in map format.
// It first maps each node to its outgoing vertices.
// Each of which then maps to the edge weights.
// Duplicate edges are to be overwritten.
// And the float64 slice always has only one element.
// We use a slice to be consistent with the package `gsd`
// that allows duplicate edges.
func GetGraphMap(fpath, gname string) map[string]map[string][]float64 {
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatalf("File error: %v\n", err)
	}
	js, err := gson.NewJSON([]byte(file))
	if err != nil {
		log.Fatalf("NewJSON error: %v\n", err)
	}
	gm, err := js.Get(gname).Map()
	if err != nil {
		log.Fatalf("Map error: %v\n", err)
	}

	// keyslice contains the keys of the map
	keyslice := []string{}
	for k := range gm {
		keyslice = append(keyslice, k)
	}

	result := []string{}
	result = append(result, keyslice...)

	for _, v := range keyslice {
		mn, err := js.Get(gname).Get(v).Map()
		if err != nil {
			log.Fatalf("keyslice err: %v", err)
		}
		for n := range mn {
			result = append(result, n)
		}
	}

	nodes := parsex.UniqElemStr(result)

	rm := make(map[string]map[string][]float64)
	for _, srcNode := range nodes {
		// each node's outgoing vertices
		mn, _ := js.Get(gname).Get(srcNode).Map()
		if err != nil {
			// log.Fatal(err)
			// in case, there is no outgoing vertices
			continue
		}
		ms := make(map[string][]float64)
		for dstNode := range mn {
			// d, _ := mn[dstNode]
			// ms[dstNode] = d.(float64)
			fs, _ := js.Get(gname).Get(srcNode).Get(dstNode).Float64Slice()

			// Overwrite the duplicate edges with the largest edge value
			// sort.Float64s(fs)
			// ms[dstNode] = []float64{fs[len(fs)-1]}

			// Overwrite with sum
			sum := 0.0
			for _, v := range fs {
				sum = sum + v
			}
			ms[dstNode] = []float64{sum}
		}
		rm[srcNode] = ms
	}

	return rm
}
