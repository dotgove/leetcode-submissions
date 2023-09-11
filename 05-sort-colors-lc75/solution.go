package sortcolorslc75

// Naive
func sortColors(nums []int) {
	var color0, color1, color2 int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			color0 += 1
		} else if nums[i] == 1 {
			color1 += 1
		} else {
			color2 += 1
		}
	}

	for i := 0; i < color0; i++ {
		nums[i] = 0
	}
	for i := color0; i < color0+color1; i++ {
		nums[i] = 1
	}
	for i := color0 + color1; i < len(nums); i++ {
		nums[i] = 2
	}
}
