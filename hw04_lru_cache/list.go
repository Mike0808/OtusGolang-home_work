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
	root ListItem
	len  int
	list *List
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *ListItem {
	if l.len != 0 {
		return l.root.Next
	}
	return nil
}

func (l *List) Back() *ListItem {
	if l.len != 0 {
		return l.root.Prev
	}
	return nil
}

func (l *List) Init() *List {
	l.root.Next = &l.root
	l.root.Prev = &l.root
	l.len = 0
	return l
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *ListItem) *ListItem {
	e.Prev = at
	e.Next = at.Next
	e.Prev.Next = e
	e.Next.Prev = e
	l.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *ListItem) *ListItem {
	return l.insert(&ListItem{Value: v}, at)
}

func (l *List) PushFront(v interface{}) *ListItem {
	if l.root.Next == nil {
		l.root.Next = &l.root
		l.root.Prev = &l.root
		l.len = 0
	}
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *ListItem {
	if l.root.Next == nil {
		l.root.Next = &l.root
		l.root.Prev = &l.root
		l.len = 0
	}
	return l.insertValue(v, l.root.Prev)
}

// remove removes e from its list, decrements l.len.
func (l *List) remove(e *ListItem) {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next = nil // avoid memory leaks
	e.Prev = nil // avoid memory leaks
	l.list = l
	l.len--
}

func (l *List) Remove(i *ListItem) {
	if l.list == l {
		l.remove(i)
	}
}

// move moves e to next to at.
func (l *List) move(i, at *ListItem) {
	if i == at {
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	i.Prev = at
	i.Next = at.Next
	i.Prev.Next = i
	i.Next.Prev = i
}

func (l *List) MoveToFront(i *ListItem) {
	if l.list != l || l.root.Next == i {
		return
	}
	l.move(i, &l.root)
}

func NewList() *List {
	return new(List)
}
