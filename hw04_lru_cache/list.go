package hw04lrucache

type IList interface {
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

type List struct {
	head *ListItem
	last *ListItem
	len  int
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *ListItem {
	if l.len != 0 {
		return l.head
	}
	return nil
}

func (l *List) Back() *ListItem {
	if l.len != 0 {
		return l.last
	}
	return nil
}

func (l *List) PushFront(v interface{}) *ListItem {
	listItem := &ListItem{v, nil, nil}
	if l.head == nil {
		l.head = listItem
		l.last = listItem
	} else {
		listItem.Next = l.head
		l.head.Prev = listItem
		l.head = listItem
	}
	l.len++
	return listItem
}

func (l *List) PushBack(v interface{}) *ListItem {
	listItem := &ListItem{v, nil, nil}
	if l.head == nil {
		l.head = listItem
		l.last = listItem
	} else {
		l.last.Next = listItem
		listItem.Prev = l.last
		l.last = listItem
	}
	l.len++
	return listItem
}

func (l *List) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}
	i.Next = nil // avoid memory leaks
	i.Prev = nil // avoid memory leaks
	l.len--
}

// move moves e to next to at.
func (l *List) move(i, head *ListItem) {
	if i == head {
		return
	}
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}
	i.Prev = nil
	i.Next = head
	i.Next.Prev = i
	l.head = i
}

func (l *List) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	l.move(i, l.head)
}

func NewList() *List {
	return new(List)
}
