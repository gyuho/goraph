// Package benchmark is to benchmark the performance
// of different types of graphs, to be run with the command
// go test -bench=. -cpu=1,2,4
package benchmark

import (
	"strconv"
	"testing"

	"github.com/gyuho/goraph/graph/gl"
	"github.com/gyuho/goraph/graph/gld"
	"github.com/gyuho/goraph/graph/gm"
	"github.com/gyuho/goraph/graph/gs"
	"github.com/gyuho/goraph/graph/gsd"
)

func Benchmark_FindVertexByID_gl(b *testing.B) {
	g := gl.NewGraph()
	for i := 0; i < 300; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		g.FindVertexByID(150)
	}
}

func Benchmark_FindVertexByID_gld(b *testing.B) {
	g := gld.NewGraph()
	for i := 0; i < 300; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		g.FindVertexByID(150)
	}
}

func Benchmark_FindVertexByID_gm(b *testing.B) {
	g := gm.NewGraph()
	for i := 0; i < 300; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		_ = g.Vertices["150"]
	}
}

func Benchmark_FindVertexByID_gs(b *testing.B) {
	g := gs.NewGraph()
	for i := 0; i < 300; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		g.FindVertexByID(150)
	}
}

func Benchmark_FindVertexByID_gsd(b *testing.B) {
	g := gsd.NewGraph()
	for i := 0; i < 300; i++ {
		s := strconv.FormatInt(int64(i), 10)
		g.CreateAndAddToGraph(s)
	}
	for i := 0; i < b.N; i++ {
		g.FindVertexByID(150)
	}
}

/*
go test -bench=. -cpu=1,2,4

Finding a vertex performs best when using map structure.

But if we want to allow(or draw) duplicate edges
which means there are more than one edges
that goes from one to the other, we need to use slice
, which still performs the second best.

Benchmark_FindVertexByID_gl	   20000	    102064 ns/op
Benchmark_FindVertexByID_gl-2	   20000	    116364 ns/op
Benchmark_FindVertexByID_gl-4	   20000	    107320 ns/op
Benchmark_FindVertexByID_gld	   20000	    100255 ns/op
Benchmark_FindVertexByID_gld-2	   20000	    114859 ns/op
Benchmark_FindVertexByID_gld-4	   20000	    109563 ns/op
Benchmark_FindVertexByID_gm	50000000	        30.4 ns/op
Benchmark_FindVertexByID_gm-2	50000000	        40.0 ns/op
Benchmark_FindVertexByID_gm-4	50000000	        38.8 ns/op
Benchmark_FindVertexByID_gs	   20000	    101181 ns/op
Benchmark_FindVertexByID_gs-2	   20000	    110709 ns/op
Benchmark_FindVertexByID_gs-4	   20000	    107284 ns/op
Benchmark_FindVertexByID_gsd	   20000	    101286 ns/op
Benchmark_FindVertexByID_gsd-2	   20000	    112706 ns/op
Benchmark_FindVertexByID_gsd-4	   20000	    107608 ns/op
ok  	github.com/gyuho/goraph/benchmark	38.678s
*/
