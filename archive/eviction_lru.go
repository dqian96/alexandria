package archive

import "container/list"

type lru struct {
	queue *list.List
}

// NewLRU create a least-recently-used EvictionPolicy.
func NewLRU() EvictionPolicy {
	return &lru{
		queue: list.New(),
	}
}

func (l *lru) Evict() (string, bool) {
	front := l.queue.Front()
	if front == nil {
		return "", false
	}
	l.queue.Remove(l.queue.Front())
	if key, ok := front.Value.(string); ok {
		// ok to prevent panic if type doesn't match
		return key, true
	}
	return "", false
}

func (l *lru) Admit(key string) {
	for n := l.queue.Front(); n != nil; n = n.Next() {
		if n.Value == key {
			l.queue.MoveToBack(n)
			return
		}
	}
	l.queue.PushBack(key)
}

func (l *lru) Disregard(key string) {
	for n := l.queue.Front(); n != nil; n = n.Next() {
		if n.Value == key {
			l.queue.Remove(n)
			return
		}
	}
}
