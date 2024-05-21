package l240

// 240. 搜索二维矩阵 II
// https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
func searchMatrix(matrix [][]int, target int) bool {
	// 逆时针选择45°

	m, n := len(matrix), len(matrix[0])

	i, j := 0, n-1

	for j >= 0 && i < m {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
