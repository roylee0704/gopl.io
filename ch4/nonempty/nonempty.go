// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty worst-case's: o(n), mutatable func
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
	data := []string{"one", "", "three"}

	fmt.Printf("%q\n", nonempty(data)) //`["one" "three"]`
	fmt.Printf("%q\n", data)           //`["one" "three" "three"]`
}
