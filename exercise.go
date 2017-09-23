package golang

func factorial(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	return sum
}
