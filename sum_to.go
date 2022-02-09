package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(checksl(7, []int{2, 1, 5, 2, 3, 2}))

}
func checksl(s int, arr []int) []int {
	total := []int{}
	sum := 0 
	for i:=0; i < len(arr); i++ {
		sum += arr[i]
		if sum >= s {
			total = append(total, sum)
			
		} 
	}
	return  total
}