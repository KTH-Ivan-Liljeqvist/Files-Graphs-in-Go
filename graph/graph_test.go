package graph_test

/*
	This is a test testing both the amtrix and hash version.
	All of the public methods are tested.

	To run these tests type the command: go test graph_test.go

	Author: Stefan Nilsson
	Small modification by Ivan Liljeqvist 22-02-2015
*/

import (
	. "."
	"fmt"
	"sort"
	"strconv"
	"testing"
)

type Grapher interface {
	NumVertices() int
	NumEdges() int
	Degree(int) int
	DoNeighbors(int, func(int, interface{}))
	HasEdge(int, int) bool
	Label(int, int) interface{}
	Add(int, int)
	AddLabel(int, int, interface{})
	AddBi(int, int)
	AddBiLabel(int, int, interface{})
	Remove(int, int)
	RemoveBi(int, int)
}

var NewFuncs = map[string]func(int) Grapher{

	//test the hash version
	"Hash": func(n int) Grapher { return NewHash(n) },

	//test the matrix version
	//"Matrix": func(n int) Grapher { return NewMatrix(n) },
}

// Constructs test graphs using the factory method f.
func setup(f func(int) Grapher) (g0, g1, g5 Grapher) {
	g0 = f(0)

	g1 = f(1)
	g1.Add(0, 0)

	g5 = f(5)
	g5.Add(0, 1)
	g5.AddLabel(2, 3, 1)
	return
}

// Checks if res and exp are different.
func diff(res, exp interface{}) (message string, diff bool) {
	switch res := res.(type) {
	case []int:
		if !arrayEq(res, exp.([]int)) {
			message = fmt.Sprintf("%v; want %v", res, exp)
			diff = true
		}
	default:
		if res != exp {
			message = fmt.Sprintf("%v; want %v", res, exp)
			diff = true
		}
	}
	return
}

// Checks if res and exp are different permutations of characters.
func diffPerm(res, exp string) (message string, diff bool) {
	r := make([]int, len(res))
	for i, ch := range res {
		r[i] = int(ch)
	}
	e := make([]int, len(exp))
	for i, ch := range exp {
		e[i] = int(ch)
	}
	sort.Ints(r)
	sort.Ints(e)
	if !arrayEq(r, e) {
		message = fmt.Sprintf("%s; want permutation of %s", res, exp)
		diff = true
	}
	return
}

// Checks if a and b have the same elements in the same order.
func arrayEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}

func TestNumVertices(t *testing.T) {
	for impl, f := range NewFuncs {
		g0, g1, g5 := setup(f)
		s := "NumVertices()"

		if mess, diff := diff(g0.NumVertices(), 0); diff {
			t.Errorf("%s: g0.%s %s", impl, s, mess)
		}
		if mess, diff := diff(g1.NumVertices(), 1); diff {
			t.Errorf("%s: g1.%s %s", impl, s, mess)
		}
		if mess, diff := diff(g5.NumVertices(), 5); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
	}
}

func TestNumEdges(t *testing.T) {
	for impl, f := range NewFuncs {
		g0, g1, g5 := setup(f)
		s := "NumEdges()"

		if mess, diff := diff(g0.NumEdges(), 0); diff {
			t.Errorf("%s: g0.%s %s", impl, s, mess)
		}
		if mess, diff := diff(g1.NumEdges(), 1); diff {
			t.Errorf("%s: g1.%s %s", impl, s, mess)
		}
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.Remove(1, 2)
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.Add(0, 1)
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.Add(2, 1)
		if mess, diff := diff(g5.NumEdges(), 3); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.Remove(0, 1)
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.Remove(1, 0)
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.RemoveBi(1, 0)
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.RemoveBi(1, 2)
		if mess, diff := diff(g5.NumEdges(), 1); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.AddBi(2, 3)
		if mess, diff := diff(g5.NumEdges(), 2); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
		g5.RemoveBi(2, 3)
		if mess, diff := diff(g5.NumEdges(), 0); diff {
			t.Errorf("%s: g5.%s %s", impl, s, mess)
		}
	}
}

