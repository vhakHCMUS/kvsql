package ports

// TransactionManager manages database transactions
type TransactionManager interface {
	Begin(readOnly bool) (Transaction, error)
	GetTransaction(id string) (Transaction, error)

	// Transaction lifecycle
	Commit(txID string) error
	Rollback(txID string) error

	// Lock management
	AcquireReadLock(table, rowID string, txID string) error
	AcquireWriteLock(table, rowID string, txID string) error
	ReleaseLocks(txID string) error

	// Deadlock detection
	DetectDeadlock() ([]string, error) // returns transaction IDs involved in deadlock

	// Transaction info
	ListActiveTransactions() ([]TransactionInfo, error)
	GetTransactionStatus(txID string) (TransactionStatus, error)

	// Cleanup
	CleanupExpiredTransactions() error
	Close() error
}

// TransactionStatus represents the current state of a transaction
type TransactionStatus int

const (
	TxActive TransactionStatus = iota
	TxCommitted
	TxRolledBack
	TxAborted
)

// TransactionInfo contains information about a transaction
type TransactionInfo struct {
	ID        string
	Status    TransactionStatus
	StartTime int64
	ReadOnly  bool
	Tables    []string // Tables involved in the transaction
}

// LockType represents different types of locks
type LockType int

const (
	ReadLock LockType = iota
	WriteLock
)

// Lock represents a database lock
type Lock struct {
	Type     LockType
	Table    string
	RowID    string
	TxID     string
	Acquired int64
}

// LockManager manages database locks
type LockManager interface {
	AcquireLock(lock *Lock) error
	ReleaseLock(table, rowID, txID string) error
	ReleaseAllLocks(txID string) error

	// Lock inspection
	GetLocks(table, rowID string) ([]*Lock, error)
	GetTransactionLocks(txID string) ([]*Lock, error)

	// Deadlock detection
	CheckForDeadlock(txID string) (bool, []string, error)
}
