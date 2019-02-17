package main

import "fmt"

type value interface{
	Increment()
	Printi()
}
type class struct{
	i int
}
func (c *class) Increment(){
	c.i++
}
func (c *class) Printi(){
	fmt.Println(c.i)
}

func main(){
	var v value
	v=&class{}
	v.Increment()
	v.Printi()
	v.Increment()
        v.Printi()
	v.Increment()
        v.Printi()
}
