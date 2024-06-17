package main

import "slices"

/*
522. 最长特殊序列 II
给定字符串列表 strs ，返回其中 最长的特殊序列 的长度。如果最长特殊序列不存在，返回 -1 。

特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。

 s 的 子序列可以通过删去字符串 s 中的某些字符实现。

例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。"aebdc"的子序列还包括"aebdc"、 "aeb" 和 "" (空字符串)。


示例 1：

输入: strs = ["aba","cdc","eae"]
输出: 3
示例 2:

输入: strs = ["aaa","aaa","aa"]
输出: -1


提示:

2 <= strs.length <= 50
1 <= strs[i].length <= 10
strs[i] 只包含小写英文字母
*/

func findLUSlength(strs []string) int {
	slices.SortFunc(strs, func(s1, s2 string) int {
		return len(s1) - len(s2)
	})

	ans := -1
	for i := 0; i < len(strs); i++ {
		isValid := true
		for j := 0; j < len(strs); j++ {
			if i == j {
				continue
			}
			if isSubStr(strs[j], strs[i]) {
				isValid = false
				break
			}
		}

		if isValid {
			ans = max(ans, len(strs[i]))
		}
	}

	return ans
}

func isSubStr(s1 string, s2 string) bool {
	// s2 是否是s1 的子序列
	if len(s1) < len(s2) {
		return false
	}

	l, r := 0, 0

	lens1 := len(s1)
	for l < lens1 {
		if s1[l] == s2[r] {
			r++
			if r == len(s2) {
				break
			}
		}
		l++
	}

	return r == len(s2)
}
