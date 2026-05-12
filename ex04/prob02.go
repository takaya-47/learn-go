// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	s := make([]int, 0, 100)
// 	for i := 0; i < cap(s); i++ {
// 		s = append(s, rand.Intn(100))
// 	}

// 	for _, v := range s {
// 		if v%2 == 0 && v%3 == 0 {
// 			fmt.Printf("Six! %v\n", v)
// 		} else if v%2 == 0 {
// 			fmt.Printf("Two! %v\n", v)
// 		} else if v%3 == 0 {
// 			fmt.Printf("Three! %v\n", v)
// 		} else {
// 			fmt.Printf("Never mind %v\n", v)
// 		}
// 	}
// }
