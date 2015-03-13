package graph

import (
	"encoding/json"
	"fmt"
	"io"
)

// fromJSON imports JSON file.
func fromJSON(reader io.Reader) (map[string]map[string]map[string]float64, error) {
	// If we want parallel edges in graph, use and define weights with []float64
	graphMap := make(map[string]map[string]map[string]float64)
	dec := json.NewDecoder(reader)
	for {
		if err := dec.Decode(&graphMap); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return graphMap, nil
}

// FromJSON creates a graph Data from JSON.
func FromJSON(reader io.Reader, graphName string) (*Data, error) {
	gmap1, err := fromJSON(reader)
	if err != nil {
		return nil, err
	}
	if _, ok := gmap1[graphName]; !ok {
		return nil, fmt.Errorf("%s does not exist", graphName)
	}
	gmap2 := gmap1[graphName]
	data := NewData()
	for vtxID1, weightToMap := range gmap2 {
		for vtxID2, weight := range weightToMap {
			data.Connect(NewNode(vtxID1), NewNode(vtxID2), weight)
		}
	}
	return data, nil
}
