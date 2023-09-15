package sortcolorslc75

func sortColors(nums []int) {
	// blue = 0 (lowPtr), white = 1 (midPtr), red = 2 (highPtr)
	var lowPtr int           // low and mid pointer start at index 0
	highPtr := len(nums) - 1 // high pointer starts at last index

	// Continue running until the midPtr goes to the right of highPtr
	for midPtr := 0; midPtr <= highPtr; {
		// if the color is blue, swap it with the low pointer
		// and increment both the mid and low pointer
		if nums[midPtr] == 0 {
			nums[midPtr], nums[lowPtr] = nums[lowPtr], nums[midPtr]
			lowPtr++
			midPtr++
			// If the color is white, Don't do any swaps and just increment the mid Pointer
		} else if nums[midPtr] == 1 {
			midPtr++
		} else {
			// if the color is red, swap it with the high pointer
			// and decrement the high pointer
			nums[midPtr], nums[highPtr] = nums[highPtr], nums[midPtr]
			highPtr--
		}
	}
}
