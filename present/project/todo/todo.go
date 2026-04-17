package todo

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle = errors.New("title cannot be empty")
	ErrNotFound   = errors.New("todo item not found")
)

type Item struct {
	ID    int
	Title string
	Done  bool
}

type List struct {
	items  []Item
	nextID int
}

// New creates an empty todo list.
func New() *List {
	return &List{
		items:  []Item{},
		nextID: 1,
	}
}

// LEVEL 1: Add + list tasks.
// Add creates a new todo item with an incremental ID.
func (l *List) Add(title string) (Item, error) {
	// TODO(Level 1):
	// 1) Trim spaces from title.
	// 2) If empty, return ErrEmptyTitle.
	// 3) Create a new Item with:
	//    - ID: nextID
	//    - Title: trimmed title
	//    - Done: false
	// 4) Increment nextID.
	// 5) Append item to l.items.
	// 6) Return the created item.
	if title == "" {
		return Item{}, ErrEmptyTitle
	}

	// Starter placeholder (intentionally incomplete for classroom use).
	return Item{}, nil
}

// LEVEL 2: Mark tasks as done and count pending work.
// Done marks the item with matching ID as completed.
func (l *List) Done(id int) error {
	// TODO(Level 2):
	// 1) Find item by ID.
	// 2) Set Done = true.
	// 3) Return nil on success.
	// 4) Return ErrNotFound if ID does not exist.

	// Starter placeholder (intentionally incomplete for classroom use).
	return nil
}

// LEVEL 3: Remove tasks from the list.
func (l *List) Remove(id int) error {
	// TODO(Level 3):
	// 1) Find the index of the item with matching ID.
	// 2) Remove it from l.items using slice append.
	// 3) Return nil on success.
	// 4) Return ErrNotFound if ID does not exist.

	// Starter placeholder (intentionally incomplete for classroom use).
	return nil
}

// LEVEL 1 helper: returns a safe copy for read-only use.
func (l *List) Items() []Item {
	// TODO(Level 1):
	// Return a copy of l.items (not the original slice).
	// Hint:
	//   out := make([]Item, len(l.items))
	//   copy(out, l.items)
	//   return out

	// Starter placeholder (intentionally incomplete for classroom use).
	return []Item{}
}

// LEVEL 2 helper: count remaining unfinished tasks.
func (l *List) PendingCount() int {
	// TODO(Level 2):
	// Count items where Done == false and return the count.

	// Starter placeholder (intentionally incomplete for classroom use).
	return 0
}
