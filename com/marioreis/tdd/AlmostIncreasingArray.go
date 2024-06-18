package tdd

import "testing"

func Test_AlmostIncreasingArray(t *testing.T) {
	if !almostIncreasingArray([]int{10, 1, 2, 3, 4}) {
		t.Error("Test A failed")
	}

	if almostIncreasingArray([]int{1, 2, 1, 2, 3}) {
		t.Error("Test B failed")
	}

	if !almostIncreasingArray([]int{1, 2, 3, 5, 4, 5}) {
		t.Error("Test C failed")
	}

	if !almostIncreasingArray([]int{1, 2, 3, 6, 5}) {
		t.Error("Test D failed")
	}

	if almostIncreasingArray([]int{1, 2, 5, 5, 5}) {
		t.Error("Test E failed")
	}

}

func almostIncreasingArray(sequence []int) bool {
	lastNum := sequence[0]
	count := 0

	if len(sequence) <= 2 {
		return true
	}

	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= lastNum {
			count++

			if i == 1 {
				lastNum = sequence[i]
				continue
			}

			if len(sequence) >= 3 && i >= 1 && sequence[i] == sequence[i-1] && sequence[i] == sequence[i+1] {
				return false
			}

			if i <= len(sequence)-2 && sequence[i-1] >= sequence[i+1] {
				if i >= 2 && sequence[i] <= sequence[i-2] {
					return false
				}

				if i >= 2 {
					lastNum = sequence[i-2]
					continue
				} else {
					lastNum = sequence[i-1]
					continue
				}

			}

			if i <= len(sequence)-2 && i >= 1 && sequence[i-1] <= sequence[i+1] {
				lastNum = sequence[i-1]
				continue
			}

		}

		lastNum = sequence[i]
	}

	return count < 2
}
