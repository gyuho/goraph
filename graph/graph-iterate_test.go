package graph

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gyuho/goraph/graph/gl"
	"github.com/gyuho/goraph/graph/gm"
	"github.com/gyuho/goraph/graph/gs"
)

func Benchmark_FindVertexByID_gl(b *testing.B) {
	g := gl.NewGraph()
	for i := 0; i < 5000; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		g.FindVertexByID(fmt.Sprintf("%v", 4500))
	}
}

func Benchmark_FindVertexByID_gs(b *testing.B) {
	g := gs.NewGraph()
	for i := 0; i < 5000; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		g.FindVertexByID(fmt.Sprintf("%v", 4500))
	}
}

func Benchmark_FindVertexByID_gm(b *testing.B) {
	g := gm.NewGraph()
	for i := 0; i < 5000; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		_ = g.Vertices["4500"]
	}
}

/*
Result:
Finding a vertex performs best when using map structure.
Note that slice performs better than `container/list` (linked list).


$ go test -bench=. -cpu=1,2,4
testing: warning: no tests to run
PASS
Benchmark_FindVertexByID_gl	       1	9967236000 ns/op
Benchmark_FindVertexByID_gl-2	       1	9833747000 ns/op
Benchmark_FindVertexByID_gl-4	       1	9943510000 ns/op
Benchmark_FindVertexByID_gs	   20000	     56835 ns/op
Benchmark_FindVertexByID_gs-2	   20000	     57174 ns/op
Benchmark_FindVertexByID_gs-4	   20000	     57121 ns/op
Benchmark_FindVertexByID_gm	100000000	        23.9 ns/op
Benchmark_FindVertexByID_gm-2	100000000	        23.9 ns/op
Benchmark_FindVertexByID_gm-4	100000000	        24.0 ns/op
ok  	github.com/gyuho/goraph/benchmark	42.944s
*/
