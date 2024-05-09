package l139

/*
139. 单词拆分
中等
相关标签
相关企业
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。



示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


提示：

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅由小写英文字母组成
wordDict 中的所有字符串 互不相同
*/

func wordBreak1(s string, wordDict []string) bool {
	hash := make(map[string]struct{})
	for _, w := range wordDict {
		hash[w] = struct{}{}
	}

	memo := make(map[string]int)

	var dfs func(str string) bool
	dfs = func(str string) bool {
		if len(str) == 0 {
			return true
		}

		if memo[str] != 0 {
			return memo[str] == 1
		}

		for w := range hash {
			n := len(w)
			if len(str) < n {
				continue
			}

			if str[:n] != w {
				continue
			}

			if dfs(str[n:]) {
				memo[str] = 1
				return true
			}
		}

		memo[str] = -1
		return false
	}
	return dfs(s)
}

func wordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, w := range wordDict {
		hash[w] = true
	}

	// dp[i] 表示s前i个是否可以拼凑
	// dp[i] = dp[j] && check(s[j:i])

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && hash[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
