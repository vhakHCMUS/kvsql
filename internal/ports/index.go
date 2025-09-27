package ports

import "kvsql/internal/domain"

// Index interface defines operations for all index types
type Index interface {
	Insert(key interface{}, value string) error // value is usually row ID
	Delete(key interface{}) error
	Search(key interface{}) ([]string, error) // returns row IDs
	Range(start, end interface{}) ([]string, error)
	Close() error
}

// BTreeIndex interface for B-Tree indexing
type BTreeIndex interface {
	Index

	// B-Tree specific operations
	GetMin() (interface{}, []string, error)
	GetMax() (interface{}, []string, error)

	// Range queries
	GreaterThan(key interface{}) ([]string, error)
	LessThan(key interface{}) ([]string, error)
	Between(start, end interface{}) ([]string, error)
}

// VectorIndex interface for vector similarity search
type VectorIndex interface {
	Index

	// Vector specific operations
	InsertVector(id string, vector []float64) error
	SearchSimilar(query []float64, k int) ([]VectorResult, error)
	SearchByThreshold(query []float64, threshold float64) ([]VectorResult, error)

	// Index maintenance
	BuildIndex() error
	GetIndexInfo() *VectorIndexInfo
}

// VectorResult represents a result from vector similarity search
type VectorResult struct {
	ID       string
	Distance float64
	Vector   []float64
}

// VectorIndexInfo contains metadata about the vector index
type VectorIndexInfo struct {
	Dimension int
	Count     int
	IndexType string // "HNSW", "IVF", etc.
}

// IndexManager manages all indexes for a table
type IndexManager interface {
	CreateIndex(table, column string, indexType domain.IndexType) error
	DropIndex(table, column string) error
	GetIndex(table, column string) (Index, error)
	ListIndexes(table string) (map[string]domain.IndexType, error)

	// Batch operations for index maintenance
	RebuildIndex(table, column string) error
	OptimizeIndexes(table string) error
}
