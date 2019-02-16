package main

import "fmt"

func palindrome(n int) bool {
	temp := n
	rev := 0
	for temp != 0 {
		rev = 10*rev + temp%10
		temp = temp / 10
	}
	if n == rev {
		return true
	} else {
		return false
	}
}
func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	if palindrome(n) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}
