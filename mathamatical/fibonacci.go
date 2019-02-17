package main
import "fmt"

func fibo(n int){
	a:=1
	b:=1
	for i:=1;i<=n;i++{
		if i==1 {
			fmt.Println(a)
		} else if i ==2 {
			fmt.Println(b)
		}else {
		        a,b = b,a+b
			fmt.Println(b)
		}
	}
}
func main() {
	var n int
	fmt.Println("Enter the number of element of the FIBONAcci series")
	fmt.Scanln(&n)
	fmt.Println("The Fibonacci series till ",n," elements is")
	fibo(n)
}
