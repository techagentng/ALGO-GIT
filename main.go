package main

import (

	"fmt"
)
    

func main() {
	arr1 := []int{}
	arrz := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//Loop through arrz in rean parse half to arr1

	for i := 0; i < len(arrz)/2; i++ {
		arr1 = append(arr1, arrz[i])
	}
	fmt.Println(arr1)
}

