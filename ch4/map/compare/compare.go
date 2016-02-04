/*
“As with slices, maps cannot be compared to each other; the only legal comparison is with nil. To test whether two maps contain the same keys and the same associated values, we must write a loop:
”

Excerpt From: Brian W. Kernighan. “The Go Programming Language (Addison-Wesley Professional Computing Series).” iBooks.
*/
package main

import "fmt"

func main() {
	mapOne := map[string]int{
		"roy": 12,
		"lee": 11,
	}

	mapTwo := map[string]int{
		"roy": 12,
		"lee": 12,
	}

	mapThree := map[string]int{
		"roy": 12,
		"lee": 11,
	}

	fmt.Printf("%t\n", equal(mapOne, mapTwo))
	fmt.Printf("%t\n", equal(mapOne, mapThree))

}

// map type is not comparable, hence we need a equal func to compare their elements.
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true

}
