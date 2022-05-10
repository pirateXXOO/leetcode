package main

////////////  dedumplicateArray
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/

// func removeDuplicates(nums []int) (int, []int) {
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i] == nums[i-1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 		}
// 	}
// 	return len(nums), nums[:]
// }

// func main() {
// 	arr := [11]int{1, 1, 2, 2, 4, 4, 5, 6, 6, 6, 6}
// 	a, barr := removeDuplicates(arr[:])
// 	fmt.Println(a)
// 	for _, k := range barr {
// 		fmt.Print(k, ",")
// 	}
// }

//////////////////// get max profit
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/

// func maxProfit(prices []int) int {
// 	var max int

// 	for i := 1; i < len(prices); i++ {
// 		if prices[i]-prices[i-1] > 0 {
// 			max += prices[i] - prices[i-1]
// 		}
// 	}

// 	return max
// }

// func main() {
// 	arr := [6]int{7, 1, 5, 3, 6, 4}
// 	sum := maxProfit(arr[:])
// 	fmt.Println(sum)
// }

/////////////////// rotate array
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/

// func rotate(nums []int, k int) {
// 	leng := len(nums) - k
// 	for j := 1; j <= k; j++ {
// 		for i := 0; i < leng; i++ {
// 			fmt.Println(i)
// 			nums[i], nums[leng] = nums[leng], nums[i]
// 		}
// 	}
// }

// func rotate(nums []int, k int) {
// 	length := len(nums)
// 	k %= length
// 	reverse(nums, 0, length-1)
// 	reverse(nums, 0, k-1)
// 	reverse(nums, k, length-1)

// }

// func reverse(nums []int, start int, end int) {
// 	for start < end {
// 		nums[start], nums[end] = nums[end], nums[start]
// 		start++
// 		end--
// 	}
// }

// func rotate(nums []int, k int) {
// 	hold := nums[0]
// 	var index int
// 	length := len(nums)
// 	var visited []int

// 	for i := 0; i < length; i++ {
// 		index = (index + k) % length
// 		f := func(visited []int, index int) bool {
// 			for i := 0; i < len(visited); i++ {
// 				if visited[i] == index {
// 					return true
// 				}
// 			}
// 			return false
// 		}
// 		if f(visited, index) {
// 			index = (index + 1) % length
// 			hold = nums[index]
// 			i--
// 		} else {
// 			visited = append(visited, index)
// 			nums[index], hold = hold, nums[index]
// 		}
// 	}

// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	k := 4

// 	rotate(nums, k)

// 	fmt.Println(nums)
// }

////////////////////// if array have the same member
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/

// func containsDuplicate(nums []int) bool {
// 	leng := len(nums)
// 	for i := 1; i < leng; i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func main() {
// 	arr := []int{7, 1, 5, 3, 1, 4}
// 	result := containsDuplicate(arr)
// 	fmt.Println(result)

// }

////////////////// get num only exist onece
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/

// func singleNumber(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		var count int
// 		for j := 0; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			}
// 			if nums[i] == nums[j] {
// 				count++
// 			}
// 		}
// 		if count == 0 {
// 			return nums[i]
// 		}
// 	}
// 	return 0
// }

// func main() {
// 	arr := []int{2, 2, 6, 2, 2, 1, 4, 3, 5, 2}
// 	result := singleNumber(arr)
// 	fmt.Println(result)
// }

/////////////////// get intersection
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/

// func intersect(nums1 []int, nums2 []int) []int {
// m1 := make(map[int]int)
// m2 := make(map[int]int)
// var nums []int
// for _, i := range nums1 {
// m1[i]++
// }
// for _, j := range nums2 {
// m2[j]++
// }
// for x := range m1 {
// if m2[x] != 0 {
// if m1[x] < m2[x] {
// for k := 0; k < m1[x]; k++ {
// nums = append(nums, x)
// }
// } else {
// for k := 0; k < m2[x]; k++ {
// nums = append(nums, x)
// }
//
// }
// }
// }
// return nums
// }

// func intersect(nums1 []int, nums2 []int) []int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)

// 	var nums []int

// 	i := 0
// 	j := 0
// app:
// 	if nums1[i] == nums2[j] {
// 		nums = append(nums, nums1[i])
// 		i++
// 		j++
// 		if i == len(nums1) {
// 			return nums

// 		}
// 		if j == len(nums2) {
// 			return nums

// 		}
// 		goto app
// 	} else if nums1[i] < nums2[j] {
// 		i++
// 		if i == len(nums1) {
// 			return nums

// 		}
// 		goto app
// 	} else {
// 		j++
// 		if j == len(nums2) {
// 			return nums

// 		}

// 		goto app
// 	}
// }

// func main() {
// 	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
// }

////////// add one
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
// func plusOne(digits []int) []int {

// 	leng := len(digits)

// 	for i := leng - 1; i >= 0; i-- {
// 		if digits[i] != 9 {
// 			digits[i]++
// 			return digits
// 		} else {
// 			digits[i] = 0
// 		}
// 	}

// 	digits = append([]int{1}, digits...)

// 	return digits
// }

// func main() {

// 	result := plusOne([]int{1, 9, 8, 9})
// 	fmt.Println(result)

// }

////////// Move 0
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2ba4i/

// func moveZeroes(nums []int) {

// 	for i := len(nums) - 1; i > 0; i-- {
// 		for j := 0; j < i; j++ {
// 			if nums[j] == 0 && nums[j+1] != 0 {
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 		}
// 	}
// }

// func moveZeroes(nums []int) {
// 	i := 0
// 	for j := 0; j < len(nums); j++ {
// 		if nums[j] == 0 {
// 			i++
// 		} else if i != 0 {
// 			nums[j-i] = nums[j]
// 			nums[j] = 0
// 		}
// 	}
// }

// func main() {

// 	result := []int{1, 2, 0, 2, 1, 0, 3, 4}

// 	moveZeroes(result)
// 	// nums := []int{1, 9, 8, 9, 0, 0}
// 	// sort.Slice(nums[0:1], func(k, j int) bool {
// 	// 	return nums[k] < nums[j]
// 	// })
// 	fmt.Println(result)

// }

////////// Get sum
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/

// func twoSum(nums []int, target int) []int {

// 	var f = func(s []int, e int) bool {
// 		for _, a := range s {
// 			if a == e {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	var checked []int
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if f(checked, i) || f(checked, j) {
// 				continue
// 			} else if nums[i]+nums[j] == target {
// 				checked = append(checked, i, j)
// 			}
// 		}
// 	}
// 	return checked
// }

// func main() {

// 	result := []int{1, 2, 0, 2, 1, 0, 3, 4}

// 	get := twoSum(result, 3)
// 	// nums := []int{1, 9, 8, 9, 0, 0}
// 	// sort.Slice(nums[0:1], func(k, j int) bool {
// 	// 	return nums[k] < nums[j]
// 	// })
// 	fmt.Println(get)

// }
////////// check valid Sudoku
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/

// func isValidSudoku(board [][]byte) bool {
// 	var prt = func(s []byte){
//         fmt.Println(string(s))
// 	}
// 	var check = func(s []byte) bool {
// 		for i := 0; i < len(s); i++ {
// 			for j := i + 1; j < len(s); j++ {
// 				if s[i] == s[j] {
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	}
// 	var toCheck []byte

