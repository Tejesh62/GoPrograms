package nodevalue
import(
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next

	}

	return dummy.Next
}

func BuildList(nums []int) *ListNode {
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
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// func main() {
// 	list := BuildList([] int{1, 2, 3, 4})
// 	swapped := SwapPairs(list)
// 	PrintList(swapped)
// }

