package main

import (

	"fmt"
)
    
func halfReverse(){
	numRes1 := []int{}
	numRes2 := []int{}
	nums := []int{22, 23, 33, 34, 35, 76, 47, 8}
	//var real = []int{1,2}
	for i := (len(nums)-1)/2; i >= 0; i-- {
		numRes1 = append(numRes1, nums[i])
	}
	for i := len(nums)-1; i > len(nums)-4; i-- {
		numRes2 = append(numRes2, nums[i])
	} 

	//fmt.Println(numRes1)
	fmt.Println("check", numRes2)
	fmt.Println("check1", numRes1)
}

func main() {
	arr1 := []int{}
	arr2 := []int{}
	arrz := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//Loop through arrz in rean parse half to arr1

	for i := 0; i < len(arrz)/2; i++ {
		arr1 = append(arr1, arrz[i])
	}
	for i := len(arrz)/2; i < len(arrz); i++ {
		arr2 = append(arr2, arrz[i])
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	halfReverse()
}