// 	for i := 0; i < 9; i++ {
// 		// row check
// 		for j := 0; j < 9; j++ {

// 			if string(board[i][j]) != "." {
// 				toCheck = append(toCheck, board[i][j])
// 			}

// 		}
// 		prt(toCheck)
// 		if check(toCheck) == false {
// 			return false
// 		}

// 		// column check
// 		toCheck = []byte{}
// 		for j := 0; j < 9; j++ {

// 			if string(board[j][i]) != "." {
// 				toCheck = append(toCheck, board[j][i])
// 			}

// 		}
// 		prt(toCheck)
// 		if check(toCheck) == false {
// 			return false
// 		}
// 		toCheck = []byte{}
// 	}

// 	for i := 0; i < 9; i += 3 {
// 		// box check
// 		for j := 0; j < 9; j += 3 {
// 			for k := 0; k < 3; k++ {
// 				for l := 0; l < 3; l++ {
// 					if string(board[i+k][j+l]) != "." {
// 						toCheck = append(toCheck, board[i+k][j+l])
// 					}
// 				}
// 			}
// 			prt(toCheck)
// 			if check(toCheck) == false {
// 				return false
// 			}
//             toCheck = []byte{}
// 		}

// 		for j := 0; j < 9; j += 3 {
// 			for k := 0; k < 3; k++ {
// 				for l := 0; l < 3; l++ {
// 					if string(board[j+k][i+l]) != "." {
// 						toCheck = append(toCheck, board[j+k][i+l])
// 					}
// 				}
// 			}
// 			prt(toCheck)
// 			if check(toCheck) == false {
// 				return false
// 			}
//             toCheck = []byte{}
// 		}
// 	}
// 	return true
// }

////////// rotate picture
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhhkv/

// func rotate(matrix [][]int) {
// 	leng := len(matrix)
// 	for i := 0; i < leng; i++ {
// 		for j := 0; j < (leng+1)/2; j++ {
// 			matrix[i][j], matrix[i][leng-1-j] = matrix[i][leng-1-j], matrix[i][j]
// 		}
// 	}
// 	for i := leng - 2; i >= 0; i-- {
// 		for j := 0; j <= leng-i-2; j++ {
// 			matrix[i][j], matrix[leng-j-1][leng-i-1] = matrix[leng-j-1][leng-i-1], matrix[i][j]
// 		}
// 	}
// }

// func main() {
// 	// var matrix [][]int
// 	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

// 	rotate(matrix)
// 	fmt.Print(matrix)
// }

////////// reverse string
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/

// func reverseString(s []byte)  {
// 	for i:=0;i< (len(s)+1)/2; i++{
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// }

////////// reverse int
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/

// func reverse(x int) int {
// 	var result int
// 	var temp int = x
// 	for i := 1; ; i++ {
// 		remainder := temp % 10
// 		result = result*10 + remainder
// 		temp = temp / 10
// 		if temp == 0 {
// 			break
// 		}
// 	}

// 	if result > 2147483647 || result < -2147483648 {
// 		return 0
// 	}

// 	return result
// }

// func main() {
// 	var intNum int = -1230012
// 	result := reverse(intNum)
// 	fmt.Println(result)
// }

////////// first uniq char
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/

// func firstUniqChar(s string) int {
// 	if len(s) == 1 {
// 		return 0
// 	}
// 	var alreadyExist []byte
// 	var f = func(alalreadyExist []byte, b byte) bool {
// 		for k := 0; k < len(alreadyExist); k++ {
// 			if b == alreadyExist[k] {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if f(alreadyExist, s[i]) {
// 			continue
// 		}
// 		if i == len(s)-1 {
// 			return i
// 		}
// 		for j := i + 1; j < len(s); j++ {
// 			if s[i] == s[j] {
// 				alreadyExist = append(alreadyExist, s[i])
// 				break
// 			} else if j == len(s)-1 {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }

// func main() {
// 	var str string = "loveleetcode"
// 	// var str2 []string = string(str)
// 	fmt.Println(firstUniqChar(str))

// }

