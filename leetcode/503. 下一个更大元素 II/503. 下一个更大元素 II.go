package main

/*
503. 下一个更大元素 II
中等
相关标签
相关企业
给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。

数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。



示例 1:

输入: nums = [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
示例 2:

输入: nums = [1,2,3,4,3]
输出: [2,3,4,-1,4]


提示:

1 <= nums.length <= 104
-109 <= nums[i] <= 109
*/

func nextGreaterElements(nums []int) []int {
	// 循环队列 [1,2,1] -> [1,2,1] + [1,2,1]  就是两个nums相加
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	var st []int // 单调栈，底大顶小

	// 从最右边元素索引开始遍历
	for i := 2*n - 1; i >= 0; i-- {
		idx := i % n // 计算在nums中的当前索引

		// 构建单调栈，移除栈内比当前值小的数据
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[idx] {
			st = st[:len(st)-1]
		}

		if len(st) > 0 && i < n {
			ans[i] = nums[st[len(st)-1]]
		}
		st = append(st, idx)
	}
	return ans
}
