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
	return -12
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

func main() {
	var nums = []int{2, 2, 1, 3, 3, 4, 4, 5}
	singleNumber(nums)
	var nums2 = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums2))
}
