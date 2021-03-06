package goque

import (
	"fmt"
	"math"
	"os"
	"testing"
	"time"
)

func TestPriorityQueueIncompatibleType(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	q, err := OpenQueue(file)
	if err != nil {
		t.Error(err)
	}
	defer q.Drop()
	q.Close()

	if _, err = OpenPriorityQueue(file, ASC); err != ErrIncompatibleType {
		t.Error("Expected priority queue to return ErrIncompatibleTypes when opening Queue")
	}
}

func TestPriorityQueueDrop(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}

	if _, err = os.Stat(file); os.IsNotExist(err) {
		t.Error(err)
	}

	pq.Drop()

	if _, err = os.Stat(file); err == nil {
		t.Error("Expected directory for test database to have been deleted")
	}
}

func TestPriorityQueueEnqueue(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue size of 50, got %d", pq.Length())
	}
}

func TestPriorityQueueDequeueAsc(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 1, got %d", pq.Length())
	}

	deqItem, err := pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if pq.Length() != 49 {
		t.Errorf("Expected queue length of 49, got %d", pq.Length())
	}

	compStr := "value for item 1"

	if deqItem.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", deqItem.Priority)
	}

	if deqItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, deqItem.ToString())
	}
}

func TestPriorityQueueDequeueDesc(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, DESC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 1, got %d", pq.Length())
	}

	deqItem, err := pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if pq.Length() != 49 {
		t.Errorf("Expected queue length of 49, got %d", pq.Length())
	}

	compStr := "value for item 1"

	if deqItem.Priority != 4 {
		t.Errorf("Expected priority level to be 4, got %d", deqItem.Priority)
	}

	if deqItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, deqItem.ToString())
	}
}

func TestPriorityQueueDequeueByPriority(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 1, got %d", pq.Length())
	}

	deqItem, err := pq.DequeueByPriority(3)
	if err != nil {
		t.Error(err)
	}

	if pq.Length() != 49 {
		t.Errorf("Expected queue length of 49, got %d", pq.Length())
	}

	compStr := "value for item 1"

	if deqItem.Priority != 3 {
		t.Errorf("Expected priority level to be 1, got %d", deqItem.Priority)
	}

	if deqItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, deqItem.ToString())
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	compStr := "value for item 1"

	peekItem, err := pq.Peek()
	if err != nil {
		t.Error(err)
	}

	if peekItem.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", peekItem.Priority)
	}

	if peekItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, peekItem.ToString())
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 50, got %d", pq.Length())
	}
}

func TestPriorityQueuePeekByOffsetAsc(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	compStrFirst := "value for item 1"
	compStrLast := "value for item 10"
	compStr := "value for item 3"

	peekFirstItem, err := pq.PeekByOffset(0)
	if err != nil {
		t.Error(err)
	}

	if peekFirstItem.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", peekFirstItem.Priority)
	}

	if peekFirstItem.ToString() != compStrFirst {
		t.Errorf("Expected string to be '%s', got '%s'", compStrFirst, peekFirstItem.ToString())
	}

	peekLastItem, err := pq.PeekByOffset(49)
	if err != nil {
		t.Error(err)
	}

	if peekLastItem.Priority != 4 {
		t.Errorf("Expected priority level to be 4, got %d", peekLastItem.Priority)
	}

	if peekLastItem.ToString() != compStrLast {
		t.Errorf("Expected string to be '%s', got '%s'", compStrLast, peekLastItem.ToString())
	}

	peekItem, err := pq.PeekByOffset(22)
	if err != nil {
		t.Error(err)
	}

	if peekItem.Priority != 2 {
		t.Errorf("Expected priority level to be 2, got %d", peekItem.Priority)
	}

	if peekItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, peekItem.ToString())
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 50, got %d", pq.Length())
	}
}

func TestPriorityQueuePeekByOffsetDesc(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, DESC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	compStrFirst := "value for item 1"
	compStrLast := "value for item 10"
	compStr := "value for item 3"

	peekFirstItem, err := pq.PeekByOffset(0)
	if err != nil {
		t.Error(err)
	}

	if peekFirstItem.Priority != 4 {
		t.Errorf("Expected priority level to be 4, got %d", peekFirstItem.Priority)
	}

	if peekFirstItem.ToString() != compStrFirst {
		t.Errorf("Expected string to be '%s', got '%s'", compStrFirst, peekFirstItem.ToString())
	}

	peekLastItem, err := pq.PeekByOffset(49)
	if err != nil {
		t.Error(err)
	}

	if peekLastItem.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", peekLastItem.Priority)
	}

	if peekLastItem.ToString() != compStrLast {
		t.Errorf("Expected string to be '%s', got '%s'", compStrLast, peekLastItem.ToString())
	}

	peekItem, err := pq.PeekByOffset(32)
	if err != nil {
		t.Error(err)
	}

	if peekItem.Priority != 1 {
		t.Errorf("Expected priority level to be 0, got %d", peekItem.Priority)
	}

	if peekItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, peekItem.ToString())
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 50, got %d", pq.Length())
	}
}

