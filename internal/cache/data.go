package cache

import (
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type CacheConfig struct {
	TTL    time.Duration // duration of persisted data per item
	Logger echo.Logger
}

type Cache[T CacheItemType] struct {
	items []CacheItem[T] // list of pets or users

	config CacheConfig
	m      sync.Mutex
}

type CacheItem[T CacheItemType] struct {
	created time.Time
	data    T
}

func NewCache[T CacheItemType](t CacheType, o CacheConfig) *Cache[T] {
	c := &Cache[T]{
		config: o,
		items:  []CacheItem[T]{},
	}

	// initialize cache cleaner based on items TTL
	go c.expireElements()

	return c
}

func (c *Cache[T]) Save(items ...T) error {
	c.m.Lock()
	defer c.m.Unlock()

	l := []CacheItem[T]{}

	for _, item := range items {
		i := CacheItem[T]{
			created: time.Now(),
			data:    item,
		}
		l = append(l, i)
	}

	c.items = append(c.items, l...)

	return nil
}

func (c *Cache[T]) GetAll() []T {
	c.m.Lock()
	defer c.m.Unlock()

	l := []T{}

	for _, item := range c.items {
		l = append(l, item.data)
	}

	return l
}

func (c *Cache[T]) expireElements() {
	for {
		c.m.Lock()
		if len(c.items) > 0 {
			filtered := []CacheItem[T]{}

			for _, item := range c.items {
				if item.created.Add(c.config.TTL).After(time.Now()) {
					filtered = append(filtered, item)
				}
			}

			c.items = filtered
		}
		c.m.Unlock()

		time.Sleep(500 * time.Millisecond)
	}
}
