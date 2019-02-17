package main

import "fmt"

type basic struct {
	a int
	b string
}

func main(){
	var b basic
	b.a = 4
	b.b = "abcd"
	fmt.Println(b.a)
	fmt.Println(b.b)
}
