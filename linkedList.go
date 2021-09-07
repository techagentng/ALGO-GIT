package main

import "fmt"

//import "fmt"

/*
Building the list and adding data in nodes
part1: a template to create a node and to create a list len(data, next)
part2: is to init the head and init the next

*/
type node struct{
	data int
	next *node
}

type llist struct{
	head *node
	length int
}
//method to add nodes to: linklist has a head to house a node and pointer to the next node
func (l *llist) prepend(n *node) {
	//fmt.Println("ooooooooooooo",n)
	//Define a place to put n
	Second := l.head
	//Place pointer n node inside its position head
	l.head = n
	//fmt.Println("xxxxxxxxxx",n)
	//Now that n is in its position, define what comes next which is second
	l.head.next = Second
	//
	l.length++
}

//A copy of llist struct
func (l *llist) printListData(){
	toPrint := l.head // head is inside toPrint
	//counting down from the highest l.length
	for l.length != 0{
		fmt.Printf("%d ", toPrint.data)
		toPrint= toPrint.next
		l.length -- 
	}
}

//Delete a particular node value
func (l llist) deleteWithVal(val int){
	//handle null list value
	prevToDelete := l.head
	//comparing values in the next node
	for prevToDelete.next.data != val{
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}
// // Given numbers in a linked list
// type Node struct {
//    value int
//    next *Node
// }
// func NewNode(value int, next *Node) *Node{
//    var n Node
//    n.value = value
//    n.next = next
//    return &n
// }
// func CountNodes(head *Node){
//    fmt.Printf("Input Linked List is: ")
//    count :=0
//    temp := head
//    for temp != nil {
//       fmt.Printf("%d ", temp.value)
//       temp = temp.next
//       count += 1
//    }
//    fmt.Printf("\nNumber of nodes in the linked list is: %d\n", count)
// }




func main(){
	myList := llist{}
    node1 := &node{data: 40}
	node2 := &node{data: 33}
	node3 := &node{data: 30}
	node4 := &node{data: 3}
	node5 := &node{data: 6}
	node6 := &node{data: 23}
	node7 := &node{data: 73}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	fmt.Printf("full list %v", myList)
	myList.printListData()

	/******************SECOND CODE CALL****************************/
	//head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	//CountNodes(head)
}