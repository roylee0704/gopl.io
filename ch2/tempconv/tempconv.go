package main

import "fmt"

// Celcius is one of temperature scale.
type Celcius float64

// Farenheit is one of temperature scale.
type Farenheit float64

func main() {

	var k Farenheit
	//	var c Celcius
	fmt.Println(k)

}

func royPrint(v interface{}) {
	fmt.Sprintf("%gC", "D")

}

func (c Celcius) String() string {
	return fmt.Sprintf("%gC", c)
}
func (f Farenheit) String() string {
	return fmt.Sprintf("%gF", f)
}
