package maxsubarraylc53

func maxSubArray(nums []int) int {
	allTimeMax := nums[0]
	for i := 0; i < len(nums); i++ {
		currentTotal := nums[i]
		currentMax := currentTotal
		for j := i + 1; j < len(nums); j++ {
			if currentTotal += nums[j]; currentTotal > currentMax {
				currentMax = currentTotal
			}
		}
		if currentMax > allTimeMax {
			allTimeMax = currentMax
		}
	}
	return allTimeMax
}
