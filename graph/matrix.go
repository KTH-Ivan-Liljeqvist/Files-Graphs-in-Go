package graph

type noEdgeType struct{}

var noEdge noEdgeType

type Matrix struct {
	// Adjaceny matrix. adj[v][w] is noEdge if v is not adjacent to w,
	// otherwise adj[v][w] is the label of the edge, or NoLabel if the
	// the edge has no label.
	adj [][]interface{}

	numEdges int // total number of directed edges in the graph
}

// NewMatrix constructs a new graph with n vertices and no edges.
func NewMatrix(n int) *Matrix {
	adj := make([][]interface{}, n)
	for i, _ := range adj {
		v := make([]interface{}, n)
		adj[i] = v
		for i, _ := range v {
			v[i] = noEdge
		}
	}
	return &Matrix{adj: adj}
}

// NumVertices returns the number of vertices in this graph.
// Time complexity: O(1).
func (g *Matrix) NumVertices() int {
	return len(g.adj)
}

// NumEdges returns the number of (directed) edges in this graph.
// Time complexity: O(1).
func (g *Matrix) NumEdges() int {
	return g.numEdges
}

// Degree returns the degree of vertex v.
// Time complexity: O(n), where n is the number of vertices.
func (g *Matrix) Degree(v int) int {
	d := 0
	for _, x := range g.adj[v] {
		if x != noEdge {
			d++
		}
	}
	return d
}

// DoNeighbors calls action for each neighbor w of v,
// with x equal to the label of the edge from v to w.
// Time complexity: O(n), where n is the number of vertices.
func (g *Matrix) DoNeighbors(v int, action func(w int, x interface{})) {
	for w, x := range g.adj[v] {
		if x != noEdge {
			action(w, x)
		}
	}
}

// HasEdge returns true if there is an edge from v to w.
// Time complexity: O(1).
func (g *Matrix) HasEdge(v, w int) bool {
	return g.adj[v][w] != noEdge
}

// Returns the label for the edge from v to w, NoLabel if the edge has no label,
// or nil if no such edge exists.
// Time complexity: O(1).
func (g *Matrix) Label(v, w int) interface{} {
	x := g.adj[v][w]
	if x == noEdge {
		return nil
	}
	return x
}

// Add inserts a directed edge.
// It removes any previous label if this edge already exists.
// Time complexity: O(1).
func (g *Matrix) Add(from, to int) {
	g.AddLabel(from, to, NoLabel)
}

// AddLabel inserts a directed edge with label x.
// It overwrites any previous label if this edge already exists.
// Time complexity: O(1).
func (g *Matrix) AddLabel(from, to int, x interface{}) {
	if g.adj[from][to] == noEdge {
		g.numEdges++
	}
	g.adj[from][to] = x
}

// AddBi inserts edges between v and w.
// It removes any previous labels if these edges already exists.
// Time complexity: O(1).
func (g *Matrix) AddBi(v, w int) {
	g.AddBiLabel(v, w, NoLabel)
}

// AddBiLabel inserts edges with label x between v and w.
// It overwrites any previous labels if these edges already exists.
// Time complexity: O(1).
func (g *Matrix) AddBiLabel(v, w int, x interface{}) {
	g.AddLabel(v, w, x)
	if v != w {
		g.AddLabel(w, v, x)
	}
}

// Remove removes an edge. Time complexity: O(1).
func (g *Matrix) Remove(from, to int) {
	x := g.adj[from][to]
	if x == noEdge {
		return
	}
	g.adj[from][to] = noEdge
	g.numEdges--
}

// RemoveBi removes all edges between v and w. Time complexity: O(1).
func (g *Matrix) RemoveBi(v, w int) {
	g.Remove(v, w)
	if v != w {
		g.Remove(w, v)
	}
}
