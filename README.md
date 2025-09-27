# KVSQL Database

A learning-oriented database implementation in Go, designed to understand database internals and architecture.

## 🏗️ Architecture

This database follows clean architecture principles with clear separation of concerns:

```
cmd/
└── sqld/main.go              # Application entry point

internal/
├── domain/                   # Core business logic & entities
│   ├── model.go             # Data structures (Table, Row, Value, etc.)
│   └── ast.go               # Abstract Syntax Tree for SQL parsing
├── usecase/                 # Business logic layer
│   ├── parser.go            # SQL parser (converts SQL → AST)
│   ├── planner.go           # Query planner (creates execution plans)
│   └── executor.go          # Query executor (executes plans)
├── ports/                   # Interface definitions
│   ├── storage.go           # Storage interfaces
│   ├── index.go             # Index interfaces (BTree, Vector)
│   └── tx.go                # Transaction manager interface
├── infra/                   # Infrastructure implementations
│   ├── storage/
│   │   ├── memory.go        # In-memory storage
│   │   └── file_store.go    # File-based storage
│   ├── wal/
│   │   └── wal.go           # Write-Ahead Log
│   └── index/
│       ├── btree.go         # B-Tree implementation
│       └── hnsw_stub.go     # Vector similarity search (stub)
└── delivery/                # External interfaces
    ├── cli.go               # Command-line interface
    └── server.go            # HTTP API server
```

## 🚀 Features (Planned)

### Core Database Features
- [x] **SQL Parser** - Convert SQL strings to AST
- [x] **Query Planner** - Optimize query execution
- [x] **Query Executor** - Execute query plans
- [x] **Storage Layer** - Both memory and file-based storage
- [x] **Indexing** - B-Tree and vector similarity search
- [x] **Transactions** - ACID compliance with locking
- [x] **WAL** - Write-Ahead Logging for durability

### Supported SQL Operations
- [ ] `SELECT` - Query data with WHERE, ORDER BY, LIMIT
- [ ] `INSERT` - Add new records
- [ ] `UPDATE` - Modify existing records
- [ ] `DELETE` - Remove records
- [ ] `CREATE TABLE` - Define new tables
- [ ] `DROP TABLE` - Remove tables

### Advanced Features
- [ ] **Vector Search** - Similarity search using HNSW algorithm
- [ ] **Concurrent Transactions** - Multi-user support with proper isolation
- [ ] **Index Management** - Automatic and manual index creation
- [ ] **Query Optimization** - Cost-based optimization

## 🛠️ Getting Started

### Prerequisites
- Go 1.21 or later

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd kvsql
```

2. Build the project:
```bash
go build -o sqld cmd/sqld/main.go
```

### Usage

#### CLI Mode (Interactive)
```bash
./sqld -mode=cli
```

This starts an interactive SQL shell:
```
Welcome to KVSQL Database
Type 'help' for available commands, 'exit' to quit
kvsql> CREATE TABLE users (id INT, name STRING);
kvsql> INSERT INTO users VALUES (1, 'Alice');
kvsql> SELECT * FROM users;
kvsql> exit
```

#### Server Mode (HTTP API)
```bash
./sqld -mode=server -port=8080
```

This starts an HTTP server with REST API endpoints:

**Execute SQL Query:**
```bash
curl -X POST http://localhost:8080/api/v1/query \
  -H "Content-Type: application/json" \
  -d '{"sql": "SELECT * FROM users WHERE id = 1"}'
```

**List Tables:**
```bash
curl http://localhost:8080/api/v1/tables
```

**Health Check:**
```bash
curl http://localhost:8080/api/v1/health
```

### Command Line Options
- `-mode`: Operation mode (`cli` or `server`)
- `-port`: Port for server mode (default: `8080`)

