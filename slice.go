package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 4, 2, 1, 2}
	fmt.Println("Lenght:", len(s), "Capability:", cap(s))

	s = append(s, 4)
	fmt.Println("Lenght:", len(s), "Capability:", cap(s))

	s = append(s, 4)
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
	fmt.Println("Lenght:", len(s), "Capability:", cap(s))
}
