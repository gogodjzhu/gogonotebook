package dynamic_programming

type LongestPalindromicSubstringSolver struct {
}

func (LongestPalindromicSubstringSolver) LongestPalindromicSubstringSolve(s string) string {
	// 在每个字符中间添加#分割，使奇偶长度的形式一致，同时返回一个对应的标识数组，用于还原
	sArr, replaceArr := makeOdd(s)
	length := len(sArr)
	// index, length 数组的两个维度分别保存[下标位置]，和构成字符串的[左右长度]
	// 数组的值表示组成的字符串是否满足回文条件
	validArr := make([][]bool, length)
	for i := 0; i < length; i++ {
		validArr[i] = make([]bool, length)
	}
	// 初始化单个字符永远满足回文条件
	for i := 0; i < length; i++ {
		validArr[i][0] = true
	}
	maxI, maxL := 0, 0
	// 从左右长度为1开始遍历(因为长度为0的已经确认为true)
	// 只需要遍历一半的长度即可，因为左右的长度不可能过半
	for l := 1; l <= len(sArr)/2; l++ {
		hasNew := false
		for i := l; i < length-l; i++ {
			// 遍历每个字符，如果左右l长度位置相同，则判为满足回文，可以在下个循环中判断l+1
			if validArr[i][l-1] && (sArr[i-l] == sArr[i+l]) {
				maxI = i
				validArr[i][l] = true
				hasNew = true
			}
		}
		if !hasNew {
			break
		}
		maxL = l
	}
	// 替换#占位符
	var retArr []byte
	for i := maxI - maxL; i < maxI+maxL; i++ {
		if !replaceArr[i] {
			retArr = append(retArr, sArr[i])
		}
	}
	return string(retArr)
}

func makeOdd(s string) ([]byte, []bool) {
	oldArr := []byte(s)
	replaceArr := []bool{true}
	sArr := []byte{'#'}
	for i := 0; i < len(oldArr); i++ {
		sArr = append(sArr, oldArr[i])
		replaceArr = append(replaceArr, false)
		sArr = append(sArr, '#')
		replaceArr = append(replaceArr, true)
	}
	return sArr, replaceArr
}
