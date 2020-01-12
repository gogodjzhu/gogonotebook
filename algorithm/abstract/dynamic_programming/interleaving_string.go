package dynamic_programming

type InterleavingStringSolver struct {
}

/**
状态转移方程:
s3[:k]字符串可以用s1[:i]和s[:j]交错构成 =	(s3[:k-1]可以由s1[:i-1]与s2[:j]交错组成 && s3[k]==s1[i]) ||
										(s3[:k-1]可以由s2[:j-1]与s1[:i]交错组成 && s3[k]==s2[j])
其中k=i+j

故，将s1与s2按字符构成数组[][]bool:
      ""  s1[0] s1[1] s1[2] ... s1[m]
""    T    T/F   T/F   T/F       T/F
s2[0] T/F  T/F   T/F   T/F       T/F
s2[1] T/F  T/F   T/F   T/F       T/F
s2[2] T/F  T/F   T/F   T/F       T/F
...
s2[n] T/F  T/F   T/F   T/F       T/F

最后返回s1[m][n]==T

*/
func (InterleavingStringSolver) InterleavingStringSolve(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	crossArr := make([][]bool, len(s2)+1)
	for i := 0; i < len(crossArr); i++ {
		crossArr[i] = make([]bool, len(s1)+1)
	}

	s1Arr := []byte(s1)
	s2Arr := []byte(s2)
	s3Arr := []byte(s3)
	for j := 0; j < len(crossArr); j++ { // j是s2的下标
		for i := 0; i < len(crossArr[j]); i++ { // i是s1的下标
			switch {
			case j == 0 && i == 0:
				crossArr[j][i] = true
			case j == 0:
				crossArr[j][i] = crossArr[j][i-1] && s1Arr[i-1] == s3Arr[i-1]
			case i == 0:
				crossArr[j][i] = crossArr[j-1][i] && s2Arr[j-1] == s3Arr[j-1]
			default:
				crossArr[j][i] = (crossArr[j-1][i] && s2Arr[j-1] == s3Arr[j+i-1]) ||
					(crossArr[j][i-1] && s1Arr[i-1] == s3Arr[j+i-1])
			}
		}
	}
	return crossArr[len(s2Arr)][len(s1Arr)]
}