////////// isAnagram
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t){
// 		return false
// 	} else {
// 		m1 := make(map[byte]int)
// 		m2 := make(map[byte]int)
// 		leng := len(s)
// 		for i:=0;i<leng;i++{
// 			m1[s[i]]++
// 			m2[t[i]]++
// 		}
// 		if len(m1) == len(m2) {
// 			for k,_ := range m1 {
// 				if m1[k] != m2[k]{
// 					return false
// 				}
// 			}
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// }

////////// isPalindrome
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/

// func isPalindrome(s string) bool {
// 	var arr []rune
// 	for _, ch := range s {
// 		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
// 			arr = append(arr, ch)
// 		} else if ch >= 'A' && ch <= 'Z' {
// 			arr = append(arr, ch+('a'-'A'))
// 		} else {
// 			continue
// 		}
// 	}

// 	for i := 0; i < (len(arr)+1)/2; i++ {
// 		if arr[i] != arr[len(arr)-1-i] {
// 			return false
// 		}
// 	}
// 	return true

// }

// func main() {
// 	var s string = "`l;`` 1o1 ??;l`"
// 	fmt.Println(isPalindrome(s))
// }

////////// myAtoi
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/

// func myAtoi(s string) int {
// 	// remove space
// 	str := strings.TrimSpace(s)
// 	l := len(str)
// 	if l == 0 {
// 		return 0
// 	}
// 	if !isCode(str[0]) && !isDigit(str[0]) {
// 		return 0
// 	}

// 	f := 1
// 	if isCode(str[0]) {
// 		if str[0] == '-' {
// 			f = -1
// 		}
// 		str = str[1:]
// 		l = len(str)
// 	}
// 	res := 0
// 	for i := 0; i < l; i++ {
// 		if !isDigit(str[i]) {
// 			break
// 		}
// 		res += int(str[i] - '0')
// 		if f*res > math.MaxInt32 {
// 			return math.MaxInt32
// 		}
// 		if f*res < math.MinInt32 {
// 			return math.MinInt32
// 		}
// 		res *= 10

// 	}
// 	res = f * res / 10
// 	return res
// }

// func isDigit(c uint8) bool {
// 	return '0' <= c && c <= '9'
// }
// func isCode(c uint8) bool {
// 	return c == '+' || c == '-'
// }

////////// strStr
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/

// func strStr(haystack string, needle string) int {
// 	if len(needle) == 0 && len(haystack) == 0 {
// 		return 0
// 	}
// 	if len(needle) == 0 {
// 		return 0
// 	}

// 	var start int
// 	for i := 0; i < len(haystack)-len(needle)+1; i++ {
// 		if haystack[i] == needle[0] {
// 			start = i
// 			for j := 0; j < len(needle); j++ {
// 				if haystack[i+j] != needle[j] {
// 					break
// 				} else if j == len(needle)-1 {
// 					return start
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

////////// countAndSay
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/

// func countAndSay(n int) string {
// 	var res int = 1

// 	for i := 1; i <= n; i++ {
// 		if i == 1 {
// 			res = 1
// 		} else {
// 			res = getRes(res)

// 		}
// 	}
// 	return strconv.Itoa(res)
// }
// func getRes(n int) int {

// 	var num []int
// 	var count []int
// 	// num := make([]int, 10)
// 	// count := make([]int, 10)
// 	num = append(num, n%10)
// 	count = append(count, 1)
// 	n = n / 10

// 	for i := 0; ; i++ {
// 		if n == 0 {
// 			break
// 		}
// 		for j := 0; ; j++ {
// 			if n%10 == num[i] {
// 				count[i]++
// 				n = n / 10
// 			} else {
// 				num = append(num, n%10)
// 				// num[i+1] = n % 10
// 				count = append(count, 1)
// 				// count[i+1] = 1
// 				n = n / 10
// 				break
// 			}
// 			if n == 0 {
// 				break
// 			}
// 		}
// 	}
// 	// fmt.Println(num)
// 	// fmt.Println(count)

// 	var res int
// 	for k := len(num) - 1; k >= 0; k-- {
// 		res = res*100 + count[k]*10 + num[k]
// 	}

// 	return res
// }
///////////////////////////////
// func countAndSay(n int) string {

// 	var arr string = "1"
// 	if n == 1 {
// 		return arr
// 	}

// 	for i := 1; i < n; i++ {
// 		word := arr[0] // current number
// 		num := 1       // count
// 		temp := ""     // string

// 		for j := 1; j < len(arr); j++ {
// 			if word == arr[j] {
// 				num++
// 			} else if word != arr[j] {
// 				num_char := num
// 				temp += strconv.Itoa(num_char)
// 				temp += string(word)

// 				word = arr[j]
// 				num = 1
// 			}
// 		}
// 		num_char := num
// 		temp += strconv.Itoa(num_char)
// 		temp += string(word)

// 		arr = temp
// 	}
// 	return arr

// }

// func main() {

// 	var n int = 4
// 	res := countAndSay(n)
// 	fmt.Println(res)

// }

////////// longestCommonPrefix
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/

// func longestCommonPrefix(strs []string) string {

// 	var resSlice []byte
// 	if len(strs) == 1 {
// 		return strs[0]
// 	}
// 	length := len(strs[0])
// 	// var shortest int
// 	for k := 0; k < len(strs); k++ {
// 		if length > len(strs[k]) {
// 			length = len(strs[k])
// 			// shortest = k
// 		}
// 	}
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < len(strs)-1; j++ {
// 			if strs[j][i] != strs[j+1][i] {
// 				return string(resSlice)
// 			} else if j == len(strs)-2 {
// 				resSlice = append(resSlice, strs[j][i])
// 			}
// 		}
// 	}
// 	return string(resSlice)
// }

// func main() {
// 	var strs = []string{"flower", "flow", "flight"}
// 	result := longestCommonPrefix(strs)
// 	fmt.Println(result)

// }

////////// deleteNode
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnarn7/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func deleteNode(node *ListNode) {
//     node.Val = node.Next.Val
// 	node.Next = node.Next.Next
// }

////////// removeNthFromEnd
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	var fast *ListNode = head
// 	var slow *ListNode = head
// 	for i := 0; i < n; i++ {
// 		fast = fast.Next
// 	}
// 	if fast == nil {
// 		return head.Next
// 	}
// 	for {
// 		if fast.Next != nil {
// 			fast = fast.Next
// 			slow = slow.Next
// 		} else {
// 			break
// 		}
// 	}
// 	slow.Next = slow.Next.Next
// 	return head
// }

////////// reverseList
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	var newHead *ListNode

// 	for head != nil {
// 		node := head.Next
// 		head.Next = newHead
// 		newHead = head
// 		head = node
// 	}
// 	return newHead

// }

// func reverseList(head *ListNode) *ListNode {
// 	var newHead *ListNode

// 	for head != nil {
// 		head, head.Next, newHead = head.Next, newHead, head
// 	}
// }
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	ret := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return ret
// }

// func main() {

// 	L5 := ListNode{
// 		Val:  5,
// 		Next: nil,
// 	}
// 	L4 := ListNode{
// 		Val:  4,
// 		Next: &L5,
// 	}
// 	L3 := ListNode{
// 		Val:  3,
// 		Next: &L4,
// 	}
// 	L2 := ListNode{
// 		Val:  2,
// 		Next: &L3,
// 	}
// 	L1 := ListNode{
// 		Val:  1,
// 		Next: &L2,
// 	}

// 	result := reverseList(&L1)
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}
// }

////////// mergeTwoLists
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnbp2/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	var head, node *ListNode

// 	if list1.Val < list2.Val {
// 		head = list1
// 		node = list1
// 		list1 = list1.Next
// 	} else {
// 		head = list2
// 		node = list2
// 		list2 = list2.Next
// 	}

// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			node.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			node.Next = list2
// 			list2 = list2.Next
// 		}
// 		node = node.Next
// 	}

// 	if list1 != nil {
// 		node.Next = list1
// 	}
// 	if list2 != nil {
// 		node.Next = list2
// 	}

// 	return head

// }

// func main() {
// 	//
// 	L5 := ListNode{
// 		Val:  5,
// 		Next: nil,
// 	}
// 	L4 := ListNode{
// 		Val:  4,
// 		Next: &L5,
// 	}
// 	L3 := ListNode{
// 		Val:  3,
// 		Next: &L4,
// 	}
// 	L2 := ListNode{
// 		Val:  2,
// 		Next: &L3,
// 	}
// 	L1 := ListNode{
// 		Val:  1,
// 		Next: &L2,
// 	}

// 	M2 := ListNode{
// 		Val:  2,
// 		Next: nil,
// 	}
// 	M1 := ListNode{
// 		Val:  1,
// 		Next: &M2,
// 	}
// 	//
// 	result := mergeTwoLists(&L1, &M1)
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}
// }

////////// isPalindrome
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnv1oc/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverse(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	for head != nil {
// 		next := head.Next
// 		head.Next = prev
// 		prev = head
// 		head = next
// 	}
// 	return prev
// }

// func isPalindrome(head *ListNode) bool {
// 	fast := head
// 	slow := head

// 	for fast != nil && fast.Next != nil {
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 	}

// 	if fast != nil {
// 		slow = slow.Next
// 	}

// 	slow = reverse(slow)

// 	fast = head

// 	for slow != nil {
// 		if fast.Val != slow.Val {
// 			return false
// 		}
// 		fast = fast.Next
// 		slow = slow.Next
// 	}

// 	return true

// }

////////// hasCycle
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnwzei/

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}

// 	slow := head
// 	fast := head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			return true
// 		}
// 	}

// 	return false

// }

///////// maxDepth
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnd69e/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	maxLeft := maxDepth(root.Left)
// 	maxRight := maxDepth(root.Right)

// 	if maxLeft > maxRight {
// 		return maxLeft + 1
// 	} else {
// 		return maxRight + 1
// 	}
// }

////////// isValidBST
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn08xg/

// func isValidBST(root *TreeNode) bool {
// 	min := -1 << (8*unsafe.Sizeof(1) - 1)
// 	max := 1<<(8*unsafe.Sizeof(1)-1) - 1
// 	return _isValidBST(root, min, max)
// }

// func _isValidBST(root *TreeNode, min int, max int) bool {
// 	if root == nil {
// 		return true
// 	}

// 	v := root.Val

// 	if root.Right != nil {
// 		if r := root.Right.Val; r <= min || r >= max || r <= v {
// 			return false
// 		}
// 	}

// 	if root.Left != nil {
// 		if l := root.Left.Val; l >= max || l <= min || l >= v {
// 			return false
// 		}
// 	}

// 	return _isValidBST(root.Right, v, max) && _isValidBST(root.Left, min, v)
// }

// func isValidBST(root *TreeNode) bool {
// 	val := -1 << (8*unsafe.Sizeof(1) - 1)
// 	cur := root
// 	for cur != nil {
// 		if cur.Left == nil {
// 			if cur.Val > val {
// 				val = cur.Val
// 			} else {
// 				return false
// 			}
// 			cur = cur.Right
// 		} else {
// 			prev := cur.Left
// 			for prev.Right != nil && prev.Right != cur {
// 				prev = prev.Right
// 			}

// 			if prev.Right == nil {
// 				prev.Right = cur
// 				cur = cur.Left
// 			} else {
// 				if cur.Val > val {
// 					val = cur.Val
// 				} else {
// 					return false
// 				}
// 				prev.Right = nil
// 				cur = cur.Right
// 			}
// 		}
// 	}
// 	return true
// }

///////// isSymmetric
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn7ihv/

// func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	}

// 	if left == nil || right == nil || left.Val != right.Val {
// 		return false
// 	}

// 	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
// }

// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return isSymmetricHelper(root.Left, root.Right)
// }

///////// levelOrder
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnldjj/

// func levelOrder(root *TreeNode) [][]int {
// 	var ret [][]int
// 	if root == nil {
// 		return ret
// 	}

// 	q := []*TreeNode{root}

// 	for i := 0; len(q) > 0; i++ {
// 		ret = append(ret, []int{})
// 		p := []*TreeNode{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			ret[i] = append(ret[i], node.Val)
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}

