// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty worst-case's: o(n)
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	strings := []string{"hello", "", "world", ",", "", "Roy"}

	fmt.Printf("%#v\n", nonempty(strings))
}
