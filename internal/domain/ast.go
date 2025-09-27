package domain

// AST (Abstract Syntax Tree) nodes for SQL parsing

// Statement represents any SQL statement
type Statement interface {
	StatementType() string
}

// SelectStatement represents a SELECT query
type SelectStatement struct {
	Columns []string
	From    string
	Where   *WhereClause
	OrderBy []OrderByClause
	Limit   *int
	GroupBy []string
	Having  *WhereClause
}

func (s *SelectStatement) StatementType() string {
	return "SELECT"
}

// InsertStatement represents an INSERT statement
type InsertStatement struct {
	Table   string
	Columns []string
	Values  [][]Value
}

func (i *InsertStatement) StatementType() string {
	return "INSERT"
}

// UpdateStatement represents an UPDATE statement
type UpdateStatement struct {
	Table string
	Set   map[string]Value
	Where *WhereClause
}

func (u *UpdateStatement) StatementType() string {
	return "UPDATE"
}

// DeleteStatement represents a DELETE statement
type DeleteStatement struct {
	From  string
	Where *WhereClause
}

func (d *DeleteStatement) StatementType() string {
	return "DELETE"
}

// CreateTableStatement represents a CREATE TABLE statement
type CreateTableStatement struct {
	Name    string
	Schema  map[string]DataType
	Indexes map[string]IndexType
}

func (c *CreateTableStatement) StatementType() string {
	return "CREATE_TABLE"
}

// WhereClause represents a WHERE condition
type WhereClause struct {
	Left     interface{} // can be column name or another clause
	Operator string      // =, !=, <, >, LIKE, IN, etc.
	Right    interface{} // can be value or another clause
	Logic    string      // AND, OR (for combining clauses)
	Next     *WhereClause
}

// OrderByClause represents an ORDER BY clause
type OrderByClause struct {
	Column string
	Desc   bool
}

// Expression represents a general expression
type Expression struct {
	Type     string      // COLUMN, LITERAL, FUNCTION
	Value    interface{} // column name, literal value, or function details
	Children []Expression
}
