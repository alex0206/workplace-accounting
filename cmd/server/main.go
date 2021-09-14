package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alex0206/workplace-accounting/config"
	"github.com/alex0206/workplace-accounting/internal/pg"
	"github.com/alex0206/workplace-accounting/internal/server"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var cfg *config.Config

func init() {
	cfg = config.Get()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if cfg.Port == "" {
		return errors.New("missing PORT in config")
	}

	// connect to Postgres
	pgDB, err := pg.NewDBConnection(context.Background(), cfg.PgURL)
	if err != nil {
		return fmt.Errorf("pgdb.Connect failed: %w", err)
	}
	defer pgDB.Close()

	// run Postgres migrations
	log.Println("Running PostgreSQL migrations")
	if err := runPgMigrations(); err != nil {
		return fmt.Errorf("runPgMigrations failed: %w", err)
	}

	srv := &http.Server{
		Addr:         net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:      server.NewAPIRouter(pgDB),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Channel to listen for errors
	errorChannel := make(chan error)

	go func() {
		log.Printf("Running HTTP server on %s\n", srv.Addr)
		errorChannel <- srv.ListenAndServe()
	}()

	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s", <-s)
	}()

	if err := <-errorChannel; err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("server Shutdown Failed:%+v", err)
		}
		log.Print("Server exited properly")
	}

	return nil
}

// runPgMigrations runs Postgres migrations
func runPgMigrations() error {
	if cfg.PgMigrationsPath == "" {
		return nil
	}

	if cfg.PgURL == "" {
		return errors.New("no cfg.PgURL provided")
	}

	m, err := migrate.New(
		cfg.PgMigrationsPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
