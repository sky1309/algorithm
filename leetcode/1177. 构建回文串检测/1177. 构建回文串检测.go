package main

import (
	"math/bits"
)

/**
https://leetcode.cn/problems/can-make-palindrome-from-substring/
1177. 构建回文串检测
给你一个字符串 s，请你对 s 的子串进行检测。

每次检测，待检子串都可以表示为 queries[i] = [left, right, k]。我们可以 重新排列 子串 s[left], ..., s[right]，并从中选择 最多 k 项替换成任何小写英文字母。

如果在上述检测过程中，子串可以变成回文形式的字符串，那么检测结果为 true，否则结果为 false。

返回答案数组 answer[]，其中 answer[i] 是第 i 个待检子串 queries[i] 的检测结果。

注意：在替换时，子串中的每个字母都必须作为 独立的 项进行计数，也就是说，如果 s[left..right] = "aaa" 且 k = 2，我们只能替换其中的两个字母。（另外，任何检测都不会修改原始字符串 s，可以认为每次检测都是独立的）

示例：

输入：s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
输出：[true,false,false,true,true]
解释：
queries[0] : 子串 = "d"，回文。
queries[1] : 子串 = "bc"，不是回文。
queries[2] : 子串 = "abcd"，只替换 1 个字符是变不成回文串的。
queries[3] : 子串 = "abcd"，可以变成回文的 "abba"。 也可以变成 "baab"，先重新排序变成 "bacd"，然后把 "cd" 替换为 "ab"。
queries[4] : 子串 = "abcda"，可以变成回文的 "abcba"。


提示：

1 <= s.length, queries.length <= 10^5
0 <= queries[i][0] <= queries[i][1] < s.length
0 <= queries[i][2] <= s.length
s 中只有小写英文字母
*/

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	// 前缀字符出现的次数
	prefix := make([][26]int, n+1)

	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		prefix[i+1][s[i]-'a']++
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		left, right, k := q[0], q[1], q[2]

		// 统计left-right种，所有出现奇数次的字母的个数
		cnt := 0
		for j := 0; j < 26; j++ {
			cnt += (prefix[right+1][j] - prefix[left][j]) % 2
		}

		ans[i] = cnt/2 <= k
	}

	return ans
}

func canMakePaliQueries_1(s string, queries [][]int) []bool {
	// 前缀异或和
	n := len(s)
	// 前缀字符出现的次数, 用一个uint32的各个位的表示26个字母出现的奇偶情况
	prefix := make([]uint32, n+1)
	for i := 0; i < n; i++ {
		// 偶数次位0  奇数次位1
		prefix[i+1] = prefix[i] ^ (1 << (s[i] - 'a')) // 00->0  01->1 10->1
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		left, right, k := q[0], q[1], q[2]

		// 计算2个前缀亦或后为1的bit的数量
		cnt := bits.OnesCount(uint(prefix[right+1] ^ prefix[left]))
		ans[i] = int(cnt/2) <= k
	}
	return ans
}
