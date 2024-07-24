package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cache := NewCache(10 * time.Second)
	testVal := "example"
	cache.Add("test", []byte(testVal))
	if len(cache.data) < 1 {
		t.Fatal("failed to add new key/value pair")
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache(10 * time.Second)
	testVal := "example"
	cache.Add("test", []byte(testVal))
	cache.Delete("test")
	if len(cache.data) > 0 {
		t.Fatal("failed to delete key/value pair")
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(10 * time.Second)
	testVal := "example"
	cache.Add("test", []byte(testVal))
	time.Sleep(10 * time.Second)
	if len(cache.data) > 0 {
		t.Fatal("failed to reap old key/value pair")
	}
}
