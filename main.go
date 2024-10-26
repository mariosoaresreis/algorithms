package main

import (
	"errors"
	"fmt"
	"sort"
)

func secondLargestUnique(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("not enough elements")
	}

	// Create a map to store unique numbers
	uniqueNums := make(map[int]struct{})

	// Populate the map with unique numbers from the slice
	for _, num := range nums {
		uniqueNums[num] = struct{}{}
	}

	// Convert the map keys to a slice
	var uniqueSlice []int
	for num := range uniqueNums {
		uniqueSlice = append(uniqueSlice, num)
	}

	// Check if there are at least two unique numbers
	if len(uniqueSlice) < 2 {
		return 0, errors.New("not enough unique elements")
	}

	// Sort the slice of unique numbers in descending order
	sort.Slice(uniqueSlice, func(i, j int) bool {
		return uniqueSlice[i] > uniqueSlice[j]
	})

	// Return the second largest unique number
	return uniqueSlice[1], nil
}

func main() {
	nums := []int{4, 1, 2, 3, 4, 2, 5}
	secondLargest, err := secondLargestUnique(nums)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The second largest unique number is:", secondLargest)
	}
}

// TreeNode structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{root.Val})

	return result
}

func addNode(lastLevel []*TreeNode, result *[][]int) {
	newNevel := make([]int, 0)
	level := make([]*TreeNode, 0)

	for _, v := range lastLevel {
		if v.Left != nil {
			newNevel = append(newNevel, v.Left.Val)
			level = append(level, v.Left)
		}

		if v.Right != nil {
			newNevel = append(newNevel, v.Right.Val)
			level = append(level, v.Right)
		}
	}

	*result = append(*result, newNevel)
	addNode(level, result)
}
