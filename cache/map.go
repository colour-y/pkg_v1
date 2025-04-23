package cache

import (
	"errors"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type MapCache[k any] struct {
	m *cmap.ConcurrentMap[string, k]
}

var KeyEmpty = errors.New("key is empty")

func NewMapCache[k any]() *MapCache[k] {
	m := cmap.New[k]()
	return &MapCache[k]{
		m: &m,
	}
}

func (c *MapCache[k]) Set(key string, value k) {
	c.m.Set(key, value)
}

func (c *MapCache[k]) Get(key string) (k, bool) {
	s, ok := c.m.Get(key)
	if ok {
		return s, ok
	} else {
		var zero k
		return zero, ok

	}
}

func (c *MapCache[k]) Update(
	key string, value k,
	f func(exist bool, valueInMap k, newValue k) k,
) k {
	return c.m.Upsert(key, value, f)
}

func (c *MapCache[k]) Remove(key string) {
	c.m.Remove(key)
}

func (c *MapCache[k]) Iter(f func(key string, e k)) {
	c.m.IterCb(f)
}

func (c *MapCache[k]) Count() int {
	return c.m.Count()
}

func (c *MapCache[k]) Clear() {
	keys := make([]string, 0, c.m.Count())
	c.Iter(func(key string, e k) {
		keys = append(keys, key)
	})
	for _, key := range keys {
		c.m.Remove(key)
	}

}