// 		q = p
// 	}

// 	return ret

// }

// func levelOrder(root *TreeNode) [][]int {
// 	res := [][]int{}
// 	if root == nil {
// 		return res
// 	}

// 	var queue = []*TreeNode{root}
// 	fmt.Println("len queue", len(queue))
// 	for _, i := range queue {
// 		fmt.Println(i.Val)
// 	}
// 	fmt.Println()
// 	fmt.Println()

// 	var level int
// 	for len(queue) > 0 {
// 		counter := len(queue)
// 		res = append(res, []int{})
// 		fmt.Println("res", res)
// 		for 0 < counter {
// 			fmt.Println()
// 			fmt.Println("res", res)
// 			for _, i := range queue {
// 				fmt.Println(i.Val)
// 			}
// 			fmt.Println("counter: ", counter)
// 			counter--
// 			if queue[0].Left != nil {
// 				queue = append(queue, queue[0].Left)
// 				for num, i := range queue {
// 					fmt.Println(num, i.Val)
// 				}
// 				fmt.Println("len queue", len(queue))
// 			}
// 			if queue[0].Right != nil {
// 				queue = append(queue, queue[0].Right)
// 				for num, i := range queue {
// 					fmt.Println(num, i.Val)
// 				}
// 				fmt.Println("len queue", len(queue))

// 			}
// 			res[level] = append(res[level], queue[0].Val)
// 			queue = queue[1:]
// 		}
// 		level++
// 	}
// 	return res
// }

// func main() {
// 	node1 := &TreeNode{3, nil, nil}
// 	node1.Left = &TreeNode{9, nil, nil}
// 	node1.Left.Left = &TreeNode{99, nil, nil}
// 	node1.Right = &TreeNode{20, nil, nil}
// 	node1.Right.Left = &TreeNode{15, nil, nil}
// 	node1.Right.Right = &TreeNode{7, nil, nil}

// 	res := levelOrder(node1)
// 	fmt.Printf("res is %v \n", res)
// }

// func main() {
// 	// val := -1 << (8*unsafe.Sizeof(1) - 1)
// 	// fmt.Println(val)
// 	// fmt.Println(math.MinInt)
// 	// fmt.Println(math.MinInt8)
// 	// fmt.Println(math.MinInt16)
// 	// fmt.Println(math.MinInt32)
// 	// fmt.Println(math.MinInt64)

// }

///////// sortedArrayToBST
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xninbt/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	return sortedArrayToBSThelp(nums, 0, len(nums)-1)
// }

// func sortedArrayToBSThelp(nums []int, start int, end int) *TreeNode {
// 	if start > end {
// 		return nil
// 	}
// 	mid := (start + end) / 2
// 	// var root TreeNode
// 	root := new(TreeNode)
// 	root.Val = nums[mid]
// 	root.Left = sortedArrayToBSThelp(nums, start, mid-1)
// 	root.Right = sortedArrayToBSThelp(nums, mid+1, end)
// 	return root
// }

// func main() {
// 	// node1 := &TreeNode{3, nil, nil}
// 	// node1.Left = &TreeNode{9, nil, nil}
// 	// node1.Left.Left = &TreeNode{99, nil, nil}
// 	// node1.Right = &TreeNode{20, nil, nil}
// 	// node1.Right.Left = &TreeNode{15, nil, nil}
// 	// node1.Right.Right = &TreeNode{7, nil, nil}
// 	var nums []int = []int{1, 3}
// 	res := sortedArrayToBST(nums)
// 	fmt.Printf("res is %v \n", res)
// }

////////// merge sorted arrays
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnumcr/

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	var i int
// 	var j int
// 	for {
// 		if j == n {
// 			break
// 		}
// 		if i == m+j {
// 			for j < n && i < m+n {
// 				nums1[i] = nums2[j]
// 				i++
// 				j++
// 			}
// 			break
// 		}
// 		if nums1[i] >= nums2[j] {
// 			for k := len(nums1) - 1; k > i; k-- {
// 				nums1[k] = nums1[k-1]
// 			}
// 			fmt.Println(nums1)
// 			nums1[i] = nums2[j]
// 			i++
// 			j++
// 			continue
// 		} else {
// 			i++
// 			continue
// 		}
// 	}
// }

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	nums2 := []int{2, 5, 6}
// 	merge(nums1, 3, nums2, 3)
// 	fmt.Println(nums1)
// }

////////// firstBadVersion
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnto1s/

// func firstBadVersion(n int) int {
// 	var min int = 1
// 	var max int = n
// 	mid := (max + min) / 2
// 	for {
// 		if max-min <= 1 {
// 			if isBadVersion(min) {
// 				return min
// 			} else {
// 				return max
// 			}
// 		}
// 		if isBadVersion(mid) {
// 			max = mid
// 			mid = (max + min) / 2
// 		} else {
// 			min = mid
// 			mid = (max + min) / 2
// 		}

// 	}
// }

////////// climbStairs
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn854d/

// func climbStairs(n int)int {
// 	if n <= 1 {
// 		return 1
// 	}
// 	if n < 3 {
// 		return n
// 	}
// 	return climbStairs(n-1)+ climbStairs(n -2)
// }

// func climbStairs(n int) int {
// 	return Fibonacci(n, 1, 1)
// }

// func Fibonacci(n, a, b int) int {
// 	if n <= 1 {
// 		return b
// 	}
// 	return Fibonacci(n-1, b, a+b)
// }

