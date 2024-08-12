package main


///////////////////////////////////////////////////

// 141. Linked List Cycle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	//Floyd's Cycle Finding Algorithm
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func arrToLinkList(array []int, pos int) *ListNode {

	if pos < 0 {
		return nil
	}

	head := &ListNode{Val: array[0]}

	current := head

	var lastNode *ListNode

	for i := 1; i < len(array); i++ {
		current.Next = &ListNode{Val: array[i]}
		current = current.Next

		if i == len(array)-1 {
			lastNode = current
		}
	}

	k := pos
	current = head

	for k > 0 {
		current = current.Next
		k--
	}

	lastNode.Next = current

	return head
}

func printLinkdList(head *ListNode, loopCnt int) {

	current := head

	for loopCnt > 0 && current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
		loopCnt--
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inp1, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp1 = strings.TrimSpace(inp1)
	inp2 = strings.TrimSpace(inp2)

	arr := strings.Split(inp1, " ")
	pos, err := strconv.Atoi(inp2)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	var arrNum []int

	for _, v := range arr {

		num, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			break
		}

		arrNum = append(arrNum, num)
	}
	loopCnt := len(arrNum)
	head := arrToLinkList(arrNum, pos)

	printLinkdList(head, loopCnt+1)

	r := hasCycle(head)

	fmt.Println(r)
}

////////////////////////////////////////////////////////

// 136. Single Number

func singleNumber(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {

}

//////////////////////////////////////////////////////////////

// 125. Valid Palindrome

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {

	s = strings.TrimSpace(s)

	realStr := ""

	for k := 0; k < len(s); k++ {
		if s[k] >= 'a' && s[k] <= 'z' {
			realStr += string(s[k])
		}
		if s[k] >= 'A' && s[k] <= 'Z' {
			realStr += strings.ToLower(string(s[k]))
		}
		if s[k] >= '0' && s[k] <= '9' {
			realStr += string(s[k])
		}
	}

	fmt.Println(realStr)

	i := 0
	j := len(realStr) - 1

	for i < j {
		if realStr[i] != realStr[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return

	}

	r := isPalindrome(inp)
	fmt.Println(r)

}

////////////////////////////////////////////////////////////////////////

121. Best Time to Buy and Sell Stock

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProfit(prices []int) int {
	minPrice := prices[0]

	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		curProfit := prices[i] - minPrice

		profit = max(curProfit, profit)

	}

	return profit
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")

	var arrNum []int

	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])

		if err != nil {
			fmt.Println("Invalid input:", arr[i])
			break
		}

		arrNum = append(arrNum, num)
	}

	r := maxProfit(arrNum)
	fmt.Println(r)

}

///////////////////////////////////////////////////////////////////////////

 // Pascal's Triangle 2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	array := make([][]int, rowIndex+1)

	array[0] = []int{1}

	for i := 1; i <= rowIndex; i++ {
		currentRow := make([]int, i+1)
		currentRow[0], currentRow[i] = 1, 1

		for j := 1; j < i; j++ {
			currentRow[j] = array[i-1][j-1] + array[i-1][j]
		}

		array[i] = currentRow
	}

	return array[rowIndex]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	rowIndex, err := strconv.Atoi(inp)

	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	r := getRow(rowIndex)

	fmt.Println(r)

}

/////////////////////////////////////////////////////////////////////////////////

// Pascal's Triangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		currentRow := []int{1}

		for j := 1; j < i; j++ {
			currentRow = append(currentRow, result[i-1][j-1]+result[i-1][j])
		}

		currentRow = append(currentRow, 1)
		result = append(result, currentRow)
	}

	return result
}

func main() {

}

