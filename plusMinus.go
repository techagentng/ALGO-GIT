package main

import "fmt"

//"strconv"
//"strings"

var arr = []int32{-4, 3, -9, 0, 4, 1, -9}
func plusMinus(arr []int32){
	ncount := float64(0)
	positive := float64(0)
	zero := float64(0)
 for i:=0; i < len(arr); i++{
	 if arr[i] < 0 {
		ncount++
	 } else if arr[i] > 0{
		 positive++
	 }else{
		 zero++
	 }
 }
  a:=float64(len(arr))
 
  b:=fmt.Sprintf("%.6f", positive/a)
  fmt.Println(b)
//   fmt.Println(zero/a)
}

func main(){

plusMinus(arr)
}