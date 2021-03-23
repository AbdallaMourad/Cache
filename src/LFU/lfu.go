package LFU

import (
	"container/list"
	"sort"
)

type LFUCache struct {
	capacity        int
	cache           map[int][]int // key - value - counter
	counter_map     map[int]*list.List
	key_element_map map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	if capacity < 0 {
		capacity = 0
	}

	return LFUCache{
		capacity:        capacity,
		cache:           make(map[int][]int),
		counter_map:     make(map[int]*list.List),
		key_element_map: make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.cache[key]

	if !ok {
		return -1
	}

	this.updateEntry(key)

	return value[0]
}

func (this *LFUCache) updateEntry(key int) {
	counter := this.cache[key][1]

	current_element := this.key_element_map[key]
	this.counter_map[counter].Remove(current_element)

	if this.counter_map[counter+1] == nil {
		this.counter_map[counter+1] = list.New()
	}

	this.counter_map[counter+1].PushFront(key)
	this.key_element_map[key] = this.counter_map[counter+1].Front()
	this.cache[key][1] = counter + 1

	if this.counter_map[counter].Len() == 0 {
		delete(this.counter_map, counter)
	}
}

func (this *LFUCache) Put(key int, value int) {
	_, ok := this.cache[key]

	if ok {
		this.updateEntry(key)
		this.cache[key][0] = value
	} else if this.capacity > 0 {
		if len(this.cache) == this.capacity {
			key_list := make([]int, 0)

			for counter := range this.counter_map {
				key_list = append(key_list, counter)
			}

			sort.Ints(key_list)

			if len(key_list) > 0 {
				counter := key_list[0]
				element := this.counter_map[counter].Back()
				this.counter_map[counter].Remove(element)

				delete(this.cache, element.Value.(int))
				delete(this.key_element_map, element.Value.(int))

				if this.counter_map[counter].Len() == 0 {
					delete(this.counter_map, counter)
				}
			}
		}

		this.cache[key] = []int{value, 1}
		if this.counter_map[1] == nil {
			this.counter_map[1] = list.New()
		}
		this.counter_map[1].PushFront(key)
		this.key_element_map[key] = this.counter_map[1].Front()
	}
}
