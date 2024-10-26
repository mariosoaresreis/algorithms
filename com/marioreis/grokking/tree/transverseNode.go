package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addNode(lastLevel []*TreeNode, result *[][]int) {
	if len(lastLevel) == 0 {
		return
	}

	newLevel := make([]int, 0)
	level := make([]*TreeNode, 0)

	for _, v := range lastLevel {
		if v.Left != nil {
			newLevel = append(newLevel, v.Left.Val)
			level = append(level, v.Left)
		}

		if v.Right != nil {
			newLevel = append(newLevel, v.Right.Val)
			level = append(level, v.Right)
		}
	}

	if len(newLevel) > 0 {
		*result = append(*result, newLevel)
		addNode(level, result)
	}
}
