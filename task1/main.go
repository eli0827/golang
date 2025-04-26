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

func main() {
	var nums = []int{2, 2, 1, 3, 3, 4, 4, 5}
	singleNumber(nums)
}
