package addtwo

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

func BuildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return dummy.Next
}

func PrintList(head *ListNode) {
	for head != nil {
		print(head.Val)
		if head.Next != nil {
			print(" -> ")
		}

		head = head.Next
	}

	println()
}
