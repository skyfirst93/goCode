package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head *node
}

func main() {
	ll := &list{}
	ll.head = &node{data:2}
	node1 := &node{data:3}
	ll.head.next=node1
	fmt.Println("list", *ll.head)
	fmt.Println("list", *ll.head.next)
}