////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum = targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func arrayToBT(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			nodes[i] = &TreeNode{Val: arr[i].(int)}
		}
	}

	var j = 0
	for i := 0; i < len(arr); i++ {
		left := 2*j + 1
		right := 2*j + 2

		if left < len(arr) && nodes[left] != nil {
			nodes[i].Left = nodes[left]
		}

		if right < len(arr) && nodes[right] != nil {
			nodes[i].Right = nodes[right]
		}

		j++

	}

	return nodes[0]

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("Error reading input: ", err)
		return
	}

	inp2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("Error reading input: ", err)
		return
	}

	inp = strings.TrimSpace(inp)
	inp2 = strings.TrimSpace(inp2)

	target, err := strconv.Atoi(inp2)

	if err != nil {
		return
	}

	arr := strings.Split(inp, " ")

	var arrNum []interface{}

	for _, val := range arr {
		if val == "null" {
			arrNum = append(arrNum, nil)
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid input:", val)
			break
		}

		arrNum = append(arrNum, num)
	}

	root := arrayToBT(arrNum)

	r := hasPathSum(root, target)

	fmt.Println(r)

}

////////////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1

}

func arrToBBT(arr []interface{}) *TreeNode {

	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			nodes[i] = &TreeNode{Val: arr[i].(int)}
		}
	}
	var j = 0

	for i := 0; i < len(arr); i++ {

		if nodes[i] == nil {
			continue
		}

		left := 2*j + 1
		right := 2*j + 2

		if left < len(nodes) && nodes[left] != nil {
			nodes[i].Left = nodes[left]
		}
		if right < len(nodes) && nodes[right] != nil {
			nodes[i].Right = nodes[right]
		}

		j++
	}
	return nodes[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	arr := strings.Split(input, " ")

	var arrNum []interface{}

	for _, v := range arr {

		if v == "null" {
			arrNum = append(arrNum, nil)
			continue
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			break
		}
		arrNum = append(arrNum, num)
	}

	root := arrToBBT(arrNum)

	result := minDepth(root)
	fmt.Println(result)
}

///////////////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func arrToBalancedBinaryTree(arr []int) *TreeNode {
	left := 0
	right := len(arr) - 1

	return helper(arr, left, right)

}

func helper(arr []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := (right + left) / 2

	root := &TreeNode{Val: arr[mid]}

	root.Left = helper(arr, left, mid-1)
	root.Right = helper(arr, mid+1, right)

	return root
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	r, _ := checkBalance(root)

	return r
}

func checkBalance(root *TreeNode) (bool, int) {

	if root == nil {
		return true, 0
	}

	leftBalance, leftHeight := checkBalance(root.Left)

	if !leftBalance {
		return false, 0
	}

	rightBalance, rightHeight := checkBalance(root.Right)

	if !rightBalance {
		return false, 0
	}

	if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
		return false, 0
	}

	return true, int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")
	var numArr []int

	for i := 0; i < len(arr); i++ {

		num, err := strconv.Atoi(arr[i])

		if err != nil {
			fmt.Println("Error converting")
			break
		}

		numArr = append(numArr, num)
	}

	r := arrToBalancedBinaryTree(numArr)

	result := isBalanced(r)

	fmt.Println(result)

}

///////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printBST(root *TreeNode) {

	if root == nil {
		return
	}

	printBST(root.Left)

	fmt.Print(root.Val, " ")

	printBST(root.Right)
}

func sortedArrayToBST(arr []int) *TreeNode {
	left := 0
	right := len(arr) - 1

	return helper(arr, left, right)

}

func helper(arr []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	var mid int = (left + right) / 2

	root := &TreeNode{Val: arr[mid]}

	root.Left = helper(arr, left, mid-1)
	root.Right = helper(arr, mid+1, right)

	return root
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")
	var numArr []int

	for i := 0; i < len(arr); i++ {

		num, err := strconv.Atoi(arr[i])

		if err != nil {
			fmt.Println("Error converting")
			break
		}

		numArr = append(numArr, num)
	}

	result := sortedArrayToBST(numArr)

	fmt.Println(result.Val)

	printBST(result)

}

////////////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)

		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}

func arrToTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			nodes[i] = &TreeNode{Val: arr[i].(int)}
		}
	}

	var j = 0
	for i := 0; i < len(arr); i++ {
		if nodes[i] == nil {
			continue
		}

		leftNode := 2*j + 1
		rightNode := 2*j + 2

		if leftNode < len(arr) && arr[leftNode] != nil {
			nodes[i].Left = nodes[leftNode]
		}

		if rightNode < len(arr) && arr[rightNode] != nil {
			nodes[i].Right = nodes[rightNode]
		}

		j++
	}

	return nodes[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	arr := strings.Split(input, " ")

	var arrNum []interface{}

	for _, numStr := range arr {

		if numStr == "null" {
			arrNum = append(arrNum, nil)
			continue
		}

		num, err := strconv.Atoi(numStr)

		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		arrNum = append(arrNum, num)
	}

	root := arrToTree(arrNum)
	result := maxDepth(root)
	fmt.Println(result)
}

