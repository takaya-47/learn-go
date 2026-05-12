package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b++
	smallI++
	bigI++

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
