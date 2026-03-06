package repository

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

func NewDB() (*pgx.Conn, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "123"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "tasktracker"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	return pgx.Connect(context.Background(), dsn)
}

func RunMigrations(db *pgx.Conn) error {
	migrationsDir := "migrations"

	return filepath.WalkDir(migrationsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".sql" {
			return nil
		}

		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = db.Exec(context.Background(), string(sqlBytes))
		if err != nil {
			log.Printf("Migration failed: %s\n", path)
			return err
		}

		log.Printf("Migration applied: %s\n", path)
		return nil
	})
}
