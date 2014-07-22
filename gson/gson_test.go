package gson

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_NewJSON(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		}
	}`))
	// fmt.Sprintf("%v%v", jso, err)
	// &{map[test:map[string_array:[A B C] array:[1 2 3] arraywithsubs:[map[skone:1] map[sktwo:2 skthree:3]] int:10 float:5.15 string:simplejson bool:true]]} <nil>
	if reflect.TypeOf(jso) != reflect.TypeOf(&JSON{}) {
		test.Errorf("Should be same but %v", reflect.TypeOf(jso))
	}
}

func Test_MarshalJSON(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		}
	}`))
	r, _ := jso.MarshalJSON()
	// fmt.Sprintf("%v%v", r, e)
	// [123 34 116 101 1 ...

	if len(r) != 167 {
		test.Errorf("expected 167 but %v", len(r))
	}
}

func Test_Encode(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		}
	}`))
	r, _ := jso.Encode()
	// fmt.Sprintf("%v%v", r, e)
	// [123 34 116 101 1 ...

	if len(r) != 167 {
		test.Errorf("expected 167 but %v", len(r))
	}
}

func Test_Map(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r, _ := jso.Map()
	// fmt.Sprintf("%v%v\b", b, len(r))
	// map[test:map[string_array:[A B C] array:[1 2 3] arraywithsubs:[map[skone:1] map[sktwo:2 skthree:3]] int:10 float:5.15 string:simplejson bool:true] test1:map[A:[A B]]]

	// fmt.Println("r['test']", r["test"])
	// map[string_array:[A B C] array:[1 2 3] arraywithsubs:[map[skone:1] map[sktwo:2 skthree:3]] int:10 float:5.15 string:simplejson bool:true]

	if len(r) != 2 {
		test.Errorf("expected 2 but %v", len(r))
	}
}

func Test_SetMap(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		}
	}`))
	jso.SetMap("test", 10)
	if "&{10}" != fmt.Sprintf("%v", jso.Get("test")) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v", jso.Get("test")))
	}
}

func Test_Get(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	if "&{map[A:[A B]]}" != fmt.Sprintf("%v", jso.Get("test1")) {
		test.Errorf("Should be same but %v", fmt.Sprintf("%v", jso.Get("test1")))
	}
	if "&{[A B C]}" != fmt.Sprintf("%v", jso.Get("test").Get("string_array")) {
		test.Errorf("Should be same but %v", fmt.Sprintf("%v", jso.Get("test").Get("string_array")))
	}
}

func Test_GetCheck(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, b1 := jso.GetCheck("test1")
	if "&{map[A:[A B]]} true" != fmt.Sprintf("%v %v", r1, b1) {
		test.Errorf("Should be same but %v", fmt.Sprintf("%v %v", r1, b1))
	}

	r2, b2 := jso.GetCheck("abc")
	if "<nil> false" != fmt.Sprintf("%v %v", r2, b2) {
		test.Errorf("Should be same but %v", fmt.Sprintf("%v %v", r2, b2))
	}
}

func Test_GetBranch(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"s1": {"s2":50},
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))
	if "&{10}" != fmt.Sprintf("%v", jso.GetBranch("test", "int")) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v", jso.GetBranch("test", "int")))
	}
	if "&{50}" != fmt.Sprintf("%v", jso.GetBranch("test", "s1", "s2")) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v", jso.GetBranch("test", "s1", "s2")))
	}
}

func Test_GetByIndex(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	if "&{1}" != fmt.Sprintf("%v", jso.Get("test").Get("array").GetByIndex(0)) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v", jso.Get("test").Get("array").GetByIndex(0)))
	}
}

func Test_Bool(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, e1 := jso.Bool()
	if "false Can't type assert with bool" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}
	ss, _ := NewJSON([]byte("true"))
	r2, e2 := ss.Bool()
	if "true <nil>" != fmt.Sprintf("%v %v", r2, e2) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r2, e2))
	}
}

func Test_String(test *testing.T) {
	jso, _ := NewJSON([]byte(`"A"`))
	r1, e1 := jso.String()
	if "A <nil>" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}
}

func Test_Byte(test *testing.T) {
	jso, _ := NewJSON([]byte(`"Hello"`))
	r1, e1 := jso.Byte()
	if "[72 101 108 108 111] <nil>" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}
}

func Test_Int(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))
	r1, e1 := jso.Get("test").Get("int").Int()
	if "-1 Can't type assert with int" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}
	js2, _ := NewJSON([]byte(`1`))
	r2, e2 := js2.Int()
	if "-1 Can't type assert with int" != fmt.Sprintf("%v %v", r2, e2) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r2, e2))
	}
}

func Test_Int64(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, e1 := jso.Int64()
	if "-1 Can't type assert with int64" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}
}

func Test_Float64(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, e1 := jso.Float64()
	if "-1 Can't type assert with float64" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}
}

func Test_Slice(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, b1 := jso.Slice()
	if "[] Can't type assert with []interface{}" != fmt.Sprintf("%v %v", r1, b1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, b1))
	}

	ss, _ := NewJSON([]byte(`["A", "B", "C"]`))
	rs1, bs1 := ss.Slice()
	if "[A B C] <nil>" != fmt.Sprintf("%v %v", rs1, bs1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", rs1, bs1))
	}
}

func Test_StringSlice(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, e1 := jso.StringSlice()
	if "[] Can't type assert with []interface{}" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}

	ss, _ := NewJSON([]byte(`["A", "B", "C"]`))
	r2, e2 := ss.StringSlice()
	if "[A B C] <nil>" != fmt.Sprintf("%v %v", r2, e2) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r2, e2))
	}
}

func Test_Float64Slice(test *testing.T) {
	jso, _ := NewJSON([]byte(`{ 
		"test": { 
			"string_array": ["A", "B", "C"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"skone": 1},
			{"sktwo": 2, "skthree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true 
		},
		"test1": {
			"A": ["A", "B"]
		}
	}`))

	r1, e1 := jso.Float64Slice()
	if "[] Can't type assert with []interface{}" != fmt.Sprintf("%v %v", r1, e1) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r1, e1))
	}

	ss, _ := NewJSON([]byte(`[1, 3, 5]`))
	r2, e2 := ss.Float64Slice()
	if "[1 3 5] <nil>" != fmt.Sprintf("%v %v", r2, e2) {
		test.Errorf("expected true but %v", fmt.Sprintf("%v %v", r2, e2))
	}
}
