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
	DoNeighbors(v int, action func(w int, x interface{}))
}

// BFS traverses the vertices of g that have not yet been visited
// in breath-first order starting at v.
// The visited array keeps track of visited vertices.
// When the algorithm arrives at a node w for which visited[w] is false,
// action(w) is called and visited[w] is set to true.
func BFS(g Iterator, v int, visited []bool, action func(w int)) {
	traverse(g, v, visited, action, bfs)
}

// DFS traverses the vertices of g that have not yet been visited
// in depth-first order starting at v.
// The visited array keeps track of visited vertices.
// When the algorithm arrives at a node w for which visited[w] is false,
// action(w) is called and visited[w] is set to true.
func DFS(g Iterator, v int, visited []bool, action func(w int)) {
	traverse(g, v, visited, action, dfs)
}

const (
	bfs = iota
	dfs
)

func traverse(g Iterator, v int, visited []bool, action func(w int), order int) {
	var queue []int

	if visited[v] {
		return
	}
	visit(v, &queue, visited, action)
	for len(queue) > 0 {
		switch order {
		case bfs: // pop from fifo queue
			v, queue = queue[0], queue[1:]
		case dfs: // pop from stack
			i := len(queue) - 1
			v, queue = queue[i], queue[:i]
		}
		g.DoNeighbors(v, func(w int, _ interface{}) {
			if !visited[w] {
				visit(w, &queue, visited, action)
			}
		})
	}
}

func visit(v int, queue *[]int, visited []bool, action func(w int)) {
	visited[v] = true
	action(v)
	*queue = append(*queue, v)
}
