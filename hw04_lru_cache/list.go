package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.head != nil {
		return l.head
	}

	return nil
}

func (l *list) Back() *ListItem {
	if l.tail != nil {
		return l.tail
	}

	return nil
}

func (l *list) PushFront(v interface{}) *ListItem {
	newNode := &ListItem{
		Value: v,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head.Prev = newNode
		l.head = newNode

	}
	l.len++

	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	newNode := &ListItem{
		Value: v,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		currentNode := l.tail
		newNode.Prev = currentNode
		currentNode.Next = newNode
		l.tail = newNode
	}
	l.len++

	return l.tail
}

func (l *list) Remove(i *ListItem) {
	if l.tail.Prev == i.Prev {
		ptr := l.tail.Prev
		ptr.Next = nil
		l.tail = ptr
		l.len--
		return
	}

	if l.head.Next == i.Next {
		ptr := l.head.Next
		ptr.Prev = nil
		l.head = ptr
		l.len--
		return
	}

	left := i.Prev
	right := i.Next
	left.Next = right
	right.Prev = left

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || l.head == i {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
}

func NewList() List {
	return new(list)
}
