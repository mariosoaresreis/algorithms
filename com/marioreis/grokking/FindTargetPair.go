package grokking

func FindTargetPair(arr []int, targetSum int) []int {
	// TODO: Write your code here
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i]+arr[j] == targetSum {
			return []int{i, j}
		}

		if arr[i]+arr[j] > targetSum {
			j--
		}

		if i >= j {
			break
		}

		if arr[i]+arr[j] < targetSum {
			i++
		}

	}

	return []int{-1, -1} // pair not found
}
