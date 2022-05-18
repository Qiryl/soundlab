package xpgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPGXPool is a PostgreSQL connection pool for pgx.
//
// Usage:
// pgPool := database.NewPGXPool(context.Background(), "", &PGXStdLogger{}, pgx.LogLevelInfo)
// defer pgPool.Close() // Close any remaining connections before shutting down your application.
//
// Instead of passing a configuration explicitly with a connString,
// you might use PG environment variables such as the following to configure the database:
// PGDATABASE, PGHOST, PGPORT, PGUSER, PGPASSWORD, PGCONNECT_TIMEOUT, etc.
// Reference: https://www.postgresql.org/docs/current/libpq-envars.html
func NewPGXPool(ctx context.Context, connString string, logger pgx.Logger, logLevel pgx.LogLevel) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(connString) // Using environment variables instead of a connection string.
	if err != nil {
		return nil, err
	}

	conf.ConnConfig.Logger = logger

	// Set the log level for pgx, if set.
	if logLevel != 0 {
		conf.ConnConfig.LogLevel = logLevel
	}

	// pgx, by default, does some I/O operation on initialization of a pool to check if the database is reachable.
	// Comment the following line if you don't want pgx to try to connect pool once the Connect function is called,
	//
	// If comment it, and your application seems stuck, you probably forgot to set up PGCONNECT_TIMEOUT,
	// and your code is hanging waiting for a connection to be established.
	conf.LazyConnect = true

	// pgxpool default max number of connections is the number of CPUs on your machine returned by runtime.NumCPU().
	// This number is very conservative, and you might be able to improve performance for highly concurrent applications
	// by increasing it.
	// conf.MaxConns = runtime.NumCPU() * 5

	pool, err := pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}
	return pool, nil
}
