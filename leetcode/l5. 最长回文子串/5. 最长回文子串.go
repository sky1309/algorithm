package main

/**
https://leetcode.cn/problems/longest-palindromic-substring
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文
子串
。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/

func longestPalindrome(s string) string {
	n := len(s)

	// 结果的起始i和长度，默认情况下分别为0和1
	ansLeft := 0
	ansLen := 1

	for i := 0; i < n; i++ {
		left := i - 1
		right := i + 1
		ln := 1

		// 向左，直到找到和当前i对应的元素不相同的left
		for left >= 0 && s[i] == s[left] {
			left--
			ln++
		}

		// 向右，直到找到和当前i对应的元素不相同的right
		for right < n && s[i] == s[right] {
			right++
			ln++
		}

		// 两边同时相反方向，知道不相同
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
			ln += 2
		}

		// 判断是否要更新结果
		if ln > ansLen {
			ansLeft = left + 1
			ansLen = ln
		}
	}

	return s[ansLeft : ansLeft+ansLen]
}
