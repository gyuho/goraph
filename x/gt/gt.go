package gt

import (
	"strconv"

	"github.com/gyuho/goraph/parsex/jsonx"
)

// Graph is a graph represented in adjacency matrix.
type Graph struct {
	Size   *int
	Matrix [][]float64

	// to look up the index of vertex
	Index map[string]int

	// to look up the string label from index
	Label map[int]string
}

// NewGraph returns a new graph in adjacency matrix.
// It returns the adjacency matrix of the size * size.
func NewGraph(size int) *Graph {
	make2DSlice := func(row, column int) [][]float64 {
		mat := make([][]float64, row)
		// for i := 0; i < row; i++ {
		for i := range mat {
			mat[i] = make([]float64, column)
		}
		return mat
	}
	mat := make2DSlice(size, size)

	// to initialize
	for r := range mat {
		for c := range mat[r] {
			mat[r][c] = 0.0
		}
	}
	lm := make(map[string]int)
	ls := make(map[int]string)
	return &Graph{&size, mat, lm, ls}
}

// GetVertices returns the slice of vertices.
func (g Graph) GetVertices() []string {
	result := []string{}
	for k := range g.Index {
		result = append(result, k)
	}
	return result
}

// GetVerticesSize returns the size of vertices.
func (g Graph) GetVerticesSize() int {
	return *g.Size
}

// GetEdgesSize returns the size of edges.
func (g Graph) GetEdgesSize() int {
	size := g.Size
	cn := 0
	// traverse the 2D slice
	for r := 0; r < *size; r++ {
		for c := 0; c < *size; c++ {
			if g.Matrix[r][c] != 0.0 {
				cn++
			}
		}
	}
	return cn
}

// GetEdgeWeight returns the value of the matrix
// which is the edge weight value from A to B.
func (g Graph) GetEdgeWeight(A, B string) float64 {
	a, _ := g.Index[A]
	b, _ := g.Index[B]
	return g.Matrix[a][b]
}

// AddEdge adds an edge with the weight value.
// The edge goes from A to B, not B to A.
func (g *Graph) AddEdge(A, B string, val float64) {
	a, _ := g.Index[A]
	b, _ := g.Index[B]
	g.Matrix[a][b] = val
}

// RemoveEdge removes an edge from A to B, not B to A.
func (g *Graph) RemoveEdge(A, B string) {
	a, _ := g.Index[A]
	b, _ := g.Index[B]
	g.Matrix[a][b] = 0.0
}

// HasEdge returns true if there is an edge
// from A to B.
func (g Graph) HasEdge(A, B string) bool {
	a, ok1 := g.Index[A]
	if !ok1 {
		// panic("The vertex does not exist in the graph.")
		return false
	}
	b, ok2 := g.Index[B]
	if !ok2 {
		// panic("The vertex does not exist in the graph.")
		return false
	}
	if g.Matrix[a][b] == 0.0 {
		return false
	}
	return true
}

// GetInVertices returns the vertices that goes into A.
func (g Graph) GetInVertices(A string) []string {
	a, ok1 := g.Index[A]
	if !ok1 {
		panic("The vertex does not exist in the graph.")
	}

	result := []string{}
	for i := 0; i < g.GetVerticesSize(); i++ {
		if g.Matrix[i][a] != 0.0 {
			result = append(result, g.Label[i])
		}
	}

	return result
}

// GetOutVertices returns the vertices that goes out of A.
func (g Graph) GetOutVertices(A string) []string {
	a, ok1 := g.Index[A]
	if !ok1 {
		panic("The vertex does not exist in the graph.")
	}

	result := []string{}
	for i := 0; i < g.GetVerticesSize(); i++ {
		if g.Matrix[a][i] != 0.0 {
			result = append(result, g.Label[i])
		}
	}

	return result
}

// Initialize initializes the elements of matrix with num.
func (g *Graph) Initialize(num float64) {
	mat := g.Matrix
	for r := range mat {
		for c := range mat[r] {
			mat[r][c] = num
		}
	}
}

// OutputMatrix outputs the matrix of Graph in string format.
func (g Graph) OutputMatrix() string {
	return Output2DSlice(g.Matrix)
}

// Output2DSlice outputs the matrix in string format.
func Output2DSlice(mat [][]float64) string {
	s := ""
	for r := range mat {
		for c := range mat[r] {
			// to exclude trailing white space chracter
			if c != len(mat[r])-1 {
				if mat[r][c] < 10 {
					s += "  " + strconv.FormatFloat(mat[r][c], 'f', -1, 64) + " "
				} else if mat[r][c] < 100 {
					s += " " + strconv.FormatFloat(mat[r][c], 'f', -1, 64) + " "
				} else {
					s += strconv.FormatFloat(mat[r][c], 'f', -1, 64) + " "
				}
			} else {
				if mat[r][c] < 10 {
					s += "  " + strconv.FormatFloat(mat[r][c], 'f', -1, 64)
				} else if mat[r][c] < 100 {
					s += " " + strconv.FormatFloat(mat[r][c], 'f', -1, 64)
				} else {
					s += strconv.FormatFloat(mat[r][c], 'f', -1, 64)
				}
			}
		}
		s += "\n"
	}
	return s
}

// FromJSON parses JSON file to a graph.
func FromJSON(fpath, gname string) *Graph {
	nodes := jsonx.GetNodes(fpath, gname)
	gmap := jsonx.GetGraphMap(fpath, gname)
	// map[string]map[string][]float64

	g := NewGraph(len(nodes))
	*g.Size = len(nodes)

	cn := 0
	for _, node := range nodes {
		g.Index[node] = cn
		g.Label[cn] = node
		cn++
	}

	for _, srcID := range nodes {
		for dstID := range gmap[srcID] {
			for _, weight := range gmap[srcID][dstID] {
				g.AddEdge(srcID, dstID, weight)
			}
		}
	}
	return g
}