////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func arrToTree(arr []interface{}) *TreeNode {

	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))

	for i := 0; i < len(arr); i++ {
		nodes[i] = &TreeNode{Val: arr[i].(int)}
	}

	var j = 0

	for i := 0; i < len(arr); i++ {
		if nodes[i] == nil {
			continue
		}

		var leftNode = 2*j + 1
		var rightNode = 2*j + 2

		if leftNode < len(arr) && arr[leftNode] != nil {
			nodes[i].Left = nodes[leftNode]
		}

		if rightNode < len(arr) && arr[rightNode] != nil {
			nodes[i].Right = nodes[rightNode]
		}

		j++
	}

	return nodes[0]

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	arrInp1, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	arrInp1 = strings.TrimSpace(arrInp1)

	arr1 := strings.Split(arrInp1, " ")

	var arrNum1 []interface{}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] == "null" {
			arrNum1 = append(arrNum1, nil)
		} else {
			num, err := strconv.Atoi(arr1[i])

			if err != nil {
				fmt.Printf("Invalid number: %s\n", arr1[i])
				break
			}
			arrNum1 = append(arrNum1, num)
		}
	}

	tree1 := arrToTree(arrNum1)

	result := isSymmetric(tree1)

	fmt.Println(result)

}

//////////////////////////////////////////////////////////////

// Same tree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	leftSame := isSameTree(p.Left, q.Left)
	rightSame := isSameTree(p.Right, q.Right)

	return leftSame && rightSame
}

func arrToTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			nodes[i] = &TreeNode{Val: arr[i].(int)}
		}
	}

	var j = 0
	for i := 0; i < len(arr); i++ {

		if nodes[i] == nil {
			continue
		}

		curNode := nodes[i]

		nodeLeft := 2*j + 1
		nodeRight := 2*j + 2

		if nodeLeft < len(arr) && arr[nodeLeft] != nil {
			curNode.Left = nodes[nodeLeft]
		}

		if nodeRight < len(arr) && arr[nodeRight] != nil {
			curNode.Right = nodes[nodeRight]
		}

		j++
	}
	return nodes[0]
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	arrInp1, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	arrInp2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	arrInp1 = strings.TrimSpace(arrInp1)
	arrInp2 = strings.TrimSpace(arrInp2)

	arr1 := strings.Split(arrInp1, " ")
	arr2 := strings.Split(arrInp2, " ")

	var arrNum1 []interface{}
	var arrNum2 []interface{}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] == "null" {
			arrNum1 = append(arrNum1, nil)
		} else {
			num, err := strconv.Atoi(arr1[i])

			if err != nil {
				fmt.Printf("Invalid number: %s\n", arr1[i])
				break
			}
			arrNum1 = append(arrNum1, num)
		}
	}

	for i := 0; i < len(arr2); i++ {
		if arr2[i] == "null" {
			arrNum2 = append(arrNum2, nil)
		} else {
			num, err := strconv.Atoi(arr2[i])

			if err != nil {
				fmt.Printf("Invalid number: %s\n", arr2[i])
				break
			}
			arrNum2 = append(arrNum2, num)
		}
	}

	tree1 := arrToTree(arrNum1)
	tree2 := arrToTree(arrNum2)

	result := isSameTree(tree1, tree2)

	fmt.Println(result)

}

//////////////////////////////////////////////////////////

Inorder traversal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	var result []int
	inorderHelper(root, &result)
	return result

}

func inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorderHelper(node.Left, result)
	*result = append(*result, node.Val)
	inorderHelper(node.Right, result)
}

