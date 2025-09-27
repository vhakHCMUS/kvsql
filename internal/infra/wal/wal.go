package wal

import (
	"errors"
	"kvsql/internal/ports"
	"os"
	"sync"
)

// WAL (Write-Ahead Log) implementation
type WAL struct {
	mu       sync.Mutex
	file     *os.File
	sequence uint64
	// TODO: Add WAL components
	// - Buffer for batching writes
	// - Checkpointing mechanism
	// - Log rotation
}

// NewWAL creates a new Write-Ahead Log
func NewWAL(filePath string) (*WAL, error) {
	// TODO: Initialize WAL
	// Should:
	// - Open or create WAL file
	// - Read last sequence number
	// - Set up buffering
	// - Enable crash recovery

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	wal := &WAL{
		file:     file,
		sequence: 0, // TODO: Read from existing WAL file
	}

	// TODO: Recovery - read existing log and get last sequence
	if err := wal.recover(); err != nil {
		file.Close()
		return nil, err
	}

	return wal, nil
}

// recover reads existing WAL and recovers state
func (w *WAL) recover() error {
	// TODO: Implement WAL recovery
	// Should:
	// - Read all log entries
	// - Find last sequence number
	// - Validate log integrity
	// - Prepare for replay if needed

	return errors.New("TODO: implement WAL recovery")
}

// WriteLog writes a log entry to WAL
func (w *WAL) WriteLog(entry *ports.LogEntry) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Implement log writing
	// Should:
	// - Assign sequence number
	// - Serialize entry to binary format
	// - Write to file with checksum
	// - Optionally flush to disk

	// Assign sequence number
	w.sequence++
	entry.Sequence = w.sequence

	// TODO: Serialize and write entry
	// For now, just return placeholder error
	return errors.New("TODO: implement WriteLog")
}

// ReadLogs reads log entries starting from a sequence number
func (w *WAL) ReadLogs(fromSequence uint64) ([]*ports.LogEntry, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Implement log reading
	// Should:
	// - Read from specified sequence
	// - Deserialize entries
	// - Validate checksums
	// - Return entries in order

	return nil, errors.New("TODO: implement ReadLogs")
}

// Checkpoint creates a checkpoint and truncates old logs
func (w *WAL) Checkpoint() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Implement checkpointing
	// Should:
	// - Ensure all data is persisted
	// - Mark checkpoint in log
	// - Allow truncation of old entries

	return errors.New("TODO: implement Checkpoint")
}

// Truncate removes log entries before specified sequence
func (w *WAL) Truncate(beforeSequence uint64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Implement log truncation
	// Should:
	// - Remove old log entries
	// - Compact log file
	// - Update metadata

	return errors.New("TODO: implement Truncate")
}

// Close closes the WAL
func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Flush any pending writes

	if w.file != nil {
		return w.file.Close()
	}

	return nil
}

// Flush ensures all pending writes are persisted
func (w *WAL) Flush() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// TODO: Flush buffer and sync file
	if w.file != nil {
		return w.file.Sync()
	}

	return nil
}

// GetLastSequence returns the last written sequence number
func (w *WAL) GetLastSequence() uint64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.sequence
}
