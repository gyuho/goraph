package bst

import (
	"fmt"
	"testing"
)

func Test_NewTree(test *testing.T) {
	tr := NewTree(10)
	v1 := fmt.Sprintf("%v", tr.Value)
	vc1 := fmt.Sprintf("%v", 10)
	if v1 != vc1 {
		test.Errorf("Should be same but %v, %v", v1, vc1)
	}
}

func Test_Inserts(test *testing.T) {
	tr0 := NewTree(5)
	for i := 0; i < 10; i++ {
		if i != 5 {
			tr0 = tr0.Insert(int64(i))
		}
	}
	check := []int64{5, 4, 3, 2, 1, 10, 4, 3, 2, 1}
	for i := 0; i < 10; i++ {
		if check[i] != tr0.GetSize(int64(i)) {
			test.Errorf("Should return true but %v, %v", i, tr0.GetSize(int64(i)))
		}
	}

	tr := NewTree(5)
	tr.Inserts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if tr.Size != 10 {
		test.Errorf("Should be size of 10 but %+v", tr.Size)
	}

	tr1 := NewTree(10)
	tr1.Inserts(13, 17, 5, 4)

	if tr1.GetSize(10) != 5 {
		test.Errorf("Should return 5 but %+v", tr1.GetSize(5))
	}
	if tr1.GetSize(13) != 2 {
		test.Errorf("Should return 2 but %+v", tr1.GetSize(2))
	}
	if tr1.GetSize(17) != 1 {
		test.Errorf("Should return 1 but %+v", tr1.GetSize(1))
	}
	if tr1.GetSize(5) != 2 {
		test.Errorf("Should return 2 but %+v", tr1.GetSize(2))
	}
	if tr1.GetSize(4) != 1 {
		test.Errorf("Should return 1 but %+v", tr1.GetSize(1))
	}
}

func Test_Insert(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr.Insert(int64(i))
		// tr = tr.Insert(int64(i))
	}
	if tr.GetSize(5) != 10 {
		test.Errorf("Should be size of 10 but %+v", tr.GetSize(5))
	}
	if tr.Size != 10 {
		test.Errorf("Should be size of 10 but %+v", tr.Size)
	}
}

func Test_Find(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	for i := 0; i < 10; i++ {
		if tr.Find(int64(i)) == nil {
			test.Errorf("Should exist but %+v", tr.Find(int64(i)))
		}
	}
	tt := NewTree(5)
	tt.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	fr := tt.Find(4)
	if fr.Left.Value != int64(2) {
		test.Errorf("Should exist but %+v", tt.Find(int64(4)))
	}
}

func Test_GetSize(test *testing.T) {
	tr1 := NewTree(5)
	tr1.Inserts(7, 8, 4, 2, 1, 6, 3)
	if tr1.GetSize(5) != 8 {
		test.Errorf("Size should be 8 but %v", tr1.GetSize(5))
	}
	if tr1.GetSize(4) != 4 {
		test.Errorf("Size should be 4 but %v", tr1.GetSize(4))
	}
	if tr1.GetSize(7) != 3 {
		test.Errorf("Size should be 3 but %v", tr1.GetSize(7))
	}
	if tr1.GetSize(2) != 3 {
		test.Errorf("Size should be 3 but %v", tr1.GetSize(2))
	}
	if tr1.GetSize(6) != 1 {
		test.Errorf("Size should be 1 but %v", tr1.GetSize(6))
	}
	if tr1.GetSize(8) != 1 {
		test.Errorf("Size should be 1 but %v", tr1.GetSize(8))
	}
	if tr1.GetSize(1) != 1 {
		test.Errorf("Size should be 1 but %v", tr1.GetSize(1))
	}
	if tr1.GetSize(3) != 1 {
		test.Errorf("Size should be 1 but %v", tr1.GetSize(3))
	}

	tr2 := NewTree(5)
	tr2.Inserts(7, 8, 4, 2, 1, 6, 3)
	if tr2.Right.Size != 3 {
		test.Errorf("Height should be 3 but %v", tr2.Right.Size)
	}
}

func Test_GetHeight(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3)
	if tr.GetHeight(5) != 3 {
		test.Errorf("Height should be 3 but %v", tr.GetHeight(5))
	}
	if tr.GetHeight(4) != 2 {
		test.Errorf("Height should be 2 but %v", tr.GetHeight(4))
	}
	if tr.GetHeight(7) != 1 {
		test.Errorf("Height should be 1 but %v", tr.GetHeight(7))
	}
	if tr.GetHeight(2) != 1 {
		test.Errorf("Height should be 1 but %v", tr.GetHeight(2))
	}
	if tr.GetHeight(1) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(1))
	}
	if tr.GetHeight(3) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(3))
	}
	if tr.GetHeight(6) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(6))
	}
	if tr.GetHeight(8) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(8))
	}
}

