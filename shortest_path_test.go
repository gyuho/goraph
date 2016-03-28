package goraph

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGraph_Dijkstra_03(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_03")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, StringID("S"), StringID("T"))
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00)" {
		t.Errorf("Expected the shortest path S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("T")] != 44.0 {
		t.Errorf("Expected 44.0 but %f", distance[StringID("T")])
	}
	fmt.Println("graph_03:", strings.Join(ts, " → "))
}

func TestGraph_Dijkstra_04(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_04")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, StringID("A"), StringID("E"))
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "A(0.00) → C(9.00) → F(11.00) → E(20.00)" {
		t.Errorf("Expected the shortest path A(0.00) → C(9.00) → F(11.00) → E(20.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("E")] != 20.0 {
		t.Errorf("Expected 20.0 but %f", distance[StringID("E")])
	}
	fmt.Println("graph_04:", strings.Join(ts, " → "))
}

func TestGraph_Dijkstra_09_0(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_09")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, StringID("A"), StringID("E"))
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "A(0.00) → C(9.00) → B(19.00) → D(34.00) → E(36.00)" {
		t.Errorf("Expected the shortest path A(0.00) → C(9.00) → B(19.00) → D(34.00) → E(36.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("E")] != 36.0 {
		t.Errorf("Expected 36.0 but %f", distance[StringID("E")])
	}
	fmt.Println("graph_09:", strings.Join(ts, " → "))
}

func TestGraph_Dijkstra_09_1(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_09")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, StringID("E"), StringID("A"))
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "E(0.00) → F(9.00) → C(11.00) → B(21.00) → A(22.00)" {
		t.Errorf("Expected the shortest path E(0.00) → F(9.00) → C(11.00) → B(21.00) → A(22.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("A")] != 22.0 {
		t.Errorf("Expected 22.0 but %f", distance[StringID("A")])
	}
	fmt.Println("graph_09:", strings.Join(ts, " → "))
}

func TestGraph_Dijkstra_10_0(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_10")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, StringID("S"), StringID("T"))
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → A(11.00) → B(16.00) → D(46.00) → E(49.00) → T(68.00)" {
		t.Errorf("Expected the shortest path S(0.00) → A(11.00) → B(16.00) → D(46.00) → E(49.00) → T(68.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("T")] != 68.0 {
		t.Errorf("Expected 68.0 but %f", distance[StringID("T")])
	}
	fmt.Println("graph_10:", strings.Join(ts, " → "))
}

func TestGraph_Dijkstra_10_1(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_10")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := Dijkstra(g, StringID("T"), StringID("S"))
	if err != nil {
		t.Error(err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "T(0.00) → D(10.00) → E(13.00) → B(31.00) → S(48.00)" {
		t.Errorf("Expected the shortest path T(0.00) → D(10.00) → E(13.00) → B(31.00) → S(48.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("S")] != 48.0 {
		t.Errorf("Expected 48.0 but %f", distance[StringID("S")])
	}
	fmt.Println("graph_10:", strings.Join(ts, " → "))
}

func TestGraph_BellmanFord_11(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_11")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := BellmanFord(g, StringID("S"), StringID("T"))
	if err != nil {
		t.Errorf("There should be no negative-weight cycle but found one with %v", err)
	}
	ts := []string{}
	for _, v := range path {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", v, distance[v]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → A(7.00) → C(4.00) → B(2.00) → T(-2.00)" {
		t.Errorf("Expected the shortest path S(0.00) → A(7.00) → C(4.00) → B(2.00) → T(-2.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("T")] != -2.0 {
		t.Errorf("Expected -2.0 but %f", distance[StringID("T")])
	}
	fmt.Println("graph_11:", strings.Join(ts, " → "))
}

func TestGraph_BellmanFord_12(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_12")
	if err != nil {
		t.Error(err)
	}
	path, distance, err := BellmanFord(g, StringID("S"), StringID("T"))
	if err == nil {
		t.Errorf("There should be negative-weight cycle but %v", err)
	}
	if path != nil || distance != nil {
		t.Errorf("Expected nil, nil but %v, %v", path, distance)
	}
}
