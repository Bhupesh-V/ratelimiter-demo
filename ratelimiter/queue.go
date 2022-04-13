package ratelimiter

import (
	"container/list"
	"fmt"
	"sync"
)

type Queue struct {
	q   list.List
	cap int64
	mu  sync.Mutex
}

func CreateQueue(capacity int64) *Queue {
	return &Queue{q: *list.New(), cap: capacity}

}
func (q *Queue) CurrentCapacity() int64 {
	return q.cap
}

func (q *Queue) Add() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.CurrentCapacity() <= q.cap {
		q.q.PushBack(1)
		q.cap = q.cap - 1
	}
	fmt.Println(q.CurrentCapacity())
}

func (q *Queue) Clear(originalCap int64) {
	// Init() clears our queue
	q.q.Init()
	// reset to original capacity
	q.cap = originalCap
}
