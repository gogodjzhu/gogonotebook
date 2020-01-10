package backtracking

type LetterCombinationsOfAPhoneNumberSolver struct {
	res []string
}

var arr = [][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func (solver *LetterCombinationsOfAPhoneNumberSolver) LetterCombinationsOfAPhoneNumberSolve(digits string) []string {
	solver.backtracking("", len(digits), digits, 0)
	return solver.res
}

func (solver *LetterCombinationsOfAPhoneNumberSolver) backtracking(tracks string, l int, digits string, idx int) {
	if idx == l {
		solver.res = append(solver.res, tracks)
		return
	}
	char := []byte(digits)[idx]
	nums := char - '2'
	for i := 0; i < len(arr[nums]); i++ {
		solver.backtracking(tracks+arr[nums][i], l, digits, idx+1)
	}
}
