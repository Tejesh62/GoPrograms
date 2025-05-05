package main

import (
	"fmt"

	"codes/addbinary"
	"codes/base7"
	"codes/binarygap"
	"codes/bitmapholes"
	"codes/colors"
	"codes/combination"
	"codes/count"
	"codes/distance"
	"codes/equalparts"
	"codes/excelsheet"
	"codes/foursum"
	"codes/graph"
	"codes/integers"
	"codes/keyboard"
	"codes/lastword"
	"codes/matrix"
	"codes/maxgap"
	"codes/minstack"
	"codes/mismatch"
	"codes/missingnumber"
	"codes/multiply"
	"codes/occurrences"
	"codes/palindrome"
	"codes/pathsum"
	"codes/permutations"
	"codes/powertwo"
	"codes/prefix"
	"codes/primenumbers"
	"codes/rectanglearea"
	"codes/reversebits"
	"codes/sametree"
	"codes/search"
	"codes/searchmatrix"
	"codes/sorting"
	"codes/spiralorder"
	"codes/subset"
	"codes/symmetric"
	"codes/transpose"
	"codes/traversal"
	"codes/tree"
	"codes/validation"
	"codes/version"
	"codes/wordbreak"
	"codes/wordmatch"
	"codes/counting"
	"codes/phone"

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

	//clone graph

	node1 := &graph.Node{Val: 1}
	node2 := &graph.Node{Val: 2}
	node1.Neighbors = []*graph.Node{node2}
	node2.Neighbors = []*graph.Node{node1}

	cloned := graph.CloneGraph(node1)

	fmt.Println("Original Node:", node1.Val)
	fmt.Println("Cloned Node:", cloned.Val)

	fmt.Println("Original Neighbor:", node1.Neighbors[0].Val)
	fmt.Println("Cloned Neighbor:", cloned.Neighbors[0].Val)

	// Longest common prefix

	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println("Common Prefix:", prefix.Prefix(strs1))

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println("Common prefix:", prefix.Prefix(strs2))

	// First Bad Version

	version.BadVer = 4
	n := 5
	fmt.Println("First bad version for n=5, bad=4:", version.FindBadVersion(n))

	version.BadVer = 1
	n = 1
	fmt.Println("First bad version for n=1, bad=1:", version.FindBadVersion(n))

	//missing number

	fmt.Println(missingnumber.MissingNumber(([]int{3, 0, 1})))
	fmt.Println(missingnumber.MissingNumber(([]int{0, 1})))
	fmt.Println(missingnumber.MissingNumber(([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})))

	//word break
	fmt.Println(wordbreak.WordBreak("leetcode", []string{"leet","code"}))
	fmt.Println(wordbreak.WordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordbreak.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

	// Symmetric
	
	treeRoot := &symmetric.TreeNode{
		Val: 1,
		Left: &symmetric.TreeNode{
			Val: 2,
			Left: &symmetric.TreeNode{Val: 3},
			Right: &symmetric.TreeNode{Val: 4},
		},
		Right: &symmetric.TreeNode{
			Val: 2,
			Left: &symmetric.TreeNode{Val: 4},
			Right: &symmetric.TreeNode{Val: 3},
		},
	}

	fmt.Println("Is tree symmetric?", symmetric.Symmetric(treeRoot))

	//primenumbers

	fmt.Println(primenumbers.CountPrimes(10)) 
	fmt.Println(primenumbers.CountPrimes(0))  
	fmt.Println(primenumbers.CountPrimes(1))  

	//mismatch
		fmt.Println(mismatch.FindErrorNums([]int{1, 2, 2, 4})) 
		fmt.Println(mismatch.FindErrorNums([]int{1, 1}))   
		
	
	//Pathsum
		pathroot := &pathsum.TreeNode{Val: 5,
			Left: &pathsum.TreeNode{Val: 4,
				Left: &pathsum.TreeNode{Val: 11,
					Left:  &pathsum.TreeNode{Val: 7},
					Right: &pathsum.TreeNode{Val: 2},
				},
			},
			Right: &pathsum.TreeNode{Val: 8,
				Left:  &pathsum.TreeNode{Val: 13},
				Right: &pathsum.TreeNode{Val: 4,
					Right: &pathsum.TreeNode{Val: 1},
				},
			},
		}
	
		target := 22
		fmt.Println("Has path sum:", pathsum.HasPathSum(pathroot, target)) 
	

	//Keyboard
	fmt.Println(keyboard.MinSteps(3))
	fmt.Println(keyboard.MinSteps(1))
	fmt.Println(keyboard.MinSteps(9))
	fmt.Println(keyboard.MinSteps(100))
	
	//ReverseBits
	fmt.Println(reversebits.ReverseBits(43261596))       
	fmt.Println(reversebits.ReverseBits(4294967293))    
	
	//rectnagle area

	fmt.Println(rectanglearea.ComputeArea(-3, 0, 3, 4, 0, -1, 9, 2)) // Output: 45
	fmt.Println(rectanglearea.ComputeArea(-2, -2, 2, 2, -2, -2, 2, 2)) // Output: 16
	
	// Transpose Matrix

	matrix1 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Transpose of matrix1:", transpose.Transpose(matrix1))

	matrix2 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Transpose of matrix2:", transpose.Transpose(matrix2))

	// Binary Gap

	fmt.Println(binarygap.BinaryGap(22)) 
	fmt.Println(binarygap.BinaryGap(8))  
	fmt.Println(binarygap.BinaryGap(5)) 

	// Occurences After Bigram

	text1 := "alice is a good girl she is a good student"
	first1 := "a"
	second1 := "good"
	fmt.Println(occurrences.FindOccurrences(text1, first1, second1))
	
	text2 := "we will we will rock you"
	first2 := "we"
	second2 := "will"
	fmt.Println(occurrences.FindOccurrences(text2, first2, second2))

	//excelsheet

	fmt.Println(excelsheet.TitleNumber(("A")))
	fmt.Println(excelsheet.TitleNumber("AB"))
	fmt.Println(excelsheet.TitleNumber("ZY"))

	// Counting Bits

	fmt.Println(counting.CountBits(2)) 
	fmt.Println(counting.CountBits(5))

	// Letter Combinations of a Phone Number

	fmt.Println(phone.Letter("23"))
	fmt.Println(phone.Letter(""))
	fmt.Println(phone.Letter("2"))
	
}
