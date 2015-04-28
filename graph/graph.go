// Package graph implements datastructures for a graph with a fixed number of vertices.
//
// The vertices are numbered from 0 to n-1.
// Edges may be added or removed from the graph.
// Each edge may have an associated label of interface{} type.
//
// graph.Hash is best suited for sparse graphs.
// The edges are represented by adjacency lists implemented as hash maps.
// Hence, space complexity is Θ(n+m), where n and m are the number of
// vertices and edges.
//
// graph.Matrix is best suited for dense graphs.
// The edges are represented by an adjacency matrix.
// Hence, space complexity is Θ(n*n), where n is the number of vertices.

/*
	Modified by Ivan Liljeqvist 28-04-2014.

	I needed the 'action' function passed in as parameter to BFS
	to have a 'from' parameter to track the parent of the node 'w'.

	I didn't see any restrictions in the task telling that it's not allowed
	to alter graph.go and therefore I changed the 'action' function.

*/

package graph

// NoLabel represents an edge with no label.
var NoLabel noLabel

type noLabel struct{} // a type with only one value

func (x noLabel) String() string { return "NoLabel" }

type Iterator interface {
	// NumVertices returns the number of vertices.
	NumVertices() int

	// DoNeighbors calls action for each neighbor w of v,
	// with x equal to the label of the edge from v to w.
	DoNeighbors(v int, action func(from, w int, x interface{}))
}

// BFS traverses the vertices of g that have not yet been visited
// in breath-first order starting at v.
// The visited array keeps track of visited vertices.
// When the algorithm arrives at a node w for which visited[w] is false,
// action(w) is called and visited[w] is set to true.
func BFS(g Iterator, v int, visited []bool, action func(from, w int)) {
	traverse(g, v, visited, action, bfs)
}

// DFS traverses the vertices of g that have not yet been visited
// in depth-first order starting at v.
// The visited array keeps track of visited vertices.
// When the algorithm arrives at a node w for which visited[w] is false,
// action(w) is called and visited[w] is set to true.
func DFS(g Iterator, v int, visited []bool, action func(from, w int)) {
	traverse(g, v, visited, action, dfs)
}

const (
	bfs = iota
	dfs
)

func traverse(g Iterator, v int, visited []bool, action func(from, w int), order int) {

	const INDEX_HAS_NO_PARENT = -1

	var queue []int

	if visited[v] {
		return
	}
	visit(INDEX_HAS_NO_PARENT, v, &queue, visited, action)
	for len(queue) > 0 {
		switch order {
		case bfs: // pop from fifo queue
			v, queue = queue[0], queue[1:]
		case dfs: // pop from stack
			i := len(queue) - 1
			v, queue = queue[i], queue[:i]
		}
		g.DoNeighbors(v, func(v, w int, _ interface{}) {
			if !visited[w] {
				visit(v, w, &queue, visited, action)
			}
		})
	}
}

func visit(from, v int, queue *[]int, visited []bool, action func(from, w int)) {
	visited[v] = true
	action(from, v)
	*queue = append(*queue, v)
}
