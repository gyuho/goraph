package jsonx

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/gyuho/goraph/gson"
	"github.com/gyuho/goraph/parsex"
)

func TestJSONX(t *testing.T) {
	file, err := ioutil.ReadFile("../../files/testgraph.json")
	if err != nil {
		t.Fatalf("%v", err)
	}
	js, err := gson.NewJSON([]byte(file))
	if err != nil {
		t.Fatalf("%v", err)
	}
	m, err := js.Map()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Printf("%v\n%v\n\n", reflect.TypeOf(m), m)
	// map[string]interface {}
	// map[testgraph.001:map ...

	fmt.Printf("%v\n", reflect.TypeOf(js.Get("testgraph.003").Get("S")))
	// *gson.JSON

	fmt.Printf("%v\n", js.Get("testgraph.003").Get("S"))
	// &{map[C:[200] A:[100] B:[6 14]]}

	fmt.Printf("%v\n", js.Get("testgraph.003").Get("S").Get("A"))
	// &{[100]}

	fmt.Printf("Duplicate B: %v\n", js.Get("testgraph.003").Get("S").Get("B"))
	// Duplicate B: &{[6 14]}

	a := js.Get("testgraph.003").Get("S").Get("A")
	fmt.Println(*a)                 // {[100]}
	fmt.Println(reflect.TypeOf(a))  // *gson.JSON
	fmt.Println(reflect.TypeOf(*a)) // gson.JSON
	fmt.Println(a.Data)             // [100]
	fmt.Println(a.Slice())          // [100] <nil>

	for k := range m {
		fmt.Println("Key", k)
	}
	/*
	   Key testgraph.001
	   Key testgraph.004
	   Key testgraph.006
	   ...
	*/

	m3, _ := js.Get("testgraph.003").Map()
	for k := range m3 {
		fmt.Println("Key for testgraph3:", k, reflect.TypeOf(m3))
	}
	/*
	   Key for testgraph3: S map[string]interface {}
	   Key for testgraph3: A map[string]interface {}
	   Key for testgraph3: B map[string]interface {}
	   ...
	*/

	println()
	m3s, _ := js.Get("testgraph.003").Get("S").Map()
	for k := range m3s {
		fmt.Println("Key for testgraph3, S:", k, reflect.TypeOf(m3s))
		fmt.Println("Value:", m3s[k], reflect.TypeOf(m3s[k]))
		println()
	}
	/*
	   Key for testgraph3, S: A map[string]interface {}
	   Value: 100 float64

	   Key for testgraph3, S: B map[string]interface {}
	   Value: 14 float64

	   Key for testgraph3, S: C map[string]interface {}
	   Value: 200 float64

	   "S": {
	   	"A": 100,
	   	"B": 6,
	   	"B": 14,
	   	"C": 200
	   },

	   Duplicate edges have been overwritten
	*/
}

func TestJG(t *testing.T) {
	m := JG("../../files/testgraph.json")
	fmt.Println(m)
	// map[testgraph.004:map[S:map[B:14] A:map[S:15 B:5 D:20 T:44] B:map[S:14 A:5 D:30 E:18] C:map[S:9 E:24] E:map[B:18 C:24 D:2 F:6 T:19] F:map[D:11 E:6 T:6] T:map[A:44 D:16 F:6 E:19]] testgraph.008:map[A:map[E:1 F:1] B:map[A:1] D:map[B:1 C:1] E:map[D:1 C:1 F:1]] testgraph.011:map[S:map[A:11 B:17 C:9] A:map[S:11 B:5 D:50 T:500] B:map[S:17 D:30] C:map[S:9] D:map[A:50 B:30 E:3 F:11] E:map[B:18 D:2 F:6 T:19] F:map[D:11 E:6 T:77] T:map[A:500 D:10 F:77 E:19]] testgraph.013:map[S:map[A:7 B:6] A:map[C:-3 T:9] B:map[A:-8 C:5 T:-4] C:map[B:-2] T:map[S:2 C:7]] testgraph.001:map[S:map[A:100 B:14 C:200] A:map[S:15 B:5 D:20 T:44] B:map[S:14 A:5 D:30 E:18] C:map[S:9 E:24] D:map[A:20 B:30 E:2 F:11 T:16] E:map[B:18 C:24 D:2 F:6 T:19] F:map[D:11 E:6 T:6] T:map[A:44 D:16 F:6 E:19]] testgraph.002:map[S:map[A:100 B:14 C:200] A:map[S:15 B:5 D:20 T:44] B:map[S:14 A:5 D:30 E:18] C:map[S:9 E:24] E:map[B:18 D:2 F:6 T:19] F:map[D:11 E:6 T:6] T:map[A:44 D:16 F:6 E:19]] testgraph.003:map[S:map[A:100 B:14 C:200] A:map[S:15 B:5 D:20 T:44] B:map[S:14 A:5 D:30 E:18] C:map[S:9 E:24] D:map[A:20 B:30 E:2 F:11 T:16] E:map[B:18 C:24 D:2 F:6 T:19] F:map[D:11 E:6 T:6] T:map[A:44 D:16 F:6 E:19]] testgraph.005:map[A:map[B:7 C:9 F:20] B:map[A:7 C:10 D:15] C:map[A:9 B:10 D:11 E:30 F:2] D:map[B:15 C:11 E:2] E:map[F:9 C:30 D:2] F:map[A:20 C:2 E:9]] testgraph.006:map[A:map[F:1] B:map[A:1] D:map[B:1 C:1] E:map[C:1 F:1]] testgraph.007:map[A:map[E:100 H:14] B:map[D:1] C:map[D:1 E:1] D:map[F:1 G:1 H:1] E:map[G:1]] testgraph.009:map[A:map[B:1 E:1 H:1] B:map[C:1 D:1] C:map[D:1 E:1] D:map[F:1 G:1 H:1] E:map[A:1 G:1] F:map[E:1] G:map[H:1] H:map[F:1]] testgraph.010:map[A:map[C:9 F:20] B:map[A:1 D:15] C:map[B:10 E:30] D:map[C:11 E:2] E:map[C:30 F:9] F:map[A:20 C:2]] testgraph.012:map[S:map[A:7 B:6] A:map[C:-3 T:9] B:map[A:8 C:5 T:-4] C:map[B:-2] T:map[C:7 S:2]]]
}

func TestGetGraph(t *testing.T) {
	m := GetGraph("../../files/testgraph.json", "testgraph.003")
	if len(m) != 8 {
		t.Fatalf("expected 8 but %v", m)
	}
}

func TestGetNodes(t *testing.T) {
	ns := GetNodes("../../files/testgraph.json", "testgraph.003")
	sl := []string{"S", "A", "B", "C", "D", "E", "F", "T"}
	if !parsex.EqualSliceElem(ns, sl) {
		t.Errorf("expected true but: %v", ns)
	}
}

func TestGetGraphMap(t *testing.T) {
	ns := GetGraphMap("../../files/testgraph.json", "testgraph.003")
	sl := []string{"S", "A", "B", "C", "D", "E", "F", "T"}
	if len(ns) != len(sl) {
		t.Errorf("expected 8 but: %v", ns)
	}
}
