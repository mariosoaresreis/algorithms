package main

import (
	"fmt"
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

type Task struct {
	Value *ChecklistYesNo
}

// deploy
func main() {
	t := Test{
		Value: ChecklistYesNoYes,
	}

	value := ChecklistYesNoYes

	task := Task{
		Value: &value,
	}

	println(*task.Value == t.Value)
}

func findSumOfDigits(num int) int {
	if num == 0 {
		return 0
	}

	sum := 0
	var digit int

	for num > 0 {
		digit = num % 10
		sum += digit
		num /= 10
	}

	return sum
}

// findElementsWithSum of k from arr of size
func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}
