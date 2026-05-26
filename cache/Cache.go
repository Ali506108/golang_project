package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrKeyFound = errors.New("cache : key not found")

type Cache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]item[V]
}

type item[V any] struct {
	value     V
	createdAt time.Time
	ttl       time.Duration
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]item[V]),
	}
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item[V]{
		value:     value,
		createdAt: time.Now(),
		ttl:       ttl,
	}
}

func (c *Cache[K, V]) Get(ctx context.Context, key K) (V, error) {
	c.mu.Lock()
	defer c.mu.RUnlock()

	if err := ctx.Err(); err != nil {
		var zero V
		return zero, err
	}

	itm, existis := c.items[key]
	if !existis {
		var zero V
		return zero, ErrKeyFound
	}

	if time.Since(itm.createdAt) > itm.ttl {
		var zero V
		return zero, ErrKeyFound
	}
	return itm.value, nil
}
