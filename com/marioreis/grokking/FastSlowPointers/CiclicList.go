package FastSlowPointers

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	// Initialize two pointers, slow and fast, to traverse the list
	slow := head
	fast := head

	// Iterate through the list until fast reaches the end or a cycle is found
	for fast != nil && fast.Next != nil {
		// Move slow one step and fast two steps
		fast = fast.Next.Next
		slow = slow.Next

		// If slow and fast meet, a cycle is detected
		if slow == fast {
			return true // found the cycle
		}
	}

	// If no cycle is found, return false
	return false
}
