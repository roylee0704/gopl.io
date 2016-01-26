// Comma insert commas every three places, e.g: "12,345".
// p/s: this package only devise a solution for integer.
package main

import "fmt"

func main() {
	fmt.Println(comma("123456789")) // "12,345"
}

func comma(s string) string {

	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
