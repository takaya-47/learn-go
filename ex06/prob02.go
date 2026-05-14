package main

import "fmt"

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println(ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss)
}

func main() {
	ss := []string{"a", "b", "c"}
	fmt.Println(ss)

	UpdateSlice(ss, "d")
	fmt.Println(ss)
	GrowSlice(ss, "e")
	fmt.Println(ss)
}
