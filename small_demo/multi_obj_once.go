package main

import (
	"context"
	"fmt"
	"sync"
)

// Item represents a generic item that needs to be processed.
type Item struct {
	ID uint64 // Unique identifier for the item.
	// Other fields that represent item details...
}

// OnceItems is a structure that ensures each item is only processed once.
type OnceItems struct {
	mutex     sync.Mutex
	processed map[uint64]bool // Tracks whether an item has been processed.
}

// NewOnceItems creates a new instance of OnceItems.
func NewOnceItems() *OnceItems {
	return &OnceItems{
		processed: make(map[uint64]bool),
	}
}

// ProcessItem processes an item if it hasn't been processed yet.
func (o *OnceItems) ProcessItem(ctx context.Context, item *Item) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if _, exists := o.processed[item.ID]; exists {
		// Item has already been processed, log and skip.
		fmt.Printf("Item with ID: %d has already been processed, skipping.\n", item.ID)
		return false
	}

	// Process the item here...
	fmt.Printf("Processing item with ID: %d\n", item.ID)
	o.processed[item.ID] = true

	return true
}

func main() {
	ctx := context.Background()
	onceItems := NewOnceItems()

	// Simulate processing multiple items, some of which may be duplicates.
	itemsToProcess := []*Item{
		{ID: 1}, {ID: 2}, {ID: 1}, // Item with ID 1 is a duplicate.
	}

	for _, item := range itemsToProcess {
		onceItems.ProcessItem(ctx, item)
	}
}

/*
在 main 函数中，我们创建了一个 OnceItems 实例，并尝试处理一系列 Item 对象，其中一些可能是重复的。
通过 ProcessItem 方法的返回值，我们可以知道哪些 Item 被处理了，哪些被跳过。
这个模式在处理批量数据时非常有用，尤其是当数据中可能包含重复项，而这些重复项只需要处理一次时。这样可以避免不必要的重复工作，提高效率。
*/
