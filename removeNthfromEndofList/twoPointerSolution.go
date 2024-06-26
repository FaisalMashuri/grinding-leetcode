package removeNthfromEndofList

type ListNode struct {
	Val  int
	Next *ListNode
}

func TwoPointerSolution(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
