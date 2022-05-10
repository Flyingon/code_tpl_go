package main

// LRUCache 结构体
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func NewLRUCache(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *DLinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

func main() {
	lRUCache := NewLRUCache(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	//spew.Dump(lRUCache.cache)
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	//spew.Dump(lRUCache.cache)
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//spew.Dump(lRUCache.cache)
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//spew.Dump(lRUCache.cache)
	lRUCache.Get(1) // 返回 -1 (未找到)
	lRUCache.Get(3) // 返回 3
	lRUCache.Get(4) // 返回 4
}
