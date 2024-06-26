package main

/*
2741. 特别的排列
给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums 的一个排列满足以下条件，我们称它是一个特别的排列：

对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i] == 0 。
请你返回特别排列的总数目，由于答案可能很大，请将它对 109 + 7 取余 后返回。


示例 1：

输入：nums = [2,3,6]
输出：2
解释：[3,6,2] 和 [2,6,3] 是 nums 两个特别的排列。
示例 2：

输入：nums = [1,4,3]
输出：2
解释：[3,1,4] 和 [4,1,3] 是 nums 两个特别的排列。


提示：

2 <= nums.length <= 14
1 <= nums[i] <= 109
*/
func specialPerm(nums []int) int {
	n := len(nums)

	// 记忆化搜索
	u := 1<<n - 1
	memo := make([][]int, u)
	for i := 0; i < u; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	// 位运算，用位的状态来表示是否已经选过了
	var dfs func(b int, i int) int
	dfs = func(b int, i int) int {
		if b == 0 {
			return 1
		}

		if memo[b][i] != -1 {
			return memo[b][i]
		}

		res := 0
		for j := 0; j < n; j++ {
			// 枚举所有没选过的情况，    (b >> j & 1) 表示第j位是否位1，如果是1表示没有选择过，可以选
			if b>>j&1 == 1 && (nums[j]%nums[i] == 0 || nums[i]%nums[j] == 0) {
				res += dfs(b^(1<<j), j)
			}
		}

		memo[b][i] = res
		return res
	}

	ans := 0
	for i := range nums {
		ans += dfs(u^(1<<i), i)
	}
	return ans % 1_000_000_007
}
