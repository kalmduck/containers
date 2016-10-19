package list

type Element struct {
	Value interface{}
	next  *Element
	prev  *Element
	l     *List
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
	return new(List)
}

// Returns the first element of the List
func (l *List) Front() *Element {
	return l.front
}

// Back returns the last element of the list
func (l *List) Back() *Element {
	return l.back
}

// Len returns the current length of the list.
// Runs in constant time.
func (l *List) Len() int {
	return l.length
}

// PushFront creates and adds an element with Value v
// to the front of the list and returns the new element.
func (l *List) PushFront(v interface{}) *Element {
	e := &Element{Value: v, l: l}

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

// PushBack adds an element with the passed value to the
// end of the list and returns the created element
func (l *List) PushBack(v interface{}) *Element {
	oldBack := l.back
	e := &Element{v, nil, oldBack, l}
	l.back = e
	if l.length == 0 {
		l.front = e
	} else {
		oldBack.next = e
	}

	l.length++
	return e
}

func (l *List) remove(e *Element) *Element {
	p, n := e.prev, e.next
	if p != nil {
		p.next = n
	} else { // removing the front
		l.front = n
	}
	if n != nil {
		n.prev = p
	} else { // removing the back
		l.back = p
	}
	e.prev = nil
	e.next = nil
	l.length--
	return e
}

func (l *List) Remove(e *Element) interface{} {
	if e.l == l {
		l.remove(e)
	}
	return e.Value
}
