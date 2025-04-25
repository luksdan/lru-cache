package main

import (
	"container/list"
	"fmt"
)

type Entry struct {
	key   string
	value interface{}
}

type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (c *LRUCache) Get(key string) interface{} {
	if l, found := c.items[key]; found {
		c.order.MoveToFront(l)
		return l.Value.(*Entry).value
	}
	return -1
}

func (c *LRUCache) Set(key string, value interface{}) {
	if l, found := c.items[key]; found {
		c.order.MoveToFront(l)
		l.Value.(*Entry).value = value
	} else {
		if c.capacity == c.order.Len() {
			oldest := c.order.Back()
			c.order.Remove(oldest)
			delete(c.items, oldest.Value.(*Entry).key)
		}
	}
	newEntry := &Entry{key, value}
	element := c.order.PushFront(newEntry)
	c.items[key] = element
}

func main() {
	cache := NewLRUCache(3)
	cache.Set("a", 1)
	cache.Set("b", 2)
	fmt.Println(cache.Get("b"))
	cache.Set("c", 3)
	cache.Set("d", 4)
	fmt.Println(cache.Get("a"))
}
