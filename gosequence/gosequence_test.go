package gosequence

import (
	"reflect"
	"testing"
)

func Test_Sequence(t *testing.T) {
	sample1 := &Sequence{}
	sample2 := NewSequence()
	if reflect.TypeOf(sample1) != reflect.TypeOf(sample2) {
		t.Errorf("sample1 should be *Sequence type: %#v", sample2)
	}
}

func Test_Len(t *testing.T) {
	// Don't be confused with []interface{1, "A", 3, -.9, "B"}
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	if len(sample) != 5 {
		t.Errorf("sample.Len() should be '5': %#v", sample)
	}
}

func Test_IsEmpty(t *testing.T) {
	sample := NewSequence()
	isempty := sample.IsEmpty()
	if isempty != true {
		t.Errorf("sample.IsEmpty() should return 'true': %#v", sample)
	}
}

func Test_Init(t *testing.T) {
	sample := NewSequence()
	sample.PushBack(1)
	sample.PushBack(2)
	sample.PushBack(3)
	isempty1 := sample.IsEmpty()
	sample.Init()
	isempty2 := sample.IsEmpty()

	if isempty1 != false && isempty2 != true {
		t.Errorf("Should return 'false' and 'true': %v, %v", isempty1, isempty2)
	}
}

func Test_Copy(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	c := sample.Copy()
	if sample.Len() != c.Len() {
		t.Errorf("Should return true but %#v / %#v", sample, c)
	}
}

func Test_Check(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	if !sample.Check("3") {
		t.Errorf("Should return true but %v", sample.Check("3"))
	}
}

func Test_IsEqual(t *testing.T) {
	s1 := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	s2 := Sequence([]interface{}{3, -.9, "B", 1, "A"})
	if !IsEqual(s1, s2) {
		t.Errorf("Should return true but %v", IsEqual(s1, s2))
	}
}

func Test_Find(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	a, b := sample.Find("A")
	if a != 1 && b != true {
		t.Errorf("Find(\"A\") should return '1, true': %#v", sample)
	}
	c, d := sample.Find(-.8)
	if c != 0 && d != false {
		t.Errorf("Find(-.8) should return '0, false': %#v", sample)
	}
}

func Test_PushFront(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	sample.PushFront("Front")
	if sample[0] != "Front" {
		t.Errorf("sample[0] should be 'Front': %#v", sample)
	}
}

func Test_PushBack(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	sample.PushBack("Back")
	if sample[5] != "Back" {
		t.Errorf("sample[5] should be 'Back': %#v", sample)
	}
}

func Test_DeepSliceDelete(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	sample.DeepSliceDelete(2)
	if sample[2] != -0.9 {
		t.Errorf("sample[2] should be '-0.9': %#v", sample)
	}
	for _ = range sample {
		sample.DeepSliceDelete(0)
		// Don't do sample.DeepSliceDelete(k)
		// the slice length decreases at the same time
	}
	if sample.Len() != 0 {
		t.Errorf("Should be empty but: %#v", sample.Len())
	}
}

func Test_FindAndDelete(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	ok := sample.FindAndDelete(3)
	if !ok || sample[2] != -0.9 {
		t.Errorf("Should return true, but %#v, and sample[2] should be '-0.9': %#v", ok, sample)
	}

	// list := sample
	// this does deep copy
	// (they are the same)
	// so we need to use Copy
	list := sample.Copy()
	for _, v := range list {
		sample.FindAndDelete(v)
	}
	l := sample.Len()
	if l != 0 {
		t.Errorf("Should be empty but: %#v / %#v / %#v", l, sample, list)
	}
}

func Test_DeepSliceCut(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	sample.DeepSliceCut(2, 3)
	if sample[2] != "B" {
		t.Errorf("sample[2] should be 'B': %#v", sample)
	}
}

func Test_Front(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	tm := sample.Front()
	if tm != 1 {
		t.Errorf("sample.Front() should be 1: %#v", sample)
	}
}

func Test_Back(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	tm := sample.Back()
	if tm != "B" {
		t.Errorf("sample.Front() should be 'B': %#v", sample)
	}
}

