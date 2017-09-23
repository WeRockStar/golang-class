package golang

import "strings"

func LettersCount(sentense string) map[string]int {
	var answer = make(map[string]int)
	var spilt = strings.Split(sentense, " ")
	for _, v := range spilt {
		answer[v] = len(v)
	}

	return answer
}
