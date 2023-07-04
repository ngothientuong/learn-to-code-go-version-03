package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	car1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "toyota",
		model: "rav4",
		doors: 4,
		color: "gray",
	}
	car2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "tesla",
		model: "s3",
		doors: 4,
		color: "white",
	}
	fmt.Printf("%v\t %v\n", car1, car1.electric)
	fmt.Printf("%v\t %v\n", car2, car2.electric)
}