func TestEdgeLabels(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		if mess, diff := diff(g1.Label(0, 0), NoLabel); diff {
			t.Errorf("%s: g1.Label(0, 0) %s", impl, mess)
		}
		if mess, diff := diff(g5.Label(0, 1), NoLabel); diff {
			t.Errorf("%s: g5.Label(0, 1) %s", impl, mess)
		}
		if mess, diff := diff(g5.Label(2, 3), 1); diff {
			t.Errorf("%s: g5.Label(2, 3) %s", impl, mess)
		}
		if mess, diff := diff(g5.Label(3, 2), nil); diff {
			t.Errorf("%s: g5.Label(3, 2) %s", impl, mess)
		}
		if mess, diff := diff(g5.Label(1, 2), nil); diff {
			t.Errorf("%s: g5.Label(1, 2) %s", impl, mess)
		}

		for i := 0; i < 2; i++ {
			g1.Remove(0, 0)
			g5.Remove(1, 0)
			g5.Remove(2, 3)

			if mess, diff := diff(g1.Label(0, 0), nil); diff {
				t.Errorf("%s: g1.Label(0, 0) %s", impl, mess)
			}
			if mess, diff := diff(g5.Label(0, 1), NoLabel); diff {
				t.Errorf("%s: g5.Label(0, 1) %s", impl, mess)
			}
			if mess, diff := diff(g5.Label(2, 3), nil); diff {
				t.Errorf("%s: g5.Label(2, 3) %s", impl, mess)
			}
			if mess, diff := diff(g5.Label(3, 2), nil); diff {
				t.Errorf("%s: g5.Label(3, 2) %s", impl, mess)
			}
			if mess, diff := diff(g5.Label(1, 2), nil); diff {
				t.Errorf("%s: g5.Label(1, 2) %s", impl, mess)
			}
		}
	}
}

// Add an edge where a Label already exists
func TestAddNewLabel(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		g1.AddLabel(0, 0, 8)
		if mess, diff := diff(g1.Label(0, 0), 8); diff {
			t.Errorf("%s: g1.Label(0, 0) %s", impl, mess)
		}
		g5.Add(2, 3)
		if mess, diff := diff(g5.Label(2, 3), NoLabel); diff {
			t.Errorf("%s: g5.Label(2, 3) %s", impl, mess)
		}
		g5.Add(3, 2)
		if mess, diff := diff(g5.Label(3, 2), NoLabel); diff {
			t.Errorf("%s: g5.Label(3, 2) %s", impl, mess)
		}
	}
}

func TestAddBiNewLabel(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		g1.AddBiLabel(0, 0, 8)
		if mess, diff := diff(g1.Label(0, 0), 8); diff {
			t.Errorf("%s: g1.Label(0, 0) %s", impl, mess)
		}
		g5.AddBi(2, 3)
		if mess, diff := diff(g5.Label(2, 3), NoLabel); diff {
			t.Errorf("%s: g5.Label(2, 3) %s", impl, mess)
		}
		g5.AddBi(3, 2)
		if mess, diff := diff(g5.Label(3, 2), NoLabel); diff {
			t.Errorf("%s: g5.Label(3, 2) %s", impl, mess)
		}
	}
}

