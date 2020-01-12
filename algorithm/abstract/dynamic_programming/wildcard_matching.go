package dynamic_programming

type WildcardMatchingSolver struct {
}

/**
状态转移方程:
s[k]能匹配p[j] = s[k-1]能匹配p[j-1] || s[k-1]能匹配p[j]
s=

*/
func (WildcardMatchingSolver) WildcardMatchingSolve(s string, p string) bool {
	return false
}
