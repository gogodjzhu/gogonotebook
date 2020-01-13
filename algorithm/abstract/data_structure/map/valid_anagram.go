package _map

type ValidAnagramSolver struct {
}

func (v ValidAnagramSolver) ValidAnagramSolve(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	byteCntMap := make(map[byte]int)
	for _, b := range []byte(s) {
		if _, ok := byteCntMap[b]; !ok {
			byteCntMap[b] = 1
		} else {
			byteCntMap[b] += 1
		}
	}
	for _, b := range []byte(t) {
		if _, ok := byteCntMap[b]; !ok {
			return false
		} else {
			byteCntMap[b] -= 1
		}
	}
	for _, cnt := range byteCntMap {
		if cnt != 0 {
			return false
		}
	}
	return true
}