func arrayToTree(arr []interface{}) *TreeNode {

	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			nodes[i] = &TreeNode{Val: arr[i].(int)}
		}
	}

	fmt.Println(nodes)

	j := 0
	for i := 0; i < len(arr); i++ {
		if nodes[i] == nil {
			continue
		}
		// node := nodes[i]

		leftIndex := 2*j + 1
		rightIndex := 2*j + 2

		if leftIndex < len(arr) && nodes[leftIndex] != nil {
			nodes[i].Left = nodes[leftIndex]
		}

		if rightIndex < len(arr) && nodes[rightIndex] != nil {
			nodes[i].Right = nodes[rightIndex]
		}

		j++
	}

	return nodes[0] // Return the root node
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")

	var arrInterface []interface{}

	for _, s := range arr {
		if s == "null" {
			arrInterface = append(arrInterface, nil)
		} else {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Invalid input:", s)
				return
			}
			arrInterface = append(arrInterface, num)
		}
	}

	root := arrayToTree(arrInterface)

	result := inorderTraversal(root)

	fmt.Println(result)
}

///////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//88. Merge Sorted Array

func merge(nums1 []int, m int, nums2 []int, n int) {

	var i, j, z int = 0, 0, 0

	var result []int = make([]int, m+n)
	var newNums1 []int = make([]int, m+20)

	for k := 0; k < m; k++ {
		newNums1[k] = nums1[k]
	}

	for i < m && j < n {

		if newNums1[i] <= nums2[j] {

			result[z] = newNums1[i]
			z++
			i++
		}
		if newNums1[i] >= nums2[j] {
			result[z] = nums2[j]
			z++
			j++
		}
	}

	for k := i; k < m; k++ {
		result[z] = nums1[k]
		z++
	}

	for k := j; k < n; k++ {
		result[z] = nums2[k]
		z++
	}

	fmt.Println(result)

	for k := 0; k < m+n; k++ {
		nums1[k] = result[k]
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp1, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp3, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	inp4, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp1 = strings.TrimSpace(inp1)
	inp2 = strings.TrimSpace(inp2)
	inp3 = strings.TrimSpace(inp3)
	inp4 = strings.TrimSpace(inp4)

	arr1 := strings.Split(inp1, " ")
	arr2 := strings.Split(inp2, " ")

	var arrNum1 []int
	var arrNum2 []int

	var m, _ = strconv.Atoi(inp3)
	var n, _ = strconv.Atoi(inp4)

	for i := 0; i < len(arr1); i++ {

		num, err := strconv.Atoi(arr1[i])

		if err != nil {
			fmt.Printf("Invalid number: %s\n", arr1[i])
			continue
		}

		arrNum1 = append(arrNum1, num)
	}

	for i := 0; i < len(arr2); i++ {

		num, err := strconv.Atoi(arr2[i])

		if err != nil {
			fmt.Printf("Invalid number: %s\n", arr1[i])
			continue
		}

		arrNum2 = append(arrNum2, num)
	}

	merge(arrNum1, m, arrNum2, n)
	fmt.Println(arrNum1)

}

///////////////////////////////////////////
83. Remove Duplicates from Sorted List

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func removeElementAtPos(head **ListNode, pos int) {

	current := *head
	if pos == 0 {
		*head = current.Next
	} else {
		for i := 0; i < pos-1; i++ {
			if current == nil {
				fmt.Println("Out of bounds")
				break
			}
			current = current.Next
		}

		current.Next = current.Next.Next
	}
}

func insertToLinkedList(head **ListNode, pos int, val int) {

	current := *head
	insertedVal := &ListNode{Val: val}

	if pos == 0 {
		*head = insertedVal
		insertedVal.Next = current
	} else {
		for i := 0; i < pos-1; i++ {
			if current == nil {
				fmt.Println("Out of bounds")
				break
			}
			current = current.Next
		}

		insertedVal.Next = current.Next
		current.Next = insertedVal
	}
}

func arrayToLinkedList(arr []int) *ListNode {

	// Init the linked list
	head := &ListNode{Val: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head

}

func printLinkdList(head *ListNode) {

	current := head
	for current != nil {

		if current.Next != nil {
			print(current.Val)
			print(" -> ")
		} else {
			print(current.Val)
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")

	var arrNum []int

	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])

		if err != nil {
			fmt.Printf("Invalid input: %s\n", arr[i])
			continue
		}

		arrNum = append(arrNum, num)
	}

	head := arrayToLinkedList(arrNum)

	// printLinkdList(head)
	// insertToLinkedList(&head, 2, 10)
	// printLinkdList(head)
	// removeElementAtPos(&head, 2)
	// printLinkdList(head)

	r := deleteDuplicates(head)
	printLinkdList(r)

}

/////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dp []int = make([]int, 10000001)

// Clim stair

func climbStairs(n int) int {

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if dp[n] > 0 {
		return dp[n]
	}

	dp[n] = climbStairs(n-1) + climbStairs(n-2)

	return dp[n]

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inputStr, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inputStr = strings.TrimSpace(inputStr)

	num, err := strconv.Atoi(inputStr)

	if err != nil {
		println("Error converting input to integer:", err)
		return
	}

	r := climbStairs(num)

	fmt.Println(r)
}

/////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func mySqrt(x int) int {
	r := int(math.Sqrt(float64(x)))

	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inputStr = strings.TrimSpace(inputStr)

	num, err := strconv.Atoi(inputStr)

	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	r := mySqrt(num)
	fmt.Println(r)

}

