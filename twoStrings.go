package main

import "fmt"

func print[T interface{}](output ...T){
	fmt.Println(output)
}
func main()  {
	//print(1,23,"4")
	print(1,2,3,4)
}