func TestNilLabel(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		g1.AddLabel(0, 0, nil)
		if mess, diff := diff(g1.Label(0, 0), nil); diff {
			t.Errorf("%s: g1.AddLabel(0, 0, nil) %s", impl, mess)
		}
		if mess, diff := diff(g1.HasEdge(0, 0), true); diff {
			t.Errorf("%s: g1.HasEdge(0, 0) %s", impl, mess)
		}
		g5.AddLabel(0, 1, nil)
		if mess, diff := diff(g5.Label(0, 1), nil); diff {
			t.Errorf("%s: g5.Label(0, 1) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(0, 1), true); diff {
			t.Errorf("%s: g5.HasEdge(0, 1) %s", impl, mess)
		}
		g5.AddLabel(1, 2, nil)
		if mess, diff := diff(g5.Label(1, 2), nil); diff {
			t.Errorf("%s: g5.Label(1, 2) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(1, 2), true); diff {
			t.Errorf("%s: g5.HasEdge(1, 2) %s", impl, mess)
		}
		g5.AddLabel(2, 3, nil)
		if mess, diff := diff(g5.Label(2, 3), nil); diff {
			t.Errorf("%s: g5.Label(2, 3) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(2, 3), true); diff {
			t.Errorf("%s: g5.HasEdge(2, 3) %s", impl, mess)
		}
	}
}

func TestDegree(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		if mess, diff := diff(g1.Degree(0), 1); diff {
			t.Errorf("%s: g1.Degree(0) %s", impl, mess)
		}
		g1.Remove(0, 0)
		if mess, diff := diff(g1.Degree(0), 0); diff {
			t.Errorf("%s: g1.Degree(0) %s", impl, mess)
		}
		if mess, diff := diff(g5.Degree(0), 1); diff {
			t.Errorf("%s: g5.Degree(0) %s", impl, mess)
		}
		if mess, diff := diff(g5.Degree(1), 0); diff {
			t.Errorf("%s: g5.Degree(1) %s", impl, mess)
		}
		g5.Add(0, 2)
		if mess, diff := diff(g5.Degree(0), 2); diff {
			t.Errorf("%s: g5.Degree(0) %s", impl, mess)
		}
		if mess, diff := diff(g5.Degree(1), 0); diff {
			t.Errorf("%s: g5.Degree(1) %s", impl, mess)
		}
		g5.Remove(0, 1)
		if mess, diff := diff(g5.Degree(0), 1); diff {
			t.Errorf("%s: g5.Degree(0) %s", impl, mess)
		}
		if mess, diff := diff(g5.Degree(1), 0); diff {
			t.Errorf("%s: g5.Degree(1) %s", impl, mess)
		}
		g5.Remove(1, 0)
		if mess, diff := diff(g5.Degree(0), 1); diff {
			t.Errorf("%s: g5.Degree(0) %s", impl, mess)
		}
		if mess, diff := diff(g5.Degree(1), 0); diff {
			t.Errorf("%s: g5.Degree(1) %s", impl, mess)
		}
		g5.Remove(4, 0)
		if mess, diff := diff(g5.Degree(4), 0); diff {
			t.Errorf("%s: g5.Degree(4) %s", impl, mess)
		}
	}
}

func TestDoNeighbors(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		count := 0
		g1.DoNeighbors(0, func(v int, c interface{}) {
			if mess, diff := diff(v, 0); diff {
				t.Errorf("%s: g1.DoNeighbors v: %s", impl, mess)
			}
			if mess, diff := diff(c, NoLabel); diff {
				t.Errorf("%s: g1.DoNeighbors c: %s", impl, mess)
			}
			count++
		})
		if mess, diff := diff(count, 1); diff {
			t.Errorf("%s: g1.DoNeighbors #it: %s", impl, mess)
		}

		count = 0
		g5.DoNeighbors(4, func(v int, c interface{}) {
			count++
		})
		if mess, diff := diff(count, 0); diff {
			t.Errorf("%s: g1.DoNeighbors #it: %s", impl, mess)
		}

		count = 0
		g5.DoNeighbors(2, func(v int, c interface{}) {
			if mess, diff := diff(v, 3); diff {
				t.Errorf("%s: g5.DoNeighbors v: %s", impl, mess)
			}
			if mess, diff := diff(c, 1); diff {
				t.Errorf("%s: g5.DoNeighbors c: %s", impl, mess)
			}
			count++
		})
		if mess, diff := diff(count, 1); diff {
			t.Errorf("%s: g5.DoNeighbors #it: %s", impl, mess)
		}

		g5.Add(0, 0)
		g5.Add(0, 1)
		g5.Add(2, 0)
		g5.Add(0, 3)
		count = 0
		set := make(map[int]int)
		g5.DoNeighbors(0, func(v int, c interface{}) {
			set[v] = 0
			count++
		})
		if mess, diff := diff(count, 3); diff {
			t.Errorf("%s: g5.DoNeighbors #it: %s", impl, mess)
		}
		for i := 0; i < 4; i++ {
			_, res := set[i]
			exp := true
			if i == 2 {
				exp = false
			}
			if mess, diff := diff(res, exp); diff {
				t.Errorf("set[%d] %s", i, mess)
			}
		}
	}
}

