package main

type NodeList struct {
	Key  int
	Prev *NodeList
	Next *NodeList
}

type LRUCache struct {
	m           map[int]int
	cap         int
	lruListHead *NodeList
	lruListCurr *NodeList

	lruMap map[int]*NodeList
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		m:      make(map[int]int, capacity),
		cap:    capacity,
		lruMap: make(map[int]*NodeList, capacity),
		lruListHead: &NodeList{
			Key:  0,
			Prev: nil,
			Next: nil,
		},
	}
	c.lruListCurr = c.lruListHead
	return c
}

func (c *LRUCache) Get(key int) int {
	val, ok := c.m[key]
	if !ok {
		return -1
	}

	c.moveToFront(key)

	return val
}

func (c *LRUCache) Put(key int, value int) {
	_, exist := c.m[key]
	if len(c.m) < c.cap || exist {
		c.add(key, value)
		return
	}

	keyToDelete := c.lruListHead.Next.Key
	c.deleteKey(keyToDelete)

	c.add(key, value)
}

func (c *LRUCache) add(key int, value int) {
	_, exist := c.m[key]
	c.m[key] = value
	if !exist {
		c.addKeyToLRUList(key)
	} else {
		c.moveToFront(key)
	}
}

func (c *LRUCache) deleteKey(key int) {
	delete(c.m, key)
	n, ok := c.lruMap[key]
	if !ok {
		return
	}

	if n == c.lruListCurr {
		c.lruListCurr = c.lruListCurr.Prev
	}

	if n.Prev != nil {
		n.Prev.Next = n.Next
		if n.Next != nil {
			n.Next.Prev = n.Prev
		}
	}

	delete(c.lruMap, key)
}

func (c *LRUCache) addKeyToLRUList(key int) {
	c.lruListCurr.Next = &NodeList{
		Key:  key,
		Next: nil,
	}
	c.lruListCurr.Next.Prev = c.lruListCurr
	c.lruListCurr = c.lruListCurr.Next

	c.lruMap[key] = c.lruListCurr
}

func (c *LRUCache) moveToFront(key int) {
	n, ok := c.lruMap[key]
	if !ok {
		return
	}

	if c.lruListCurr == n {
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	c.lruListCurr.Next = n
	n.Next = nil
	n.Prev = c.lruListCurr
	c.lruListCurr = n
}
