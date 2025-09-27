package domain

import "errors"

// Common domain errors
var (
	ErrNotFound     = errors.New("record not found")
	ErrDuplicate    = errors.New("duplicate key")
	ErrInvalidQuery = errors.New("invalid query")
	ErrTransaction  = errors.New("transaction error")
)

// DataType represents the supported data types
type DataType int

const (
	TypeString DataType = iota
	TypeInt
	TypeFloat
	TypeBool
	TypeVector
)

// Value represents a value with its type
type Value struct {
	Type DataType
	Data interface{}
}

// Row represents a database row
type Row struct {
	ID     string
	Values map[string]Value
}

// Table represents a database table structure
type Table struct {
	Name    string
	Schema  map[string]DataType
	Indexes map[string]IndexType
}

// IndexType represents different index types
type IndexType int

const (
	BTreeIndex IndexType = iota
	VectorIndex
)

// QueryResult represents the result of a query execution
type QueryResult struct {
	Rows     []Row
	Affected int
	Error    error
}

// Transaction represents a database transaction
type Transaction struct {
	ID        string
	ReadOnly  bool
	StartTime int64
	// TODO: Add transaction state management
}
