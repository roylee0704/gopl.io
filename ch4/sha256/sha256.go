package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("hex char (64): %x\nhex char (64): %x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(len(c1))
}
