package main

import "fmt"
import "math"

func main(){
	var area,r float64
	fmt.Println("Enter the radius of the circle")
	fmt.Scanf("%f",&r)
	a := float64(math.Pi)
        fmt.Println("Pi value",a)
        circum := float64(2*a*r)
	area = r*r
	fmt.Println("area",area,"circum",circum)
}
