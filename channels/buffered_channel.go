package main


func main(){
	buff := make(chan int,5)
	buff <- 1
	buff <- 2
	buff <- 3
	buff <- 4
	buff <- 6
}
