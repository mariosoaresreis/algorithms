package FastSlowPointers

func ReverseLinkedList() {
	node5 := &ListNode{
		Val:  5,
		Next: nil,
	}

	node4 := &ListNode{
		Val:  4,
		Next: node5,
	}

	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}

	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}

	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}

	currentNode := node1
	var prev *ListNode
	array := make([]*ListNode, 0)

	for currentNode != nil {
		array = append(array, currentNode)
		next := currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next

	}

}
