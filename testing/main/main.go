package main

import (
	"fmt"
	"github.com/techagentng/testPrac/testPackage"

)



func main(){
		s:= []int{1,2,3,4,45,6,7,8,67,4}
		d := testPackage.MyAverage(s...) //unfurling
		fmt.Println(d)

		sum := testPackage.MySum(s...)
		fmt.Println(sum)
}