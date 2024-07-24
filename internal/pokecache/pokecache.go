package pokecache

import (
	"time"
)

type Cache struct {
	data map[uint]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
