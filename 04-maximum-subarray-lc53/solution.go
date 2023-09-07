package maxsubarraylc53

// Kadane's Algo
func maxSubArray(nums []int) int {
	allTimeMax := nums[0]
	currentMax := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentMax += nums[i]; currentMax < nums[i] {
			currentMax = nums[i]
		}
		if currentMax > allTimeMax {
			allTimeMax = currentMax
		}
	}
	return allTimeMax
}
