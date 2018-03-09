package cache

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, timeout time.Duration) error
	Get(key string) (interface{}, bool)
}
