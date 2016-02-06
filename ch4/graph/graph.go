package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	fmt.Println(hasEdge("from", "to string"))
}

func addEdge(from, to string) {
	edges := graph[from]

	// assignment to nil map will panic
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {

	// accessing nil map is a safe operation
	return graph[from][to]
}
