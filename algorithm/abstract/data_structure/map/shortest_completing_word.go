package _map

import (
	"math"
	"strings"
)

type ShortestCompletingWordSolver struct {
}

func (s ShortestCompletingWordSolver) ShortestCompletingWordSolve(licensePlate string, words []string) string {
	char2CntMap := make(map[byte]int)
	arr := []byte(strings.ToLower(licensePlate))
	for _, c := range arr {
		if c < 'a' || c > 'z' {
			continue
		}
		if _, ok := char2CntMap[c]; !ok {
			char2CntMap[c] = 0
		}
		char2CntMap[c]++
	}
	idx := -1
	length := math.MaxInt32
	for i, w := range words {
		fitAll := true
		for ch, cnt := range char2CntMap {
			if strings.Count(w, string(ch)) < cnt {
				fitAll = false
				break
			}
		}
		if fitAll && len(w) < length {
			length = len(w)
			idx = i
		}
	}
	return words[idx]
}
