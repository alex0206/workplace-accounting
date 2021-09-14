package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DB is a shortcut structure to a Postgres DB
type DB struct {
	Conn *pgxpool.Pool
}

// Close close db connection
func (c *DB) Close() {
	c.Conn.Close()
}

// NewDBConnection creates new database connection to postgres
func NewDBConnection(ctx context.Context, dsn string) (*DB, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &DB{Conn: pool}, nil
}
