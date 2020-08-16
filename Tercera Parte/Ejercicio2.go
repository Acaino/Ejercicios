package main

import (
	"fmt"
)

func main() {
	var i int
	var in int = 10
	var j int
	var jn int = 4

	for j = 1; j <= jn; j++ {
		for i = 1; i <= in; i++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
