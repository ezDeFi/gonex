package core

import (
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"

	"golang.org/x/time/rate"
)

// MaxConcurentLimiter defines the maximum number of concurent limiter
const MaxConcurentLimiter = 32

// A Throttler controls how often events are allowed to happen per string key
type Throttler struct {
	m *sync.Mutex

	limit    rate.Limit
	burst    int
	limiters *lru.ARCCache
}

// NewThrottler returns a new Throttler with up to rate r and burts of at most b
func NewThrottler(limit rate.Limit, burst int) *Throttler {
	limiters, err := lru.NewARC(MaxConcurentLimiter)
	if err != nil {
		return nil
	}
	return &Throttler{
		m:        &sync.Mutex{},
		limit:    limit,
		burst:    burst,
		limiters: limiters,
	}
}

// t.m locked assumed
func (t *Throttler) get(key string) *rate.Limiter {
	limiter, ok := t.limiters.Get(key)
	if !ok {
		limiter = rate.NewLimiter(t.limit, t.burst)
		t.limiters.Add(key, limiter)
	}
	return limiter.(*rate.Limiter)
}

// l.m locked assumed
func (t *Throttler) peek(key string) (*rate.Limiter, bool) {
	limiter, ok := t.limiters.Peek(key)
	if !ok {
		return nil, false
	}
	return limiter.(*rate.Limiter), true
}

// Allow reports whether one event may happen at time now for the provided key
func (t *Throttler) Allow(key string) bool {
	t.m.Lock()
	defer t.m.Unlock()

	return t.get(key).Allow()
}

// AllowN reports whether n events may happen at time now for the provided key
func (t *Throttler) AllowN(key string, n int) bool {
	t.m.Lock()
	defer t.m.Unlock()

	return t.get(key).AllowN(time.Now(), n)
}

// SetLimit sets a new Limit for all keys.
func (t *Throttler) SetLimit(newLimit rate.Limit) {
	t.m.Lock()
	defer t.m.Unlock()

	t.limit = newLimit

	for _, key := range t.limiters.Keys() {
		limiter, _ := t.peek(key.(string))
		limiter.SetLimit(newLimit)
	}
}

// SetBurst sets a new Burst for all keys.
func (t *Throttler) SetBurst(newBurst int) {
	t.m.Lock()
	defer t.m.Unlock()

	t.burst = newBurst

	for _, key := range t.limiters.Keys() {
		limiter, _ := t.peek(key.(string))
		limiter.SetBurst(newBurst)
	}
}
