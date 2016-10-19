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

// Returns the first of the List
func (l *List) Front() *Element {
	return l.front
}

func (l *List) Back() *Element {
	return l.back
}

func (l *List) Len() int {
	return l.length
}

func (l *List) PushFront(v interface{}) *Element {
	e := &Element{v, nil, nil}

	if l.length == 0 { // empty List
		l.back = e
		l.front = e
	} else { // non-empty List
		oldFront := l.front
		e.next = oldFront
		oldFront.prev = e
		l.front = e
	}

	l.length++
	return e
}

func (l *List) PushBack(v interface{}) *Element {
	oldBack := l.back
	e := &Element{v, nil, oldBack}
	l.back = e
	if l.length == 0 {
		l.front = e
	} else {
		oldBack.next = e
	}

	l.length++
	return e
}
