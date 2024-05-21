package main

/**
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长
子串
 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
func lengthOfLongestSubstring(s string) int {
	// 滑动窗口+hash表
	hash := make(map[byte]struct{})

	ans := 0
	left := 0
	n := len(s)

	for right := 0; right < n; right++ {
		b := s[right]
		for {
			if _, ok := hash[b]; !ok {
				break
			}

			// 移除左边的节点数据， 直到不存在当前元素为止
			delete(hash, s[left])
			left++
		}
		// 存储已经出现的数据
		hash[b] = struct{}{}

		// 计算结果
		ans = max(ans, right-left+1)
	}
	return ans
}
