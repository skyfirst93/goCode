package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64 
	fmt.Println("Enter the number ")
	fmt.Scanln(&n)
	for i := 2; i <= int(math.Sqrt(n)); i++ {
		if int(n)%i == 0 {
			fmt.Println("NOt PRIME")
			return
		}
	}
	fmt.Println("Prime")
}