func TestHasEdge(t *testing.T) {
	for impl, f := range NewFuncs {
		_, g1, g5 := setup(f)

		if mess, diff := diff(g1.HasEdge(0, 0), true); diff {
			t.Errorf("%s: g1.HasEdge(0, 0) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(0, 0), false); diff {
			t.Errorf("%s: g5.HasEdge(0, 0), false %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(0, 1), true); diff {
			t.Errorf("%s: g5.HasEdge(0, 1) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(1, 0), false); diff {
			t.Errorf("%s: g5.HasEdge(1, 0) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(1, 2), false); diff {
			t.Errorf("%s: g5.HasEdge(1, 2) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(3, 2), false); diff {
			t.Errorf("%s: g5.HasEdge(3, 2) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(2, 3), true); diff {
			t.Errorf("%s: g5.HasEdge(2, 3) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(4, 0), false); diff {
			t.Errorf("%s: g5.HasEdge(4, 0) %s", impl, mess)
		}
		g5.Add(2, 1)
		if mess, diff := diff(g5.HasEdge(1, 2), false); diff {
			t.Errorf("%s: g5.HasEdge(1, 2) %s", impl, mess)
		}
		if mess, diff := diff(g5.HasEdge(2, 1), true); diff {
			t.Errorf("%s: g5.HasEdge(2, 1) %s", impl, mess)
		}
	}
}

func TestDFS(t *testing.T) {
	for impl, f := range NewFuncs {
		g := f(8)
		edges := []struct {
			v, w int
		}{
			{0, 1}, {0, 3}, {1, 4}, {3, 4}, {5, 3},
			{2, 2},
			{6, 7},
		}
		for _, e := range edges {
			g.AddBi(e.v, e.w)
		}

		dfs := ""
		state := make([]bool, g.NumVertices())
		action := func(v int) { dfs += strconv.Itoa(v) }
		for v, visited := range state {
			if !visited {
				DFS(g, v, state, action)
				dfs += "#"
			}
		}
		if mess, diff := diffPerm(dfs[0:5], "01345"); diff {
			t.Errorf("%s: dfs[0:5] %s", impl, mess)
		}
		if mess, diff := diffPerm(dfs[5:6], "#"); diff {
			t.Errorf("%s: dfs[5:6] %s", impl, mess)
		}
		if mess, diff := diffPerm(dfs[6:7], "2"); diff {
			t.Errorf("%s: dfs[6:7] %s", impl, mess)
		}
		if mess, diff := diffPerm(dfs[7:8], "#"); diff {
			t.Errorf("%s: dfs[7:8] %s", impl, mess)
		}
		if mess, diff := diffPerm(dfs[8:10], "67"); diff {
			t.Errorf("%s: dfs[8:10] %s", impl, mess)
		}
		if mess, diff := diffPerm(dfs[10:11], "#"); diff {
			t.Errorf("%s: dfs[10:11] %s", impl, mess)
		}
	}
}

func TestBFS(t *testing.T) {
	for impl, f := range NewFuncs {
		g := f(10)
		edges := []struct {
			v, w int
		}{
			{0, 1}, {0, 4}, {0, 7}, {0, 9},
			{4, 2}, {7, 5}, {7, 8},
			{2, 3}, {5, 6},
			{3, 6}, {8, 9}, {4, 4},
		}
		for _, e := range edges {
			g.AddBi(e.v, e.w)
		}

		bfs := ""
		state := make([]bool, g.NumVertices())
		action := func(v int) { bfs += strconv.Itoa(v) }
		BFS(g, 0, state, action)
		if mess, diff := diffPerm(bfs[0:1], "0"); diff {
			t.Errorf("%s: bfs[0:1] %s", impl, mess)
		}
		if mess, diff := diffPerm(bfs[1:5], "1479"); diff {
			t.Errorf("%s: bfs[1:5] %s", impl, mess)
		}
		if mess, diff := diffPerm(bfs[5:8], "258"); diff {
			t.Errorf("%s: bfs[5:8] %s", impl, mess)
		}
		if mess, diff := diffPerm(bfs[8:10], "36"); diff {
			t.Errorf("%s: bfs[8:10] %s", impl, mess)
		}
	}
}
