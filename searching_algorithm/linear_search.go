package main

import "fmt"

func main(){
	var n int
	var a []int = []int{1,2,3,4,5}
	fmt.Println("enter the number to be searched ")
	fmt.Scanln(&n)
	for i:=0;i<=len(a)-1;i++ {
		if n == a[i] {
			fmt.Println("found at",i)
			return 
		}
	}
	fmt.Println("Not found")
}
