package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie information
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //if value is zero-value, do not marshal.
	Actors []string
}

var movies = []Movie{
	{Title: "ABC", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func main() {

	// no indent
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// with indent , prefix and indent (level)
	data, err = json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// unmarshaling
	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
