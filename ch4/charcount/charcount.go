// Charcount computes counts of Unicode characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    //counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int //count of lengths of UTF-8 encodings
	invalid := 0                    //count of invalid UTF-8 characters

	input := bufio.NewReader(os.Stdin)

	for {
		r, n, err := input.ReadRune() // returns rune, nbytes, error

		if err == io.EOF {
			fmt.Println("EOF looks like this:-->", err)
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	// frequency table of unicode-characters
	fmt.Printf("rune\tcount\n")
	for r, n := range counts {
		fmt.Printf("%q\t%d\n", r, n)
	}

	// distribution of the lengths of all the UTF-8 encodings.
	fmt.Printf("\nlen\tcount\n")
	for l, n := range utflen {
		fmt.Printf("%d\t%d\n", l, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
