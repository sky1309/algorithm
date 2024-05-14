package main

/**
https://leetcode.cn/problems/longest-consecutive-sequence
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109``

*/

func longestConsecutive(nums []int) int {
	// hash 存储所有出现过的数据
	hash := make(map[int]struct{})
	for _, num := range nums {
		hash[num] = struct{}{}
	}

	ans := 0
	// 直接遍历去重后的集合的key
	for num := range hash {
		// 这个可以作为起始节点，num-1之前没有出现过
		if _, ok := hash[num-1]; ok {
			continue
		}

		// 从当前节点向后，直到没有出现数据的为止
		k := 1
		for i := num + 1; ; i++ {
			if _, ok := hash[i]; !ok {
				break
			}
			k++
		}
		ans = max(ans, k)
	}

	return ans
}
