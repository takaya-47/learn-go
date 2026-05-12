package main

import (
	"fmt"
)

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total += i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("total=%v\n", total)
}
