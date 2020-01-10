package dynamic_programming

type DecodeWaysSolver struct {
}

/**
状态转移方程:
直到第i个数字的解码方式 = 直到第i-1个数字的解码方式(解码为一位数) +
					  直到第i-2个数字的解码方式(解码为二位数)
难点在边界条件的处理
*/
func (DecodeWaysSolver) DecodeWaysSolve(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) <= 1 {
		return 1
	}
	sArr := []byte(s)
	waysArr := make([]int, len(sArr))
	waysArr[0] = 1
	for i := 1; i < len(sArr); i++ {
		switch {
		case int(sArr[i]-'0') == 0:
			if int(sArr[i-1]-'0') == 1 || int(sArr[i-1]-'0') == 2 {
				// 只有合法的20/10能正确解析
				if i-2 < 0 {
					waysArr[i] = 1
				} else {
					waysArr[i] = waysArr[i-2]
				}
			} else {
				return 0
			}
		case int(sArr[i]-'0')+int(sArr[i-1]-'0')*10 <= 26 && sArr[i-1] != '0':
			if sArr[i] == '0' {
				if i-2 < 0 {
					waysArr[i] = 1
				} else {
					waysArr[i] = waysArr[i-2]
				}
			} else {
				if i-2 < 0 {
					waysArr[i] = 2
				} else {
					waysArr[i] = waysArr[i-1] + waysArr[i-2]
				}
			}
		default: // greater than 20+
			waysArr[i] = waysArr[i-1]
		}
	}
	return waysArr[len(waysArr)-1]
}
