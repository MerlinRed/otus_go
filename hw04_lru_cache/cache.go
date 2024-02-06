package hw04lrucache

type Key string

type KeyValue struct {
	Key   Key
	Value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		item.Value = KeyValue{Key: key, Value: value}
		c.queue.MoveToFront(item)
		return true
	}

	if c.queue.Len() >= c.capacity {
		oldest := c.queue.Back()
		delete(c.items, oldest.Value.(KeyValue).Key)
		c.queue.Remove(oldest)
	}

	newItem := c.queue.PushFront(KeyValue{Key: key, Value: value})
	c.items[key] = newItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(KeyValue).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
