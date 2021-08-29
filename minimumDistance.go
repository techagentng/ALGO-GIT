package main

//O(n) time complexity 
//O(n) space complexity
//Please post your solution here if you achieved O(n) time and O(1) space 
func minimumDistances(a []int32) int32 {
    minimum := int32(math.MaxInt32)
    myMap := map[int32]int{}
    for i, v := range a {
        pastIdx := myMap[v]
        if pastIdx != 0{
            if minimum > int32(i+1-pastIdx){
                minimum = int32(i+1-pastIdx)
            }
        }else {
            myMap[v] = i+1
        }
    }
    if minimum == math.MaxInt32{
        return -1
    }
    return minimum
}
