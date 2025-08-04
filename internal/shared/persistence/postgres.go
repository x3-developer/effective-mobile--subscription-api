package persistence

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"subscriptions/config"
	"time"
)

type Postgres struct {
	*sql.DB
}

func NewPostgres(cfg *config.Config) *Postgres {
	dsn := GetDSN(cfg)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &Postgres{db}
}

func (p *Postgres) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := p.DB.ExecContext(ctx, query, args...)
	duration := time.Since(start)

	log.Printf("[SQL EXEC] %s | args: %v | duration: %v | err: %v", query, args, duration, err)

	return result, err
}

func (p *Postgres) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := p.DB.QueryContext(ctx, query, args...)
	duration := time.Since(start)

	log.Printf("[SQL QUERY] %s | args: %v | duration: %v | err: %v", query, args, duration, err)

	return rows, err
}

func (p *Postgres) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	start := time.Now()
	log.Printf("[SQL ROW] %s | args: %v | started at: %v", query, args, start)
	return p.DB.QueryRowContext(ctx, query, args...)
}

func GetDSN(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbSsl,
	)
}
