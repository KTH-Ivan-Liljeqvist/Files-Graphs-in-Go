package graph_test

import (
	graph "."
	"fmt"
)

func ExampleDFS() {

	g := graph.NewMatrix(5)

	g.AddBi(1, 2)
	g.AddBi(2, 3)
	g.AddBi(4, 4)

	// Print the vertices of each component of g in separate groups.
	state := make([]bool, g.NumVertices())
	for v, visited := range state {
		if !visited {
			fmt.Print("{")
			graph.DFS(g, v, state, func(w int) { fmt.Print(w) })
			fmt.Print("}")
		}
	}
	// Output: {0}{123}{4}
}
