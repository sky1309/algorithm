package main

import "sort"

/**
15. 三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {

	// 从小到达排序
	sort.Ints(nums)

	n := len(nums)
	var ans [][]int

	// 枚举第一个元素，到n-2结束
	for i := 0; i < n-2; i++ {

		// 当前值比0大了，那么 nums[i] + nums[j] + nums[k] > 0 是一定成立的，后面不可能在出现和为0的情况了
		if nums[i] > 0 {
			break
		}

		// 当前和前面的值相同，则是相同的答案，跳过
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s < 0 {
				left++
			} else if s > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right+1] == nums[right] {
					right--
				}
			}
		}
	}

	return ans
}
