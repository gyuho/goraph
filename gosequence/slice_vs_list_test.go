// go test -bench=. -cpu=1,2,4
package gosequence

import (
	"container/list"
	"fmt"
	"testing"
)

func Benchmark_Sequence_Find(b *testing.B) {
	s := NewSequence()
	for i := 0; i < 300000; i++ {
		s.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		s.Find(150000)
	}
}

func Benchmark_Linked_List_Find(b *testing.B) {
	l := list.New()
	for i := 0; i < 300000; i++ {
		l.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		for elem := l.Front(); elem != nil; elem = elem.Next() {
			if fmt.Sprintf("%v", elem.Value) == fmt.Sprintf("%v", 150000) {
				// Done
				break
			}
		}
	}
}

/*
Not much difference to Find
but if we know the index, Find of Sequence is much faster

Benchmark_Sequence_Find	      10	 106576858 ns/op
Benchmark_Sequence_Find-2	      10	 107994465 ns/op
Benchmark_Sequence_Find-4	      10	 109859550 ns/op
Benchmark_Linked_List_Find	      10	 114807781 ns/op
Benchmark_Linked_List_Find-2	      10	 114894408 ns/op
Benchmark_Linked_List_Find-4	      10	 114953431 ns/op
*/
