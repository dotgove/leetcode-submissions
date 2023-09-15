package nextpermutationlc31

func nextPermutation(nums []int) {
	// Comparing consecutive values from the right, find the index i with first smallest value
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	// If a value is found (i.e. i != -1), compare the values from the right of i with its value and
	// find the index j with first highest value. Swap these two values.
	if i >= 0 {
		j := len(nums) - 1
		for ; j > i && nums[j] <= nums[i]; j-- {
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	// Sort the sub array from the right index i in ascending order
	for j := len(nums[i+1:]); j > 0; j-- {
		for k := 0; k < j-1; k++ {
			if nums[i+1:][k] > nums[i+1:][k+1] {
				nums[i+1:][k], nums[i+1:][k+1] = nums[i+1:][k+1], nums[i+1:][k]
			}
		}
	}
}
