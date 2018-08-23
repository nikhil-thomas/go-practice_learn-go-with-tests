package arrays

// Sum computes sum of [5]int
func Sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// SumAll calculates sum of elements of mutliple slices
func SumAll(ns ...[]int) []int {
	sums := []int{}
	for _, v := range ns {
		sums = append(sums, Sum(v))
	}
	return sums
}

// SumAllTails calculates sum of tails of mutliple slices
func SumAllTails(ns ...[]int) []int {
	sums := []int{}
	for _, v := range ns {
		if len(v) > 1 {
			sums = append(sums, Sum(v[1:]))
		} else {
			sums = append(sums, 0)
		}

	}
	return sums
}
