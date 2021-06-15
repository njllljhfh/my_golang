package main

import "fmt"

type Vertex4 struct {
	Lat, Long float64
}

// Map literals are like struct literals, but the keys are required.
var m4 = map[string]Vertex4{
	"Bell Labs": Vertex4{
		40.68433, -74.39967,
	},
	"Google": Vertex4{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m4)

	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
	for k, v := range m4 {
		fmt.Printf("k=%s,  v=%g\n", k, v)
	}
}
