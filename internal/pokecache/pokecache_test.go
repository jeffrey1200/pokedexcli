package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://exmaple.com",
			val: []byte("testdata"),
		},
		{
			key: "https//example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("EXpected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTIme = 5 * time.Millisecond
	const waitTime = baseTIme + 5*time.Millisecond
	cache := NewCache(baseTIme)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("Expected to find key")
		return
	}
	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("Expected to not find key")
		return
	}
}
