/* Innitialize newlist to || followed by || iinit of a new node*/
//Node
//Prepend method
// list result
package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type linkedlist struct{
	head *Node 
	length int
}

//Linked list
func (l *linkedlist) prepend(n *Node){
	second := l.head
	l.head = n 
	l.head.next = second
	l.length ++
	//fmt.Println("xxxxxxx", n) 
}

//Print data
//Check if len 0 val
//set var to var.next
// decreament 
func (l linkedlist) printList(){
	newlist := l.head
	for l.length != 0 {
		fmt.Printf("%d ", newlist.data)
		newlist = newlist.next
		l.length --
	}
	fmt.Printf("\n")
}

func main(){
	myList := linkedlist{}
	node1 := &Node{data: 90}
	node2 := &Node{data: 10}
	node3 := &Node{data: 40}
	node4 := &Node{data: 30}
	//fmt.Println(node1.next)
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.printList()
}