package main

import (
	"fmt"

	"codes/addbinary"
	"codes/base7"
	"codes/bitmapholes"
	"codes/colors"
	"codes/combination"
	"codes/count"
	"codes/distance"
	"codes/equalparts"
	"codes/foursum"
	"codes/integers"
	"codes/lastword"
	"codes/matrix"
	"codes/maxgap"
	"codes/minstack"
	"codes/multiply"
	"codes/palindrome"
	"codes/permutations"
	"codes/sametree"
	"codes/search"
	"codes/searchmatrix"
	"codes/sorting"
	"codes/spiralorder"
	"codes/subset"
	"codes/traversal"
	"codes/tree"
	"codes/validation"
	"codes/wordmatch"
	"codes/powertwo"
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

	//min distance

	fmt.Println(distance.MinDistance("horse", "ros"))
	fmt.Println(distance.MinDistance("intention", "execution"))

	// odd even tree

	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 10,
			Left: &tree.TreeNode{
				Val:   3,
				Left:  &tree.TreeNode{Val: 12},
				Right: &tree.TreeNode{Val: 8},
			},
		},
		Right: &tree.TreeNode{
			Val: 4,
			Left: &tree.TreeNode{
				Val:  7,
				Left: &tree.TreeNode{Val: 6},
			},
			Right: &tree.TreeNode{
				Val:   9,
				Right: &tree.TreeNode{Val: 2},
			},
		},
	}

	fmt.Println(tree.Tree(root))

	//base7
	fmt.Println(base7.ConvertBase(100))
	fmt.Println(base7.ConvertBase(-7))
	fmt.Println(base7.ConvertBase(0))

	//UTF-8 Validate

	fmt.Println(validation.Validate([]int{197, 130, 1}))
	fmt.Println(validation.Validate([]int{235, 140, 4}))

	//foursum
	nums1 := []int{1, 0, -1, 0, -2, 2}
	target1 := 0
	fmt.Println("Input:", nums1)
	fmt.Println("Target:", target1)
	fmt.Println("Output:", foursum.FourSum(nums1, target1))

	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	fmt.Println("\nInput:", nums2)
	fmt.Println("Target:", target2)
	fmt.Println("Output:", foursum.FourSum(nums2, target2))

	//count and say

	fmt.Println(count.CountAndSay(1))
	fmt.Println(count.CountAndSay(4))
	fmt.Println(count.CountAndSay(6))

	//divide two integers
	fmt.Println(integers.Divide(10, 3))
	fmt.Println(integers.Divide(7, -3))
	fmt.Println(integers.Divide(-2147483648, 3))

	// multiply strings
	fmt.Println(multiply.Multiply("2", "3"))
	fmt.Println(multiply.Multiply("123", "456"))
	fmt.Println(multiply.Multiply("0", "123456"))

	//Same Tree

	p := &sametree.TreeNode{Val: 1}
	p.Left = &sametree.TreeNode{Val: 2}
	p.Right = &sametree.TreeNode{Val: 3}

	q := &sametree.TreeNode{Val: 1}
	q.Left = &sametree.TreeNode{Val: 2}
	q.Right = &sametree.TreeNode{Val: 3}

	fmt.Println(sametree.SameTree(p, q))

	r := &sametree.TreeNode{Val: 1}
	r.Right = &sametree.TreeNode{Val: 2}

	fmt.Println(sametree.SameTree(p, r))

	// search 2d matrix

	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchmatrix.SearchMatrix(matrix, 5))
	fmt.Println(searchmatrix.SearchMatrix(matrix, 20))

	//Inorder Traversal

	inorderRoot := &traversal.TreeNode{Val: 1}
	inorderRoot.Right = &traversal.TreeNode{Val: 2}
	inorderRoot.Right.Left = &traversal.TreeNode{Val: 3}

	fmt.Println(traversal.OrderTraversal(inorderRoot))

	
	//word match

	fmt.Println(wordmatch.WordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordmatch.WordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordmatch.WordPattern("aaaa", "dog cat cat dog"))

	//add binary

	fmt.Println(addbinary.AddBinary("11", "1"))
	fmt.Println(addbinary.AddBinary("1010", "1011"))

	//sort colors
	nums := []int{2, 0, 2, 1, 1, 0}
	colors.SortColors(nums)
	fmt.Println(nums)

	//power of two
	fmt.Println(powertwo.PowerTwo(1))   
	fmt.Println(powertwo.PowerTwo(16))  
	fmt.Println(powertwo.PowerTwo(3))   

}
