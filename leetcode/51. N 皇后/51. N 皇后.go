package main

/*
51. N 皇后
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]


提示：

1 <= n <= 9
*/

func solveNQueens(n int) [][]string {

	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}

	var ans [][]string

	var dfs func(x int)
	dfs = func(x int) {
		if x == n {
			temp := make([]string, n)
			for i, row := range path {
				temp[i] = string(row)
			}
			ans = append(ans, temp)
			return
		}

		// 枚举列
		for y := 0; y < n; y++ {
			if !isValid(path, x, y) {
				continue
			}

			path[x][y] = 'Q'
			dfs(x + 1) // 递归下一层
			path[x][y] = '.'
		}
	}

	dfs(0)

	return ans
}

func isValid(path [][]byte, x, y int) bool {
	// (x, y) 位置能否放皇后

	// 上
	for i := x - 1; i >= 0; i-- {
		if path[i][y] == 'Q' {
			return false
		}
	}

	// 左上
	for i, j := x-1, y-1; i >= 0 && j >= 0; {
		if path[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}

	// 右上
	for i, j := x-1, y+1; i >= 0 && j < len(path); {
		if path[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}

	return true
}
