package subsets

func FindSubsetsWithDuplicates(nums []int) [][]int {
	subsets := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			subsets = append(subsets, nums[i:j+1])
		}
	}
	return subsets
}
