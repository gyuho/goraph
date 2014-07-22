package parsex

import "testing"

func TestUniqElemStrANDEqualSliceElem(t *testing.T) {
	slice := []string{"A", "B", "C", "D", "E", "E", "E", "E"}
	// fmt.Println(UniqElemStr(slice))
	// [A B C D E]
	sl := []string{"A", "B", "C", "D", "E"}
	if !EqualSliceElem(sl, UniqElemStr(slice)) {
		t.Errorf("expected true but %v", UniqElemStr(slice))
	}
}

func TestCheckStr(t *testing.T) {
	slice := []string{"A", "B", "C", "D", "E", "E", "E", "E"}
	if !CheckStr("A", slice) {
		t.Errorf("expected true but %v", CheckStr("A", slice))
	}
	if CheckStr("F", slice) {
		t.Errorf("expected false but %v", CheckStr("F", slice))
	}
}

func TestStrToFloat64(t *testing.T) {
	num1 := 123.123
	num2 := StrToFloat64("123.123")
	if num1 != num2 {
		t.Fatalf("expected\n%v\n%v", num1, num2)
	}
}
