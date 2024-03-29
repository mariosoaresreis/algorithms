package domain

func RemoveDuplicates(nums []int) int {

	n := len(nums)

	if n == 0 {
		return 0
	}

	result := make([]int, len(nums))
	copy(result, nums)

	i := 0
	prevNumber := -101

	for i < len(result) {
		if prevNumber == result[i] {
			//result = slices.Concat(result[0:i], result[i+1:])
		} else {
			i++
		}

		if i >= 1 {
			prevNumber = result[i-1]
		}
	}

	for index, _ := range result {
		nums[index] = result[index]
	}

	nums = nums[0:len(result)]

	return len(nums)
}
