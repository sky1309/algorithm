package main

import "math/bits"

/**
1542. 找出最长的超赞子字符串
困难
给你一个字符串 s 。请返回 s 中最长的 超赞子字符串 的长度。

「超赞子字符串」需满足满足下述两个条件：

该字符串是 s 的一个非空子字符串
进行任意次数的字符交换后，该字符串可以变成一个回文字符串


示例 1：

输入：s = "3242415"
输出：5
解释："24241" 是最长的超赞子字符串，交换其中的字符后，可以得到回文 "24142"
示例 2：

输入：s = "12345678"
输出：1
示例 3：

输入：s = "213123"
输出：6
解释："213123" 是最长的超赞子字符串，交换其中的字符后，可以得到回文 "231132"
示例 4：

输入：s = "00"
输出：2


提示：

1 <= s.length <= 10^5
s 仅由数字组成
*/

func longestAwesome_1(s string) int {
	// 前缀和暴力解法，超时！！！
	n := len(s)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] ^ (1 << (s[i] - '0'))
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if bits.OnesCount(uint(sum[i+1]^sum[j])) <= 1 {
				ans = max(ans, i-j+1)
				break
			}
		}
	}
	return ans
}

func longestAwesome(s string) int {
	n := len(s)

	/*
		前缀异或，当i->j是超赞子字符串时
			1. 偶数个数的情况，则    pre[i] ^ pre[j] == 0   pre[i]==pre[j]
			2. 只存在一个位位1，     pre[i] ^ pre[j] = 2**k

	*/

	const D = 10 // 每位元素的情况数0-9
	pos := [1 << D]int{}
	for i := 0; i < len(pos); i++ {
		pos[i] = n
	}

	pos[0] = -1

	ans := 0
	pre := 0
	for i := 0; i < n; i++ {
		// 统计[0,i]前缀出现的次数
		pre = pre ^ (1 << (s[i] - '0'))

		// 偶数的情况，找之前出现过和当前pre相同的索引
		ans = max(ans, i-pos[pre])

		// 奇数的情况，遍历所有的位0-9，枚举所有位为0的情况
		for j := 0; j < D; j++ {
			ans = max(ans, i-pos[pre^(1<<j)])
		}

		// 更新最左边的索引
		if pos[pre] == n {
			pos[pre] = i
		}
	}

	return ans
}
