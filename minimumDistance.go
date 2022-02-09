// Given an array of positive numbers and a positive number
// ‘S,’ find the length of the smallest contiguous subarray
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
	minL := 0
	start := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		for sum >= 8 {
			if minL < i-start+1 {
				minL = i - start + 1
			}
			sum -= arr[start]
			start++
		}
		if minL == 0 {
			return 0
		}
	}
	return minL
}
