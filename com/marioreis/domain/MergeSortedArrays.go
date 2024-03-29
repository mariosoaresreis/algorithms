package domain

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, len(nums1))
	i := 0
	j := 0
	index := 0

	if len(nums1) == 0 && len(nums2) > 0 {
		copy(nums1, nums2)
		return
	}

	if len(nums1) > 0 && len(nums2) == 0 {
		return
	}

	if len(nums1) == 0 && len(nums2) == 0 {
		return
	}

	if len(nums1) == 2 && len(nums2) == 1 {
		if nums2[0] >= nums1[0] {
			nums1[1] = nums2[0]
		} else {
			nums[0] = nums2[0]
			nums[1] = nums1[0]
			copy(nums1, nums)
		}
		return
	}

	for index < len(nums1) {
		if nums1[i] == 0 {
			nums[index] = nums2[j]
			j++
			index++
		} else if nums1[i] < nums2[j] {
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
