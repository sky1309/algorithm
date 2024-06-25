package main

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

var mp = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var st []byte

	n := len(s)
	for i := 0; i < n; i++ {
		if len(st) > 0 && st[len(st)-1] == mp[s[i]] {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}

	return len(st) == 0
}
