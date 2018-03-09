package imp

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

// Cache is just an super simple example, nothing is being invalidated.
type Cache struct {
	data *goCache.Cache
}

func NewCache() *Cache {
	return &Cache{
		data: goCache.New(5*time.Second, time.Second),
	}
}

func (s *Cache) Set(key string, value interface{}, timeout time.Duration) error {
	s.data.Set(key, value, goCache.DefaultExpiration)
	return nil
}

func (s *Cache) Get(key string) (interface{}, bool) {
	return s.data.Get(key)
}
