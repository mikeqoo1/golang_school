package leetcode

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

//ListNode linklist
type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue []*ListNode

func (h Queue) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Queue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Queue) Len() int {
	return len(h)
}

func (h *Queue) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Queue) Pop() interface{} {
	head := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return head
}

//FindMedianSortedArrays 尋找中位數
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	nums3 := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i = i + 1
		} else {
			nums3 = append(nums3, nums2[j])
			j = j + 1
		}
	}
	if i < len(nums1) {
		for ; i < len(nums1); i++ {
			nums3 = append(nums3, nums1[i])
		}
	}
	if j < len(nums2) {
		for ; j < len(nums2); j++ {
			nums3 = append(nums3, nums2[j])
		}
	}
	lenSums3 := len(nums3)
	if lenSums3%2 == 0 {
		return (float64(nums3[lenSums3/2-1]) + float64(nums3[lenSums3/2])) / 2.0
	} else {
		return float64(nums3[(lenSums3+1)/2-1])
	}
}

//Reverse 反轉整數 123 -> 321 -123 -> -321
func Reverse(x int) int {
	reverseans := 0
	if x < 0 {
		x = -x
		for x > 0 {
			fmt.Println("0.")
			reverseans *= 10
			fmt.Println("1.", "x =", x)
			fmt.Println("1.", "reverseans =", reverseans)
			reverseans += x % 10
			fmt.Println("2.", "x =", x)
			fmt.Println("2.", "reverseans =", reverseans)
			x /= 10
			fmt.Println("3.", "x =", x)
			fmt.Println("3.", "reverseans =", reverseans)
		}

		reverseans = -reverseans
		if reverseans < -2147483648 {
			return 0
		}

		return reverseans
	}

	for x > 0 {
		fmt.Println("0.")
		reverseans *= 10
		fmt.Println("1.", "x =", x)
		fmt.Println("1.", "reverseans =", reverseans)
		reverseans += x % 10
		fmt.Println("2.", "x =", x)
		fmt.Println("2.", "reverseans =", reverseans)
		x /= 10
		fmt.Println("3.", "x =", x)
		fmt.Println("3.", "reverseans =", reverseans)
	}

	if reverseans > 2147483647 {
		return 0
	}

	return reverseans
}

//MyAtoi MyAtoi
func MyAtoi(str string) int {
	//先判斷是不是正常的字串, 再判斷正負, 再合成新的字串, 接下來就可以轉了
	//假設 "-123"
	//找出有效起始位置, 開始組裝

	if len(str) == 0 {
		return 0
	}
	strByte := []byte(str)
	Head := true
	var newStr []byte
	for i := 0; i < len(strByte); i++ {
		if strByte[i] == ' ' {
			if Head {
				continue
			} else {
				break
			}
		}
		if strByte[i] >= '0' && strByte[i] <= '9' || (strByte[i] == '-' || strByte[i] == '+') && Head {
			Head = false
			newStr = append(newStr, strByte[i])
			continue
		}
		break
	}

	res, _ := strconv.Atoi(string(newStr))

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res

}

//IsPalindrome 判斷回文數
func IsPalindrome(x int) bool {
	// 先把原數字%10 在 /10 依照此步驟來反轉
	if x < 0 {
		return false
	}
	temp := x
	var x2 int
	for temp != 0 {
		pop := temp % 10
		//fmt.Println("pop", pop)
		temp = temp / 10
		//fmt.Println("temp", temp)
		x2 = x2*10 + pop
		//fmt.Println("x2", x2)
	}
	if x2 == x {
		return true
	}
	return false
}

