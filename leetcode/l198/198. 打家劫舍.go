package l198

/**
https://leetcode.cn/problems/house-robber
198. 打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/

// 动态规划
func rob(nums []int) int {
	// dp[i] 表示偷窃到第i个时的最大金额数
	// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// dfs 选和不选 + cache
func rob1(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}

	// dfs(i) 表示前到达第i个最大的收益数
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		// 选i
		var res int
		res = max(res, dfs(i-2)+nums[i])

		//不选i
		res = max(res, dfs(i-1))

		memo[i] = res
		return res
	}

	return dfs(n - 1)
}
