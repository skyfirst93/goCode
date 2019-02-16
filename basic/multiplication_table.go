package main

import "fmt"

func main(){
	var n int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	for i:=1;i<=10;i++ {
		fmt.Println(n," X ",i," = ",n*i)
	}
}
