package main

import (
	"Project5/Ex2/Area"
)

func main() {

	triangle := Area.Triangle{
		Base:   10,
		Height: 20,
	}
	Area.PrintArea(triangle)
}
