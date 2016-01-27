// Spinner provide the user with a visual indication that the program is still
// running, by displaying an animated textual "spinner"

package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond) // animate 0.1s per frame.

	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonnaci(%d) = %d\n", n, fibN)
}

// fib computes fibonacci number at nth position. (in-efficient)
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r Crunching... %c", r)
			time.Sleep(delay)
		}
	}

}
