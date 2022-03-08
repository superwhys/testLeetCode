package main

import (
	"gitlab.momoso.com/mms2/utils/lg"
	"sort"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		lg.Info(carry)
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		lg.Info(carry)
		ans = strconv.Itoa(carry%2) + ans
		lg.Infof("ans: %v", ans)
		carry /= 2
		lg.Info(carry)
		lg.Info("==========")
	}

	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func sortedSquares(nums []int) []int {
	var newNums []int
	for _, num := range nums {
		newNums = append(newNums, num*num)
	}
	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i] < newNums[j]
	})
	return newNums
}

func sortedSquares2(nums []int) []int {
	var res []int
	numLen := len(nums)
	neg := -1
	for i := 0; i < numLen && nums[i] < 0; i++ {
		neg = i
	}

	for i, j := neg, neg+1; i >= 0 || j < numLen; {
		if i < 0 {
			// all item is > 0
			res = append(res, nums[j]*nums[j])
			j++
		} else if j == numLen {
			// all item is < 0
			res = append(res, nums[i]*nums[i])
			i++
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			res = append(res, nums[i]*nums[i])
		} else {
			res = append(res, nums[j]*nums[j])
		}
	}
	return res
}

func rotate(nums []int, k int) {
	numLen := len(nums)
	start := numLen - k
	end := numLen
	for i := 0; i < start; i++ {
		nums = append(nums, nums[i])
		end++
	}
	nums = append([]int{}, nums[start:end]...)
	lg.Info(nums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func moveZeroes(nums []int) {
	left, right, numsLen := 0, 0, len(nums)
	for right < numsLen {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l <= r {
			mid := l + (r-l)/2
			if numbers[i]+numbers[mid] == target {
				return []int{i + 1, mid + 1}
			} else if numbers[i]+numbers[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	wordSplit := strings.Split(s, " ")
	for idx, item := range wordSplit {
		itemSplit := []byte(item)
		left, right := 0, len(itemSplit)-1
		for left < right {
			itemSplit[left], itemSplit[right] = itemSplit[right], itemSplit[left]
			left++
			right--
		}
		wordSplit[idx] = string(itemSplit)
	}
	return strings.Join(wordSplit, " ")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	itemList := []*ListNode{head}
	item := head.Next
	for item != nil {
		itemList = append(itemList, item)
		item = item.Next
	}
	lstLen := len(itemList)
	return itemList[lstLen / 2]
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 1 && head.Next == nil {
		return nil
	}
	itemList := []*ListNode{head}
	item := head.Next
	for item != nil {
		itemList = append(itemList, item)
		item = item.Next
	}
	lstLen := len(itemList)
	if n == 1 {
		itemList[lstLen-2].Next = nil
	} else if n == lstLen{
		return itemList[1]
	} else {
		needDel := lstLen - n
		itemList[needDel-1].Next = itemList[needDel+1]
	}
	return itemList[0]
}

// [1, 2] 2
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	node := &ListNode{0, head}
	first, second := head, node
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ;first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return node.Next
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	strLst := []byte(s)
	strLen := len(strLst)
	maxLen := 0
	if strLen == 1 {
		return 1
	}

	rk := -1
	hashMap := make(map[byte]bool)
	for i := 0; i < strLen; i++ {
		if i != 0 {
			delete(hashMap, strLst[i-1])
		}
		for rk+1 < strLen && !hashMap[strLst[rk+1]] {
			hashMap[strLst[rk+1]] = true
			lg.Info(rk, i)
			rk++
		}
		maxLen = Max(rk - i + 1, maxLen)
		//lg.Info(maxLen)
	}
	return maxLen
}

func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}

	s1ItemCount := [26]int{}
	s2SubItemCount := [26]int{}
	for idx, item := range s1 {
		s1ItemCount[item - 'a']++
		s2SubItemCount[s2[idx] - 'a']++
	}

	if s1ItemCount == s2SubItemCount {
		return true
	}

	right := s1Len
	for right < s2Len {
		s2SubItemCount[s2[right - s1Len] - 'a']--
		s2SubItemCount[s2[right] - 'a']++
		if s1ItemCount == s2SubItemCount {
			return true
		}
		right++
	}
	return false
}

func main() {
	s1 := "adc"
	s2 := "dcda"
	lg.Info(checkInclusion(s1, s2))
}
