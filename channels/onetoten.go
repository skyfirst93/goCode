package main

import (
	"fmt"
	"time"
)

func main() {
	ic := make(chan int)
	go increment(ic)
	for i := range ic {
		fmt.Println(i)
	}
}

func increment(ic chan int) {
	i := 0
	for i <= 10 {
		time.Sleep(3 * time.Second)
		ic <- i
		i++
	}
	close(ic)
}
