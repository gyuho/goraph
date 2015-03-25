// Package radix implements a radix tree(trie). Trie(reTRIEval) is also called a digital or prefix tree:
//	* Ordered tree data structure to store an associative array.
//	* If the radix r of the radix trie, the radix trie is binary.
//	* Binary radix trie minimizes the sparseness at the expense of maximizing the trie depth.
//	* If r is greater than 2, then the radix trie becomes an r-ary trie.
//	* Values exist in leaves. Unlike regular trees.
//	* In a trie, all comparisons require constant time.
//	* Unlike balanced trees, the time complexities of lookup, insertion, and deletion are O(k), when k is the number of members.
//	* Although constant, k is normally greater than log n.
//	* This can be slow with long common prefixes.
//
// Reference:
//	* http://en.wikipedia.org/wiki/Radix_tree
//
package radix // import "github.com/gyuho/goraph/radix"
