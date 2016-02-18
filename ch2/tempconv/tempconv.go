package main

import "fmt"

// Celcius is one of temperature scale.
type Celcius float64

// Farenheit is one of temperature scale.
type Farenheit float64

func main() {

	//var k Farenheit
	var c Celcius
	royPrint(c)

}

func royPrint(v interface{}) {
	//
	c := v.(Celcius)

	//var _ = &c
	fmt.Print(c.String())
}

func (c *Celcius) String() string {
	return fmt.Sprintf("%gC\n", *c)
}

// func (f Farenheit) String() string {
// 	return fmt.Sprintf("%gF", f)
// }
