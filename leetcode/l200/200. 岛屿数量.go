package l200

/**
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
				setFlag(grid, i, j)
			}
		}
	}

	return ans
}

func setFlag(grid [][]int, x, y int) {
	if !inArea(grid, x, y) {
		return
	}
	setFlag(grid, x-1, y)
	setFlag(grid, x+1, y)
	setFlag(grid, x, y-1)
	setFlag(grid, x, y+1)
}

func inArea(grid [][]int, x, y int) bool {
	return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0])
}
