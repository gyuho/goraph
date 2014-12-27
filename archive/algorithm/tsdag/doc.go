// Package tsdag finds the topological sort.
package tsdag

/*
http://en.wikipedia.org/wiki/Topological_sorting

L ← Empty list that will contain the sorted nodes

while there are unmarked nodes do
  select an unmarked node n
	visit(n)

function visit(node n)
	if n has a temporary mark then stop (not a DAG)
	if n is not marked (i.e. has not been visited yet) then
	  mark n temporarily
	  for each node m with an edge from n to m do
	    visit(m)
		mark n permanently
		add n to head of L
*/