//RegularExpressionMatching 10. Regular Expression Matching
func RegularExpressionMatching(s string, p string) bool {
	for true {
		r1, w1 := utf8.DecodeRuneInString(p)
		if w1 <= 0 {
			break
		}
		p = p[w1:]

		r2, w2 := utf8.DecodeRuneInString(p)
		if r2 == '*' {
			// Zero or more
			p = p[w2:]
			for true {
				if ok := RegularExpressionMatching(s, p); ok {
					return true
				}
				r3, w3 := utf8.DecodeRuneInString(s)
				if (r1 == '.' || r1 == r3) && w3 > 0 {
					s = s[w3:]
				} else {
					break
				}
			}
		} else {
			// One
			r3, w3 := utf8.DecodeRuneInString(s)
			if w3 <= 0 || !(r1 == '.' || r1 == r3) {

				return false
			}
			s = s[w3:]
		}
	}

	return len(s) <= 0
	//快速版
	// n, m := len(s), len(p)
	// dp := make([][]bool, n+1)
	// for i:=0;i<=n;i++ {
	//     dp[i] = make([]bool, m+1)
	// }
	// dp[0][0] = true
	// for j:=2; j<=m; j++ {
	//     if p[j-1] == '*' {
	//         dp[0][j] = dp[0][j-2]
	//     }
	// }
	// // fmt.Printf("%#v\n",dp[0])
	// for i:=1; i<=n; i++ {
	//     for j:=1; j<=m; j++ {
	//         if s[i-1] == p[j-1] || p[j-1] == '.' {
	//             dp[i][j] = dp[i-1][j-1]
	//             continue
	//         }
	//         if p[j-1] == '*' {
	//             if p[j-2] != '.' && p[j-2] != s[i-1] {
	//                 dp[i][j] = dp[i][j-2]
	//             } else {
	//                 dp[i][j] = dp[i][j-2] || dp[i-1][j]
	//             }
	//         }
	//     }
	// }
	// // for _, r := range dp{
	// //     fmt.Printf("%v\n",r)
	// // }
	// return dp[n][m]
}

//MaxArea leetCode11
func MaxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	var left = 0
	var right = len(height) - 1
	var result = 0
	var temp = 0
	for left < right {
		temp = (right - left) * min(height[left], height[right])
		result = max(result, temp)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//IntToRoman 整數轉羅馬數字
func IntToRoman(num int) string {
	if num <= 0 {
		return ""
	}
	var ans string
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(nums); i++ { //舉例: 進來的數字 3978,
		for num >= nums[i] { //第一次迴圈 3978 > 1000 第二次迴圈 2978 > 1000
			num -= nums[i]    //3978 - 1000 = 2978  2978 - 1000 = 1978
			ans += symbols[i] // 加一個M , 再加一個M 以此類推...
		}
	}
	return ans
}

//RomanToInt 羅馬數字轉數字
func RomanToInt(s string) int {
	symbolToI := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	n := len(s)
	number := 0
	i := 0
	var currentBase int

	for i < n {
		currentSum := 0
		prevBase := symbolToI[string(s[i])]
		for j := i; j < n; j++ {
			currentBase = symbolToI[string(s[j])]
			if currentBase != prevBase {
				break
			}
			i++
			currentSum += currentBase
		}

		//例如: IV or IX
		if prevBase < currentBase {
			currentSum = currentBase - prevBase
			i++
		}

		number += currentSum
	}

	return number
}

//LongestCommonPrefix 最長共同前綴
func LongestCommonPrefix(strs []string) string {
	//strings.HasPrefix 檢查前墜是否存在
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	//[aim air aimless]

	prefix := strs[0]                //把第一個給到前墜字串 aim
	for i := 1; i < len(strs); i++ { //開始逐一訪問, 從1開始就好, 跳過自己
		for strings.HasPrefix(strs[i], prefix) == false { //當這個字串沒有找到時, (air, aim) = false
			prefix = string([]rune(prefix)[:len(prefix)-1]) // aim 往前一位 prefix = ai, 繼續找, 直到true
		}
	}
	return prefix
}

//ThreeSum LeetCodeNo15
func ThreeSum(nums []int) [][]int {
	//input: [-1,0,1,2,-1,-4]
	ans := [][]int{}
	sort.Ints(nums) //先排序,  [-4 -1 -1 0 1 2]
	fmt.Println("排序後", nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { //過濾重複的 -1
			continue
		}
		j, k := i+1, len(nums)-1 //一開始 i代表-4 j代表-1 k代表2, 固定i之後 接下來就類似2Sum
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}

//ThreeSumClosest LeetCodeNo16
//首先, 對Array排序。
//接著初始化一個整數變數。就是要return回去的答案。
//訪問每一array中每一個數字。
//每找一個數字, 可以在array中找剩下的數字, 這2個數字和現在的數字加起來應該樣在target附近。當騙尋结束, 就得到结果。
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ret int
	length := len(nums)
	sum := nums[0] + nums[1] + nums[2]
	ret = sum
	diff := abs(sum - target)

	for i := 0; i < length-2; i++ {
		for j, k := i+1, length-1; j < k; {
			sum = nums[i] + nums[j] + nums[k]
			if sum-target < 0 {
				j++
			} else {
				k--
			}
			if diff > abs(sum-target) {
				ret = sum
				diff = abs(sum - target)
			}
		}
	}
	return ret
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

//LetterCombinations LeetCodeNo17
//用遞迴的方式實作
func LetterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	if len(digits) == 1 {
		switch digits {
		case "2":
			return []string{"a", "b", "c"}
		case "3":
			return []string{"d", "e", "f"}
		case "4":
			return []string{"g", "h", "i"}
		case "5":
			return []string{"j", "k", "l"}
		case "6":
			return []string{"m", "n", "o"}
		case "7":
			return []string{"q", "p", "r", "s"}
		case "8":
			return []string{"t", "u", "v"}
		case "9":
			return []string{"w", "x", "y", "z"}
		case "0":
			return []string{" "}
		default:
			return []string{""}
		}
	}

	r := make([]string, 0)

	for _, front := range LetterCombinations(digits[:1]) {
		fmt.Println("digits[:1]", digits[:1], "value", front)
		for _, back := range LetterCombinations(digits[1:]) {
			fmt.Println("digits[1:]", digits[1:], "value", back)
			r = append(r, front+back)
		}
	}

	return r

}

/*FourSum 4數之和
將array從小到大排列, 雙重for loop, 固定nums[i]和nums[j],
雙指針left和right分别為j+1, len(nums) - 1,
令sum = nums[i] + nums[j] + nums[left] + nums[right],
因為sum要等於target, 所以sum偏小時, left向左移動, sum偏大時, right向右移動, 如果出現sum == target, 存储该四元组组合
*/
func FourSum(nums []int, target int) [][]int {
	count := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < count-3; i++ {
		//去掉重複的值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//此時最小的sum都大於target, 後面就不用查詢了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		//nums[i]固定, 最大的sum小於target, 增大nums[i]繼續查詢
		if nums[i]+nums[count-3]+nums[count-2]+nums[count-1] < target {
			continue
		}
		for j := i + 1; j < count-2; j++ {
			//去掉重複的值
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			//此時最小的sum都大於target, 後面就不用查詢了
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			//nums[i]和nums[j]固定, 最大的sum小於target, 增大nums[j]繼續查詢
			if nums[i]+nums[j]+nums[count-2]+nums[count-1] < target {
				continue
			}
			left, right := j+1, count-1
			for left < right {
				//去掉重複的值
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				//去掉重複的值
				if right < count-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					temp := make([]int, 0)
					temp = append(temp, nums[i], nums[j], nums[left], nums[right])
					result = append(result, temp)
					left++
					right--
				} else if sum < target {
					left++
				} else if sum > target {
					right--
				}
			}
		}
	}
	return result
}

