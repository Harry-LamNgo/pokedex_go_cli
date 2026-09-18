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
			key: "https://zenquotes.io/",
			val: []byte(`"quote":"random quote test"}`),
		},
		{
			key: "https://zenquotes.io/",
			val: []byte(`"quote":"Different quote testing time"}`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			value, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			} else {
				t.Logf("Actual key %v\n", c.key)
			}

			if string(value) != string(c.val) {
				t.Errorf("Expected to find value")
				return
			} else {
				t.Logf("Expected value: %v -- Actual value: %v \n", c.val, value)
			}
		})
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("https://zenquotes.io/", []byte("testing data"))

	// when cache.Add url -> there is data in cache
	// -> cache.Get must be true (because there is data cache)

	_, ok := cache.Get("https://zenquotes.io/")
	if !ok {
		t.Errorf("Expected to find key")
		return
	}

	// Wait for interval time between two times cache
	time.Sleep(waitTime)

	// Now the old cache from Add has been delete because time is greater than interval time of cache
	// -> The entry added by Add has now been removed from the map, because its age exceeded the cache's interval
	// -> "cache.Get cannot find the old key of old cache -> so cache.Get must be false
	_, ok = cache.Get("https://zenquotes.io/")
	if ok {
		t.Errorf("Expected to Not find key")
		return
	}
}
