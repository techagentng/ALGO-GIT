package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(checksl(3, []int{2, 1, 5, 1, 3, 2}))

}
func checksl(k int, arr []int) int {
	var sum int
	var total []int
	winStart := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i >= k-1 {
			total = append(total, sum)
			sum -= arr[winStart]
			winStart += 1
		}
	}
	//loop thru total
	max := arr[0]
	min := arr[0]
	for _, v := range total {
		if v < max {
			min = v
		}
		if v > min {
			max = v
		}
	}
	return max
}