// func climbStairs(n int) int {
// 	if n <= 1 {
// 		return 1
// 	}
// 	dp := make([]int, 3)
// 	dp[1] = 1
// 	dp[2] = 2
// 	for i := 3; i <= n; i++ {
// 		dp = append(dp, dp[i-1]+dp[i-2])
// 	}
// 	return dp[n]
// }

// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	first := 1
// 	second := 2
// 	sum := 0
// 	for n > 2 {
// 		sum = first + second
// 		first = second
// 		second = sum
// 		n--
// 	}
// 	return sum
// }

// *	func climbStairs(n int) int {
// *		sqrt := math.Sqrt(5)
// *		return int((math.Pow((1+sqrt)/2, float64(n+1)) - math.Pow((1-sqrt)/2, float64(n+1))) / sqrt)
// *	}

////////// maxProfit
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn8fsh/

// func maxProfit(prices []int) int {
// if prices == nil || len(prices) == 0 {
// 	return 0
// }
// 	var res int
// 	for i := 0; i < len(prices)-1; i++ {
// 		for j := i + 1; j < len(prices); j++ {
// 			if prices[j]-prices[i] > res {
// 				res = prices[j] - prices[i]
// 			}
// 		}
// 	}
// 	return res
// }

// func maxProfit(prices []int) int {
// 	if prices == nil || len(prices) == 0 {
// 		return 0
// 	}
// 	length := len(prices)
// 	hold := -prices[0]
// 	var nohold int
// 	for i:=1;i<length;i++{
// 		nohold = int(math.Max(float64(nohold), float64(hold+prices[i])))
// 		hold = int(math.Max(float64(hold), float64(-prices[i])))
// 	}
// 	return nohold
// }

// func maxProfit(prices []int) int {
// 	if prices == nil || len(prices) == 0 {
// 		return 0
// 	}
// 	var maxPro int
// 	min := prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] < min {
// 			min = prices[i]
// 		}
// 		// min = int(math.Min(float64(min), float64(prices[i])))
// 		if prices[i] - min > maxPro{
// 			maxPro = prices[i] - min
// 		}
// 		// maxPro = int(math.Max(float64(prices[i]-min), float64(maxPro)))

// 	}
// 	return maxPro
// }

///////// maxSubArray
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn3cg3/

// func maxSubArray(nums []int) int {
// 	length := len(nums)
// 	cur := nums[0]
// 	max := cur
// 	for i := 1; i < length; i++ {
// 		if cur > 0 {
// 			cur += nums[i]
// 		} else {
// 			cur = nums[i]
// 		}
// 		if max < cur {
// 			max = cur
// 		}
// 	}
// 	return max
// }

////////// rob
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnq4km/

// func rob(nums []int) int {
// 	if nums == nil || len(nums) == 0 {
// 		return 0
// 	}
// 	length := len(nums)
// 	dp0 := 0
// 	dp1 := nums[0]

// 	for i := 1; i < length; i++ {
// 		temp := math.Max(float64(dp0), float64(dp1))
// 		dp1 = dp0 + nums[i]
// 		dp0 = int(temp)
// 	}
// 	return int(math.Max(float64(dp0), float64(dp1)))
// }

////////// shuffleArray
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn6gq1/

// type Solution struct {
// 	nums []int
// }

// func Constructor(nums []int) Solution {
// 	return Solution{nums}
// }

// func (this *Solution) Reset() []int {
// 	return this.nums
// }

// func (this *Solution) Shuffle() []int {
// 	nums := make([]int, len(this.nums))
// 	copy(nums, this.nums)
// 	rand.Shuffle(len(nums), func(i, j int) {
// 		nums[i], nums[j] = nums[j], nums[i]
// 	})
// 	return nums
// }

////////// minStack
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnkq37/

// type MinStack struct {
// 	sta []int
// }

// func Constructor() MinStack {
// 	return MinStack{}
// }

// func (this *MinStack) Push(val int) {
// 	this.sta = append(this.sta, val)
// }

// func (this *MinStack) Pop() {
// 	this.sta = this.sta[:len(this.sta)-1]
// }

// func (this *MinStack) Top() int {
// 	return this.sta[len(this.sta)-1]
// }

// func (this *MinStack) GetMin() int {
// 	min := this.sta[0]
// 	for i:=1;i< len(this.sta);i++{
// 		if this.sta[i]< min {
// 			min = this.sta[i]
// 		}
// 	}
// 	return min
// }

// *type MinStack struct{
// *	elems []int
// *	mins []int
// *}
//
// *func Constructor() MinStack{
// *	return MinStack{
// *		elems: make([]int, 0),
// *		mins: make([]int, 0),
// *	}
// *}
//
// *func (this *MinStack) Push (x int){
// *	this.elems = append(this.elems, x)
// *	if len(this.mins) == 0 {
// *		this.mins = append(this.mins, x)
// *	} else {
// *		if this.elems[len(this.elems)-1]>= x {
// *			this.mins = append(this.mins, x)
// *		}
// *	}
// *}
//
// *func (this *MinStack) Pop() {
// *	elem := this.elems[len(this.elems)-1]
// *	this.elems = this.elems[:len(this.elems)-1]
// *	if elem == this.mins[len(this.mins)-1] {
// *		this.mins = this.mins[:len(this.mins)-1]
// *	}
// *}
//
// *func (this *MinStack) Top() int {
// *	return this.elems[len(this.elems)-1]
// *}
//
// *func (this *MinStack) GetMin() int {
// *	return this.mins[len(this.mins)-1]
// *}

////////// fizzBuzz
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xngt85/

// func fizzBuzz(n int) []string {
// 	var res []string
// 	for i:=1;i<= n ; i++{
// 		if i % 3 == 0 && i % 5 == 0 {
// 			res = append(res, "FizzBuzz")
// 		} else if i % 3 == 0 {
// 			res = append(res, "Fizz")
// 		} else if i %5 ==0 {
// 			res = append(res, "Buzz")
// 		} else {
// 			res = append(res, strconv.Itoa(i))
// 		}
// 	}
// 	return res
// }

////////// countPrimes
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnzlu6/

// func isPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func countPrimes(n int) int {
// 	if n < 3 {
// 		return 0
// 	}
// 	var count int
// 	for i := 2; i < n; i++ {
// 		if isPrime(i) {
// 			fmt.Println("i: ", i)
// 			count++
// 		}
// 	}
// 	fmt.Println("Count: ", count)
// 	return count
// }

// func countPrimes(n int) int {
// 	arr := make([]bool, n)
// 	var cnt int
// 	for i := 2; i < n; i++ {
// 		if arr[i] {
// 			continue
// 		}
// 		fmt.Println("i: ", i)
// 		cnt++
// 		for j := i; j < n; j = j + i {
// 			arr[j] = true
// 		}
// 	}
// 	fmt.Println(cnt)
// 	return cnt
// }

