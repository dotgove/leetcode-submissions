package twonumbersumlc48

// Naive
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
loop:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0], result[1] = i, j
				break loop
			}
		}
	}
	return result
}
