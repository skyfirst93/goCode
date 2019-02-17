package main

import "fmt"

type hello interface {
	Sayhello()
	Saymsg(string)
}

type greet struct{
}

func (gr greet) Sayhello(){
	fmt.Println("hello")
}

func (gr greet) Saymsg(msg string){
	fmt.Println(msg)
}

func main() {
	var inf hello
	inf = greet{}
	inf.Sayhello()
	inf.Saymsg("hi")
}
