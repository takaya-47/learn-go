package main

import "fmt"

type SinglyLinkedList[T comparable] struct {
	Value T
	Next  *SinglyLinkedList[T]
}

func (s *SinglyLinkedList[T]) Add(v T) {
	// 最後のノードまで移動する
	for s.Next != nil {
		s = s.Next
	}

	// 最後のノードのNextに新しいノードを作成する
	s.Next = &SinglyLinkedList[T]{Value: v}
}

func (s *SinglyLinkedList[T]) Insert(v T, pos int) {
	// Insertしたい場所までノードを移動する
	for i := 0; i < pos-1; i++ {
		if s.Next == nil {
			break // posがノードの数より大きい場合は最後に追加するため、移動を終了する
		}
		s = s.Next
	}

	// Insertしたい場所に新しいノードを作成する
	new := &SinglyLinkedList[T]{Value: v}
	new.Next = s.Next
	s.Next = new
}

func (s *SinglyLinkedList[T]) Index(v T) int {
	i := 0
	for s != nil {
		if s.Value == v {
			return i
		}
		s = s.Next
		i++
	}
	return -1
}

func main() {
	node := &SinglyLinkedList[int]{Value: 1}
	node.Add(2)
	node.Add(3)
	node.Add(4)
	node.Add(5)

	node.Insert(99, 3)
	node.Insert(999, 6)
	node.Insert(777, 13)

	fmt.Println(node.Index(5))
	fmt.Println(node.Index(777))
}
