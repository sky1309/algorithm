package main

/*
419. 甲板上的战舰
已解答
中等
相关标签
相关企业
给你一个大小为 m x n 的矩阵 board 表示甲板，其中，每个单元格可以是一艘战舰 'X' 或者是一个空位 '.' ，返回在甲板 board 上放置的 战舰 的数量。

战舰 只能水平或者垂直放置在 board 上。换句话说，战舰只能按 1 x k（1 行，k 列）或 k x 1（k 行，1 列）的形状建造，其中 k 可以是任意大小。两艘战舰之间至少有一个水平或垂直的空位分隔 （即没有相邻的战舰）。



示例 1：


输入：board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
输出：2
示例 2：

输入：board = [["."]]
输出：0


提示：

m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] 是 '.' 或 'X'


进阶：你可以实现一次扫描算法，并只使用 O(1) 额外空间，并且不修改 board 的值来解决这个问题吗？
*/

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])

	// mask(x, y) 表示从(x, y) 开始递归遍历，标记已经走过的点
	var mask func(x, y int)
	mask = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}

		if board[x][y] == '.' {
			return
		}

		// 标记为空
		board[x][y] = '.'

		// 四方形递归
		mask(x-1, y)
		mask(x+1, y)
		mask(x, y-1)
		mask(x, y+1)
	}

	ans := 0
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] == 'X' {
				mask(x, y)
				ans++
			}
		}
	}
	return ans
}
