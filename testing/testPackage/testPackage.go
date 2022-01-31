package testPackage

func MyAverage(s ...int) int{
	slen := len(s)
	total := 0
	for _, v := range s {
		total += v
	}
	d := total/slen
	return d
}

func MySum(n ...int) int{
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
