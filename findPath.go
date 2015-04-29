// Stefan Nilsson 2014-04-20
// Radically modified and completed by Ivan Liljeqvist 2015-04-27

/*
	You give this program a file containing a graph and two nodes in that graph.
	The program will print the shortes path (with as few nodes as possible) between those nodes.
	If no path is found, an empty row will be printed.
*/

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
	DoNeighbors(int, func(int, int, interface{}))
	HasEdge(int, int) bool
	Label(int, int) interface{}
	Add(int, int)
	AddLabel(int, int, interface{})
	AddBi(int, int)
	AddBiLabel(int, int, interface{})
	Remove(int, int)
	RemoveBi(int, int)
}

/*
	Main function.
	Checks the parameters from the Terminal.
	Constructs the graph from file.
	Prints the path from 'FROM'-node to 'TO'-node.
*/

func main() {

	//check so that the arguments aare correct
	if len(os.Args) != 4 {
		log.Fatalf("usage: FROM TO FILE\n")
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

	if from < 0 {
		panic("FROM argument can't be less than zero.")
	}

	if to < 0 {
		panic("TO argument can't be less than zero.")
	}

	//construct the graph from the file
	g := constructGraph(filepath)

	if g == nil {
		panic("Couldn't build the graph.")
	}

	getPath(from, to, g)
}

/*
	Finds the shortest (as few nodes as possible) way from 'from' node to 'to' node
	in the 'g' graph.

	Prints out the results if a path is found.
	Prints out an empty line if there is no path between the nodes.
*/

func getPath(from, to int, g Grapher) {

	const INDEX_HAS_NO_PARENT = -1

	visited := make([]bool, g.NumVertices())

	stack := make([]int, g.NumEdges())
	for i, _ := range stack {
		stack[i] = INDEX_HAS_NO_PARENT
	}

	graph.BFS(g, from, visited, func(parent, w int) {

		//save the parent of each node
		stack[w] = parent

	})

	//this array will contain the path to our goal
	var path []int

	//put the goal as the first value in the path
	//we're building this backwards, this will be reversed at the end
	index := to
	path = append(path, index)

	//flag that will be set to false when we've found the goal-node
	searching := true

	if index > len(stack) {
		panic("THE NODE YOU'RE SEARCHING FOR DOESNT EVEN EXIST IN THE GRAPH!")
	}

	for stack[index] != INDEX_HAS_NO_PARENT {
		//get parent of the index
		//in the first iteration it will get the parent of the goal
		//in the second - the grandparent of the goal, until we reach starting point (from)
		prev := stack[index]
		index = prev
		//save parent to the path
		path = append(path, index)

		//abort search if we've reached the starting point
		//this means we've went backwards from the goal to starting point
		if index == from {
			searching = false
			break
		}
	}

	//if the node is found - we're not longer searching
	if searching == false {
		reverseSlice(path)
		fmt.Println("Path from ", from, " to ", to, ": ", path)
	} else {
		//node not found - just print empty line.
		fmt.Println(" ")
	}

}

/*
	Reverses the slice passed in as the parameter.
*/

func reverseSlice(slice []int) {
	length := len(slice)

	//go through the array and switch from both sides.
	for i := 0; i < length/2; i++ {
		temp := slice[i]
		slice[i] = slice[length-i-1]
		slice[length-i-1] = temp
	}
}

/*
	This methods constructs an undirected graph from the file found
	at the file path passed in as the parameter.

	Returns a Grapher object with a hash graph containing all
	the connections as in the file.
*/

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
					panic("Error constructing the graph. File is structured in a wrong way. Couldn't convert a value to int.")
				}

				currentWord++

			}

			if shouldAttachToGraph {
				//attach the information from this line to the graph
				g.AddBiLabel(from, to, label)
			}

		}

	}

	file.Close()

	return g
}
