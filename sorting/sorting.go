package sorting

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMiddle(head)
	right := mid.Next
	mid.Next = nil

	left := SortList(head)
	right = SortList(right)

	return merge(left, right)
}

func getMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}

func CreateList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range nums {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

func PrintList(head *ListNode) {
	for head != nil {
		print(head.Val, " ")
		head = head.Next
	}
	println()
}

// func main() {
// 	head := sorting.CreateList([]int{4, 2, 1, 3})
// 	sorted := sorting.SortList(head)
// 	sorting.PrintList(sorted)

// 	head2 := sorting.CreateList([]int{-1, 5, 3, 4, 0})
// 	sorted2 := sorting.SortList(head2)
// 	sorting.PrintList(sorted2)

// 	head3 := sorting.CreateList([]int{})
// 	sorted3 := sorting.SortList(head3)
// 	sorting.PrintList(sorted3)

// }