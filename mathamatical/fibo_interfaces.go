package main

import "fmt"

func main(){
	n:=190
	a:=1
	b:=1
	for i:=1;i<=n;i++{
		if i==1{
			fmt.Print(a)
		}else if i==2 {
			fmt.Print(" ",b)
		}else{
			a,b=b,a+b
			fmt.Print(" ",b)
		}
	}

}
