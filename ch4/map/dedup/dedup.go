// Go doesn't provide a set type, but since the key of maps are distinct, a map
// can serve this purpose.
// Dedup uses map whose keys represent the `set` of lines that have already appeared
// to ensure that subsequent occurances are not present
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	seen := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "exit" {
			break
		}

		if !seen[line] {
			seen[line] = true
			fmt.Printf("added: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v", err)
		os.Exit(1)
	}

}
