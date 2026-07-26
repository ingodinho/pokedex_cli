package pokecache

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestCacheGet(t *testing.T) {
	tests := map[string]struct {
		data []byte
		key  string
	}{
		"happy": {data: []byte("test1data"), key: "test1"},
	}

	cache := NewPokeCache(time.Second * 10)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache.Add(tc.key, tc.data)
			got, inCache := cache.Get(tc.key)

			if !inCache {
				t.Fatalf("value not in cache: %v", inCache)
			}

			diff := cmp.Diff(got, tc.data)
			if diff != "" {
				t.Fatalf("%v", diff)
			}
		})
	}
}
