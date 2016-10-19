package list

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	if front := l.Front(); front != nil {
		t.Error("Front is not nil")
	}
	if back := l.Back(); back != nil {
		t.Error("Back is not nil")
	}
	if length := l.Len(); length != 0 {
		t.Error("Length is not 0")
	}
}

// Creates a list and performs a couple of PushFront operations
// to ensure that it's working as expected
func TestPushFront(t *testing.T) {
	l := New()
	e1 := l.PushFront(5)
	if e1 == nil {
		t.Error("Failure: PushFront returned nil.")
		return
	}
	if l.Len() != 1 {
		t.Error("PushFront did not increment length.")
	}

	if v, ok := e1.Value.(int); !ok || v != 5 {
		t.Errorf("Expected: 5\tGot: %d\n", v)
	}
	if e1.Next() != nil || e1.Prev() != nil {
		t.Error("e1 should be singleton list.")
	}
	if l.Front() != e1 {
		t.Error("PushFront didn't update l.front")
	}
	if l.Back() != e1 {
		t.Error("PushFront didn't update l.back")
	}

	e2 := l.PushFront(4)

	if e2 == nil {
		t.Error("Failure: PushFront returned nil.")
		return
	}
	if v, ok := e2.Value.(int); !ok || v != 4 {
		t.Errorf("Expected: 4\tGot: %d\n", v)
	}
	if l.Front() != e2 || l.Back() != e1 {
		t.Error("PushFront didn't update front/back")
	}
	if e1.Prev() != e2 {
		t.Error("PushFront didn't update prevFront's prev")
	}
	if e2.Next() != e1 {
		t.Error("PushFront didn't update element pointers")
	}
}

func TestPushBack(t *testing.T) {
	l := New()
	e1 := l.PushBack(5)
	if e1 == nil {
		t.Error("Failure: PushBack returned nil.")
		return
	}
	if l.Len() != 1 {
		t.Error("PushBack did not increment length.")
	}

	if v, ok := e1.Value.(int); !ok || v != 5 {
		t.Errorf("Expected: 5\tGot: %d\n", v)
	}
	if e1.Next() != nil || e1.Prev() != nil {
		t.Error("e1 should be singleton list.")
	}
	if l.Front() != e1 {
		t.Error("PushBack didn't update l.front")
	}
	if l.Back() != e1 {
		t.Error("PushBack didn't update l.back")
	}

	e2 := l.PushBack(4)

	if e2 == nil {
		t.Error("Failure: PushBack returned nil.")
		return
	}
	if v, ok := e2.Value.(int); !ok || v != 4 {
		t.Errorf("Expected: 4\tGot: %d\n", v)
	}
	if l.Back() != e2 || l.Front() != e1 {
		t.Error("PushBack didn't update front/back")
	}
	if e1.Next() != e2 {
		t.Error("PushBack didn't update prevBack's next")
	}
	if e2.Prev() != e1 {
		t.Errorf("PushBack didn't update element pointers\n%v\n", e2.Prev())
	}
}

func Example_forwardIteration() {
	l := New()
	for i := 0; i < 5; i++ {
		_ = l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func Example_reverseIteration() {
	l := New()
	for i := 0; i < 5; i++ {
		_ = l.PushFront(i)
	}
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}