// func isPrime(n int) bool {
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func countPrimes(n int) int {
// 	m := make([]int, n+1)
// 	count := 0

// 	for i := 0; i <= n; i++ {
// 		m[i] = 1
// 	}

// 	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
// 		if isPrime(i) {
// 			for j := i * 2; j < n; j = j + i {
// 				m[j] = 0
// 			}
// 		}
// 	}
// 	for i := 2; i < n; i++ {
// 		if m[i] == 1 {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {
// 	var num int = 8
// 	countPrimes(num)
// }

////////// isPowerOfThree
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnsdi2/

// func isPowerOfThree(n int) bool {
// 	if n == 1 {
// 		return true
// 	}
// 	seed := 3
// 	for i := 0; seed <= n; i++ {
// 		if seed == n {
// 			return true
// 		}
// 		seed = seed * 3
// 	}
// 	return false
// }

///////// romanToInt
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn4n7c/

// func romanToInt(s string) int {
// 	m := make(map[string]int)
// 	m["I"] = 1
// 	m["IV"] = 4
// 	m["V"] = 5
// 	m["IX"] = 9
// 	m["X"] = 10
// 	m["XL"] = 40
// 	m["L"] = 50
// 	m["XC"] = 90
// 	m["C"] = 100
// 	m["CD"] = 400
// 	m["D"] = 500
// 	m["CM"] = 900
// 	m["M"] = 1000
// 	var res int
// 	for i := 0; i < len(s); i++ {
// 		if i < len(s)-1 {
// 			numString := string(s[i]) + string(s[i+1])
// 			if m[numString] != 0 {
// 				res += m[numString]
// 				i++
// 			} else {
// 				res += m[string(s[i])]
// 			}
// 		} else {
// 			res += m[string(s[i])]
// 		}
// 	}
// 	return res
// }

// func main() {
// 	fmt.Println(romanToInt("XLL"))
// }

////////// hammingWeight
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn1m0i/

// func hammingWeight(num uint32) int {
// 	var count int
// 	for num != 0 {
// 		count += int(num) & 1
// 		num = num >> 1
// 	}
// 	return count
// }

////////// hammingDistance
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnyode/

// func hammingWeight(num uint32) int {
// 	var count int
// 	for num != 0 {
// 		count += int(num) & 1
// 		num = num >> 1
// 	}
// 	return count
// }

// func hammingDistance(x, y int) int {
// 	num := x ^ y
// 	return hammingWeight(uint32(num))
// }

///////// generate Pascal's Triangle
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xncfnv/

// func generate(numRows int) [][]int {
// 	res := make([][]int, numRows)
// 	for j := 1; j <= numRows; j++ {
// 		res[j-1] = make([]int, j)
// 	}
// 	if numRows == 1 {
// 		res[0][0] = 1
// 		return res
// 	} else if numRows == 2 {
// 		res[0][0] = 1
// 		res[1][0] = 1
// 		res[1][1] = 1
// 		return res
// 	} else {
// 		res[0][0] = 1
// 		res[1][0] = 1
// 		res[1][1] = 1
// 		for i := 2; i < numRows; i++ {
// 			for j := 0; j < numRows; j++ {
// 				if j == 0 {
// 					res[i][j] = 1
// 					continue
// 				}
// 				if j == i {
// 					res[i][j] = 1
// 					break
// 				}
// 				res[i][j] = res[i-1][j-1] + res[i-1][j]
// 			}
// 		}
// 		return res
// 	}
// }

// func generate(numRows int) [][]int {
// 	res := make([][]int, numRows)
// 	for j := 1; j <= numRows; j++ {
// 		res[j-1] = make([]int, j)
// 	}
// 	if numRows == 1 {
// 		res[0][0] = 1
// 		return res
// 	} else if numRows == 2 {
// 		res[0][0] = 1
// 		res[1][0] = 1
// 		res[1][1] = 1
// 		return res
// 	} else {
// 		res[0][0] = 1
// 		res[1][0] = 1
// 		res[1][1] = 1
// 		for i := 2; i < numRows; i++ {
// 			for j := 0; j < numRows; j++ {
// 				if j == 0 {
// 					res[i][j] = 1
// 					continue
// 				}
// 				if j == i {
// 					res[i][j] = 1
// 					break
// 				}
// 				if j > (i+1)/2 {
// 					res[i][j] = res[i][i-j]
// 				}
// 				res[i][j] = res[i-1][j-1] + res[i-1][j]
// 			}
// 		}
// 		return res
// 	}
// }

// func main() {
// 	for _, j := range generate(6) {
// 		fmt.Println(j)
// 	}
// 	fmt.Println(generate(6))
// }

////////// isValidBracket
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnbcaj/

// func isValid(s string) bool {
// 	var res1 []string
// 	for i := 0; i < len(s); i++ {
// 		switch string(s[i]) {
// 		case "(":
// 			res1 = append(res1, ")")
// 		case "[":
// 			res1 = append(res1, "]")
// 		case "{":
// 			res1 = append(res1, "}")
// 		case ")":
// 			if len(res1) == 0 {
// 				return false
// 			} else if res1[len(res1)-1] == ")" {
// 				res1 = res1[:len(res1)-1]
// 				continue
// 			} else {
// 				return false
// 			}
// 		case "]":
// 			if len(res1) == 0 {
// 				return false
// 			} else if res1[len(res1)-1] == "]" {
// 				res1 = res1[:len(res1)-1]
// 				continue
// 			} else {
// 				return false
// 			}
// 		case "}":
// 			if len(res1) == 0 {
// 				return false
// 			} else if res1[len(res1)-1] == "}" {
// 				res1 = res1[:len(res1)-1]
// 				continue
// 			} else {
// 				return false
// 			}
// 		}
// 	}
// 	return len(res1) == 0
// }

////////// missingNumber
// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnj4mt/

// func missingNumber(nums []int) int {
// 	var res1 int
// 	var res2 int
// 	for _, i := range nums {
// 		res1 = res1 + i
// 	}
// 	for i := 1; i <= len(nums); i++ {
// 		res2 = res2 + i
// 	}
// 	return res2 - res1
// }

////////// threeSum
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvpj16/
// func IfEqual(n11, n22 []int) bool {
// 	// fmt.Println("n11n22", n11, n22)
// 	n1 := make([]int, len(n11))
// 	n2 := make([]int, len(n22))
// 	copy(n1, n11)
// 	copy(n2, n22)
// 	// fmt.Println(n1, n2)
// 	if len(n1) != len(n2) {
// 		return false
// 	}
// 	for i := len(n1) - 1; i >= 0; i-- {
// 		for j := len(n2) - 1; j >= 0; j-- {
// 			if n1[i] == n2[j] {
// 				n1 = append(n1[:i], n1[i+1:]...)
// 				n2 = append(n2[:j], n2[j+1:]...)
// 				break
// 			} else if j == 0 {
// 				return false
// 			}
// 		}
// 	}
// 	// fmt.Println("n11n22", n11, n22)
// 	return len(n1)+len(n2) == 0
// }
// func IfExist(numsSlices [][]int, numSlice []int) bool {
// 	// fmt.Println("res: ", numsSlices)
// 	for _, i := range numsSlices {
// 		// fmt.Println(i, numSlice)
// 		if IfEqual(i, numSlice) {
// 			return true
// 		}
// 	}
// 	return false
// }
// func ThreeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	var res [][]int
// 	if len(nums) < 3 {
// 		return [][]int{}
// 	}
// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				// fmt.Println("nums: ", nums)
// 				// fmt.Printf("i,j,k: %v, %v, %v\n", i, j, k)
// 				a := nums[i]
// 				b := nums[j]
// 				c := nums[k]

// 				if a+b+c == 0 {
// 					// fmt.Printf("a, b, c, : %v, %v, %v\n", a, b, c)
// 					compTemp := []int{a, b, c}
// 					if !IfExist(res, compTemp) {
// 						res = append(res, compTemp)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func threeSum(nums []int) [][]int {
// 	var res [][]int
// 	if len(nums) < 3 {
// 		return res
// 	}

// 	sort.Ints(nums)
// 	for i := 0; i < len(nums)-2; i++ {
// 		if nums[i] > 0 {
// 			break
// 		}
// 		if i > 0 && nums[i-1] == nums[i] {
// 			continue
// 		}
// 		for j, k := i+1, len(nums)-1; j < k; {
// 			sum := nums[i] + nums[j] + nums[k]
// 			if sum == 0 {
// 				tmp := []int{nums[i], nums[j], nums[k]}
// 				res = append(res, tmp)
// 				for j+1 < k && nums[j] == nums[j+1] {
// 					j++
// 				}
// 				for k-1 > j && nums[k] == nums[k-1] {
// 					k--
// 				}
// 				j++
// 				k--
// 			} else if sum > 0 {
// 				k--
// 			} else {
// 				j++
// 			}

// 		}
// 	}
// 	return res
// }

// func main() {
// 	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
// 	fmt.Println(res)
// }

////////// setZeros
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvmy42/

// func setZeroes(matrix [][]int) {
// 	col := make([]bool, len(matrix[0]))
// 	row := make([]bool, len(matrix))

// 	for i, rowRes := range matrix {
// 		for j, colRes := range rowRes {
// 			if colRes == 0 {
// 				col[j] = true
// 				row[i] = true
// 			}
// 		}
// 	}
// 	for i, rowBool := range row {
// 		if rowBool {
// 			for j := 0; j < len(matrix[i]); j++ {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}

// 	for j, colBool := range col {
// 		if colBool {
// 			for i := 0; i < len(matrix); i++ {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// }

// func setZeroes(matrix [][]int) {
// 	if matrix == nil || len(matrix) == 0 && len(matrix[0]) == 0 {
// 		return
// 	}

// 	var firstRow bool
// 	var firstCol bool
// 	for i := 0; i < len(matrix[0]); i++ {
// 		if matrix[0][i] == 0 {
// 			firstRow = true
// 			break
// 		}
// 	}
// 	for i := 0; i < len(matrix); i++ {
// 		if matrix[i][0] == 0 {
// 			firstCol = true
// 			break
// 		}
// 	}

// 	if len(matrix) == 1 {
// 		if firstRow {
// 			for i := 0; i < len(matrix[0]); i++ {
// 				matrix[0][i] = 0
// 			}
// 		} else {
// 			return
// 		}
// 	}
// 	if len(matrix[0]) == 1 {
// 		if firstCol {
// 			for i := 0; i < len(matrix); i++ {
// 				matrix[i][0] = 0
// 			}
// 		} else {
// 			return
// 		}
// 	}

// 	for i := 1; i < len(matrix); i++ {
// 		for j := 1; j < len(matrix[0]); j++ {
// 			if matrix[i][j] == 0 {
// 				matrix[0][j] = 0
// 				matrix[i][0] = 0
// 			}
// 		}
// 	}

// 	for i := 1; i < len(matrix[0]); i++ {
// 		if matrix[0][i] == 0 {
// 			for j := 1; j < len(matrix); j++ {
// 				matrix[j][i] = 0
// 			}
// 		}
// 	}

// 	for i := 1; i < len(matrix); i++ {
// 		if matrix[i][0] == 0 {
// 			for k := 1; k < len(matrix[0]); k++ {
// 				matrix[i][k] = 0
// 			}
// 		}
// 	}
// 	if firstRow {
// 		for i := 0; i < len(matrix[0]); i++ {
// 			matrix[0][i] = 0
// 		}
// 	}
// 	if firstCol {
// 		for i := 0; i < len(matrix); i++ {
// 			matrix[i][0] = 0
// 		}
// 	}
// }

////////// groupAnagrams
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvaszc/

// func IsEqual(a, b string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	aa := strings.Split(a, "")
// 	bb := strings.Split(b, "")
// 	sort.Strings(aa)
// 	sort.Strings(bb)

// 	for i := 0; i < len(aa); i++ {
// 		if aa[i] != bb[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func GroupAnagrams(strs []string) [][]string {
// 	var res [][]string
// 	res = append(res, []string{strs[0]})
// 	if len(strs) < 2 {
// 		return res
// 	}

// 	for i := 1; i < len(strs); i++ {
// 		for j := 0; j < len(res); j++ {

// 			if IsEqual(res[j][0], strs[i]) {
// 				res[j] = append(res[j], strs[i])
// 				break
// 			} else {
// 				if j == len(res)-1 {
// 					res = append(res, []string{strs[i]})
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func main() {
// 	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	res := GroupAnagrams(input)
// 	fmt.Println(res)
// }
// func aSort(s string) string {
// 	x := []byte(s)
// 	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
// 	return string(x)
// }
// func groupAnagrams(strs []string) [][]string {
// 	if len(strs) == 0 {
// 		return [][]string{}
// 	}
// 	rMap := make(map[string][]int, 0)
// 	for i, v := range strs {
// 		str := aSort(v)
// 		if _, ok := rMap[str]; !ok {
// 			rMap[str] = []int{i}
// 		} else {
// 			rMap[str] = append(rMap[str], i)
// 		}
// 	}
// 	res := make([][]string, 0)
// 	for _, v := range rMap {
// 		a := make([]string, 0)
// 		for _, x := range v {
// 			a = append(a, strs[x])
// 		}
// 		res = append(res, a)
// 	}
// 	return res
// }

////////// lengthOfLongestSubstring
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xv2kgi/

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) < 2 {
// 		return len(s)
// 	}
// 	var res1 []byte
// 	// var res2 []byte
// 	res1 = append(res1, s[0])
// 	length := 1
// 	for i := 1; i < len(s); i++ {
// 		for j := 0; j < len(res1); j++ {
// 			if res1[j] == s[i] {
// 				// fmt.Println("equal")
// 				// fmt.Println(string(res1[j]), string(s[i]))
// 				res1 = append(res1[j+1:], s[i])
// 				// fmt.Println("***", string(res1))
// 				break
// 			} else {
// 				if j == len(res1)-1 {
// 					// fmt.Println("unequal")
// 					// fmt.Println(string(res1), string(s[i]))
// 					res1 = append(res1, s[i])
// 					if length < len(res1) {
// 						length = len(res1)
// 					}
// 					break
// 				}
// 			}
// 			// fmt.Println("***", string(res1))
// 		}
// 	}
// 	return length
// }

// func main() {
// 	res := lengthOfLongestSubstring("dvdf")
// 	fmt.Println(res)
// }

////////// longestPalindrome
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvn3ke/
// func longestPalindrome(s string) string {
// 	if len(s) < 2 {
// 		return s
// 	}

// 	var start int
// 	var maxLen int
// 	length := len(s)

// 	for i := 0; i < length; {
// 		if length-i <= maxLen/2 {
// 			break
// 		}

// 		left := i
// 		right := i

// 		for right < length-1 && s[right+1] == s[right] {
// 			right++
// 		}
// 		i = right + 1
// 		for right < length-1 && left > 0 && s[right+1] == s[left-1] {
// 			right++
// 			left--
// 		}
// 		if right-left+1 > maxLen {
// 			start = left
// 			maxLen = right - left + 1
// 		}

// 	}
// 	return s[start : start+maxLen]
// }

////////// increasingTriplet
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvvuqg/

// func increasingTriplet(nums []int) bool {
// 	var min int = math.MaxInt
// 	var mid int = math.MaxInt

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] <= min {
// 			min = nums[i]
// 		} else if nums[i] <= mid {
// 			mid = nums[i]
// 		} else {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	var nums []int = []int{2, 1, 5, 0, 4, 6}
// 	increasingTriplet(nums)
// }

////////// addTwoNumbers
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvw73v/
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	dummyNode := new(ListNode)
// 	preNode := dummyNode

// 	var carry int

// 	for l1 != nil || l2 != nil || carry != 0 {
// 		sum := carry
// 		if l1 != nil {
// 			sum += l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			sum += l2.Val
// 			l2 = l2.Next
// 		}
// 		preNode.Next = new(ListNode)
// 		preNode.Next.Val = sum % 10

// 		carry = sum / 10
// 		preNode = preNode.Next
// 	}

// 	return dummyNode.Next
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	dummyNode := new(ListNode)
// 	helper(dummyNode, l1, l2, 0)
// 	return dummyNode.Next
// }

// func helper(preNode, l1, l2 *ListNode, carry int) {
// 	if l1 == nil && l2 == nil && carry == 0 {
// 		return
// 	}
// 	sum := carry

// 	if l1 != nil {
// 		sum += l1.Val
// 		l1 = l1.Next
// 	}
// 	if l2 != nil {
// 		sum += l2.Val
// 		l2 = l2.Next
// 	}
// 	preNode.Next = new(ListNode)
// 	preNode.Next.Val = sum % 10
// 	carry = sum / 10
// 	helper(preNode.Next, l1, l2, carry)
// }

////////// oddEvenList
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvdwtj/
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func oddEvenList(head *ListNode) *ListNode {

// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	oddTemp := head
// 	evenTemp := head.Next
// 	evenHead := head.Next

// 	for evenTemp != nil && evenTemp.Next != nil {
// 		oddTemp.Next = oddTemp.Next.Next

// 		evenTemp.Next = evenTemp.Next.Next

// 		oddTemp = oddTemp.Next
// 		evenTemp = evenTemp.Next
// 	}

// 	oddTemp.Next = evenHead
// 	return head
// }

////////// getIntersectionNode
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xv02ut/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	currA := headA
// 	currB := headB

// 	var countA int
// 	var countB int

// 	for currA != nil {
// 		countA++
// 		currA = currA.Next
// 	}
// 	for currB != nil {
// 		countB++
// 		currB = currB.Next
// 	}
// 	currA = headA
// 	currB = headB

// 	for i := 0; i < countA-countB; i++ {
// 		currA = currA.Next
// 	}
// 	for i := 0; i < countB-countA; i++ {
// 		currB = currB.Next
// 	}

// 	for currA != nil {
// 		if currA == currB {
// 			return currA
// 		}
// 		currA = currA.Next
// 		currB = currB.Next
// 	}
// 	return nil
// }

////////// inorderTraversal
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xv7pir/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func f(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	var tmp []int
// 	if root.Left != nil {
// 		tmp = append(tmp, f(root.Left)...)
// 	}

// 	tmp = append(tmp, root.Val)

// 	if root.Right != nil {
// 		tmp = append(tmp, f(root.Right)...)
// 	}

// 	return tmp
// }

// func inorderTraversal(root *TreeNode) []int {
// 	return f(root)
// }

// func inorderTraversal(root *TreeNode) (tmp []int) {
// 	if root == nil {
// 		return []int{}
// 	}
// 	if root.Left != nil {
// 		tmp = append(tmp, inorderTraversal(root.Left)...)
// 	}

// 	tmp = append(tmp, root.Val)

// 	if root.Right != nil {
// 		tmp = append(tmp, inorderTraversal(root.Right)...)
// 	}

// 	return tmp
// }

// type Stack struct {
// 	items []*TreeNode
// }

// func (m *Stack) Len() int {
// 	return len(m.items)
// }

// func (m *Stack) Push(s *TreeNode) {
// 	m.items = append(m.items, s)
// }

// func (m *Stack) Pop() *TreeNode {
// 	last := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return last
// }

// func inorderTraversal(root *TreeNode) (ret []int) {
// 	node := root
// 	stack := &Stack{
// 		items: make([]*TreeNode, 0),
// 	}
// 	for node != nil || stack.Len() > 0 {
// 		if node != nil {
// 			stack.Push(node)
// 			node = node.Left
// 		} else {
// 			node = stack.Pop()
// 			ret = append(ret, node.Val)
// 			node = node.Right
// 		}
// 	}
// 	return
// }

////////// zigzagLevelOrder
// https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvle7s/

// func levelOrder(root *TreeNode) []int {
// 	var res []int

// 	if root == nil {
// 		return res
// 	}

// 	var arr []*TreeNode
// 	arr = append(arr, root)

// 	for len(arr) > 0 {
// 		length := len(arr)
// 		for i := 0; i < length; i++ {
// 			node := arr[0]
// 			if node.Left != nil {
// 				arr = append(arr, node.Left)
// 			}
// 			if node.Right != nil {
// 				arr = append(arr, node.Right)
// 			}
// 			res = append(res, node.Val)
// 			arr = arr[1:]
// 		}
// 	}
// 	return res
// }

// func zigzagLevelOrder(root *TreeNode) [][]int {
// 	var res [][]int
// 	isReverse := true
// 	if root == nil {
// 		return res
// 	}

// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		k := len(queue)
// 		ans := make([]int, k)
// 		for i := 0; i < k; i++ {
// 			node := queue[i]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 			if isReverse {
// 				ans[i] = node.Val
// 			} else {
// 				ans[k-i-1] = node.Val
// 			}
// 		}
// 		res = append(res, ans)
// 		isReverse = !isReverse
// 		queue = queue[k:]
// 	}
// 	return res
// }

////////// buildTree
// https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvix0d/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	if len(inorder) == 0 {
// 		return nil
// 	}
// 	// First of preorder is the root
// 	res := new(TreeNode)
// 	res.Val = preorder[0]
// 	var index int
// 	// Devide inorder to left and right
// 	for ; index < len(inorder); index++ {
// 		if preorder[0] == inorder[index] {
// 			break
// 		}
// 	}

// 	// If first of preorder in left, root.Left = preorder[0]
// 	if index > 0 {
// 		res.Left = buildTree(preorder[1:index+1], inorder[0:index])
// 	}
// 	// If first of preorder in right, root.Right = preorder[0]
// 	if index < len(inorder)-1 {
// 		res.Right = buildTree(preorder[index+1:], inorder[index+1:])
// 	}
// 	return res
// }

// func main() {
// 	preorder := []int{3, 9, 20, 15, 7}
// 	inorder := []int{9, 3, 15, 20, 7}
// 	buildTree(preorder, inorder)
// }
