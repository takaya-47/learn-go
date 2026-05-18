package main

import "fmt"

type doubler interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func double[T doubler](d T) T {
	return 2 * d
}

func main() {
	fmt.Println((double(3)))
	fmt.Println((double(-3)))
	fmt.Println((double(3.14)))
}
