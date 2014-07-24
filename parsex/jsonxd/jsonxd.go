package jsonxd

import (
	"io/ioutil"
	"log"
	"sort"

	"github.com/gyuho/goraph/gson"
	"github.com/gyuho/goraph/parsex"
)

// JGD returns the data in map[string]interface{} format.
func JGD(fpath string) map[string]interface{} {
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

// GetGraphMapDupl returns the graph data in map format.
// It first maps each node to its outgoing vertices.
// Each of which then maps to the edge weights.
// It uses slices in order to allow duplicate edges.
func GetGraphMapDupl(fpath, gname string) map[string]map[string][]float64 {
	// func GetGraphMap(fpath, gname string) map[string]map[string]float64 {
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

	for _, v := range keyslice {
		mn, err := js.Get(gname).Get(v).Map()
		if err != nil {
			log.Fatal(err)
		}
		for n := range mn {
			result = append(result, n)
		}
	}

	nodes := parsex.UniqElemStr(result)

	rm := make(map[string]map[string][]float64)
	for _, nodeName := range nodes {
		// each node's outgoing vertices
		mn, _ := js.Get(gname).Get(nodeName).Map()
		ms := make(map[string][]float64)
		for vtx := range mn {
			fs, _ := js.Get(gname).Get(nodeName).Get(vtx).Float64Slice()
			sort.Float64s(fs)
			ms[vtx] = fs
		}
		rm[nodeName] = ms
	}

	return rm
}
