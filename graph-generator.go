package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang-collections/collections/set"
)

// graph represents an undirected graph
type graph struct {
	numberOfVertices int              // number of vertices in the graph
	edges            map[int]*set.Set // map of edge to the set of edges
}

// newGraph creates a new graph with a given number of vertices
func newGraph(numberVertices int) *graph {

	// Preconditions
	if numberVertices < 0 {
		log.Fatalf("Invalid number of vertices: %v\n", numberVertices)
	}

	g := graph{
		numberOfVertices: numberVertices,
		edges:            make(map[int]*set.Set),
	}

	// Initialise the graph
	for i := 0; i < numberVertices; i++ {
		g.edges[i] = set.New()
	}

	// Return a pointer to the graph
	return &g
}

// minMax returns the (minimum, maximum) value of a pair of integers
func minMax(v1 int, v2 int) (int, int) {
	if v1 <= v2 {
		return v1, v2
	}
	return v2, v1
}

// addEdge adds an undirected edge to the graph
func (g *graph) addEdge(v1 int, v2 int) {

	// Preconditions
	if v1 == v2 {
		log.Fatalf("Can't add self-loops")
	}

	if v1 < 0 || v1 >= g.numberOfVertices {
		log.Fatalf("Invalid vertex: %v\n", v1)
	}

	if v2 < 0 || v2 >= g.numberOfVertices {
		log.Fatalf("Invalid vertex: %v\n", v2)
	}

	// Sort the vertices
	lower, upper := minMax(v1, v2)

	// Add the edge
	g.edges[lower].Insert(upper)
}

// hasEdge returns whether an undirected edge exists between two vertices
func (g *graph) hasEdge(v1 int, v2 int) bool {

	// Preconditions
	if v1 == v2 {
		log.Fatalf("Can't add self-loops")
	}

	if v1 < 0 || v1 >= g.numberOfVertices {
		log.Fatalf("Invalid vertex: %v\n", v1)
	}

	if v2 < 0 || v2 >= g.numberOfVertices {
		log.Fatalf("Invalid vertex: %v\n", v2)
	}

	// Sort the vertices
	lower, upper := minMax(v1, v2)

	// Check if the edge exists
	return g.edges[lower].Has(upper)
}

// buildRandomGraphFixedNumEdges builds a random graph with a given number of edges
func buildRandomGraphFixedNumEdges(numberVertices int, numberEdgesRequired int) *graph {

	// Preconditions
	if numberVertices < 2 {
		log.Fatalf("Invalid number of vertices: %v\n", numberVertices)
	}

	// Initialise the graph
	fmt.Printf("[>] Initialising graph with %v vertices\n", numberVertices)
	g := newGraph(numberVertices)

	// Set the seed of the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	numberEdges := 0
	lastNumberEdgesDisplay := 0

	fmt.Printf("[>] Generating random edges ...\n")
	for numberEdges < numberEdgesRequired {

		if numberEdges%500000 == 0 && lastNumberEdgesDisplay != numberEdges {
			fmt.Printf("[>] Number of edges created: %v\n", numberEdges)
			lastNumberEdgesDisplay = numberEdges
		}

		// Generate two random vertex IDs
		v1 := rand.Intn(numberVertices)
		v2 := rand.Intn(numberVertices)

		if v1 == v2 {
			continue
		}

		// If the edge already exists, ignore it
		if g.hasEdge(v1, v2) {
			continue
		}

		// Add the edge to the graph
		g.addEdge(v1, v2)
		numberEdges++
	}

	return g
}

// buildRandomGraph builds a random graph
func buildRandomGraph(numberOfVertices int, probConnection float64) *graph {

	// Initialise the graph
	fmt.Printf("[>] Initialising graph with %v vertices\n", numberOfVertices)
	g := newGraph(numberOfVertices)

	// Set the seed of the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Walk through each pair of vertices
	fmt.Printf("[>] Generating random edges ...\n")
	for i := 0; i < (numberOfVertices - 1); i++ {

		if i%5000 == 0 {
			fmt.Printf("[>] Building from vertex %v of %v\n", i, numberOfVertices)
		}

		for j := i + 1; j < numberOfVertices; j++ {

			// Generate a uniformly distributed random number
			r := rand.Float64()

			if r < probConnection {
				g.addEdge(i, j)
			}
		}
	}

	// Return a pointer to the graph
	return g
}

