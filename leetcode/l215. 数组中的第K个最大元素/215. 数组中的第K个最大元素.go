package l215

import "math/rand"

/*
https://leetcode.cn/problems/kth-largest-element-in-an-array/
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4


提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func findKthLargest(nums []int, k int) int {

	// 快速选择
	var quickSelect func(arr []int, i int) int
	quickSelect = func(arr []int, i int) int {
		n := len(arr)
		// 随机选择一个作为标杆
		flag := arr[rand.Intn(n)]

		// 记录小的和大的数的情况
		var small []int
		var big []int

		for _, num := range arr {
			if num < flag {
				small = append(small, num)
			} else if num > flag {
				big = append(big, num)
			}
		}

		// 如果比标杆大数的 数目大于k 则结果肯定在big里面
		if i <= len(big) {
			return quickSelect(big, i)
		}

		// 如果big+small的数目比k小，则在small里
		if len(arr)-len(small) < i {
			return quickSelect(small, i-len(arr)+len(small))
		}

		// 在相等的里面
		return flag
	}

	return quickSelect(nums, k)
}
