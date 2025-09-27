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

## 📚 Learning Path

This project is designed as a learning resource. Here's a suggested implementation order:

### Phase 1: Basic Foundation
1. **Domain Models** - Complete the data structures in `domain/model.go`
2. **Memory Storage** - Implement basic CRUD operations in `internal/infra/storage/memory.go`
3. **Simple Parser** - Start with basic SELECT/INSERT parsing in `usecase/parser.go`
4. **CLI Interface** - Get basic SQL execution working in `delivery/cli.go`

### Phase 2: Core Database Features
1. **B-Tree Index** - Implement indexing in `infra/index/btree.go`
2. **Query Planner** - Add query optimization in `usecase/planner.go`
3. **Query Executor** - Implement execution engine in `usecase/executor.go`
4. **File Storage** - Add persistence in `infra/storage/file_store.go`

### Phase 3: Advanced Features
1. **Transactions** - Implement ACID transactions
2. **WAL** - Add Write-Ahead Logging for durability
3. **Concurrency** - Multi-user support with proper locking
4. **Vector Search** - Complete HNSW implementation

### Phase 4: Production Features
1. **HTTP API** - Complete REST interface in `delivery/server.go`
2. **Query Optimization** - Advanced planner features
3. **Monitoring** - Add metrics and logging
4. **Performance** - Optimization and benchmarking

## 🧑‍💻 Development Guide

### Project Structure
- All placeholder methods throw errors with "TODO:" messages
- Interfaces are defined in `ports/` package
- Business logic goes in `usecase/` package
- Infrastructure implementations go in `infra/` package

### Adding New Features
1. Start with the interface definition in `ports/`
2. Add domain models in `domain/` if needed
3. Implement business logic in `usecase/`
4. Create infrastructure implementation in `infra/`
5. Wire everything together in `cmd/sqld/main.go`

### Testing
Each package should have comprehensive tests:
```bash
go test ./...
```

## 🎯 Current Status

This is a **skeleton implementation** with:
- ✅ Complete project structure
- ✅ All interfaces defined
- ✅ Placeholder implementations
- ❌ **No actual logic implemented yet**

All methods currently return `"TODO: implement ..."` errors. This is intentional - the goal is to provide a complete framework for learning database implementation.

## 📖 Learning Resources

To implement the TODOs in this project, study these topics:
- **Database Systems Concepts** by Silberschatz, Galvin, and Gagne
- **Architecture of a Database System** by Hellerstein, Stonebraker, and Hamilton
- **Designing Data-Intensive Applications** by Martin Kleppmann

## 🤝 Contributing

This is a learning project! Feel free to:
- Implement any of the TODO methods
- Add tests
- Improve documentation
- Share your learning experience

## 📄 License

MIT License - feel free to use this for learning and education.

---

**Happy Learning! 🚀**

*Remember: The best way to understand databases is to build one yourself.*