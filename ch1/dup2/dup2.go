// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {

			fd, err := os.Open(file)

			if err != nil {
				fmt.Println("error", err)
				continue
			}
			countLines(fd, counts)
			fd.Close()

		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		if line == "." {
			break
		} //for stdin

		counts[line]++
	}
}
