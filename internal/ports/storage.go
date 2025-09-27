package ports

import "kvsql/internal/domain"

// Storage interface defines the storage layer operations
type Storage interface {
	// Table operations
	CreateTable(table *domain.Table) error
	DropTable(name string) error
	GetTable(name string) (*domain.Table, error)
	ListTables() ([]string, error)

	// Row operations
	Insert(table string, row *domain.Row) error
	Update(table string, id string, values map[string]domain.Value) error
	Delete(table string, id string) error
	Get(table string, id string) (*domain.Row, error)

	// Batch operations
	InsertBatch(table string, rows []*domain.Row) error

	// Query operations
	Scan(table string, filter func(*domain.Row) bool) ([]*domain.Row, error)

	// Transaction support
	BeginTransaction() (Transaction, error)

	// Persistence
	Flush() error
	Close() error
}

// Transaction interface defines transaction operations
type Transaction interface {
	GetID() string
	IsReadOnly() bool

	// Transaction-specific operations
	Get(table string, id string) (*domain.Row, error)
	Insert(table string, row *domain.Row) error
	Update(table string, id string, values map[string]domain.Value) error
	Delete(table string, id string) error

	// Transaction control
	Commit() error
	Rollback() error
}

// WAL (Write-Ahead Log) interface
type WAL interface {
	WriteLog(entry *LogEntry) error
	ReadLogs(fromSequence uint64) ([]*LogEntry, error)
	Checkpoint() error
	Truncate(beforeSequence uint64) error
	Close() error
}

// LogEntry represents a single WAL entry
type LogEntry struct {
	Sequence  uint64
	Timestamp int64
	TxID      string
	Operation string // INSERT, UPDATE, DELETE, BEGIN, COMMIT, ROLLBACK
	Table     string
	Data      interface{} // Operation-specific data
}
