package main

import (
	"flag"
	"fmt"
	"log"

	"kvsql/internal/delivery"
	"kvsql/internal/usecase"
)

func main() {
	// Command line flags
	var (
		mode = flag.String("mode", "cli", "Mode to run: cli or server")
		port = flag.String("port", "8080", "Port for server mode")
	)
	flag.Parse()

	// TODO: Initialize logging
	// Should set up structured logging with different levels
	log.Println("Starting KVSQL Database...")

	// TODO: Initialize storage layer
	// For now, we'll skip storage initialization as it's not fully implemented
	log.Println("Storage layer initialization skipped - implement when needed")

	// TODO: Initialize components as they get implemented:
	// - Storage (memory or file-based)
	// - Index manager
	// - Transaction manager
	// - WAL for durability

	// TODO: Initialize index manager
	// indexManager := index.NewIndexManager()

	// TODO: Initialize transaction manager
	// txManager := tx.NewTransactionManager()

	// Initialize use cases
	parser := usecase.NewParser()
	planner := usecase.NewPlanner(nil, nil)        // TODO: Pass actual storage and index manager
	executor := usecase.NewExecutor(nil, nil, nil) // TODO: Pass actual dependencies

	// Run in specified mode
	switch *mode {
	case "cli":
		runCLI(parser, planner, executor)
	case "server":
		runServer(parser, planner, executor, *port)
	default:
		log.Fatalf("Unknown mode: %s. Use 'cli' or 'server'", *mode)
	}
}

// runCLI starts the interactive CLI
func runCLI(parser *usecase.Parser, planner *usecase.Planner, executor *usecase.Executor) {
	cli := delivery.NewCLI(parser, planner, executor)

	if err := cli.Start(); err != nil {
		log.Fatalf("CLI error: %v", err)
	}
}

// runServer starts the HTTP server
func runServer(parser *usecase.Parser, planner *usecase.Planner, executor *usecase.Executor, port string) {
	server := delivery.NewServer(parser, planner, executor, port)

	fmt.Printf("Starting HTTP server on port %s...\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST /api/v1/query   - Execute SQL queries")
	fmt.Println("  GET  /api/v1/tables  - List tables")
	fmt.Println("  GET  /api/v1/health  - Health check")

	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// TODO: Add additional utility functions

// setupLogging configures structured logging
func setupLogging(level string) {
	// TODO: Configure logging with different levels
	// Should support: DEBUG, INFO, WARN, ERROR
	// Should include timestamps and source locations
}

// loadConfig loads configuration from file or environment
func loadConfig(configPath string) error {
	// TODO: Load configuration from file
	// Should support:
	// - Database settings
	// - Storage configuration
	// - Security settings
	// - Performance tuning
	return nil
}

// gracefulShutdown handles graceful shutdown
func gracefulShutdown() {
	// TODO: Implement graceful shutdown
	// Should:
	// - Close database connections
	// - Flush pending writes
	// - Save state
	// - Clean up resources
}