/////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addBinary(a string, b string) string {

	var i, j int

	i = len(a) - 1
	j = len(b) - 1

	var result string = ""
	var carry int = 0

	for i >= 0 && j >= 0 {
		if a[i] == '1' && b[j] == '1' && carry == 1 {
			carry = 1
			result = "1" + result
		}

		if a[i] == '1' && b[j] == '1' && carry == 0 {
			carry = 1
			result = "0" + result
		}

		if (a[i] == '0' && b[j] == '1' && carry == 1) || (a[i] == '1' && b[j] == '0' && carry == 1) {
			carry = 1
			result = "0" + result
		}

		if (a[i] == '0' && b[j] == '1' && carry == 0) || (a[i] == '1' && b[j] == '0' && carry == 0) {
			carry = 0
			result = "1" + result
		}

		if a[i] == '0' && b[j] == '0' && carry == 0 {
			carry = 0
			result = "0" + result
		}

		if a[i] == '0' && b[j] == '0' && carry == 1 {
			carry = 0
			result = "1" + result
		}

		// fmt.Println("--- ", result)
		i--
		j--
	}

	if i >= 0 {
		for k := i; k >= 0; k-- {
			if a[k] == '1' && carry == 1 {
				carry = 1
				result = "0" + result
			} else if a[k] == '1' && carry == 0 {
				carry = 0
				result = "1" + result
			} else if a[k] == '0' && carry == 0 {
				carry = 0
				result = "0" + result
			} else if a[k] == '0' && carry == 1 {
				carry = 0
				result = "1" + result
			}

			// fmt.Println("???: ", result)
		}
		if carry == 1 {
			result = "1" + result
			carry = 0
		}
	}

	if j >= 0 {

		for k := j; k >= 0; k-- {
			if b[k] == '1' && carry == 1 {
				carry = 1
				result = "0" + result
			} else if b[k] == '1' && carry == 0 {
				carry = 0
				result = "1" + result
			} else if b[k] == '0' && carry == 0 {
				carry = 0
				result = "0" + result
			} else if b[k] == '0' && carry == 1 {
				carry = 0
				result = "1" + result
			}
		}
		if carry == 1 {
			result = "1" + result
			carry = 0
		}
	}

	if carry == 1 {
		result = "1" + result
	}

	return result

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	inp1, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp1 = strings.TrimSpace(inp1)
	inp2 = strings.TrimSpace(inp2)

	result := addBinary(inp1, inp2)
	fmt.Println(result)

}

///////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Plus one
func plusOne(digits []int) []int {

	var carry = 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry

		if sum == 10 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] = sum
			carry = 0
			break
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")
	fmt.Println(arr)

	var arrNum []int
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}
		arrNum = append(arrNum, num)

	}

	arrNum = plusOne(arrNum)
	fmt.Println(arrNum)
}

