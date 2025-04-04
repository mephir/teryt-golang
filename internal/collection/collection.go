package collection

import (
	"fmt"
	"sync"
)

type Collectable interface {
	Identifier() uint
}

type Collection[C Collectable] struct {
	Items map[uint]*C
	mu    sync.RWMutex
}

func NewCollection[C Collectable]() *Collection[C] {
	return &Collection[C]{Items: make(map[uint]*C)}
}

func (c *Collection[C]) Add(items ...*C) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range items {
		if item == nil {
			return fmt.Errorf("item cannot be nil")
		}

		if _, exists := c.Items[(*item).Identifier()]; exists {
			return fmt.Errorf("item with id %d already exists", (*item).Identifier())
		}

		c.Items[(*item).Identifier()] = item
	}

	return nil
}

func (c *Collection[C]) Get(id uint) *C {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if item, exists := c.Items[id]; exists {
		return item
	}

	return nil
}

func (c *Collection[C]) Remove(id uint) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Items, id)
}

func (c *Collection[C]) All() []*C {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make([]*C, 0, len(c.Items))
	for _, item := range c.Items {
		items = append(items, item)
	}

	return items
}

func (c *Collection[C]) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.Items)
}

func (c *Collection[C]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Items = make(map[uint]*C)
}

func (c *Collection[C]) Contains(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.Items[id]
	return exists
}

func (c *Collection[C]) Iterator() <-chan *C {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ch := make(chan *C)
	go func() {
		defer close(ch)
		for _, item := range c.Items {
			ch <- item
		}
	}()

	return ch
}
