package golang

func fibonacci(x int) []int {
	var answer = []int{}
	for i := 0; i < x; i++ {
		if i == 0 || i == 1 {
			answer = append(answer, i)
			continue
		}
		answer = append(answer, (answer[i-2] + answer[i-1]))
	}
	return answer
}
