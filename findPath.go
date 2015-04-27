// Stefan Nilsson 2014-04-20
// Radically modified and completed by Ivan Liljeqvist 2015-04-27

// Grep searches the input file for lines containing the given pattern and
// prints these lines. It is a simplified version of the Unix grep command.
package main

import (
	graph "./graph"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	This interface describes the hash and matrix classes.
	It declares all the methods.
	These methods are matched in matrix.go and hash.go
*/

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

const prog = "grep"

func init() {
	log.SetPrefix(prog + ": ")
	log.SetFlags(0) // no extra info in log messages
}

func main() {

	//check so that the arguments aare correct
	if len(os.Args) != 4 {
		log.Fatalf("usage: %s FROM TO FILE\n", prog)
	}
	from, from_error := strconv.Atoi(os.Args[1])

	if from_error != nil {
		panic("FROM argument is wrong!")
	}

	to, to_error := strconv.Atoi(os.Args[2])

	if to_error != nil {
		panic("TO argument is wrong!")
	}

	//get the pattern from the arguments
	filepath := os.Args[3]

	//construct the graph from the file
	g := constructGraph(filepath)

	if g == nil {
		panic("Couldn't build the graph.")
	}

	getPath(from, to, g)
}

func getPath(from, to int, g Grapher) {

	visited := make([]bool, g.NumVertices())

	count := 0

	graph.BFS(g, from, visited, func(w int) {
		count++

		if to == w {
			fmt.Println(count)
		}

	})

}

func constructGraph(filepath string) (g Grapher) {
	//open file
	file, errors := os.Open(filepath)

	if errors != nil {
		panic("error opening file!")
	}

	graphSize := 0
	firstLine := true

	//go through each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		//ignore all commented lines
		if strings.HasPrefix(line, "//") == false {

			//split the string
			words := strings.Fields(line)

			currentWord := 0

			var from, to, label int

			shouldAttachToGraph := false

			//go through all words n this line
			for _, s := range words {

				//convert string to int
				i, err := strconv.Atoi(s)

				if err == nil {
					//if first line - just get the size
					if firstLine == true {
						graphSize = i
						firstLine = false

						//init the graph
						g = graph.NewHash(graphSize)
					} else {

						//if not first line

						//get from and to and label
						if currentWord == 0 {
							from = i
						} else if currentWord == 1 {
							to = i
						} else if currentWord == 2 {
							label = i

							//if we've come this far, everything worked fine
							//we can attach
							shouldAttachToGraph = true
						}

					}
				} else {
					//string counldt be converted to int
				}

				currentWord++

			}

			if shouldAttachToGraph {
				//attach the information from this line to the graph
				g.AddLabel(from, to, label)
			}

		}

	}

	file.Close()

	return g
}
