package domain

import "fmt"

type Tree struct {
	Left  *Tree
	Right *Tree
	Data  int32
}

func NewTree(left *Tree, right *Tree, data int32) (result Tree) {
	return Tree{
		Left:  left,
		Right: right,
		Data:  data,
	}
}

func FindTree(toFind int32, tree *Tree) (result int) {
	if tree == nil {
		return 0
	}

	if tree.Left == nil && tree.Right == nil {
		return 0
	}

	if tree.Left.Data == toFind {
		return 1
	}

	if tree.Right.Data == toFind {
		return 1
	}

	if FindTree(toFind, tree.Left) == 1 {
		return 1
	} else {
		return FindTree(toFind, tree.Right)
	}

	return 0
}

func TestBinarySearch() {
	fmt.Println("Testing binary search...")
	var tree5 Tree = NewTree(nil, nil, 5)
	var tree8 Tree = NewTree(nil, nil, 8)
	var tree15 Tree = NewTree(nil, nil, 15)
	var tree13 Tree = NewTree(nil, nil, 13)
	var tree10 Tree = NewTree(&tree5, &tree8, 10)
	var tree20 Tree = NewTree(&tree15, &tree13, 20)
	var tree1 Tree = NewTree(&tree10, &tree20, 1)

	fmt.Println("Finding data = 2")
	fmt.Printf("Result: %d", FindTree(2, &tree1))
	fmt.Println("Finding data = 20")
	fmt.Printf("Result: %d", FindTree(20, &tree1))
	fmt.Println("Finding data = 15")
	fmt.Printf("Result: %d", FindTree(15, &tree1))
	fmt.Println("Finding data = 16")
	fmt.Printf("Result: %d", FindTree(16, &tree1))
}
