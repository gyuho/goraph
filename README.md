goraph [![Build Status](https://travis-ci.org/gyuho/goraph.svg?branch=master)](https://travis-ci.org/gyuho/goraph) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/gyuho/goraph)
==========

Package goraph implements graph data structures and algorithms with visualization.


<kbd>goraph</kbd> is a set of packages for graph database analytics in Go, with handcrafted testing data and visualizations. It was my first coding project written from scratch. This is an ongoing project with no time-line.

For fast query and retrieval, please check out  <a href="http://google-opensource.blogspot.co.uk/2014/06/cayley-graphs-in-go.html" target="_blank">Cayley</a>.



## YouTube Tutorial by <kbd>gyuho</kbd>


<a href="http://www.youtube.com/watch?v=ImMnYq2zP4Y" target="_blank"><img src="http://img.youtube.com/vi/ImMnYq2zP4Y/0.jpg"></a>

- <a href="https://www.youtube.com/channel/UCWzSgIp_DYRQnEsJuH32Fww" target="_blank">Please visit my YouTube Channel</a>
- <a href="https://www.youtube.com/watch?v=NdfIfxTsVDo&list=PLT6aABhFfinvsSn1H195JLuHaXNS6UVhf" target="_blank"><kbd>Tree</kbd>, <kbd>Graph</kbd> Theory Algorithms (Playlist)</a>
- <a href="https://www.youtube.com/watch?v=ImMnYq2zP4Y&list=PLT6aABhFfinvsSn1H195JLuHaXNS6UVhf&index=4" target="_blank"><kbd>Graph</kbd> : BFS, DFS</a>


<hr><b>Tree (a graph <kbd>G</kbd> with <kbd>V</kbd> vertices)  if and only if it satisfies any of the following 5 conditions:</b>
<ul><li>G has V-1 edges and no cycles</li><li>G has V-1 edges and is connected</li><li>G is connected, and removing any edge disconnects the - G</li><li>G is acyclic, and adding any edge creates a cycle in G</li><li>Exactly one simple path connects each pair of vertices in G</li></ul><hr><b>Degree of a vertex(node):</b> number of edges adjacent to the vertex(loop counts as 2)
<b>Adjacency List vs. Adjacency Matrix</b>
<ul><li>When Graph <kbd>G = (V, E)</kbd> = (Vertex, Edge)</li><li>|V| = # of nodes(vertices), |E| = # of edges</li><li><b>Sparse graph:</b> |E| is much less than |V|^2, Relatively few edges present</li><li><b>Dense graph:</b> |E| is close to |V|^2, Relatively few edges missing</li><li><kbd>Adjacency List</kbd> is good for <b>Sparse graph</b></li><li><kbd>Adjacency List</kbd> uses memory in proportion to |E| so fast to iterate with fewer |E| (but slower for lookup)</li><li><kbd>Adjacency Matrix</kbd> is good for <b>Dense graph</b></li><li><kbd>Adjacency Matrix</kbd> uses O(|V|^2) memory so fast lookup of edge presence but slow to iterate</li></ul>




<i>README.md Updated at 2015-03-14 21:52:53</i>
