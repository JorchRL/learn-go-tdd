package array

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		// cannot take the tail of an empty slice. Return 0.
		if len(slice) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := slice[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}