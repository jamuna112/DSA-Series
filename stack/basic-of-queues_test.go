package stack

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	que := &Queue{}

	que.Enqueue(2)
	que.Enqueue(3)
	que.Enqueue(4)

	expected := []int{2, 3, 4}

	if !reflect.DeepEqual(que.items, expected) {
		t.Errorf("test failed, expected %v but got %v", expected, que.items)
	}
}

func TestDequeue(t *testing.T) {
	que := &Queue{}

	que.Enqueue(2)
	que.Enqueue(3)
	got := que.Dequeue()

	expected := 2

	if expected != got {
		t.Errorf("test failed, expected %v but got %v", expected, got)
	}

}

func TestFront(t *testing.T) {
	que := &Queue{}

	que.Enqueue(5)
	que.Enqueue(6)
	got := que.Front()

	expected := 5

	if expected != got {
		t.Errorf("test failed, expected %v but got %v", expected, got)
	}

}

func TestFrontIsEmpty(t *testing.T) {
	que := &Queue{}

	got := que.Front()

	expected := -1

	if got != expected {
		t.Errorf("test failed, expected %v but got %v", expected, got)
	}

}

func TestIsQueueEmpty(t *testing.T) {
	que := &Queue{}

	if !que.isEmpty() {
		t.Errorf("test failed, queue is empty but got %v", que.isEmpty())
	}

	que.Enqueue(1)
	if que.isEmpty() {
		t.Errorf("test failed, expected que to be non-empty but got %v", que.isEmpty())
	}

}
