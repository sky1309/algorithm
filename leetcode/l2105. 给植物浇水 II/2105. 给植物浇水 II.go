package l2105

// 2105. 给植物浇水 II
// https://leetcode.cn/problems/watering-plants-ii/description

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	left := 0
	right := len(plants) - 1

	curA, curB := capacityA, capacityB
	ans := 0
	for left <= right {
		if left == right {
			// 到最后其中一个人不满足则再增加一次
			if max(curA, curB) < plants[left] {
				ans++
			}
			break
		}

		if curA < plants[left] {
			curA = capacityA
			ans++
		}
		curA -= plants[left]

		if curB < plants[right] {
			curB = capacityB
			ans++
		}
		curB -= plants[right]

		left++
		right--
	}
	return ans
}