func Test_PopFront(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	tm := sample.PopFront()
	if tm != 1 {
		t.Errorf("sample.PopFront() should return 1: %#v", sample)
	}
	if sample[0] != "A" {
		t.Errorf("sample[0] should be 'A': %#v", sample)
	}
}

func Test_PopBack(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	tm := sample.PopBack()
	if tm != "B" {
		t.Errorf("sample.PopBack() should return 'B': %#v", sample)
	}
	if sample[3] != -0.9 {
		t.Errorf("sample[3] should be -0.9: %#v", sample)
	}
}

func Test_GetElements(t *testing.T) {
	sample := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	slice := sample.GetElements()
	if len(slice) != 5 {
		t.Errorf("len(slice) should return 5: %#v", sample)
	}
	if slice[3] != -0.9 {
		t.Errorf("slice[3] should be -0.9: %#v", sample)
	}
}

func Test_CommonPrefix(t *testing.T) {
	s1 := Sequence([]interface{}{1, "A", 3, -.9, "B", "e", "f", "G"})
	s2 := Sequence([]interface{}{1, "A", 3, -.9, "B"})
	s3 := Sequence([]interface{}{1, "A", 3, -.9, "H", 2, 3, 4})
	s4 := Sequence([]interface{}{1, "A", 3, -.9, "H", 2, 3, 4})
	s5 := Sequence([]interface{}{1, "A", 3, -.9, "B", "e", "f"})
	result := CommonPrefix(s1, s2, s3, s4, s5)
	if len(result) != 4 {
		t.Errorf("len(result) should return 4: %#v", result)
	}
	if result[3] != -0.9 {
		t.Errorf("result[3] should be -0.9: %#v", result)
	}
}

var sm1 = map[string]int64{
	"California":    -10,
	"Japan":         15,
	"Korea":         30,
	"Hello":         -100,
	"USA":           30,
	"San Francisco": 2,
	"Ohio":          56,
	"New York":      14,
	"Los Angeles":   23,
	"Mountain View": 70,
}

/*
Sorted data:
[{Hello -100} {California -10} {San Francisco 2} {New York 14} {Japan 15} {Los Angeles 23} {USA 30} {Korea 30} {Ohio 56} {Mountain View 70}]
*/
func Test_SortMapSIByValue(t *testing.T) {
	result := SortMapSIByValue(sm1)
	tm := MapSI{key: "Ohio", value: 56}
	if result[len(sm1)-2] != tm {
		t.Errorf("result[len(sm1)-2] should return '{Ohio 56}': %#v", sm2)
	}
}

var sm2 = map[string]float64{
	"California":    9.9,
	"Japan":         7.23,
	"Korea":         -.3,
	"Hello":         1.5,
	"USA":           8.4,
	"San Francisco": 8.4,
	"Ohio":          -1.10,
	"New York":      1.23,
	"Los Angeles":   23.1,
	"Mountain View": 9.9,
}

/*
Sorted data:
[{Ohio -1.1} {Korea -0.3} {New York 1.23} {Hello 1.5} {Japan 7.23} {USA 8.4} {San Francisco 8.4} {California 9.9} {Mountain View 9.9} {Los Angeles 23.1}]
*/
func Test_SortMapSFByValue(t *testing.T) {
	result := SortMapSFByValue(sm2)
	tm := MapSF{key: "Ohio", value: -1.1}
	if result[0] != tm {
		t.Errorf("result[0] should return '{Ohio -1.1}': %#v", sm2)
	}
}

func Test_SwitchMap(t *testing.T) {
	// to convert to map[interface{}]interface{}
	temp := make(map[interface{}]interface{})
	for key, value1 := range sm2 {
		temp[key] = value1
	}
	result := SwitchMap(temp)
	value2, exist := result[1.23]
	if value2 != "New York" {
		t.Errorf("result[1.23] should return '\"New York\"': %#v", sm2)
	}
	if exist != true {
		t.Errorf("result[1.23] should exist: %#v", sm2)
	}
}
