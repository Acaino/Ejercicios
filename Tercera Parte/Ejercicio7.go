package main

import (
	"fmt"
)

func main() {
	var i int
	var in int = 5
	var j int
	var jn int = 5

	for j = 1; j <= jn; j++ {
		for i = 1; i <= in; i++ {
			if i >= jn+1-j {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
