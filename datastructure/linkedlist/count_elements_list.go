package main

import "fmt"

type Node struct{
	data int
	next *Node
}
type List struct{
	head *Node
}

func insert(ll *List,val int){
	if ll.head == nil {
		ll.head=&Node{data:val,next:nil}
		return
	}
	cnode := ll.head
	for cnode.next != nil {
		cnode=cnode.next
	}
	cnode.next=&Node{data:val,next:nil}
}

func counter(ll *List){
	count:=0
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}
	cnode := ll.head
	for cnode.next != nil {
		cnode=cnode.next
		count++
	}
	fmt.Println("No of element is the list are",count+1)
}
func main() {
	llist := &List{}
	counter(llist)
	insert(llist,1)
	insert(llist,2)
	insert(llist,3)
	counter(llist)
}
