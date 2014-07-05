goraph [![Build Status](https://travis-ci.org/gyuho/goraph.svg?branch=master)](https://travis-ci.org/gyuho/goraph) [![GoDoc](https://godoc.org/github.com/gyuho/goraph?status.png)](http://godoc.org/github.com/gyuho/goraph) [![Project Stats](http://www.ohloh.net/p/714468/widgets/project_thin_badge.gif)](http://www.ohloh.net/p/714468)
==========

Package goraph provides graph visualizing tools and algorithm implementations.

- [Getting Started](https://github.com/gyuho/goraph#getting-started)
- [Package Hierarchy](https://github.com/gyuho/goraph#package-hierarchy)
- [Example](https://github.com/gyuho/goraph#example)
- [Testing Graphs](https://github.com/gyuho/goraph#testing-graphs)
- [List(Linked List) vs. Slice(Array)](https://github.com/gyuho/goraph#listlinked-list-vs-slicearray)
- [What is Graph? (YouTube Clips)](https://github.com/gyuho/goraph#what-is-graph-youtube-clips)
- [Adjacency List vs. Adjacency Matrix](https://github.com/gyuho/goraph#adjacency-list-vs-adjacency-matrix)
- [Channel](https://github.com/gyuho/goraph#channel)
- [C++ Version](https://github.com/gyuho/goraph#c-version)
- [package gotree](https://github.com/gyuho/goraph#package-gotree)


Getting Started
==========
- [godoc.org](http://godoc.org/github.com/gyuho/goraph)
- [gowalker.org](http://gowalker.org/github.com/gyuho/goraph#_index)

```go
go get github.com/gyuho/goraph
```
<img src="./example_files/sample.gif" alt="sample" width="450px" height="350px"/>

<img src="./example_files/testgraph.004_spd.png" alt="testgraph.004_spd"/>


[↑ top](https://github.com/gyuho/goraph#goraph---)


Package Hierarchy
==========
```go
goraph/							# Top Package
	
	algorithm/					# Package: Graph Algorithms
		bfs/					# Package: Breadth First Search Algorithm
		dfs/					# Package: Depth First Search Algorithm
		sp/						# Package: Shortest Path Algorithm (Dijkstra, Bellman-Ford)
		spbf/					# Package: Shortest Path Algorithm for Negative Edges (Bellman-Ford)
		spd/					# Package: Shortest Path Algorithm for Positive Edges (Dijkstra)
		spfw/					# Package: Shortest Path Algorithm for Positive Edges (Floyd-Warshall)
		tsdag/					# Package: Topological Sort, Detects whether it is a DAG
		tsdfs/					# Package: Topological Sort using DFS, Not Detecting DAG
		tskahn/					# Package: Topological Sort by Arthur Kahn(1962), Detects DAG
		mst/					# Package: Minimum Spanning Tree (Kruskal, Prim)
			kruskal/			# Package: Kruskal Minimum Spanning Tree Algorithm
			prim/				# Package: Prim Minimum Spanning Tree Algorithm
		scc/					# Package: Strongly Connected Component
			tarjan/				# Package: Tarjan Strongly Connected Component Algorithm
			kosaraju/			# Package: Kosaraju Strongly Connected Component Algorithm
		maxflow/				# Package: Maximum Network Flow
			fdfk/				# Package: Ford-Fulkerson Maximum Network Flow Algorithm

	benchmark/					# Package: Benchmark, Comparison of graph representations

	example_files/				# Learn DOT
	example_files/				# Files for Example
	example_with_script/		# Example Script (Program)
	example_with_testing/		# Example Code

	goroup/						# Package: Disjoint Set for Graph Nodes
		gsdset/					# Package: Set Operation with the package graph/gsd

	graph/						# Package: Graph Data Structure
		gl/						# Package: Adjacency List, Linked List(container/list), No Duplicate Edges
		gld/					# Package: Adjacency List, Linked List(container/list), Allow Duplicate Edges
		gm/						# Package: Adjacency List, Map, No Duplicate Edges
		gs/						# Package: Adjacency List, Slice, No Duplicate Edges
		gsd/					# Package: Adjacency List, Slice, Allow Duplicate Edges
		gsdflow/				# Package: Customized gsd for Network Flow Algorithm
		gt/						# Package: Adjacency Matrix, Map, No Duplicate Edges

	viz/						# Package: Graph Visualization (Graphviz)
		dot/					# Package: Convert JSON graph data to DOT
```

##### External Package
- <a href="https://github.com/gyuho/gson" target="_blank"><b>gson</b></a>: JSON Import Package
- <a href="https://github.com/gyuho/gosequence" target="_blank"><b>gosequence</b></a>: Customized Slice(Array) Data Structure
- <a href="http://www.graphviz.org" target="_blank"><b>Graphviz</b></a>: Graph Visualization

[↑ top](https://github.com/gyuho/goraph#goraph---)


Example
==========

##### Shortest Path
```go
fmt.Println("Dijkstra Shortest Path on testgraph10:")
g10 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.010")
fmt.Println(spd.SPD(g10, "A", "E"))
fmt.Println(g10.ShowPrev("E"))
fmt.Println(g10.ShowPrev("D"))
fmt.Println(g10.ShowPrev("B"))
fmt.Println(g10.ShowPrev("C"))
fmt.Println(g10.ShowPrev("A"))
/*
   A(=0) → C(=9) → B(=19) → D(=34) → E(=36)
   Prev of E:  C D
   Prev of D:  B
   Prev of B:  C
   Prev of C:  A
   Prev of A:

   BackTrack keeps adding the Prev vertex
   with the biggest StampD, recursively
   until we reach the source
*/

println()
fmt.Println("Dijkstra Shortest Path on testgraph10o:")
g10o := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.010")
fmt.Println(spd.SPD(g10o, "E", "A"))
fmt.Println(g10o.ShowPrev("A"))
fmt.Println(g10o.ShowPrev("B"))
fmt.Println(g10o.ShowPrev("C"))
fmt.Println(g10o.ShowPrev("F"))
fmt.Println(g10o.ShowPrev("E"))
/*
   E(=0) → F(=9) → C(=11) → B(=21) → A(=22)
   Prev of A:  F B
   Prev of B:  C
   Prev of C:  E F
   Prev of F:  E
   Prev of E:

   BackTrack keeps adding the Prev vertex
   with the biggest StampD, recursively
   until we reach the source
*/
```
[↑ top](https://github.com/gyuho/goraph#goraph---)

```go
func Test_JSON_ShowSPD(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	fmt.Println(ShowSPD(g4, "S", "T", "testgraph.004_spd.dot"))
}
```

<img src="./example_files/testgraph.004_spd.png" alt="testgraph.004_spd"/>

[↑ top](https://github.com/gyuho/goraph#goraph---)

<hr>

##### BFS, DFS
```go
func Test_JSON_ShowBFS(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	fmt.Println(ShowBFS(g4, g4.FindVertexByID("S"), "testgraph.004_bfs.dot"))
}

func Test_JSON_ShowDFS(test *testing.T) {
	g4 := gsd.JSONGraph("../../example_files/testgraph.json", "testgraph.004")
	fmt.Println(ShowDFS(g4, "testgraph.004_dfs.dot"))
}

```
<img src="./example_files/testgraph.004_bfs.png" alt="testgraph.004_bfs"/>
<img src="./example_files/testgraph.004_dfs.png" alt="testgraph.004_dfs"/>

[↑ top](https://github.com/gyuho/goraph#goraph---)

<hr>


##### Minimum Spanning Tree
**Kruskal Algorithm**
```go
func Test_JSON_ShowMST(test *testing.T) {
	g14 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.014")
	ShowMST(g14, "g14mst_kruskal.dot")
}
```
<img src="./example_files/g14mst_kruskal.png" alt="g14mst_kruskal"/>

[↑ top](https://github.com/gyuho/goraph#goraph---)

<hr>

**Prim Algorithm**
```go
func Test_JSON_ShowMST(test *testing.T) {
	g14 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.014")
	ShowMST(g14, "g14mst_prim.dot")
}
```
<img src="./example_files/g14mst_prim.png" alt="g14mst_prim"/>

[↑ top](https://github.com/gyuho/goraph#goraph---)


[Testing Graphs](https://github.com/gyuho/goraph/tree/master/example_files)
==========
<img src="./example_files/testgraph.001.png" alt="testgraph.001" width="260px" height="220px"/>
<img src="./example_files/testgraph.002.png" alt="testgraph.002" width="260px" height="220px"/>
<img src="./example_files/testgraph.003.png" alt="testgraph.003" width="260px" height="220px"/>
<img src="./example_files/testgraph.004.png" alt="testgraph.004" width="260px" height="220px"/>
<img src="./example_files/testgraph.005.png" alt="testgraph.005" width="260px" height="220px"/>
<img src="./example_files/testgraph.006.png" alt="testgraph.006" width="260px" height="220px"/>
<img src="./example_files/testgraph.007.png" alt="testgraph.007" width="260px" height="220px"/>
<img src="./example_files/testgraph.008.png" alt="testgraph.008" width="260px" height="220px"/>
<img src="./example_files/testgraph.009.png" alt="testgraph.009" width="260px" height="220px"/>
<img src="./example_files/testgraph.010.png" alt="testgraph.010" width="260px" height="220px"/>
<img src="./example_files/testgraph.011.png" alt="testgraph.011" width="260px" height="220px"/>
<img src="./example_files/testgraph.012.png" alt="testgraph.012" width="260px" height="220px"/>
<img src="./example_files/testgraph.013.png" alt="testgraph.013" width="260px" height="220px"/>
<img src="./example_files/testgraph.014.png" alt="testgraph.014" width="260px" height="220px"/>
<img src="./example_files/testgraph.015.png" alt="testgraph.015"/>
<img src="./example_files/testgraph.016.png" alt="testgraph.016"/>
<img src="./example_files/testgraph.017.png" alt="testgraph.017"/>

[↑ top](https://github.com/gyuho/goraph#goraph---)


List(Linked List) vs. Slice(Array)
========
Goraph mainly uses customized slice(array) data structure implemented in package gsd, instead of using linked list. To store vertices and edges, we can either use "container/list" or slice(array) data structure. It depends on the number of elements in the lists. Linked list will be more efficient, when we need to do many deletions in the 'middle' of the list. The more elements we have in graph , the less efficient a slice becomes. When the ordering of the elements isn't important, then it is most efficient to use a slice and deleting an element by replacing it by the last element and reslicing to shrink the size(len) by 1. We can mitigate the deletion problem using this slice trick but there is no way to mitigate the slowness of traversing linked list. Both ways are implemented, but mainly this will be run with slice. 

<b>Reference</b>
<ul>
	<li><a href="https://groups.google.com/d/msg/golang-nuts/mPKCoYNwsoU/tLefhE7tQjMJ" target="_blank">Go(Golang) Slice vs. List?</a></li>
	<li><a href="http://www.youtube.com/watch?v=YQs6IC-vgmo" target="_blank">Bjarne Stroustrup: Why you should avoid Linked Lists (C++)</a></li>
	<li><a href="http://www.codeproject.com/Articles/340797/Number-crunching-Why-you-should-never-ever-EVER-us" target="_blank">Why you should never use linked-list</a></li>
</ul>

[↑ top](https://github.com/gyuho/goraph#goraph---)


What is Graph? (YouTube Clips)
========

<a href="http://www.youtube.com/watch?v=s4l_0sXpsBM" target="_blank"><img src="http://img.youtube.com/vi/s4l_0sXpsBM/0.jpg"></a>
<ul>
	<li class="special"><a href="http://www.youtube.com/watch?v=NdfIfxTsVDo&list=PLT6aABhFfinvsSn1H195JLuHaXNS6UVhf" target="_blank">Tree, Heap, Graph (Playlist)</a></li>
</ul>


- **Graph**: Data Structure with Nodes and Edges

- There are various **ways to connect nodes**
	- Doubly Connected Directed Graph (Undirected Graph)
	- Singly Connected Directed Graph.

- **Path**: sequence of vertices connected by edges

- **Simple Path**: path with NO repeated vertices

- **Cycle**: path with at least one edge whose first and last vertices are the same

- **Simple Cycle**: cycle with NO repeated edges or vertices

- **Length of path, cycle**: its number of edges

- **Connectivity**: Graph is connected if there is a path from every vertex to every other vertex

- **Acyclic Graph**: graph with NO cycles

- **Acyclic Connected Graph**: Tree is an Acyclic Connected Graph

- **Forest**: Disjoint Set of Trees (have no vertices in common)

- **Spanning Tree of a Connected Graph**
	-	subgraph that contains all of that graph’s vertices subgraph that is a single tree

- **Spanning Forest of a Graph**
	-	the union of spanning trees of its connected components

- **Spanning Tree of a Connected Graph**
	- subgraph that contains all of that graph’s vertices
subgraph that is a single tree

- **Degree of a Vertex**
	- number of edges incident to the vertex(loop counts as 2).

- **Predecessor of a Vertex**
 - edge(u, v), then vertex v is the descendant of u. Vertex u is the predecessor, or parent/ancestor, of vertex v. v.d is the distance from the source; s.d is 0 when s is the source node. u.d is 1 when the distance from source to u is 1. This is implemented as <b>InVertices</b> in goraph.

[↑ top](https://github.com/gyuho/goraph#goraph---)


Adjacency List vs. Adjacency Matrix
========
- When Graph **G = (V, E)** = (Vertex, Edge)
	- |V| is the number of vertices in a graph
	- |E| is the number of edges in a graph

- **Sparse Graph**
	- |E| is much less than |V|^2
	- Relatively few edges present

- **Dense Graph**
	- |E| is close to |V|^2
	- Relatively few edges missing

- **Adjacency List**: good for Sparse Graph
	- Use memory in proportion to |E|
	- So save memory when G is sparse
	- Fast to iterate over all edges
	- Slightly slower lookup to check for an edge

- **Adjacency Matrix**: good for Dense Graph
	- Use O(n^2) memory
	- Fast lookups to check for presence of an edge
	- Slow to iterate over all edges

[↑ top](https://github.com/gyuho/goraph#goraph---)


Channel
=========
```go
func (v *VertexT) GetEdgeTsChannelFromThisVertex() chan *EdgeT {
	edgechan := make(chan *EdgeT)

	go func() {
		defer close(edgechan)
		for e := v.OutGetEdge().Front(); e != nil; e = e.Next() {
			edgechan <- e.Value.(*EdgeT)
		}
	}()
	return edgechan
}

```
It's not idiomatic Go style to use channels, simply for the ability to iterate over them. It's not efficient, and it can easily lead to an accumulation of idle goroutines: Consider what happens when the caller of GetEdgeTsChannelFromThisVertex discards the channel before reading to the end. It's better to use container/list rather than channel.

[↑ top](https://github.com/gyuho/goraph#goraph---)


C++ Version
=========
I have another Graph Algorithm project written in C++. It is **NOT** maintained anymore, but if interested check out <a href="https://github.com/gyuho/learn/tree/master/learn_cpp/code/cpp_graph_algorithm" target="_blank">HERE</a>.

[↑ top](https://github.com/gyuho/goraph#goraph---)


package gotree
=========
Tree is just another kind of graph data structure. If interested check out <a href="https://github.com/gyuho/gotree" target="_blank">gotree</a>.

[↑ top](https://github.com/gyuho/goraph#goraph---)


Tree
========
Tree (a graph G with V vertices)  if and only if it satisfies any of the following 5 conditions.

- G has V-1 edges and no cycles
- G has V-1 edges and is connected
- G is connected, and removing any edge disconnects the - G
- G is acyclic, and adding any edge creates a cycle in G
- Exactly one simple path connects each pair of vertices in G

[↑ top](https://github.com/gyuho/goraph#goraph---)

