package delivery

import (
	"bufio"
	"errors"
	"fmt"
	"kvsql/internal/usecase"
	"os"
	"strings"
)

// CLI provides command-line interface for the database
type CLI struct {
	parser   *usecase.Parser
	planner  *usecase.Planner
	executor *usecase.Executor
	// TODO: Add CLI components
	// - History management
	// - Auto-completion
	// - Command aliases
}

// NewCLI creates a new CLI interface
func NewCLI(parser *usecase.Parser, planner *usecase.Planner, executor *usecase.Executor) *CLI {
	return &CLI{
		parser:   parser,
		planner:  planner,
		executor: executor,
	}
}

// Start starts the interactive CLI
func (cli *CLI) Start() error {
	fmt.Println("Welcome to KVSQL Database")
	fmt.Println("Type 'help' for available commands, 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("kvsql> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if err := cli.processCommand(input); err != nil {
			if err.Error() == "exit" {
				break
			}
			fmt.Printf("Error: %v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	fmt.Println("Goodbye!")
	return nil
}

// processCommand processes a single command
func (cli *CLI) processCommand(input string) error {
	// Handle special commands
	switch strings.ToLower(input) {
	case "exit", "quit", "\\q":
		return errors.New("exit")
	case "help", "\\h":
		cli.showHelp()
		return nil
	case "\\dt": // List tables (PostgreSQL style)
		return cli.listTables()
	}

	// Check if it's a SQL command
	if cli.isSQLCommand(input) {
		return cli.executeSQL(input)
	}

	return fmt.Errorf("unknown command: %s. Type 'help' for available commands", input)
}

// isSQLCommand checks if the input is a SQL command
func (cli *CLI) isSQLCommand(input string) bool {
	input = strings.ToUpper(strings.TrimSpace(input))
	sqlKeywords := []string{"SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER"}

	for _, keyword := range sqlKeywords {
		if strings.HasPrefix(input, keyword) {
			return true
		}
	}

	return false
}

// executeSQL executes a SQL command
func (cli *CLI) executeSQL(sql string) error {
	// TODO: Implement SQL execution
	// Should:
	// - Parse the SQL using parser
	// - Create execution plan using planner
	// - Execute using executor
	// - Display results

	fmt.Printf("Executing SQL: %s\n", sql)

	// TODO: Parse SQL
	stmt, err := cli.parser.Parse(sql)
	if err != nil {
		return fmt.Errorf("parsing SQL: %w", err)
	}

	// TODO: Create execution plan
	plan, err := cli.planner.Plan(stmt)
	if err != nil {
		return fmt.Errorf("creating plan: %w", err)
	}

	// TODO: Execute plan
	result, err := cli.executor.Execute(plan, nil)
	if err != nil {
		return fmt.Errorf("executing query: %w", err)
	}

	// TODO: Display results
	cli.displayResults(result)

	return errors.New("TODO: implement SQL execution in CLI")
}

// displayResults displays query results
func (cli *CLI) displayResults(result interface{}) {
	// TODO: Implement result display
	// Should:
	// - Format results in table format
	// - Handle different result types
	// - Show row counts
	// - Handle large result sets with pagination

	fmt.Println("TODO: implement result display")
}

// listTables lists all tables
func (cli *CLI) listTables() error {
	// TODO: Implement table listing
	// Should:
	// - Query storage for table names
	// - Display in formatted table
	// - Show additional metadata if available

	fmt.Println("TODO: implement table listing")
	return nil
}

// showHelp displays help information
func (cli *CLI) showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  SQL Commands:")
	fmt.Println("    SELECT * FROM table_name;           - Query data")
	fmt.Println("    INSERT INTO table (col) VALUES (val); - Insert data")
	fmt.Println("    UPDATE table SET col=val WHERE ...;  - Update data")
	fmt.Println("    DELETE FROM table WHERE ...;         - Delete data")
	fmt.Println("    CREATE TABLE name (columns...);      - Create table")
	fmt.Println("    DROP TABLE name;                     - Delete table")
	fmt.Println()
	fmt.Println("  Meta Commands:")
	fmt.Println("    \\dt                                  - List tables")
	fmt.Println("    \\d table_name                       - Describe table")
	fmt.Println("    help, \\h                            - Show this help")
	fmt.Println("    exit, quit, \\q                      - Exit the program")
	fmt.Println()
	fmt.Println("  Note: SQL commands should end with semicolon (;)")
}

// ExecuteFile executes SQL commands from a file
func (cli *CLI) ExecuteFile(filename string) error {
	// TODO: Implement file execution
	// Should:
	// - Read SQL commands from file
	// - Handle multiple statements
	// - Show progress for large files
	// - Support transaction blocks

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("opening file %s: %w", filename, err)
	}
	defer file.Close()

	fmt.Printf("Executing commands from %s...\n", filename)

	// TODO: Read and execute commands
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}

		// TODO: Handle multi-line statements
		// TODO: Execute each statement
		fmt.Printf("Line %d: %s\n", lineNum, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	return errors.New("TODO: implement file execution")
}
