package cache

import "time"

type CachedString struct {
	Value      string
	Expiration time.Time
}

type Cache struct {
	Cached map[string]CachedString
}

func NewCache() Cache {
	return Cache{
		Cached: make(map[string]CachedString),
	}
}

func (c *Cache) expire() {
	for key, cachedString := range c.Cached {
		if !cachedString.Expiration.IsZero() {
			if time.Now().After(cachedString.Expiration) {
				delete(c.Cached, key)
			}
		}
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.expire()
	cachedString, ok := c.Cached[key]
	return cachedString.Value, ok
}

func (c *Cache) Put(key, value string) {
	c.Cached[key] = CachedString{
		Value: value,
	}
}

func (c *Cache) Keys() []string {
	c.expire()
	keys := make([]string, 0, len(c.Cached))
	for key := range c.Cached {
		keys = append(keys, key)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Cached[key] = CachedString{
		Value:      value,
		Expiration: deadline,
	}
}
