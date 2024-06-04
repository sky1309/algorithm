package main

/*
3067. 在带权树网络中统计可连接服务器对数目
给你一棵无根带权树，树中总共有 n 个节点，分别表示 n 个服务器，服务器从 0 到 n - 1 编号。同时给你一个数组 edges ，其中 edges[i] = [ai, bi, weighti] 表示节点 ai 和 bi 之间有一条双向边，边的权值为 weighti 。再给你一个整数 signalSpeed 。

如果两个服务器 a ，b 和 c 满足以下条件，那么我们称服务器 a 和 b 是通过服务器 c 可连接的 ：

a < b ，a != c 且 b != c 。
从 c 到 a 的距离是可以被 signalSpeed 整除的。
从 c 到 b 的距离是可以被 signalSpeed 整除的。
从 c 到 b 的路径与从 c 到 a 的路径没有任何公共边。
请你返回一个长度为 n 的整数数组 count ，其中 count[i] 表示通过服务器 i 可连接 的服务器对的 数目 。

示例 1：
输入：edges = [[0,1,1],[1,2,5],[2,3,13],[3,4,9],[4,5,2]], signalSpeed = 1
输出：[0,4,6,6,4,0]
解释：由于 signalSpeed 等于 1 ，count[c] 等于所有从 c 开始且没有公共边的路径对数目。
在输入图中，count[c] 等于服务器 c 左边服务器数目乘以右边服务器数目。

示例 2：
输入：edges = [[0,6,3],[6,5,3],[0,3,1],[3,2,7],[3,1,6],[3,4,2]], signalSpeed = 3
输出：[2,0,0,0,0,0,2]
解释：通过服务器 0 ，有 2 个可连接服务器对(4, 5) 和 (4, 6) 。
通过服务器 6 ，有 2 个可连接服务器对 (4, 5) 和 (0, 5) 。
所有服务器对都必须通过服务器 0 或 6 才可连接，所以其他服务器对应的可连接服务器对数目都为 0 。


提示：

2 <= n <= 1000
edges.length == n - 1
edges[i].length == 3
0 <= ai, bi < n
edges[i] = [ai, bi, weighti]
1 <= weighti <= 106
1 <= signalSpeed <= 106
输入保证 edges 构成一棵合法的树。
*/

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	// 节点数为长度+1
	n := len(edges) + 1
	// 构建图
	graph := make([][][]int, n)
	for _, edge := range edges {
		x, y, w := edge[0], edge[1], edge[2]
		graph[x] = append(graph[x], []int{y, w})
		graph[y] = append(graph[y], []int{x, w})
	}

	// dfs(i, root, sum) 表示以root节点开始，到子节点i 的边长能被 signalSpeed 整除的节点个数
	var dfs func(i int, root int, sum int) int
	dfs = func(i int, root int, sum int) int {
		res := 0
		if sum%signalSpeed == 0 {
			res++
		}

		// 递归遍历i的子节点
		for _, arr := range graph[i] {
			x, w := arr[0], arr[1]
			// 由于i和root是连接的关系，这里跳过root节点相同的情况
			if x != root {
				res += dfs(x, i, sum+w)
			}
		}

		return res
	}

	ans := make([]int, n)
	for i, childrens := range graph {
		pre := 0
		for _, ch := range childrens {
			x, w := ch[0], ch[1]
			cnt := dfs(x, i, w)
			ans[i] += pre * cnt
			pre += cnt
		}

	}

	return ans
}
