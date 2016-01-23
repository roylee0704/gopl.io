// String values are immutable.
// Byte Sequence aka byte array contained in a String value can never be changed.
package main

import "fmt"

func main() {

	// byte array [...,'l','e','f','t',...] is created
	// malloc(s) and it has a struct of pointer-to-byte-array & len
	s := "left foot"

	// malloc(t) and it has a struct of the same pointer-to-byte-array & len
	t := s

	fmt.Printf("address of s: %v\n", &s)
	fmt.Printf("address of t: %v\n", &t)

	// byte array [...,'l', 'e', 'f', 't', ',' ,'r', 'i',...] is created
	// assign(s) and s has a struct of new pointer-to-byte-array & len
	s += ", right foot"

	fmt.Println("after string concat on s")
	fmt.Printf("address of s: %v\n", &s)
	fmt.Printf("address of t: %v\n", &t)

}
