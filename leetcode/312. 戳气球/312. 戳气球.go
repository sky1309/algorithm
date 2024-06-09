package main

/*
312. 戳气球
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

求所能获得硬币的最大数量。



示例 1：
输入：nums = [3,1,5,8]
输出：167
解释：
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
示例 2：

输入：nums = [1,5]
输出：10


提示：

n == nums.length
1 <= n <= 300
0 <= nums[i] <= 100
*/

func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	val[0] = 1
	val[n+1] = 1
	copy(val[1:], nums)

	// 记忆化搜索
	memo := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		memo[i] = make([]int, n+2)
		for j := 0; j < n+2; j++ {
			memo[i][j] = -1
		}
	}

	// dfs(i, j) 表示[i, j] 中间插入一个元素可以达到最大结果的值
	// 枚举所有一个插入的位置mid    dfs(i, j) = val[i] * val[j] * val[mid] + dfs(i, mid) + dfs(mid, j)  for mid in range(i+1, j)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := 0
		// 枚举所有的第一次插入的位置
		for x := i + 1; x < j; x++ {
			res = max(res, val[i]*val[j]*val[x]+dfs(i, x)+dfs(x, j))
		}

		memo[i][j] = res
		return res
	}

	return dfs(0, n+1)
}
