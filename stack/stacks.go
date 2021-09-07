package main

import "fmt"

//Stack that holds a slice of integers
type Stack struct{
	items []int
}

//Pop function to remove items from the stack
func (s *Stack) push(i int){
	s.items = append(s.items, i)
}

//Push function to remove items from the stack
func (s *Stack) pop() int{
	l:= len(s.items)-1
	toRemove:= s.items[l]
	s.items = s.items[:l]
	fmt.Println("torem", toRemove)
	return toRemove
}

func main(){
	myStack := Stack{}
	myStack.push(300)
	myStack.push(100)
	myStack.push(500)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
}