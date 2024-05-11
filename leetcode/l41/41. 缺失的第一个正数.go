package l41

/**
https://leetcode.cn/problems/first-missing-positive/description
41. 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。


示例 1：

输入：nums = [1,2,0]
输出：3
解释：范围 [1,2] 中的数字都在数组中。
示例 2：

输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。
示例 3：

输入：nums = [7,8,9,11,12]
输出：1
解释：最小的正数 1 没有出现。

提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
*/

func firstMissingPositive1(nums []int) int {
	hash := make(map[int]struct{})
	for _, num := range nums {
		hash[num] = struct{}{}
	}

	n := len(nums)
	for i := 1; i <= n; i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}
	return n + 1
}

// 原数组上做的hash
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 把复数都改为n+1
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		// 把num对应的位置的值修改为负数，表示以这个num为index的位置数存在过
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	// 第一个为正数的位置的i就是要求的
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
