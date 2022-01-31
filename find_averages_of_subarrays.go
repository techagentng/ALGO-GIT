package main
import ("fmt")
func find_averages_of_subarrays(k int, s []int)[]int{
	var slc []int
	for i := 0; i <= len(s)-k; i++{
		sum := 0
		for j :=i; j < i+k; j++{
			sum += s[j]
		}
		slc = append(slc, sum/k)
	}
	return slc
}
func main() {
	fmt.Println(find_averages_of_subarrays(5, []int{1, 3, 2, 6, -1, 4, 1, 8, 2}))
}