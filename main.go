package main

import (
	"fmt"

	"codes/bitmapholes"
	"codes/combination"
	"codes/equalparts"
	"codes/lastword"
	"codes/maxgap"
	"codes/minstack"
	"codes/palindrome"
	"codes/permutations"
	"codes/sorting"
	"codes/spiralorder"
	"codes/subset"
	"codes/matrix"
	"codes/search"
)

func main() {
	// BitmapHoles test
	bitmapInput := []string{
		"01111",
		"01001",
		"01001",
		"01111",
	}
	fmt.Println("Number of holes:", bitmapholes.BitmapHoles(bitmapInput))

	// Longest Palindrome test
	fmt.Println("Longest palindrome in 'babad':", palindrome.LongestPalindrome("babad"))
	fmt.Println("Longest palindrome in 'cbbd':", palindrome.LongestPalindrome("cbbd"))

	// Permutations test
	fmt.Println("Permutations of [7, 9, 15]:", permutations.Permute([]int{7, 9, 15}))
	fmt.Println("Permutations of [0, 1]:", permutations.Permute([]int{0, 1}))

	//lastword
	fmt.Println(lastword.LengthOfLastWord("Hello World"))
	fmt.Println(lastword.LengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lastword.LengthOfLastWord("luffy is still joyboy"))

	//combination
	fmt.Println(combination.Combination([]int{2, 3, 6, 7}, 7))
	fmt.Println(combination.Combination([]int{2, 3, 5}, 8))
	fmt.Println(combination.Combination([]int{2}, 1))

	//subset
	fmt.Println(subset.Subset([]int{1, 2, 3}))
	fmt.Println(subset.Subset([]int{0}))

	//minstack
	ms := minstack.Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())

	//spiralorder
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println("Output for matrix1:", spiralorder.SpiralOrder(matrix1))
	fmt.Println("Output for matrix2:", spiralorder.SpiralOrder(matrix2))

	//equalparts
	arr1 := []int{1, 0, 1, 0, 1}
	fmt.Println("Input:", arr1, "Output:", equalparts.ThreeEqualParts(arr1))

	arr2 := []int{1, 1, 0, 1, 1}
	fmt.Println("Input:", arr2, "Output:", equalparts.ThreeEqualParts(arr2))

	arr3 := []int{1, 1, 0, 0, 1}
	fmt.Println("Input:", arr3, "Output:", equalparts.ThreeEqualParts(arr3))

	//maxgap
	fmt.Println(maxgap.MaximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maxgap.MaximumGap([]int{10}))

	//sorting

	head := sorting.CreateList([]int{4, 2, 1, 3})
	sorted := sorting.SortList(head)
	sorting.PrintList(sorted)

	head2 := sorting.CreateList([]int{-1, 5, 3, 4, 0})
	sorted2 := sorting.SortList(head2)
	sorting.PrintList(sorted2)

	head3 := sorting.CreateList([]int{})
	sorted3 := sorting.SortList(head3)
	sorting.PrintList(sorted3)

	//matrix
	mat1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	mat2 := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}

	fmt.Println("Output 1:")
	matrix.PrintMatrix(matrix.UpdateMatrix(mat1))

	fmt.Println("Output 2:")
	matrix.PrintMatrix(matrix.UpdateMatrix(mat2))

	//search
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(search.Exist(board, "ABCCED")) 
	fmt.Println(search.Exist(board, "SEE"))    
	fmt.Println(search.Exist(board, "ABCB"))   

}
