package main

/*
739. 每日温度
已解答
中等
相关标签
相关企业
提示
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。


示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]


提示：

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)

	// st 是一个从大到小的栈，用slice模拟
	var st []int

	// 从右向左遍历
	for i := n - 1; i >= 0; i-- {
		// 把小于当前温度的出栈
		for len(st) > 0 && temperatures[st[len(st)-1]] <= temperatures[i] {
			st = st[:len(st)-1]
		}

		// 如果栈不为空，说明右边存在比当前大的温度的时间
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}

		// 假如到栈中取
		st = append(st, i)
	}

	return ans
}
