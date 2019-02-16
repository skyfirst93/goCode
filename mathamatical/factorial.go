package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number for which factorial is to be found ")
	fmt.Scanln(&n)
	temp := n
	fact := 1
	fmt.Println("n value is ", n)
	for temp > 0 {
		fact = fact * temp
		fmt.Println("temp is ", temp)
		temp -= 1
	}
	fmt.Println("factorial of ", n, " is ", fact)

}
