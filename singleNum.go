
/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.
url: https://leetcode.com/problems/single-number/

solution: A better solution is to use hashing. 
1) Traverse all elements and put them in a map table. Element is used as key and the count of occurrences is 
used as the value in the hash table.
XOR url: https://play.golang.org/p/NhY-nDHDNF4 
https://play.golang.org/p/h0WSrT-5eFA
https://play.golang.org/p/gaRXS5JmuLG ......Checking for the key in the map
https://play.golang.org/p/LDQ19Vdcl3c ......Checking for the key in the map
https://play.golang.org/p/QRwMI0eoYYz  .....A map outputed by its key and value
https://divrhino.com/articles/build-command-line-tool-go-cobra/...... Set up cobra article
*/


package main
//import "fmt"
func singleNumber(nums []int) int32{
    a := int32(0)
    for _, i2 := range nums {
        a ^= i2
    }
    return a
 }

 