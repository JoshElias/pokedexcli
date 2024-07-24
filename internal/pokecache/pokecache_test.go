package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cache := NewCache(10 * time.Second)
	defer cache.Destroy()
	testVal := "example"
	cache.Add("test", []byte(testVal))
	if len(cache.data) < 1 {
		t.Fatal("failed to add new key/value pair")
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache(10 * time.Second)
	defer cache.Destroy()
	testVal := "example"
	cache.Add("test", []byte(testVal))
	cache.Delete("test")
	if len(cache.data) > 0 {
		t.Fatal("failed to delete key/value pair")
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(5 * time.Second)
	defer cache.Destroy()
	testVal := "example"
	cache.Add("test", []byte(testVal))
	time.Sleep(12 * time.Second)
	if _, exists := cache.Get("test"); exists {
		t.Fatal("failed to reap old key/value pair")
	}
}
