package commands

type LinkedList[T any] struct {
	Value T
	Next  *LinkedList[T]
}

func (ll *LinkedList[T]) Add(value T) *LinkedList[T] {
	if ll == nil {
		return &LinkedList[T]{
			Value: value,
			Next:  ll,
		}
	}

	// Recursive call until we reach the end.
	ll.Next = ll.Next.Add(value)

	return ll
}
