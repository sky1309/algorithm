package main

import "sort"

/*
881. 救生艇
给定数组 people 。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。

每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。

返回 承载所有人所需的最小船数 。



示例 1：

输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)
示例 2：

输入：people = [3,2,2,1], limit = 3
输出：3
解释：3 艘船分别载 (1, 2), (2) 和 (3)
示例 3：

输入：people = [3,5,3,4], limit = 5
输出：4
解释：4 艘船分别载 (3), (3), (4), (5)


提示：

1 <= people.length <= 5 * 104
1 <= people[i] <= limit <= 3 * 104
*/

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)-1
	ans := 0
	for l <= r {
		sm := people[l] + people[r]
		if sm <= limit {
			// 如果最大的和最小的能一起过，先匹配上
			l++
			r--
		} else {
			// 如果最大的和最小的不能一起，那么先让最大的过，最小的等待下一轮的匹配
			r--
		}
		ans++
	}
	return ans
}
