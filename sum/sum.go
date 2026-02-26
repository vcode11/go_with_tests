package sum

func Sum(numArray []int) int {
	var total int
	for _, num := range numArray {
		total += num
	}
	return total
}

func SumAll(slices [][]int) []int{
	output := make([]int, len(slices))
	for i, slice := range slices {
		output[i] = Sum(slice)
	}
	return output
}

func SumAllTails(slices ...[]int) []int {
	output := SumAll(slices)
	for i := range output {
		if len(slices[i]) > 0 {
			output[i] -= slices[i][0]
		}
	}
	return output
}