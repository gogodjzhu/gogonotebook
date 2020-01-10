package dynamic_programming

import "strings"

type WildcardMatchingSolver struct {
}

/**
状态转移方程:
s[k]能匹配p[j] = s[k-1]能匹配p[j-1] || s[k-1]能匹配p[j]
s=

*/
func (WildcardMatchingSolver) WildcardMatching(s string, p string) bool {
	matchIdx := make([][]int, len(p))
	for i := 0; i < len(matchIdx); i++ {
		matchIdx[i] = make([]int, 2)
	}
	for i := 0; i < len(p); i++ {
		switch p[i] {
		case '*':
			if i == 0 {
				matchIdx[i] = []int{0, len(s)}
			} else {
				matchIdx[i] = []int{matchIdx[i-1][1], len(s)}
			}
		case '?':
			if i == 0 {
				matchIdx[i] = []int{0, 1}
			} else {
				matchIdx[i] = []int{matchIdx[i-1][1], matchIdx[i-1][1] + 1}
			}
			if matchIdx[i][1]-matchIdx[i][0] <= 0 {
				return false
			}
		default:
			if i == 0 {
				if p[0] != s[0] {
					return false
				} else {
					matchIdx[i] = []int{0, 1}
				}
			} else {
				if p[i-1] == '*' {
					matchIdx[i] = []int{strings.LastIndexByte(s, p[i]), strings.LastIndexByte(s, p[i]) + 1}
				} else {
					last := matchIdx[i-1][1]
					matchIdx[i] = []int{matchIdx[i-1][1],
						matchIdx[i-1][1] + strings.IndexByte(s[last:], p[i]) + 1}
				}
			}
			if matchIdx[i][1]-matchIdx[i][0] <= 0 {
				return false
			}
		}
	}
	if len(p) > 0 {
		return matchIdx[len(p)-1][1] == len(s)
	} else {
		return len(s) == 0
	}
}
