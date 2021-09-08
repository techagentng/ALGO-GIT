package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type lList struct{
	length int
	head *Node
}

func (l *lList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length ++
}

//Append Element at the end by transversing till we get to the end
func (l *lList) Append(val int)  {
	node := &Node{}
	node.data = val
	if l.length == 0{
		l.head = node
		l.length++
		return
	}

	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.next == nil {
			ptr.next = node
			l.length++
			return
		}
		ptr = ptr.next
	}
}

func main(){
	//node1 := &Node{data: 30}
	//node2 := &Node{data: 40}
	//node3 := &Node{data: 60}
	newList := lList{}
	newList.Append(600)
	newList.Append(80000)
	fmt.Println(newList)
	//newList.Prepend(node1)
	//newList.Prepend(node2)
	//newList.Prepend(node3)
	}