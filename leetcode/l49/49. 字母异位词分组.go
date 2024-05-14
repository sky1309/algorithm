package l49

import "sort"

/**
https://leetcode.cn/problems/group-anagrams
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]


提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母*/

func groupAnagrams(strs []string) [][]string {
	// 使用hash来做
	hash := make(map[string][]string)
	for _, s := range strs {
		sArr := []byte(s)
		sort.Slice(sArr, func(i, j int) bool {
			return sArr[i] < sArr[j]
		})

		sortS := string(sArr)
		hash[sortS] = append(hash[sortS], s)
	}

	ans := make([][]string, 0, len(hash))
	for _, arr := range hash {
		ans = append(ans, arr)
	}
	return ans
}
