package index

import (
	"errors"
	"sync"
)

// BTree implements a B-Tree index
type BTree struct {
	mu   sync.RWMutex
	root *BTreeNode
	// TODO: Add B-Tree parameters
	// - Order (max children per node)
	// - Key comparison function
	// - Persistence layer
}

// BTreeNode represents a node in the B-Tree
type BTreeNode struct {
	Keys     []interface{} // Indexed keys
	Values   [][]string    // Row IDs for each key (can have duplicates)
	Children []*BTreeNode  // Child nodes
	IsLeaf   bool
	// TODO: Add more node metadata
	// - Parent pointer for easier navigation
	// - Node ID for persistence
}

// NewBTree creates a new B-Tree index
func NewBTree() *BTree {
	return &BTree{
		root: &BTreeNode{
			Keys:     make([]interface{}, 0),
			Values:   make([][]string, 0),
			Children: make([]*BTreeNode, 0),
			IsLeaf:   true,
		},
	}
}

// Insert adds a key-value pair to the B-Tree
func (bt *BTree) Insert(key interface{}, value string) error {
	bt.mu.Lock()
	defer bt.mu.Unlock()

	// TODO: Implement B-Tree insertion
	// Should:
	// - Find correct position for key
	// - Handle duplicates appropriately
	// - Split nodes when they overflow
	// - Update tree height if needed

	return errors.New("TODO: implement B-Tree Insert")
}

// Delete removes a key from the B-Tree
func (bt *BTree) Delete(key interface{}) error {
	bt.mu.Lock()
	defer bt.mu.Unlock()

	// TODO: Implement B-Tree deletion
	// Should:
	// - Find key in tree
	// - Remove key-value pair
	// - Rebalance tree if needed
	// - Handle underflow conditions

	return errors.New("TODO: implement B-Tree Delete")
}

// Search finds all values for a given key
func (bt *BTree) Search(key interface{}) ([]string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Implement B-Tree search
	// Should:
	// - Navigate to correct leaf node
	// - Return all values for the key
	// - Handle key not found case

	return nil, errors.New("TODO: implement B-Tree Search")
}

// Range returns all values within a key range
func (bt *BTree) Range(start, end interface{}) ([]string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Implement range query
	// Should:
	// - Find starting position
	// - Collect all keys in range
	// - Return all associated values

	return nil, errors.New("TODO: implement B-Tree Range")
}

// GetMin returns the minimum key and its values
func (bt *BTree) GetMin() (interface{}, []string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Find minimum key in tree
	// Should navigate to leftmost leaf

	return nil, nil, errors.New("TODO: implement B-Tree GetMin")
}

// GetMax returns the maximum key and its values
func (bt *BTree) GetMax() (interface{}, []string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Find maximum key in tree
	// Should navigate to rightmost leaf

	return nil, nil, errors.New("TODO: implement B-Tree GetMax")
}

// GreaterThan returns all values with keys greater than the given key
func (bt *BTree) GreaterThan(key interface{}) ([]string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Implement greater than search
	return nil, errors.New("TODO: implement B-Tree GreaterThan")
}

// LessThan returns all values with keys less than the given key
func (bt *BTree) LessThan(key interface{}) ([]string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Implement less than search
	return nil, errors.New("TODO: implement B-Tree LessThan")
}

// Between returns all values with keys between start and end (inclusive)
func (bt *BTree) Between(start, end interface{}) ([]string, error) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	// TODO: Implement between search
	// This can reuse Range implementation
	return bt.Range(start, end)
}

// Close closes the B-Tree and frees resources
func (bt *BTree) Close() error {
	bt.mu.Lock()
	defer bt.mu.Unlock()

	// TODO: Clean up resources
	// - Flush to disk if persistent
	// - Free memory
	bt.root = nil

	return nil
}

// Utility methods for B-Tree maintenance

// split splits a full node
func (bt *BTree) split(node *BTreeNode) error {
	// TODO: Implement node splitting
	// Should:
	// - Find median key
	// - Create new node with right half
	// - Update parent references
	// - Propagate split up if needed

	return errors.New("TODO: implement node splitting")
}

// merge merges two adjacent nodes
func (bt *BTree) merge(left, right *BTreeNode) error {
	// TODO: Implement node merging
	// Should:
	// - Combine keys and values
	// - Update parent references
	// - Handle children if internal nodes

	return errors.New("TODO: implement node merging")
}

// findLeaf finds the leaf node where a key should be inserted
func (bt *BTree) findLeaf(key interface{}) (*BTreeNode, error) {
	// TODO: Navigate from root to appropriate leaf
	return nil, errors.New("TODO: implement findLeaf")
}

// compareKeys compares two keys
func (bt *BTree) compareKeys(a, b interface{}) int {
	// TODO: Implement key comparison
	// Should handle different data types
	// Return: -1 if a < b, 0 if a == b, 1 if a > b
	return 0
}
