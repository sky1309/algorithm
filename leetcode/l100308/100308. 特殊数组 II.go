package main

/**
https://leetcode.cn/problems/special-array-ii/
100308. 特殊数组 II
如果数组的每一对相邻元素都是两个奇偶性不同的数字，则该数组被认为是一个 特殊数组 。

周洋哥有一个整数数组 nums 和一个二维整数矩阵 queries，对于 queries[i] = [fromi, toi]，请你帮助周洋哥检查子数组 nums[fromi..toi] 是不是一个 特殊数组 。

返回布尔数组 answer，如果 nums[fromi..toi] 是特殊数组，则 answer[i] 为 true ，否则，answer[i] 为 false 。



示例 1：

输入：nums = [3,4,1,2,6], queries = [[0,4]]

输出：[false]

解释：

子数组是 [3,4,1,2,6]。2 和 6 都是偶数。

示例 2：

输入：nums = [4,3,1,6], queries = [[0,2],[2,3]]

输出：[false,true]

解释：

子数组是 [4,3,1]。3 和 1 都是奇数。因此这个查询的答案是 false。
子数组是 [1,6]。只有一对：(1,6)，且包含了奇偶性不同的数字。因此这个查询的答案是 true。


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 105
1 <= queries.length <= 105
queries[i].length == 2
0 <= queries[i][0] <= queries[i][1] <= nums.length - 1
*/

func isArraySpecial(nums []int, queries [][]int) []bool {
	/*
	   定义一个长度为n-1的数组a，a[i] 表示第i个逗号所在位置的左右是否是不同状态的值，相同为1  不相同为0

	   从from -> to 是否是有效的，可以转换为  a[from] - > a[to-1] 的和为0  转换为前缀和的概念
	*/
	n := len(nums)
	a := make([]int, n)
	for i := 1; i < n; i++ {
		a[i] = a[i-1]

		// 如果当前和前面的状态相同，则前缀和+1
		if nums[i]%2 == nums[i-1]%2 {
			a[i]++
		}
	}

	m := len(queries)
	ans := make([]bool, m)
	for i, q := range queries {
		ans[i] = a[q[1]] == a[q[0]]
	}

	return ans
}
