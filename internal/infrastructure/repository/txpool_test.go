package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/hatch-studio/pgtools/sqltest"
)

var mainURL = "postgresql://postgres:querty@localhost:5432/postgres?sslmode=disable&search_path=public"

func TestMain(m *testing.M) {
	if flag := os.Getenv("INTEGRATION_TESTDB"); flag != "true" {
		log.Printf("Skipping tests that require database connection, cuz flag: %s", flag)
		return
	}
	os.Exit(m.Run())
}

func TestTransactionContext(t *testing.T) {
	t.Parallel()

	migration := sqltest.New(t, sqltest.Options{
		Path: "../../../migrations",
	})
	pool := migration.Setup(context.Background(), mainURL)
	repo := &txPool{
		Pool: pool,
	}

	ctx, err := repo.TransactionContext(context.Background())
	if err != nil {
		t.Errorf("cannot create transaction context: %v", err)
	}
	defer repo.Rollback(ctx)

	if err := repo.Commit(ctx); err != nil {
		t.Errorf("cannot commit: %v", err)
	}
}

func TestTransactionContextCanceled(t *testing.T) {
	t.Parallel()

	migration := sqltest.New(t, sqltest.Options{
		Path: "../../../migrations",
	})

	pool := migration.Setup(context.Background(), mainURL)
	repo := &txPool{
		Pool: pool,
	}

	canceledCtx, immediateCancel := context.WithCancel(context.Background())
	immediateCancel()

	if _, err := repo.TransactionContext(canceledCtx); err != context.Canceled {
		t.Errorf("unexpected error value: %v", err)
	}
}

func TestCommitNoTransaction(t *testing.T) {
	t.Parallel()

	db := &txPool{}
	if err := db.Commit(context.Background()); err.Error() != "context has no transaction" {
		t.Errorf("unexpected error value: %v", err)
	}
}

func TestRollbackNoTransaction(t *testing.T) {
	t.Parallel()

	db := &txPool{}
	if err := db.Rollback(context.Background()); err.Error() != "context has no transaction" {
		t.Errorf("unexpected error value: %v", err)
	}
}

func TestWithAcquire(t *testing.T) {
	t.Parallel()
	migration := sqltest.New(t, sqltest.Options{
		Path: "../../../migrations",
	})

	pool := migration.Setup(context.Background(), mainURL)
	db := &txPool{
		Pool: pool,
	}

	// Reuse the same connection for executing SQL commands.
	dbCtx, err := db.WithAcquire(context.Background())
	if err != nil {
		t.Fatalf("unexpected DB.WithAcquire() error = %v", err)
	}
	defer db.Release(dbCtx)

	// Check if we can acquire a connection only for a given context.
	defer func() {
		want := "context already has a connection acquired"
		if r := recover(); r != want {
			t.Errorf("expected panic %v, got %v instead", want, r)
		}
	}()
	db.WithAcquire(dbCtx)
}

func TestWithAcquireClosedPool(t *testing.T) {
	t.Parallel()
	migration := sqltest.New(t, sqltest.Options{
		// Opt out of automatic tearing down migration as we want to close the connection pool before t.Cleanup() is called.
		SkipTeardown: true,

		Path: "../../../migrations",
	})
	pool := migration.Setup(context.Background(), mainURL)
	db := &txPool{
		Pool: pool,
	}
	migration.Teardown(context.Background())
	if _, err := db.WithAcquire(context.Background()); err == nil {
		t.Errorf("expected error acquiring pgx connection for context, got nil")
	}
}
