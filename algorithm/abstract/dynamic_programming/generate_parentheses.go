package dynamic_programming

type GenerateParenthesesSolver struct {
}

// 动态规划解括号生成器，很容易想到的是使用数组来保存每个括号序列的左括号数量，关键在于如何根据左括号的数量来判断新的括号生成策略
func (GenerateParenthesesSolver) GenerateParenthesesSolve(n int) []string {
	if n <= 0 {
		return []string{}
	}
	result := []string{"("}
	// 未配对的左括号数
	leftCount := []int{1}

	for i := 1; i < 2*n; i++ {
		tmpResult := make([]string, 0)
		tmpLeftCount := make([]int, 0)
		for j := 0; j < len(result); j++ {
			// 新增左括号的条件是保证新增之后，还有足够的空间来放下目前所有的(包括新增的这一个)未配对的左括号对应的右括号
			// 2 * n - i = 剩余空间数
			// leftCount[i] = 已有未配对的左括号数
			if 2*n-i > leftCount[j] {
				tmpResult = append(tmpResult, result[j]+"(")
				tmpLeftCount = append(tmpLeftCount, leftCount[j]+1)
			}
			// 新增右括号的条件是有未配对的左括号
			if leftCount[j] > 0 {
				tmpResult = append(tmpResult, result[j]+")")
				tmpLeftCount = append(tmpLeftCount, leftCount[j]-1)
			}
		}
		result = tmpResult
		leftCount = tmpLeftCount
	}
	return result
}
