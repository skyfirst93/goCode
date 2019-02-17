package main

import "fmt"

type node struct {
	a    int
	next *node
}

type llist struct {
	head *node
}

func insert(ll *llist, val int) {
	if ll.head == nil {
		ll.head = &node{a:val,next:nil}
		return
	}
	if ll.head.next==nil {
		ll.head.next = &node{a:val,next:nil}
		return
	}
	currentnode := ll.head
	for currentnode.next != nil {
		currentnode = currentnode.next
	}
	currentnode.next = &node{a:val,next:nil}
}
func printlist(ll *llist) {
	if ll.head == nil {
		fmt.Println("list is empty")
		return
	}
	currentnode := ll.head
	for currentnode.next != nil {
		fmt.Println(currentnode.a)
		currentnode=currentnode.next
	}
	fmt.Println(currentnode)
}
func main() {
	ll := llist{}
	insert(&ll, 1)
	insert(&ll, 2)
	insert(&ll, 3)
	insert(&ll, 4)
	printlist(&ll)
}
