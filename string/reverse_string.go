package main

import "fmt"

func main() {
	var a string
	var b [10]int
	fmt.Println("Enter the string")
	fmt.Scanln(&a)
	fmt.Println(len(a))
	i:=len(a)-1
	for j := range a {
		b[i]=int(a[j])
		i--
	}
	for i:=0;i<len(a);i++ {
		fmt.Print(string(b[i]))
	}
}