/*RemoveNthFromEnd leetcode19
[1,2,3,4,5] n=3等 輸出 [1,2,4,5], 從後面數來的第3個要消失
宣告2個指標 都從頭開始
當x走n步, 會指到4的地方
接下來y也開始走, 不過不是n步, 而是x走到結尾的步數
4->5 要走1步
所以y也走1步到2的地方, 代表下一個元素3是不要的, 所以把2的下一個再指到下一個, 回傳

4走1步到結尾       4->5
2走3步到結尾 2->3->4->5
代表3就是要刪掉的直,所以把2->->4
*/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	x, y := head, head
	for ; n > 0; n-- {
		x = x.Next
	}
	if nil == x {
		return head.Next
	}
	for nil != x.Next {
		x = x.Next
		y = y.Next
	}
	y.Next = y.Next.Next
	return head
}

/*IsValid leetcode20
用stack來處理, 分成左右2邊
1. 當前括號是左括號时, 壓入stack。
2. 當前括號是右括號时, stack的最上面如果不是對應的左括號, 則為無效組合。否則, pop掉stack里的左括號。
3. 所有字串都判斷處理過後, stack應該是空, 否則則無效。
*/
func IsValid(s string) bool {
	para := map[string]string{")": "(", "}": "{", "]": "["} //右邊當key, 左邊當value
	stack := []string{}                                     //建立stack
	for _, value := range s {                               //遍尋字串
		key := string(value)
		if _, ok := para[key]; !ok { //把字串的值,丟到map尋找對應的value
			stack = append(stack, key) //(Push)
		} else {
			//stack[len(stack)-1] (Top element)
			if len(stack) == 0 || stack[len(stack)-1] != para[key] {
				return false
			}
			stack = stack[:len(stack)-1] //(Pop)
		}
	}
	return len(stack) == 0
}

