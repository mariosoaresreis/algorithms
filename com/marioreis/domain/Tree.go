package domain

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
