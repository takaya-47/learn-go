package main

import "fmt"

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // Hello Bob
	fmt.Println(helloPrefix("Maria")) // Hello Maria
}

func prefixer(prefix string) func(string) string {
	return func(st string) string {
		return prefix + " " + st
	}
}
