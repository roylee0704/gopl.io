// A `better` version of ../embed_verbose/embed_verbose.go. Less verbose.
// Let you refer to the names at the leaves of the implicit tree without giving
// the intervening names.
package main

import "fmt"

// Point can be constructed with x-y coordinates.
type Point struct {
	X, Y int
}

// Circle can be constructed with point and radius.
type Circle struct {
	Point
	Radius int
}

// Wheel can be constructed with point, radius (circle) and spokes.
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	X := w.X | w.Circle.Point.X // both are equivalents
	Y := w.Y | w.Circle.Point.Y

	fmt.Println(X, Y)

}
