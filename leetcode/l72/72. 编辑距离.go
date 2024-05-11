package l72

/**
https://leetcode.cn/problems/edit-distance/description
72. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/
func minDistance(word1 string, word2 string) int {

	// dfs(i, j) 表示 前i,j个最少的操作数
	// dfs(i, j) = min(dfs(i-1,j)删除, dfs(i,j-1)插入,dfs(i-1,j-1)替换)+1
	m, n := len(word1), len((word2))
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		var res int
		if word1[i] == word2[j] {
			res = dfs(i-1, j-1)
			memo[i][j] = res
			return res
		}

		// dfs(i-1, j) 删除
		// dfs(i, j-1) 插入
		// dfs(i-1, j-1) 替换
		res = min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
		memo[i][j] = res
		return res
	}

	return dfs(m-1, n-1)
}
