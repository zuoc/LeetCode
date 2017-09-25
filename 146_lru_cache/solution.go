package _146_lru_cache

import "container/list"

type LRUCache struct {
	capacity int

	container *list.List

	bucket map[int]*list.Element
}

type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,

		container: list.New(),

		bucket: make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.bucket[key]; ok {
		this.container.MoveToFront(ele)
		return ele.Value.(*entry).val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if existVal := this.Get(key); existVal != -1 {
		this.bucket[key].Value.(*entry).val = value
	} else {
		ele := this.container.PushFront(&entry{key, value})
		this.bucket[key] = ele
	}

	if this.container.Len() > this.capacity {
		removeEle := this.container.Back()
		this.container.Remove(removeEle)
		delete(this.bucket, removeEle.Value.(*entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
