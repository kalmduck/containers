package list

import "testing"

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
