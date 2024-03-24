package domain

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, len(nums1))
	i := 0
	j := 0
	index := 0

	for index < m {
		if nums1[i] < nums2[j] {
			nums[index] = nums1[i]
			i++
			index++
		} else if nums2[j] < nums1[i] {
			nums[index] = nums2[j]
			index++
			j++
		} else {
			nums[index] = nums2[j]
			index++
			nums[index] = nums1[i]
			index++
			i++
			j++
		}

	}

	copy(nums1, nums)
}