func Test_Parent(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	if tr.Parent(int64(6)).Value != 7 {
		test.Errorf("Parent should be 7 but\n%v", tr.Parent(int64(6)).Value)
	}
	if tr.Parent(int64(8)).Value != 7 {
		test.Errorf("Parent should be 7 but\n%v", tr.Parent(int64(8)).Value)
	}
	if tr.Parent(int64(7)).Value != 5 {
		test.Errorf("Parent should be 5 but\n%v", tr.Parent(int64(7)).Value)
	}
	if tr.Parent(int64(1)).Value != 2 {
		test.Errorf("Parent should be 2 but\n%v", tr.Parent(int64(1)).Value)
	}
	if tr.Parent(int64(5)).Value != 5 {
		test.Errorf("Parent should be 5 but\n%v", tr.Parent(int64(5)).Value)
	}
}

func Test_IsLeaf(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	if tr.IsLeaf(int64(5)) {
		test.Errorf("IsLeaf should return false but\n%v", tr.IsLeaf(int64(5)))
	}
	if !tr.IsLeaf(int64(1)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(1)))
	}
	if !tr.IsLeaf(int64(3)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(3)))
	}
	if !tr.IsLeaf(int64(6)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(6)))
	}
	if !tr.IsLeaf(int64(8)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(8)))
	}
}

func Test_IsRoot(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	if !tr.IsRoot(int64(5)) {
		test.Errorf("IsRoot should return true but\n%v", tr.IsRoot(int64(5)))
	}
	if tr.IsRoot(int64(1)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(1)))
	}
	if tr.IsRoot(int64(3)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(3)))
	}
	if tr.IsRoot(int64(6)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(6)))
	}
	if tr.IsRoot(int64(8)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(8)))
	}
}

func Test_Delete1(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	tr.Delete(int64(6))
	st := tr.Find(int64(7))
	if st.Left != nil {
		test.Errorf("Left should be nil but %+v", st.Left)
	}
	if st.Right.Value != 8 {
		test.Errorf("Right should be 8 but %+v", st)
	}
}

func Test_Delete2(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 3)
	tr.Delete(int64(7))
	st := tr.Find(int64(5))
	if st.Right.Value != 8 {
		test.Errorf("Right should be 8 but %+v", st.Right)
	}
}

func Test_FindMinMax(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	if tr.FindMin() != 1 {
		test.Errorf("Should be 1 but\n%v", tr.FindMin())
	}
	if tr.FindMax() != 8 {
		test.Errorf("Should be 8 but\n%v", tr.FindMax())
	}
}

func Test_Same(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	t2 := NewTree(5)
	for i := 0; i < 10; i++ {
		t2 = t2.Insert(int64(i))
	}
	if !Same(tr, t2) {
		test.Errorf("Should be same but %v, %v", tr, t2)
	}
}

func Test_StringInOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	sc1 := "0 1 2 3 4 5 6 7 8 9 "
	ch1 := make(chan int64)
	s1 := StringInOrder(tr, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringPreOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	sc1 := "5 0 1 2 3 4 6 7 8 9 "
	ch1 := make(chan int64)
	s1 := StringPreOrder(tr, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringPostOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	sc1 := "4 3 2 1 0 9 8 7 6 5 "
	ch1 := make(chan int64)
	s1 := StringPostOrder(tr, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringLevelOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	sc1 := "5 0 6 1 7 2 8 3 9 4 "
	s1 := StringLevelOrder(tr)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_ValuePreOrder(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 3, 4, 2, 1, 6)
	ss := []int64{5, 3, 2, 1, 4, 7, 6, 8}
	ch := make(chan int64)
	slice := ValuePreOrder(tr, ch)
	if !SameSlice(ss, slice) {
		test.Errorf("Should be true but\n%v", slice)
	}
}

func Test_Construct(test *testing.T) {
	slice := []int64{5, 7, 8, 3, 4, 2, 1, 6}
	tr := Construct(5, slice)
	ss := []int64{7, 3, 2, 1, 4, 6, 8}
	ch := make(chan int64)
	vs := ValuePreOrder(tr, ch)
	if !SameSlice(ss, vs) {
		test.Errorf("Should be true but\n%v", vs)
	}
}
