package main

import "fmt"

func main() {

	//uint8 is acronym to byte
	var x uint8 = 1<<1 | 1<<5 // Union {1} and {5}
	var y uint8 = 1<<1 | 1<<2 // Union {1} and {2}

	//08 is adverb of verb %b, it modifies %b
	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y) // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y) // "00100110", the union {1, 2, 5}

	fmt.Printf("%08b\n", x^y) // "00100100", XOR, the symmetric (of both sides) diff {2, 5}

	fmt.Printf("%08b\n", x&^y) // "00100000", ANDNOT, the one-sided (x) diff {5}
	fmt.Printf("%08b\n", y&^x) // "00000100", ANDNOT, the one-sided (y) diff {2}

}
