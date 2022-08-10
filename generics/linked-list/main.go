package main

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
