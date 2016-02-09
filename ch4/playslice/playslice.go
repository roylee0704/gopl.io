package main

import "fmt"

func main() {

	months := []string{
		"",
		"January",
		"Febuary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	} // len(current): 13 items , cap(total): 13 items

	Q2 := months[4:7]     // len:(current) 3, cap(total): 9
	summer := months[6:9] // len(current): 3, cap(total): 7
	alvi := summer[0:1]   // len(current): 1, cap(total): 7

	fmt.Printf("%q\n", months)
	fmt.Printf("%q\n", Q2)
	fmt.Printf("%q %d %d\n", summer, len(summer), cap(summer))
	fmt.Printf("%q %d %d\n", alvi, len(alvi), cap(alvi))

}
