package LRU

import "testing"

func TestConstructor(t *testing.T) {
	lru_cache := []LRUCache{Constructor(-1), Constructor(3)}
	lru_cache_capacity := []int{0, 3}

	for i := 0; i < len(lru_cache_capacity); i++ {
		if lru_cache[i].capacity != lru_cache_capacity[i] {
			t.Errorf("Capacity should be %d found %d", lru_cache_capacity[i], lru_cache[i].capacity)
		}
	}
}
