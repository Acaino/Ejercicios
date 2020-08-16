package main

import (
	"fmt"
	"math"
)

func main() {
	const distanciaCentro float64 = 4
	var i float64
	var imax float64 = distanciaCentro
	var imin float64 = -distanciaCentro
	var j float64
	var jmax float64 = distanciaCentro
	var jmin float64 = -distanciaCentro

	for j = jmin; j <= jmax; j++ {
		for i = imin; i <= imax; i++ {
			if math.Abs(i)+math.Abs(j) <= imax {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
