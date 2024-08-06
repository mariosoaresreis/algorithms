package subsets

import "fmt"

func FindPermutation(nums []int, pos int, result *[][]int) {
	if pos == len(nums) {
		res := make([]int, len(nums))
		copy(res, nums)
		*result = append(*result, res)
		println(fmt.Sprintf("%v", nums))
	}

	for i := pos; i < len(nums); i++ {
		swap(nums, i, pos)
		FindPermutation(nums, pos+1, result)
		swap(nums, i, pos)
	}
}

func swap(nums []int, i int, pos int) {
	nums[i], nums[pos] = nums[pos], nums[i]
}