//////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Length of Last Word
func lengthOfLastWord(s string) int {

	s = strings.ReplaceAll(s, `"`, " ")

	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	lastWord := arr[len(arr)-1]

	return len(lastWord)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Wrong input")
		return
	}

	inp = strings.TrimSpace(inp)

	result := lengthOfLastWord(inp)

	fmt.Println(result)
}

////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Search Insert Position

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, left int, right int) int {

	if left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			return binarySearch(nums, target, mid+1, right)
		}
		return binarySearch(nums, target, left, mid-1)
	}
	return left
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	targetStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)
	targetStr = strings.TrimSpace(targetStr)

	target, err := strconv.Atoi(targetStr)

	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	var arr []int

	for _, v := range strings.Split(inp, " ") {
		num, _ := strconv.Atoi(v)
		arr = append(arr, num)
	}

	result := searchInsert(arr, target)

	fmt.Println(result)

}

////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Find the Index of the First Occurrence in a String

func strStr(haystack string, needle string) int {

	// Handle the case where needle is longer than haystack
	if len(needle) > len(haystack) {
		return -1
	}

	k := len(needle)
	window := haystack[0:k]

	if window == needle {
		return 0
	}

	for i := k; i < len(haystack); i++ {

		window = string(haystack[i-k+1:i-k+len(needle)]) + string(haystack[i])

		if window == needle {
			return i - k + 1
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	haystack, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	needle, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	haystack = strings.TrimSpace(haystack)
	needle = strings.TrimSpace(needle)

	result := strStr(haystack, needle)

	fmt.Println(result)
}

////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Remove element

func removeElement(nums []int, val int) int {
	var k int = 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}

	return k
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp2, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp2 = strings.TrimSpace(inp2)
	inp = strings.TrimSpace(inp)

	arr := strings.Split(inp, " ")
	target, err := strconv.Atoi(inp2)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	var arrNum []int
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		arrNum = append(arrNum, num)
	}

	result := removeElement(arrNum, target)
	fmt.Println(result)
	fmt.Println(arrNum[:result])
}

/////////////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Remove duplicated from sorted array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var k int = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			k++
			nums[k-1] = nums[i]
		}
	}

	return k
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)
	arr := strings.Split(inp, " ")

	var arrNum []int
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err == nil {
			arrNum = append(arrNum, num)
		}
	}

	results := removeDuplicates(arrNum)

	fmt.Printf("Length of the unique array: %d\n", results)

}

/////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Merge two sorted lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to help easily return the merged list
	dummy := &ListNode{}
	current := dummy

	// Traverse both lists and link nodes in sorted order
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach the remaining part of list1 or list2, if any
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the next node after dummy, which is the head of the merged list
	return dummy.Next
}

func arrToNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	// Init the node
	head := &ListNode{Val: arr[0]}

	// the first element in the node list
	current := head

	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	current := head

	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	arrInp1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	arrInp2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	arrInp1 = strings.TrimSpace(arrInp1)
	arrInp2 = strings.TrimSpace(arrInp2)

	arr1 := strings.Split(arrInp1, ",")
	arr2 := strings.Split(arrInp2, ",")

	var arrNum1 []int
	var arrNum2 []int

	for _, v := range arr1 {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", v)
			continue
		}
		arrNum1 = append(arrNum1, num)
	}

	for _, v := range arr2 {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", v)
			continue
		}
		arrNum2 = append(arrNum2, num)
	}

	nodeLis1 := arrToNode(arrNum1)
	nodeLis2 := arrToNode(arrNum2)

	results := mergeTwoLists(nodeLis1, nodeLis2)

	printList(results)

	// for nodeLis1 != nil {
	// 	fmt.Print(nodeLis1.Val, " ")
	// 	nodeLis1 = nodeLis1.Next
	// }

	// fmt.Println()

	// for nodeLis2 != nil {
	// 	fmt.Print(nodeLis2.Val, " ")
	// 	nodeLis2 = nodeLis2.Next
	// }
}

////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Valid Parenthesis

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val string) {

	// Add the new value to top of the stack
	*s = append(*s, val)
}

