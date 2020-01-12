package _map

import "fmt"

type GroupAnagramsSolver struct {
}

/**
字母异位词分组, 关键在于将字符的关系映射成出现次数的数组
构造一个长度为26的[]int, 每个位置保存一个字符的出现次数, 所有映射成相同数组的字符串,
属于字母异位词
*/
func (g GroupAnagramsSolver) GroupAnagramsSolve(strs []string) [][]string {
	m := make(map[string][]int)

	for i, str := range strs {
		arr := make([]int, 26) // 字母表
		bs := []byte(str)
		for i := 0; i < len(bs); i++ {
			arr[bs[i]-'a'] += 1
		}
		alphaBitStr := ""
		for _, a := range arr {
			alphaBitStr += fmt.Sprintf("%d,", a)
		}
		if idxes, ok := m[alphaBitStr]; ok {
			m[alphaBitStr] = append(idxes, i)
		} else {
			m[alphaBitStr] = []int{i}
		}
	}

	var ret [][]string
	for _, v := range m {
		var group []string
		for _, i := range v {
			group = append(group, strs[i])
		}
		ret = append(ret, group)
	}
	return ret
}
