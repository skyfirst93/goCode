package main

import "fmt"

func main(){
	var s string
	fmt.Println("Enter the string")
	fmt.Scanln(&s)
	for i:=0;i<len(s);i++ {
		for j:=0;j<len(s);j++ {
			for k:=i;k<= j;k++ {
				fmt.Print(string(s[k]))
			}
		fmt.Println("")
		}
	}
}
