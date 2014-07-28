package fdfk

import (
	"github.com/gyuho/goraph/graph/gt"
)

const (
	wHITE    int     = 0
	gRAY     int     = 1
	bLACK    int     = 2
	maxNODES int     = 1000
	oo       float64 = 1000000000.0
)

// MaxFlow returns the Maximum Network Flow.
func MaxFlow(g *gt.Graph, src, dst string) float64 {
	size := g.GetVerticesSize()
	// esz := g.GetEdgesSize()
	flow := [maxNODES][maxNODES]float64{}
	capacity := g.Matrix
	color := [maxNODES]int{}
	pred := [maxNODES]int{}

	head, tail := 0, 0
	q := [maxNODES + 2]int{}
	enqueue := func(x int) {
		q[tail] = x
		tail++
		color[x] = gRAY
	}
	dequeue := func() int {
		x := q[head]
		head++
		color[x] = bLACK
		return x
	}

	bfs := func(start, target int) bool {
		u, v := 0, 0
		for u = 0; u < size; u++ {
			color[u] = wHITE
		}
		head, tail = 0, 0
		enqueue(start)
		pred[start] = -1
		for head != tail {
			u := dequeue()
			for v = 0; v < size; v++ {
				if color[v] == wHITE && capacity[u][v]-flow[u][v] > 0 {
					enqueue(v)
					pred[v] = u
				}
			}
		}
		return color[target] == bLACK
	}

	srcIdx := g.Index[src]
	dstIdx := g.Index[dst]
	i, j := 0, 0
	maxflow := 0.0
	for i = 0; i < size; i++ {
		for j = 0; j < size; j++ {
			flow[i][j] = 0.0
		}
	}

	for bfs(srcIdx, dstIdx) {
		increment := oo
		for u := size - 1; pred[u] >= 0; u = pred[u] {
			increment = min(increment, capacity[pred[u]][u]-flow[pred[u]][u])
		}
		for u := size - 1; pred[u] >= 0; u = pred[u] {
			flow[pred[u]][u] += increment
			flow[u][pred[u]] -= increment
		}
		maxflow += increment
	}
	return maxflow
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
