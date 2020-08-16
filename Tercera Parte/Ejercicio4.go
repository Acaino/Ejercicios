package main

import (
	"fmt"
)

func main() {
	var i int
	var in int = 10
	var j int
	var jn int = 7

	for j = 1; j <= jn; j++ {
		for i = 1; i <= in; i++ {
			if j%2 == 0 && i > 1 && i < in {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}

}
