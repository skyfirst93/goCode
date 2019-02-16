package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
