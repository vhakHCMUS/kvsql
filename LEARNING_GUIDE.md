## KVSQL Learning TODO List

Đây là danh sách các tính năng cần implement theo thứ tự ưu tiên, được thiết kế để học từ cơ bản đến nâng cao.

### 🏁 Phase 1: Foundation (Cơ bản nhất)

#### 1.1 Domain Models (`internal/domain/model.go`)
- [ ] Implement `Value` type với type conversion
- [ ] Implement `Row` operations (get/set values)
- [ ] Implement `Table` schema validation
- [ ] Add data type validation functions

#### 1.2 Memory Storage (`internal/infra/storage/memory.go`) 
**Bắt đầu từ đây - dễ nhất để học:**
- [ ] `CreateTable()` - tạo table trong memory
- [ ] `Insert()` - thêm row vào table
- [ ] `Get()` - lấy row theo ID
- [ ] `Scan()` - duyệt tất cả rows
- [ ] `ListTables()` - liệt kê tables

#### 1.3 Basic Parser (`internal/usecase/parser.go`)
**Chỉ implement những SQL đơn giản trước:**
- [ ] `ParseInsert()` - parse "INSERT INTO table VALUES (...)"
- [ ] `ParseSelect()` - parse "SELECT * FROM table"
- [ ] `Tokenize()` - tách SQL thành tokens cơ bản

#### 1.4 CLI Interface (`internal/delivery/cli.go`)
- [ ] `executeSQL()` - kết nối parser + storage
- [ ] `displayResults()` - hiển thị kết quả dạng table
- [ ] `listTables()` - show tất cả tables

### 🎯 Phase 2: Core Features (Tính năng cốt lõi)

#### 2.1 Query Executor (`internal/usecase/executor.go`)
- [ ] `executeTableScan()` - scan toàn bộ table
- [ ] `executeFilter()` - filter rows theo condition
- [ ] `executeInsert()` - thực hiện insert
- [ ] `executeProjection()` - chọn columns cụ thể

#### 2.2 Query Planner (`internal/usecase/planner.go`)
- [ ] `planSelect()` - tạo plan cho SELECT
- [ ] `planInsert()` - tạo plan cho INSERT
- [ ] Basic cost estimation

#### 2.3 WHERE Clause Support
- [ ] Parse WHERE conditions
- [ ] Support operators: =, !=, <, >, LIKE
- [ ] Support AND/OR logic

### 🚀 Phase 3: Indexing (Cải thiện performance)

#### 3.1 B-Tree Index (`internal/infra/index/btree.go`)
**Học cấu trúc dữ liệu quan trọng nhất:**
- [ ] `Insert()` - thêm key vào B-tree
- [ ] `Search()` - tìm key trong B-tree
- [ ] `split()` - chia node khi đầy
- [ ] `Range()` - range queries

#### 3.2 Index Integration
- [ ] Automatic index updates on INSERT/UPDATE/DELETE
- [ ] Query planner sử dụng indexes
- [ ] CREATE INDEX command

### 💾 Phase 4: Persistence (Lưu trữ lâu dài)

#### 4.1 File Storage (`internal/infra/storage/file_store.go`)
- [ ] Basic file I/O operations
- [ ] Page-based storage
- [ ] Buffer pool management

#### 4.2 Write-Ahead Log (`internal/infra/wal/wal.go`)
- [ ] `WriteLog()` - ghi transaction log
- [ ] `ReadLogs()` - đọc logs để recovery
- [ ] Crash recovery mechanism

### 🔒 Phase 5: Transactions (ACID Properties)

#### 5.1 Basic Transactions (`internal/ports/tx.go`)
- [ ] `BeginTransaction()`
- [ ] `Commit()`
- [ ] `Rollback()`

#### 5.2 Locking & Concurrency
- [ ] Row-level locking
- [ ] Deadlock detection
- [ ] Transaction isolation levels

### 🌐 Phase 6: HTTP API (`internal/delivery/server.go`)

- [ ] `handleQuery()` - execute SQL via HTTP
- [ ] `handleListTables()` - REST endpoint for tables
- [ ] JSON request/response handling
- [ ] Error handling & status codes

### 🚀 Phase 7: Advanced Features

#### 7.1 Vector Search (`internal/infra/index/hnsw_stub.go`)
- [ ] HNSW algorithm implementation
- [ ] Vector similarity functions
- [ ] Integration with SQL (SELECT ... ORDER BY vector_distance())

#### 7.2 Query Optimization
- [ ] Join operations
- [ ] Subqueries
- [ ] Advanced WHERE conditions (IN, EXISTS, etc.)

---

## 🎯 Lộ trình học tập được khuyến nghị:

### Tuần 1-2: Foundation
Bắt đầu với Memory Storage - dễ debug, không phải lo về file I/O.
Implement CREATE TABLE, INSERT, SELECT đơn giản.

### Tuần 3-4: Parser & Executor  
Học cách parse SQL và execute basic queries.
Focus on SELECT với WHERE conditions đơn giản.

### Tuần 5-6: Indexing
Implement B-Tree - cấu trúc dữ liệu cực kỳ quan trọng trong database.
Học về query optimization với indexes.

### Tuần 7-8: File Storage
Chuyển từ memory sang disk storage.
Học về buffer management và caching.

### Tuần 9-10: Transactions
Implement ACID properties, locking mechanisms.

### Tuần 11-12: Advanced Features
HTTP API, vector search, query optimization.

---

## 💡 Tips cho việc implement:

1. **Bắt đầu simple nhất có thể** - đừng optimize sớm
2. **Viết test cho mỗi function** - database code rất dễ bug
3. **Log mọi thứ** - debug database rất khó
4. **Đọc code của database khác** - SQLite, PostgreSQL source
5. **Benchmark performance** - measure before optimize

## 📚 Tài liệu tham khảo:
- "Database System Concepts" - Silberschatz (Bible của database)
- "Architecture of a Database System" - paper hay nhất về kiến trúc DB
- SQLite source code - database đơn giản nhất để học
- PostgreSQL documentation - document tốt nhất

**Chúc bạn học tốt! 🚀**