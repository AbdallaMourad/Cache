package LFU

import "testing"

func TestConstructor(t *testing.T) {
	lfu_cache := []LFUCache{Constructor(-1), Constructor(3)}
	lfu_cache_capacity := []int{0, 3}

	for i := 0; i < len(lfu_cache_capacity); i++ {
		if lfu_cache[i].capacity != lfu_cache_capacity[i] {
			t.Errorf("Capacity should be %d found %d", lfu_cache_capacity[i], lfu_cache[i].capacity)
		}
	}
}
