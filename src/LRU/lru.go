package LRU

import "container/list"

type LRUCache struct {
	capacity  int
	cache     map[int]int
	lru_map   map[int]*list.Element
	lru_queue *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		cache:     make(map[int]int),
		lru_map:   make(map[int]*list.Element),
		lru_queue: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.cache[key]
	if !ok {
		return -1
	}

	current_node := this.lru_map[key]
	this.lru_queue.MoveToFront(current_node)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		current_node := this.lru_map[key]
		this.lru_queue.Remove(current_node)
	} else if len(this.cache) == this.capacity {
		key_to_remove := this.lru_queue.Remove(this.lru_queue.Back()).(int)
		delete(this.cache, key_to_remove)
		delete(this.lru_map, key_to_remove)
	}

	current_element := this.lru_queue.PushFront(key)
	this.cache[key] = value
	this.lru_map[key] = current_element
}
