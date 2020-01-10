package dynamic_programming

type LetterCombinationsOfAPhoneNumberSolver struct {
}

func (LetterCombinationsOfAPhoneNumberSolver) LetterCombinationsOfAPhoneNumberSolve(digits string) []string {
	arr := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	digitsArr := []byte(digits)
	results := []string{""}
	for i := 0; i < len(digitsArr); i++ {
		numIndex := digitsArr[i] - '2'
		var tmpResult []string
		for _, prev := range results {
			for _, char := range arr[numIndex] {
				tmpResult = append(tmpResult, prev+char)
			}
		}
		results = tmpResult
	}
	return results
}
