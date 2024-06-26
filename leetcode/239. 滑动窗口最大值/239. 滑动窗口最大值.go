package main

/*
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length

*/

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	var q []int // 单调队列，存放 nums 从大到小的索引i
	for i, x := range nums {
		// 出，从q的尾部，移除所有比x小的数
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}

		// 入队列
		q = append(q, i)

		// 如果 (当前索引-头部最大值的索引) >= k 则表示当前头部的最大值已经不在窗口内了，则移除掉
		if i-q[0] >= k {
			q = q[1:]
		}

		// 开始记录答案
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}
