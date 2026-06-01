// package main

// import "fmt"

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 		for _, v := range s {
// 			ch1 <- v
// 		}
// 		close(ch1)
// 	}()

// 	go func() {
// 		s := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
// 		for _, v := range s {
// 			ch2 <- v
// 		}
// 		close(ch2)
// 	}()

// 	for i := 0; i < 2; {
// 		select {
// 		case v, ok := <-ch1:
// 			if !ok {
// 				ch1 = nil // nilにすることで、以降のselectでこのケースは選択されなくなる
// 				i++
// 				continue
// 			}

// 			fmt.Printf("%v from ch1\n", v)
// 		case v, ok := <-ch2:
// 			if !ok {
// 				ch2 = nil
// 				i++
// 				continue
// 			}
// 			fmt.Printf("%v from ch2\n", v)
// 		}
// 	}
// }
