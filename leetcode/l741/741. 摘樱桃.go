package main

import "math"

/**
741. 摘樱桃
困难
给你一个 n x n 的网格 grid ，代表一块樱桃地，每个格子由以下三种数字的一种来表示：

0 表示这个格子是空的，所以你可以穿过它。
1 表示这个格子里装着一个樱桃，你可以摘到樱桃然后穿过它。
-1 表示这个格子里有荆棘，挡着你的路。
请你统计并返回：在遵守下列规则的情况下，能摘到的最多樱桃数：

从位置 (0, 0) 出发，最后到达 (n - 1, n - 1) ，只能向下或向右走，并且只能穿越有效的格子（即只可以穿过值为 0 或者 1 的格子）；
当到达 (n - 1, n - 1) 后，你要继续走，直到返回到 (0, 0) ，只能向上或向左走，并且只能穿越有效的格子；
当你经过一个格子且这个格子包含一个樱桃时，你将摘到樱桃并且这个格子会变成空的（值变为 0 ）；
如果在 (0, 0) 和 (n - 1, n - 1) 之间不存在一条可经过的路径，则无法摘到任何一个樱桃。


示例 1：


输入：grid = [[0,1,-1],[1,0,-1],[1,1,1]]
输出：5
解释：玩家从 (0, 0) 出发：向下、向下、向右、向右移动至 (2, 2) 。
在这一次行程中捡到 4 个樱桃，矩阵变成 [[0,1,-1],[0,0,-1],[0,0,0]] 。
然后，玩家向左、向上、向上、向左返回起点，再捡到 1 个樱桃。
总共捡到 5 个樱桃，这是最大可能值。
示例 2：

输入：grid = [[1,1,-1],[1,-1,1],[-1,1,1]]
输出：0


提示：

n == grid.length
n == grid[i].length
1 <= n <= 50
grid[i][j] 为 -1、0 或 1
grid[0][0] != -1
grid[n - 1][n - 1] != -1
*/

func cherryPickup(grid [][]int) int {
	// A B 两个同时从(0, 0) 开始走t步，dfs(t, j, k) 表示A(t-j, j)  B(t-k, k) 时的最大樱桃树

	// A向下(t-1-j, j) B向下(t-1-k, k)
	// A向下(t-1-j, j) B向右(t-k, k-1)
	// A向右(t-j, j-1) B向下(t-1-k, k)
	// A向右(t-j, j-1) B向右(t-k, k-1)
	n := len(grid)
	memo := make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, 2*n)
			for k := 0; k < 2*n; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	var dfs func(t, j, k int) int
	dfs = func(t, j, k int) int {
		if j < 0 || k < 0 || t < j || t < k || grid[t-j][j] < 0 || grid[t-k][k] < 0 {
			return math.MinInt
		}

		// 原点直接返回0的值
		if t == 0 {
			return grid[0][0]
		}

		if memo[j][k][t] != -1 {
			return memo[j][k][t]
		}

		res := max(dfs(t-1, j, k), dfs(t-1, j, k-1), dfs(t-1, j-1, k), dfs(t-1, j-1, k-1)) + grid[t-j][j]

		// 如果走到了不同的点，则再统计下B点的樱桃
		if j != k {
			res += grid[t-k][k]
		}

		// 记忆化搜索
		memo[j][k][t] = res

		return res
	}

	return max(dfs(2*n-2, n-1, n-1), 0)
}
