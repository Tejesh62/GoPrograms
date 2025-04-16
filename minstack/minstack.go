package minstack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)

	if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	} else {
		
		ms.minStack = append(ms.minStack, ms.minStack[len(ms.minStack)-1])
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return -1 
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return -1 
	}
	return ms.minStack[len(ms.minStack)-1]
}


// func main() {
// 	ms := minstack.Constructor()
// 	ms.Push(-2)
// 	ms.Push(0)
// 	ms.Push(-3)
// 	fmt.Println(ms.GetMin()) 
// 	ms.Pop()
// 	fmt.Println(ms.Top())    
// 	fmt.Println(ms.GetMin())
// }