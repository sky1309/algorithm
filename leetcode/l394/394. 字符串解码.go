package l394

import "strings"

// 394. 字符串解码
// https://leetcode.cn/problems/decode-string/description

func decodeString(s string) string {

	multiStack := make([]int, 0)
	resStack := make([]string, 0)

	res := strings.Builder{}
	multi := 0

	n := len(s)
	for i := 0; i < n; i++ {
		w := s[i]
		if w == '[' {
			multiStack = append(multiStack, multi)
			resStack = append(resStack, res.String())
			multi = 0
			res = strings.Builder{}
		} else if w == ']' {

			// pop
			curMulti := multiStack[len(multiStack)-1]
			multiStack = multiStack[:len(multiStack)-1]

			tmp := strings.Repeat(res.String(), curMulti)
			if len(resStack) > 0 {
				tmp = resStack[len(resStack)-1] + tmp
				resStack = resStack[:len(resStack)-1]
			}

			res = strings.Builder{}
			res.WriteString(tmp)

		} else if w >= '0' && w <= '9' {
			multi = multi*10 + int(w-'0')
		} else {
			res.WriteByte(w)
		}
	}
	return res.String()
}
