package queue

import (
	"testing"
)

func TestNewQueue_SizeIsZero(t *testing.T) {
	q := NewQueue()
	if q.Size() != 0 {
		t.Errorf("Expected size to be %d but was %d", 0, q.Size())
	}
}

func TestPush_SizeOne(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	if q.Size() != 1 {
		t.Errorf("Expected size to be %d but was %d", 1, q.Size())
	}
}

func TestPush_SizeTwo(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(1)
	if q.Size() != 2 {
		t.Errorf("Expected size to be %d but was %d", 2, q.Size())
	}
}

func TestPop_SizeOne(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Pop()
	if q.Size() != 0 {
		t.Errorf("Expected size to be %d but was %d", 0, q.Size())
	}
}

func TestPop_SizeTwo(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(1)
	q.Pop()
	if q.Size() != 1 {
		t.Errorf("Expected size to be %d but was %d", 1, q.Size())
	}
}

func TestPop_ZeroElements(t *testing.T) {
	q := NewQueue()
	x := q.Pop()

	if x != nil {
		t.Errorf("Expected value to be nil but was %d", q.Size())
	}
}

func TestPop_OneElement(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	x := q.Pop().(int)

	if x != 1 {
		t.Errorf("Expected value to be %d but was %d", 1, q.Size())
	}
}

func TestPop_TwoElements(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)

	x := q.Pop().(int)
	if x != 1 {
		t.Errorf("Expected value to be %d but was %d", 1, q.Size())
	}

	x = q.Pop().(int)
	if x != 2 {
		t.Errorf("Expected value to be %d but was %d", 2, q.Size())
	}
}
