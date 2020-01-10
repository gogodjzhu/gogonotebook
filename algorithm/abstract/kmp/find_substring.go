package kmp

type FindSubstringSolver struct {
}

func (FindSubstringSolver) FindSubstringSolver(p, t string) int {
	if len(t) < 1 || len(p) < 1 {
		return -1 // empty t return -1
	}
	tArr, pArr := []byte(t), []byte(p)
	i, j := 0, 0
	next := getNext(tArr)
	for i < len(pArr) {
		for j < len(tArr) && i < len(pArr) {
			if tArr[j] != pArr[i] {
				if j == 0 {
					i++ // T(j) has reset to front, let P(i) move next
					break
				}
				j = next[j]
				break
			}
			// P(i) = T(i), so both go to next
			i += 1
			j += 1
		}
		if j+1 >= len(tArr) { // reach end of j
			return i - j
		}
	}
	return -1
}

func getNext(tArr []byte) []int {
	next := make([]int, len(tArr))
	next[0] = 0 // force 0 to start a check round at front of tArr
	if len(tArr) <= 1 {
		return next // single byte
	}
	next[1] = 0 // set the first, than each othor depends on it
	for j := 2; j < len(tArr); j++ {
		if tArr[j-1] == tArr[next[j-1]] { // depends on the number of same bytes of the prefix and suffix
			next[j] = next[j-1] + 1
		}
		// else default to zero
	}
	return next
}
