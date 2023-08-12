package datastructures

type ListNode[T comparable] struct {
	next  *ListNode[T]
	value T
}

type SingleLinkedList[T comparable] struct {
	head   *ListNode[T]
	tail   *ListNode[T]
	length int
}

func (l *SingleLinkedList[T]) prepend(node *ListNode[T]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		l.length = 1
	} else {
		node.next = l.head
		l.head = node
		l.length++
	}

}

func (l *SingleLinkedList[T]) append(node *ListNode[T]) {
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	l.length++
}

func (l *SingleLinkedList[T]) getSize() int {
	return l.length
}

func (l *SingleLinkedList[T]) getData() []T {
	data := []T{}
	node := l.head
	data = append(data, node.value)

	for node.next != nil {
		node = node.next
		data = append(data, node.value)
	}
	return data
}

func (l *SingleLinkedList[T]) delete(val T) {
	node := l.head
	if node.value == val {
		l.head = node.next
		node = nil
		l.length--
		return
	}
	prev_node := &ListNode[T]{}
	for node.next != nil {
		prev_node = node
		node = node.next
		if node.value == val {
			prev_node.next = node.next
			node = nil
			l.length--
			return
		}
	}

}

func (l *SingleLinkedList[T]) find(val T) *ListNode[T] {
	node := l.head

	for node.next != nil {
		if node.value == val {
			return node
		}
		node = node.next
	}

	if node.value == val {
		return node
	}

	return nil
}

func (l *SingleLinkedList[T]) contains(val T) bool {
	node := l.find(val)
	if node != nil {
		return true
	}

	return false
}
