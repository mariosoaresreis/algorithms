package main

import (
	"math/rand/v2"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Test struct {
	Value interface{} `json:"id"`
	Scale int         `json:"scale"`
}

type ChecklistYesNo string

const (
	ChecklistYesNoYes ChecklistYesNo = "YES"
	ChecklistYesNoNo  ChecklistYesNo = "NO"
	ChecklistYesNoNA  ChecklistYesNo = "NOT_APPLICABLE"
)

func areEqualArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[string]int)

	for _, char := range a {
		counts[char]++
	}

	for _, char := range b {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

/*
func equalArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[string]int)
	mapB := make(map[string]int)

	for i := range a {
		mapA[a[i]]++
	}

	for i := range b {
		mapB[b[i]]++
	}

	if len(mapA) != len(mapB) {
		return false
	}

	for k := range mapA {
		if mapB[k] != mapA[k] {
			return false
		}
	}

	return true
}

*/

type Task struct {
	Value *ChecklistYesNo
}

func isAlphanumeric(str string) bool {
	var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return alphanumeric.MatchString(str)

}

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	leftWindow := 0

	for i := len(s) - 1; i >= 0; {
		println(string(rune(s[i])))
		println(string(rune(s[leftWindow])))

		for !isAlphanumeric(string(rune(s[i]))) {
			i--
		}

		for !isAlphanumeric(string(rune(s[leftWindow]))) {
			leftWindow++
		}

		if strings.ToUpper(string(rune(s[i]))) != strings.ToUpper(string(rune(s[leftWindow]))) {
			return false
		}

		i--
		leftWindow++
	}

	return true
}
func longestSubarray(nums []int) int {
	maxOneCount := 0
	left := 0
	maxResult := 0
	zeroCount := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 1 {
			maxOneCount++
		} else {
			zeroCount++
		}

		if right-left+1-maxOneCount > 1 {
			if nums[left] == 1 {
				maxOneCount--
			} else {
				zeroCount--
			}

			left++
		}

		sumWindow := right - left + 1

		if sumWindow > maxResult {
			maxResult = sumWindow
		}

	}

	if zeroCount > 0 {
		return maxResult
	}

	return maxResult - 1
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// deploy
func main() {
	result := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				work := DoWork()
				result <- work

			}()
		}

		wg.Wait()
		close(result)
	}()

	for i := range result {
		println(i)
	}
}

func DoWork() int {
	time.Sleep(1 * time.Second)
	result := rand.IntN(100)
	return result
}

func maxVowels(s string, k int) int {
	mapVowels := make(map[rune]int)
	mapVowels['a'] = 1
	mapVowels['e'] = 1
	mapVowels['i'] = 1
	mapVowels['o'] = 1
	mapVowels['u'] = 1
	left := 0
	maxCount := 0

	for right := 0; right < len(s) && left < len(s); {
		windowSize := right - left + 1

		if windowSize > k {
			left++

			if right < left {
				right = left
			}

			continue
		}

		if !isVowel(rune(s[left]), mapVowels) {
			left++
			if right < left {
				right = left
			}
			continue
		}

		if !isVowel(rune(s[right]), mapVowels) {

			left = right
			continue
		}

		maxCount = max(maxCount, windowSize)
		right++
	}

	return maxCount
}

func isVowel(char rune, m map[rune]int) bool {
	_, ok := m[char]
	return ok
}
