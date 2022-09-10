package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}}, {Title: "Coll Hand Luke", Year: 1967, Color: true, Actors: []string{"Pau Newman"}}}
	marshal, err := json.Marshal(&movies)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(marshal))

	indent, err := json.MarshalIndent(&movies, "", "      ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(indent))

	var titles []struct{ Title string }
	err = json.Unmarshal(marshal, &titles)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(titles)
}