func TestPriorityQueuePeekByPriorityID(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	compStr := "value for item 3"

	peekItem, err := pq.PeekByPriorityID(1, 3)
	if err != nil {
		t.Error(err)
	}

	if peekItem.Priority != 1 {
		t.Errorf("Expected priority level to be 1, got %d", peekItem.Priority)
	}

	if peekItem.ToString() != compStr {
		t.Errorf("Expected string to be '%s', got '%s'", compStr, peekItem.ToString())
	}

	if pq.Length() != 50 {
		t.Errorf("Expected queue length of 50, got %d", pq.Length())
	}
}

func TestPriorityQueueUpdate(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	item, err := pq.PeekByPriorityID(0, 3)
	if err != nil {
		t.Error(err)
	}

	oldCompStr := "value for item 3"
	newCompStr := "new value for item 3"

	if item.ToString() != oldCompStr {
		t.Errorf("Expected string to be '%s', got '%s'", oldCompStr, item.ToString())
	}

	if err = pq.Update(item, []byte(newCompStr)); err != nil {
		t.Error(err)
	}

	if item.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", item.Priority)
	}

	if item.ToString() != newCompStr {
		t.Errorf("Expected current item value to be '%s', got '%s'", newCompStr, item.ToString())
	}

	newItem, err := pq.PeekByPriorityID(0, 3)
	if err != nil {
		t.Error(err)
	}

	if newItem.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", newItem.Priority)
	}

	if newItem.ToString() != newCompStr {
		t.Errorf("Expected new item value to be '%s', got '%s'", newCompStr, item.ToString())
	}
}

func TestPriorityQueueUpdateString(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 0; p <= 4; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	item, err := pq.PeekByPriorityID(0, 3)
	if err != nil {
		t.Error(err)
	}

	oldCompStr := "value for item 3"
	newCompStr := "new value for item 3"

	if item.ToString() != oldCompStr {
		t.Errorf("Expected string to be '%s', got '%s'", oldCompStr, item.ToString())
	}

	if err = pq.UpdateString(item, newCompStr); err != nil {
		t.Error(err)
	}

	if item.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", item.Priority)
	}

	if item.ToString() != newCompStr {
		t.Errorf("Expected current item value to be '%s', got '%s'", newCompStr, item.ToString())
	}

	newItem, err := pq.PeekByPriorityID(0, 3)
	if err != nil {
		t.Error(err)
	}

	if newItem.Priority != 0 {
		t.Errorf("Expected priority level to be 0, got %d", newItem.Priority)
	}

	if newItem.ToString() != newCompStr {
		t.Errorf("Expected new item value to be '%s', got '%s'", newCompStr, item.ToString())
	}
}

func TestPriorityQueueHigherPriorityAsc(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 5; p <= 9; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	item, err := pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority level to be 5, got %d", item.Priority)
	}

	err = pq.Enqueue(NewPriorityItemString("value", 2))
	if err != nil {
		t.Error(err)
	}

	higherItem, err := pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if higherItem.Priority != 2 {
		t.Errorf("Expected priority level to be 5, got %d", higherItem.Priority)
	}
}

func TestPriorityQueueHigherPriorityDesc(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, DESC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	for p := 5; p <= 9; p++ {
		for i := 1; i <= 10; i++ {
			item := NewPriorityItemString(fmt.Sprintf("value for item %d", i), uint8(p))
			if err = pq.Enqueue(item); err != nil {
				t.Error(err)
			}
		}
	}

	item, err := pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if item.Priority != 9 {
		t.Errorf("Expected priority level to be 9, got %d", item.Priority)
	}

	err = pq.Enqueue(NewPriorityItemString("value", 12))
	if err != nil {
		t.Error(err)
	}

	higherItem, err := pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	if higherItem.Priority != 12 {
		t.Errorf("Expected priority level to be 12, got %d", higherItem.Priority)
	}
}

func TestPriorityQueueEmpty(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	err = pq.Enqueue(NewPriorityItemString("value for item", 0))
	if err != nil {
		t.Error(err)
	}

	_, err = pq.Dequeue()
	if err != nil {
		t.Error(err)
	}

	_, err = pq.Dequeue()
	if err != ErrEmpty {
		t.Errorf("Expected to get queue empty error, got %s", err.Error())
	}
}

func TestPriorityQueueOutOfBounds(t *testing.T) {
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		t.Error(err)
	}
	defer pq.Drop()

	err = pq.Enqueue(NewPriorityItemString("value for item", 0))
	if err != nil {
		t.Error(err)
	}

	_, err = pq.PeekByOffset(2)
	if err != ErrOutOfBounds {
		t.Errorf("Expected to get queue out of bounds error, got %s", err.Error())
	}
}

func BenchmarkPriorityQueueEnqueue(b *testing.B) {
	// Open test database
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		b.Error(err)
	}
	defer pq.Drop()

	// Create dummy data for pushing
	item := NewPriorityItemString("value", 0)

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = pq.Enqueue(item)
	}
}

func BenchmarkPriorityQueueDequeue(b *testing.B) {
	// Open test database
	file := fmt.Sprintf("test_db_%d", time.Now().UnixNano())
	pq, err := OpenPriorityQueue(file, ASC)
	if err != nil {
		b.Error(err)
	}
	defer pq.Drop()

	// Fill with dummy data
	for n := 0; n < b.N; n++ {
		if err := pq.Enqueue(NewPriorityItemString("value", uint8(math.Mod(float64(n), 255)))); err != nil {
			b.Error(err)
		}
	}

	// Start benchmark
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_, _ = pq.Dequeue()
	}
}
