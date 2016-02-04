package main

import "fmt"

func main() {
	fmt.Println(naivePopCount(1))
	fmt.Println(naivePopCount(255))

	fmt.Println(naiveButSmartPopCount(1))
	fmt.Println(naiveButSmartPopCount(255))
}

func naivePopCount(n uint8) int {
	var count uint8
	for places := uint8(0); places < uint8(8); places++ {

		if n&(1<<places) != 0 { //{1,2,8} & {1} != empty set.
			count++
		}
	}
	return int(count)
}

func naiveButSmartPopCount(b uint8) int {
	var count uint8
	for places := uint8(0); places < 8; places++ {
		count += b >> places & 1 //{7} -> {6} -> {5}
	}
	return int(count)
}
