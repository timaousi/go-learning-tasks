package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SingleNumber 找出只出现一次的数字，其他数字都出现两次")
	fmt.Println(SingleNumber([]int{1, 1, 2, 2, 3}))

	fmt.Println("IsPalindrome 判断整数是否是回文数")
	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(-121))
	fmt.Println(IsPalindrome(10))
	fmt.Println(IsPalindrome(12321))

	fmt.Println("IsValid 判断括号字符串是否有效")
	fmt.Println(IsValid("()"))
	fmt.Println(IsValid("()[]{}"))
	fmt.Println(IsValid("(]"))
	fmt.Println(IsValid("([)]"))
	fmt.Println(IsValid("{[]}"))

	fmt.Println("LongestCommonPrefix 查找字符串数组中的最长公共前缀")
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("PlusOne 给表示整数的数组加一")
	fmt.Println(PlusOne([]int{1, 2, 9}))
	fmt.Println(PlusOne([]int{9, 9, 9}))

	fmt.Println("RemoveDuplicates 原地删除排序数组中的重复元素")
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	k := RemoveDuplicates(nums)
	fmt.Println(k)
	fmt.Println(nums[:k])

	fmt.Println("Merge 合并重叠区间")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(Merge(intervals))

	fmt.Println("TwoSum 找出数组中两数之和为目标值的下标")
	nums1 := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(TwoSum(nums1, target))
}

// SingleNumber 找出只出现一次的数字，其他数字都出现两次
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// IsPalindrome 判断整数是否是回文数
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return x == reverted || x == reverted/10
}

// IsValid 判断括号字符串是否有效
func IsValid(s string) bool {
	// 用切片模拟栈
	stack := []rune{}
	// 定义括号对应关系
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		if ch == ')' || ch == '}' || ch == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pairs[ch] != top {
				return false
			}
		} else {
			// 左括号入栈
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

// LongestCommonPrefix 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && (len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// PlusOne 给表示整数的数组加一
func PlusOne(digits []int) []int {
	n := len(digits)

	// 从末尾开始加1
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits // 没进位，直接返回
		}
		// 进位，当前位变0，继续下一位
		digits[i] = 0
	}

	// 全部进位，扩容最高位
	newDigits := make([]int, n+1)
	newDigits[0] = 1
	return newDigits
}

// RemoveDuplicates 原地删除排序数组中的重复元素，返回唯一元素数量
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0 // 慢指针，指向唯一元素的最后一个位置
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

// Merge 合并重叠区间，返回合并后的区间数组
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] <= last[1] {
			// 重叠，更新结束点
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			// 不重叠，直接追加
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

// TwoSum 找出数组中两数之和为目标值的下标
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int) // key: 数字 value: 下标

	for i, num := range nums {
		complement := target - num
		if idx, found := m[complement]; found {
			return []int{idx, i}
		}
		m[num] = i
	}

	return nil
}