// graphSummary represents the summary statistics of the graph
type graphSummary struct {
	numberVertices int // number of vertices with connections
	numberEdges    int // number of edges
}

// calcGraphSummary calculates summary statistics of the graph
func calcGraphSummary(g *graph) graphSummary {

	connectedVertices := set.New()
	numberEdges := 0

	for source, destinations := range g.edges {

		// If there are no connections, move to the next vertex
		if destinations.Len() == 0 {
			continue
		}

		connectedVertices.Insert(source)

		destinations.Do(func(s interface{}) {
			// Destination as an integer
			d := s.(int)

			if !connectedVertices.Has(d) {
				connectedVertices.Insert(d)
			}

			numberEdges++

			if numberEdges%5000000 == 0 {
				fmt.Printf("[>] Processed %v edges\n", numberEdges)
			}
		})
	}

	return graphSummary{
		numberVertices: connectedVertices.Len(),
		numberEdges:    numberEdges,
	}
}

// writeGraphToFile writes the graph to file as a list of edges
func writeGraphToFile(g *graph, filepath string) {

	// Open the output CSV file for writing
	outputFile, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("Unable to open output file %v for writing: %v\n", filepath, err)
	}
	defer outputFile.Close()

	// Walk through the source vertices
	for source, destinations := range g.edges {

		// If there are no connections, move to the next vertex
		if destinations.Len() == 0 {
			continue
		}

		// Walk through the set of destination vertices
		destinations.Do(func(s interface{}) {

			// Destination as an integer
			d := s.(int)

			fmt.Fprintf(outputFile, "%v,%v\n", source, d)
		})
	}
}

// buildRandomGraphFile creates a graph and saves the edges to a file
func buildRandomGraphFile(
	numberOfVertices int,
	probConnection float64,
	numberEdges int,
	outputFilepath string) {

	// Preconditions
	if numberOfVertices < 0 {
		log.Fatalf("Number of vertices is invalid: %v\n", numberOfVertices)
	}

	fixedNumberEdges := numberEdges > 0

	// Create the random graph
	t0 := time.Now()
	fmt.Printf("[>] Building random graph ...\n")

	var g *graph
	if fixedNumberEdges {
		g = buildRandomGraphFixedNumEdges(numberOfVertices, numberEdges)
	} else {
		g = buildRandomGraph(numberOfVertices, probConnection)
	}

	fmt.Printf("[>] Time taken to build graph: %v\n", time.Now().Sub(t0))

	// Summarise the graph
	t1 := time.Now()
	fmt.Printf("[>] Calculating graph summary ...\n")
	summary := calcGraphSummary(g)
	fmt.Printf("[>] Graph summary:\n")
	fmt.Printf("    Number of vertices: %v\n", summary.numberVertices)
	fmt.Printf("    Number of edges:    %v\n", summary.numberEdges)
	fmt.Printf("[>] Time taken to calculate summary: %v\n", time.Now().Sub(t1))

	// Write the graph to file
	t2 := time.Now()
	fmt.Printf("[>] Writing graph to %v ...\n", outputFilepath)
	writeGraphToFile(g, outputFilepath)
	fmt.Printf("[>] Time taken to write graph to file: %v\n", time.Now().Sub(t2))

	fmt.Printf("[>] Total time taken: %v\n", time.Now().Sub(t0))
}

func main() {

	// Command line arguments
	numberOfVertices := flag.Int("n", 100, "Number of vertices")
	probConnection := flag.Float64("p", -1.0, "Probability of a connection between vertices")
	numberEdges := flag.Int("e", -1, "Number of edges")
	outputFilepath := flag.String("output", "results.csv", "File of edges")
	flag.Parse()

	// Build the graph
	fmt.Println("Random graph generator")
	buildRandomGraphFile(*numberOfVertices, *probConnection, *numberEdges, *outputFilepath)
}
