package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {

	// Ping the database
	// It returns an error if the connection is not made
	Ping() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service

	ErrEmailAlreadyExists = errors.New("a user with this email already exists")
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Ping() map[string]string {
	resp := make(map[string]string)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	resp["message"] = err.Error()
	if err != nil {
		log.Panicf("Error while ping data: %s", err)
	}

	resp["message"] = "Database Connected Successfully"

	return resp
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}