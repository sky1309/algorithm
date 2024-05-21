package l994

/**
https://leetcode.cn/problems/rotting-oranges/description/
994. 腐烂的橘子
在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。



示例 1：



输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：

输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
示例 3：

输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] 仅为 0、1 或 2

*/

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func orangesRotting(grid [][]int) int {
	q := make([][]int, 0, 100)

	m, n := len(grid), len(grid[0])

	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				// bfs队列的起始元素
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				// 统计最开始新鲜的水果数量
				cnt++
			}
		}
	}

	round := 0
	for {
		qn := len(q)
		if qn == 0 {
			break
		}

		for i := 0; i < qn; i++ {
			x, y := q[i][0], q[i][1]

			// 遍历4个方向
			for _, dir := range dirs {
				newX, newY := x+dir[0], y+dir[1]
				if isValidPos(m, n, newX, newY) && grid[newX][newY] == 1 {
					// 感染这个这个位置的橘子
					grid[newX][newY] = 2
					// 假如到队列中去
					q = append(q, []int{newX, newY})
					// 剩余新鲜的橘子数减少
					cnt--
				}
			}
		}
		q = q[qn:]

		// 只有本次感染了橘子，才计算时间
		if len(q) > 0 {
			round++
		}
	}

	// 说明不能全部都感染
	if cnt != 0 {
		return -1
	}

	return round
}

func isValidPos(m, n int, x, y int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	return true
}
