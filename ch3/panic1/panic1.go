// Attempting to access a byte outside of this range results in a panic.
package main

import "fmt"

func main() {
	s := "hello, world"     // 12 ascii chacracters
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 109"  ('h' and 'w')

	c := s[len(s)] // panic: index out of range
	fmt.Println(c)
}
