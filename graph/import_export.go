package graph

import (
	"encoding/json"
	"fmt"
	"io"
)

// fromJSON imports JSON file.
func fromJSON(reader io.Reader) (map[string]map[string]map[string]float32, error) {
	// If we want parallel edges in graph, use and define weights with []float32
	graphMap := make(map[string]map[string]map[string]float32)
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

// FromJSON creates a graph Data from JSON. Here's the sample JSON data:
//
//	{
//	    "test_graph_01": {
//	        "S": {
//	            "A": 100,
//	            "B": 14,
//	            "C": 200
//	        },
//	        "A": {
//	            "S": 15,
//	            "B": 5,
//	            "D": 20,
//	            "T": 44
//	        },
//	        "B": {
//	            "S": 14,
//	            "A": 5,
//	            "D": 30,
//	            "E": 18
//	        },
//	        "C": {
//	            "S": 9,
//	            "E": 24
//	        },
//	        "D": {
//	            "A": 20,
//	            "B": 30,
//	            "E": 2,
//	            "F": 11,
//	            "T": 16
//	        },
//	        "E": {
//	            "B": 18,
//	            "C": 24,
//	            "D": 2,
//	            "F": 6,
//	            "T": 19
//	        },
//	        "F": {
//	            "D": 11,
//	            "E": 6,
//	            "T": 6
//	        },
//	        "T": {
//	            "A": 44,
//	            "D": 16,
//	            "F": 6,
//	            "E": 19
//	        }
//	    },
//	}
//
func FromJSON(reader io.Reader, graphName string) (*Data, error) {
	gmap1, err := fromJSON(reader)
	if err != nil {
		return nil, err
	}
	if _, ok := gmap1[graphName]; !ok {
		return nil, fmt.Errorf("%s does not exist", graphName)
	}
	gmap2 := gmap1[graphName]
	data := New()
	for vtxID1, weightToMap := range gmap2 {
		for vtxID2, weight := range weightToMap {
			data.Connect(NewNode(vtxID1), NewNode(vtxID2), weight)
		}
	}
	return data, nil
}
