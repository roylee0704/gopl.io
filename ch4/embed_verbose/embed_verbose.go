// Clearer but verbose version of struct usage.
package main

import "fmt"

// Point can be constructed with x-y coordinates.
type Point struct {
	X, Y int
}

// Circle can be constructed with point and radius.
type Circle struct {
	Center Point
	Radius int
}

// Wheel can be constructed with point, radius (circle) and spokes.
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel

	w.Circle.Center.X = 8
	w.Circle.Center.Y = 9
	w.Circle.Radius = 20
	w.Spokes = 100

	fmt.Println(w)
}
