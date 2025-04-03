package collection

import (
	"fmt"
	"sync"

	"github.com/mephir/teryt-golang/internal/dataset/model"
)

type Collection[M model.Model] struct {
	Items map[uint]*M
	mu    sync.RWMutex
}

func NewCollection[M model.Model]() *Collection[M] {
	return &Collection[M]{Items: make(map[uint]*M)}
}

func (c *Collection[M]) Add(item *M) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item == nil {
		return fmt.Errorf("item cannot be nil")
	}

	if _, exists := c.Items[(*item).Identifier()]; exists {
		return fmt.Errorf("item with id %d already exists", (*item).Identifier())
	}

	c.Items[(*item).Identifier()] = item

	return nil
}

func (c *Collection[M]) Get(id uint) *M {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if item, exists := c.Items[id]; exists {
		return item
	}

	return nil
}

func (c *Collection[M]) Remove(id uint) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Items, id)
}

func (c *Collection[M]) All() []*M {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make([]*M, 0, len(c.Items))
	for _, item := range c.Items {
		items = append(items, item)
	}

	return items
}

func (c *Collection[M]) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.Items)
}

func (c *Collection[M]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Items = make(map[uint]*M)
}

func (c *Collection[M]) Contains(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.Items[id]
	return exists
}
