package ratelimiter

import (
	"time"
)

type LeakyBucket struct {
	q         *Queue
	cap       int64
	timestamp time.Time
}

func CreateBucket(cap int64) *LeakyBucket {
	queue := CreateQueue(cap)
	return &LeakyBucket{q: queue, cap: cap, timestamp: time.Now()}
}

func (l *LeakyBucket) RefreshBucket() {
	timeDiff := time.Since(l.timestamp)
	if timeDiff.Milliseconds() >= 1000 {
		l.q.Clear(l.cap)
		l.timestamp = time.Now()
	}
}

func (l *LeakyBucket) GrantAccess() bool {
	l.RefreshBucket()
	if l.q.CurrentCapacity() > 0 {
		l.q.Add()
		return true
	}
	return false
}