/*MergeTwoLists leetcode21
建一個新的ListNode, 接著 l1 和 l2 比大小, 小的先放
最佳解法: 遞迴
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists(l1, l2.Next)
		return l2
	}
}

/*GenerateParenthesis leetcode22
第一步, 是得到n對括號的排列组合方式, 這個必需將左右括號分開排列组合,不然就只有一種
這樣的話只需要通過遞迴的方式將n個左右括號排列组合的可能全部形成String值, 裝到List中即可
第二步, 是得到有效的括號, 即先有一個左括號, 才能有一個右括號。所以將遞迴方法中的判斷條件
*/
func GenerateParenthesis(n int) []string {
	str := make([]string, 0)
	position := make([]int, n*2) //第i個小標示示位置i上是左括號还是右括號，0表示左括號1表示右括號
	placeParentheses(position, &str, 0, n)
	return str
}

func placeParentheses(position []int, str *[]string, i, n int) {
	if i == 2*n { //當所有的括號都放完了
		var s string
		for _, v := range position {
			if v == 0 {
				s += "("
			} else {
				s += ")"
			}
		}
		*str = append(*str, s)
		return
	}
	if isValid(position, i, 0, n) { //放左括號是否合法
		position[i] = 0
		placeParentheses(position, str, i+1, n)
	}
	if isValid(position, i, 1, n) { //放右括號是否合法
		position[i] = 1
		placeParentheses(position, str, i+1, n)
	}
}

func isValid(position []int, cur int, LR int, n int) bool {
	var num_left, num_right int
	for i := 0; i < cur; i++ {
		if position[i] == 0 {
			num_left++
		} else {
			num_right++
		}
	}
	if LR == 0 { //如果當前放入的是左括號
		if num_left >= n {
			return false
		}
	} else { //如果當前放入的是右括號
		if num_left <= num_right {
			return false
		}
	}
	return true
}

/*MergeKLists leetcode23
先排序List,再合併
有幾種方法
法1 把K個list, 變成leetcode21的問題,需要不停的合併, 產生2個list, 需要(k-1)次的合併
法2 沿用法1,但是進行優化, 將K個List, 配對並將同一對中的list合併 重複這一個動作, 就好了
法3 優先佇列, 先把list上的所有值放進優先佇列生成最小堆, 然后最小堆直接poll就可以了
heap里面存k個(length), 一共n個數, while n次, 每次poll時間複雜度logk, 所以總時間複雜度nlogk
*/
func MergeKLists(lists []*ListNode) *ListNode {
	//法1 法2 快速精簡版
	// if lists == nil || len(lists) == 0 {
	// 	return nil
	// }

	// for len(lists) > 1 {
	// 	// pop 2 lists
	// 	l1 := lists[0]
	// 	l2 := lists[1]
	// 	lists = lists[2:]

	// 	merged := MergeTwoLists(l1, l2)
	// 	lists = append(lists, merged)
	// }

	// return lists[0]
	//法2
	// if len(lists) == 0 {
	// 	return nil
	// }
	// begin, end := 0, len(lists)-1
	// for begin < end {
	// 	mid := (begin + end - 1) / 2
	// 	for i := 0; i <= mid; i++ {
	// 		lists[i] = MergeTwoLists(lists[i], lists[end-i])
	// 	}
	// 	end = (begin + end) / 2
	// }
	// return lists[0]

	//法3
	if len(lists) == 0 {
		return nil
	}
	q := new(Queue)
	for _, node := range lists {
		if node != nil {
			heap.Push(q, node)
		}
	}
	res := new(ListNode)
	head := res
	for q.Len() > 0 {
		node := heap.Pop(q).(*ListNode)
		head.Next = node
		head = head.Next
		if node.Next != nil {
			heap.Push(q, node.Next)
		}
	}
	return res.Next
}

/*SwapPairs leetcode24
交換
*/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	node := head
	next := node.Next

	for (node != nil) && (next != nil) {
		node.Val, next.Val = next.Val, node.Val
		node = next.Next
		if (node == nil) || (node.Next == nil) {
			break
		} else {
			next = node.Next
		}
	}

	return head
}
