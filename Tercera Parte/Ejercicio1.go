package main

import (
	"fmt"
)

func main() {
	var j int
	var jn int = 11 //a veces es medio antiintuitivo que arranque el loop en 0, así que arranco en 1 y uso 11 como cota
	for j = 1; j <= jn; j++ {
		if j == (jn+1)/2 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("#")
		}
	}
}
