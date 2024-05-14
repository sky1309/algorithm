package main

/*
70. 爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶


提示：

1 <= n <= 45
*/

func climbStairs1(n int) int {
	/*
		动态规划
		假设dp[i] 表示i个台阶的方案数
		dp[i] = dp[i-1] + dp[i-2]  (i>=2)
	*/
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs(n int) int {
	/*
		动态规划
		假设dp[i] 表示i个台阶的方案数
		dp[i] = dp[i-1] + dp[i-2]  (i>=2)
	*/
	if n <= 2 {
		return n
	}

	pre1 := 1
	pre2 := 2
	for i := 3; i <= n; i++ {
		s := pre1 + pre2
		pre1 = pre2
		pre2 = s
	}

	return pre2
}
