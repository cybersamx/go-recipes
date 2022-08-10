package main

import (
	"fmt"
)

type LinkedList[T any] struct {
	Val  T
	Next *LinkedList[T]
}

func (ll *LinkedList[T]) Insert(i int, val T) *LinkedList[T] {
	// Call recursive until we reach the right position. As we advance to the next
	// element, we also decrement pos. When i == 0, we will be at the right position.
	if i == 0 || ll == nil {
		return &LinkedList[T]{
			Val:  val,
			Next: ll,
		}
	}

	ll.Next = ll.Next.Insert(i-1, val)

	return ll
}

func (ll *LinkedList[T]) Add(val T) *LinkedList[T] {
	if ll == nil {
		return &LinkedList[T]{
			Val:  val,
			Next: ll,
		}
	}

	// Call recursively until we reach the end.
	// Alternatively, we can also call Insert with an index at the end. We need to
	// get and pass the length of the linked list as the index to Insert.
	ll.Next = ll.Next.Add(val)

	return ll
}

func (ll *LinkedList[T]) String() string {
	var str string
	for elem := ll; elem != nil; elem = elem.Next {
		if elem == ll {
			str = fmt.Sprintf("%v", elem.Val)
			continue
		}
		str = fmt.Sprintf("%s %v", str, elem.Val)
	}

	return str
}

func main() {
	numlist := LinkedList[int]{
		Val: 1,
		Next: &LinkedList[int]{
			Val:  3,
			Next: nil,
		},
	}

	fmt.Println(numlist)

	numlist.Insert(1, 2)
	numlist.Add(4)

	fmt.Println(numlist)
}
