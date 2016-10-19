package list

type Element struct {
	Value interface{}
	next  *Element
	prev  *Element
}

type List struct {
	front, back *Element
	length      int
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

func New() *List {
	l := new(List)
	return l
}

func (l *List) Front() *Element {
	return l.front
}

func (l *List) Back() *Element {
	return l.back
}

func (l *List) Len() int {
	return l.length
}
