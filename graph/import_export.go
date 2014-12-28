package graph

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// fromJSON imports JSON file.
func fromJSON(fpath string) (map[string]map[string]map[string]float64, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	jsonStream, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// If we want parallel edges in graph, use and define weights with []float64
	// graphMap := make(map[string]map[string]map[string][]float64)
	//
	graphMap := make(map[string]map[string]map[string]float64)
	dec := json.NewDecoder(bytes.NewReader(jsonStream))
	for {
		if err := dec.Decode(&graphMap); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return graphMap, nil
}

// ToJSON exports a graph Data to JSON file.
func (d Data) ToJSON(fpath string) error {
	return nil
}

// FromDOT constructs Data from DOT file.
func FromDOT(fpath string) (map[string]map[string]map[string]float64, error) {
	return nil, nil
}

// ToDOT exports a graph Data to DOT file.
func (d Data) ToDOT(fpath string) error {
	return nil
}