func (s *Stack) Pop() (string, bool) {

	// Remove the top value from the stack
	if s.isEmpty() {
		return "", false
	} else {
		// The top of the stack
		index := len(*s) - 1

		// Take the top value from the stack
		element := (*s)[index]

		// Remove the top value from the stack
		*s = (*s)[:index]

		return element, true
	}
}

func isValid(s string) bool {
	stack := Stack{}

	for _, char := range s {
		strChar := string(char)
		if strChar == "(" || strChar == "{" || strChar == "[" {
			stack.Push(strChar)
		} else if strChar == ")" || strChar == "}" || strChar == "]" {
			if stack.isEmpty() {
				return false
			}
			top, _ := stack.Pop()
			if (strChar == ")" && top != "(") || (strChar == "}" && top != "{") || (strChar == "]" && top != "[") {
				return false
			}
		} else {
			return false
		}
	}
	return stack.isEmpty()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	inp = strings.TrimSpace(inp)

	value := isValid(inp)

	fmt.Println(value)
}

//////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Longest Common Prefix

// Using divide and conquer approach

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return solve(strs, 0, len(strs)-1)
}

func solve(arr []string, left int, right int) string {
	if left == right {
		return arr[left]
	}

	mid := left + (right-left)/2

	leftPrefix := solve(arr, left, mid)
	rightPrefix := solve(arr, mid+1, right)

	return commonPrefix(leftPrefix, rightPrefix)

}

func commonPrefix(str1 string, str2 string) string {

	// Find the smaller length prefix
	length := min(len(str1), len(str2))

	for i := 0; i < length; i++ {

		if str1[i] != str2[i] {
			return str1[0:i]
		}
	}
	return str1[0:length]
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	strs := strings.TrimSpace(inp)
	strsArr := strings.Split(strs, ",")

	for i := range strsArr {
		strsArr[i] = strings.TrimSpace(strsArr[i])
	}

	result := longestCommonPrefix(strsArr)

	fmt.Println(result)
}

/////////////////////////////////////////////////////////////

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Roman to integer

func integerToRoman(inp string) string {

	romanMap := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	number, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return ""
	}

	var result = ""
	for _, symbol := range romanMap {
		for number >= symbol.Value {
			number -= symbol.Value
			result += symbol.Symbol
		}
	}

	return result

}

// roman to integer conversion
func solve(inp string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int = 0
	for i := 0; i <= len(inp)-1; i++ {

		if i <= len(inp)-2 && romanMap[string(inp[i])] < romanMap[string(inp[i+1])] {
			result -= romanMap[string(inp[i])]
		} else {
			result += romanMap[string(inp[i])]
		}

	}
	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)

	// result := solve(inp)
	result := integerToRoman(inp)
	fmt.Println(result)
}

/////////////////////////////////////////////////////////////

// Palindrome number

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPalindrome(val int) bool {

	if val < 0 {
		return false
	}

	inp := strconv.Itoa(val)

	var reverseInp string

	for i := len(inp) - 1; i >= 0; i-- {
		reverseInp += string(inp[i])
	}

	if inp == reverseInp {
		return true
	} else {
		return false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inp = strings.TrimSpace(inp)
	val, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	result := isPalindrome(val)

	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

/////////////////////////////////////////////////////////
// Two Sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(numArr []int, target int) []int {

	resultMap := make(map[int]int)
	for i := 0; i < len(numArr); i++ {
		complement := target - numArr[i]
		val, ok := resultMap[complement]

		if !ok {
			resultMap[numArr[i]] = i
		} else {
			return []int{val, i}
		}
	}
	return nil

}

func solve(input string, target string) {

	// Format the input
	input = strings.TrimSpace(input)
	input = strings.Trim(input, "[]")
	target = strings.TrimSpace(target)
	targetNum, err := strconv.Atoi(target)

	if err != nil {
		fmt.Println("Invalid target:", target)
		return
	}

	var arrayStr = strings.Split(input, ",")
	var arrNum = make([]int, 0, len(arrayStr))

	for _, v := range arrayStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid input:", v)
			return
		}
		arrNum = append(arrNum, num)

	}

	// Solve problems
	results := twoSum(arrNum, targetNum)
	fmt.Println(results)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	target, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	solve(input, target)

}
