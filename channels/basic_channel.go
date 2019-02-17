package main
import "fmt"

func main(){
	qs := make(chan bool)
	go func(){
		qs <- true
		fmt.Println("Inside go routine")
	}()
	fmt.Println("outside")
	value,ok := <-qs
	fmt.Println("value:", value," is ",ok)
}
