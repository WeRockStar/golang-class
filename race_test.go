package golang

import "testing"

func TestRaceCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		go scrum()
	}
}
