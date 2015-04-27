package graph

/*
	This is a class representing the hash version of the Graph.
	A Graph consists of edges connected to each other by verices.

	Author: Stefan Nilsson
	Completed and modified by Ivan Liljeqvist 22-04-2015
*/

const initialMapSize = 4

type Hash struct {
	// The map edges[v] contains the mapping {w:x} if there is an edge
	// from v to w; x is the label assigned to this edge.
	// The maps may be nil and are allocated only when needed.
	edges []map[int]interface{}

	numEdges int // total number of directed edges in the graph
}

// NewList constructs a new graph with n vertices and no edges.
func NewHash(n int) *Hash {
	return &Hash{edges: make([]map[int]interface{}, n)}
}

// NumVertices returns the number of vertices in this graph.
// Time complexity: O(1).
func (g *Hash) NumVertices() int {
	return len(g.edges)
}

// NumEdges returns the number of (directed) edges in this graph.
// Time complexity: O(1).
func (g *Hash) NumEdges() int {
	return g.numEdges
}

// Degree returns the degree of vertex v. Time complexity: O(1).
func (g *Hash) Degree(v int) int {
	//the degree is how many neighbours the node v has

	neighbours_of_v := g.edges[v]
	num_of_neighbours_of_v := len(neighbours_of_v)

	return num_of_neighbours_of_v
}

// DoNeighbors calls action for each neighbor w of v,
// with x equal to the label of the edge from v to w.
// Time complexity: O(m), where m is the number of neighbors.
func (g *Hash) DoNeighbors(v int, action func(w int, x interface{})) {

	//first we need to get all the neighbours w of v
	//we get the neighbours by looking in the edges
	w_plural := g.edges[v]

	//w_plural is a map with mapping {w:label}
	//we take out w and label and call action
	for w, label := range w_plural {
		action(w, label)
	}

}

// HasEdge returns true if there is an edge from v to w.
// Time complexity: O(1).
func (g *Hash) HasEdge(v, w int) bool {

	//get all the neighbours of v
	neighbours_v := g.edges[v]

	for neighbour, _ := range neighbours_v {
		//if we have a match - then w and v are connected
		if neighbour == w {
			return true
		}
	}

	//we've looped through all the neighbours and found no match
	return false
}

// Returns the label for the edge from v to w, NoLabel if the edge has no label,
// or nil if no such edge exists.
// Time complexity: O(1).
func (g *Hash) Label(v, w int) interface{} {

	//neighbours_v is a map with mapping {w:label}
	neighbours_v := g.edges[v]

	//check if map contains value and return
	if label, hasValue := neighbours_v[w]; hasValue {

		//this will be NoLabel if nobody has inserted a label.
		//NoLabel is by default inserted in the Add method
		return label

	} else {
		return nil
	}

}

// Add inserts a directed edge.
// It removes any previous label if this edge already exists.
// Time complexity: O(1).
func (g *Hash) Add(from, to int) {

	neighbours_from := g.edges[from]

	//if the map doesnt contain value
	if _, hasValue := neighbours_from[to]; !hasValue {
		//increase edges
		g.numEdges += 1
	}

	//if no neighbours - init the slice
	if neighbours_from == nil {
		neighbours_from = make(map[int]interface{}, initialMapSize)
		//save to the edges
		g.edges[from] = neighbours_from
	}

	//add connection from 'from' to 'to' with empty label 'NoLabel'
	g.edges[from][to] = NoLabel

}

// AddLabel inserts a directed edge with label x.
// It overwrites any previous label if this edge already exists.
// Time complexity: O(1).
func (g *Hash) AddLabel(from, to int, x interface{}) {
	m := g.edges[from]
	if m == nil {
		m = make(map[int]interface{}, initialMapSize)
		g.edges[from] = m
	}
	if _, ok := m[to]; !ok {
		g.numEdges++
	}
	m[to] = x
}

// AddBi inserts edges between v and w.
// It removes any previous labels if these edges already exists.
// Time complexity: O(1).
func (g *Hash) AddBi(v, w int) {
	//use add function to add edges in both directions
	g.Add(w, v)
	g.Add(v, w)
}

// AddBiLabel inserts edges with label x between v and w.
// It overwrites any previous labels if these edges already exists.
// Time complexity: O(1).
func (g *Hash) AddBiLabel(v, w int, x interface{}) {

	//use addLabel function to add labels in both directions
	g.AddLabel(w, v, x)
	g.AddLabel(v, w, x)

}

// Remove removes an edge. Time complexity: O(1).
func (g *Hash) Remove(from, to int) {

	//check if the edge exists
	if _, hasEdge := g.edges[from][to]; hasEdge {
		//if it exists - remove
		g.numEdges -= 1
		delete(g.edges[from], to)
	}

}

// RemoveBi removes all edges between v and w. Time complexity: O(1).
func (g *Hash) RemoveBi(v, w int) {

	//use the Remove function above to delete edges in both directions

	g.Remove(w, v)
	g.Remove(v, w)

}
