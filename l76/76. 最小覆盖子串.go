package l76

// 76. 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring
/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/

func check(cnt1 []int, cnt2 []int) bool {
	for i, c := range cnt2 {
		if cnt1[i] < c {
			return false
		}
	}
	return true
}

func minWindow1(s string, t string) string {
	ns := len(s)
	nt := len(t)

	cntT := [128]int{}
	for i := 0; i < nt; i++ {
		cntT[t[i]]++
	}

	ansLeft, ansRight := -1, ns
	left := 0

	cntS := [128]int{}
	for right := 0; right < ns; right++ {
		cntS[s[right]]++

		for check(cntS[:], cntT[:]) {

			if right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}
			cntS[s[left]]--
			left++
		}
	}

	if ansLeft < 0 {
		return ""
	}

	return s[ansLeft : ansRight+1]
}

func minWindow(s string, t string) string {
	ns := len(s)
	nt := len(t)

	less := 0 // 当前不达标的字符数
	cntS := [128]int{}
	cntT := [128]int{}
	for i := 0; i < nt; i++ {
		if cntT[t[i]] == 0 {
			less++ // 次数为0时记录一次
		}
		cntT[t[i]]++
	}

	ansLeft, ansRight := -1, ns
	left := 0

	for right := 0; right < ns; right++ {
		cntS[s[right]]++
		if cntS[s[right]] == cntT[s[right]] {
			less--
		}

		// 当全部覆盖的时候
		for less == 0 {
			if right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}

			if cntS[s[left]] == cntT[s[left]] {
				less++
			}

			cntS[s[left]]--
			left++
		}
	}

	if ansLeft < 0 {
		return ""
	}

	return s[ansLeft : ansRight+1]
}
