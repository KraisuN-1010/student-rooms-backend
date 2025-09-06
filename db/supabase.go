package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is the global database connection pool
var DB *pgxpool.Pool

// Connect initializes the connection to Supabase/Postgres using pgx
func Connect() error {
	dbURL := os.Getenv("SUPABASE_DB_URL")
	if dbURL == "" {
		log.Fatal("SUPABASE_DB_URL is not set in environment")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return err
	}

	// Optional: set max connections and timeout
	config.MaxConns = 10
	config.MaxConnLifetime = time.Hour

	// Create a connection pool
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = DB.Ping(ctx)
	if err != nil {
		return err
	}

	log.Println("✅ Connected to Supabase/Postgres successfully using pgx")
	return nil
}

// Close closes the database connection pool
func Close() {
	if DB != nil {
		DB.Close()
		log.Println("🔒 Database connection closed")
	}
}
