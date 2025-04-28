package main

import "fmt"

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	var m1 = make(map[int]int)
	for _, v := range nums {
		count, ok := m1[v]
		if ok {
			count += 1
		} else {
			count = 1
		}
		m1[v] = count
	}
	//遍历map获取统计值为1的数值
	for key, val := range m1 {
		if val == 1 {
			fmt.Println(key)
			return key
		}
	}
	return -1
}

// 198. 打家劫舍(https://leetcode.cn/problems/house-robber/)：
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，
// 一夜之内能够偷窃到的最高金额。这道题可以使用动态规划的思想，通过 `for` 循环遍历数组，利用 `if` 条件判断来决定是否选择当前房屋进行抢劫，
// 状态转移方程为 `dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])`。
func rob(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	if size == 2 {
		return max(nums[0], nums[1])
	}
	var dp = make([]int, size)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[size-1]
}

// 定义一个链表结构体
type LinkNode struct {
	Val  int
	Next *LinkNode
}

// 21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，
// 直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
func mergeTwoLists(list1 *LinkNode, list2 *LinkNode) *LinkNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// 46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
// 可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，
// 使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(int)
	backtrack = func(first int) {
		if first == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
		}
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return res
}

// 344. 反转字符串：编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。不要给另外的数组分配额外的空间，
// 你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
// 可以使用 for 循环和两个指针，一个指向字符串的开头，一个指向字符串的结尾，然后交换两个指针所指向的字符，直到两个指针相遇。
func resveString(bytes []byte) []byte {
	from := 0
	to := len(bytes) - 1
	for from < to {
		temp := bytes[from]
		bytes[from] = bytes[to]
		bytes[to] = temp
		from++
		to--
	}
	return bytes
}

// 69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
// 可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件的平方根或者确定不存在精确的平方根。
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {
	// var nums = []int{2, 2, 1, 3, 3, 4, 4, 5}
	// singleNumber(nums)
	// var nums2 = []int{2, 7, 9, 3, 1}
	// fmt.Println(rob(nums2))
	// var nums2 = []int{2, 7, 9}
	// result := permute(nums2)
	// fmt.Println(result)
	fmt.Println(resveString([]byte{'a', 'b', 'c'}))
}
