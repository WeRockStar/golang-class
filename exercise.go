package golang

func concat(first, second []string) []string {
	return append(first, second...)
}

func delFirst(data []string) []string {
	return data[1:]
}

func delLast(data []string) []string {
	return data[0 : len(data)-1]
}
func delSecond(data []string) []string {
	var answer = []string{data[0]}
	answer = append(answer, data[2:]...)
	return answer
}

func chooseOdd(data []int) []int {
	var answer = []int{}
	for _, v := range data {
		if v%2 != 0 {
			answer = append(answer, v)
		}
	}

	return answer
}
