package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	//_ "github.com/bmizerany/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

func ContextSelectExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// If the context is canceled or timed out, the query execution will be stopped.
	// If the query is INSERT or UPDATE you can use function ExecContext.
	_, err = db.QueryContext(ctx, "SELECT pg_sleep(15)")
	if err != nil {
		log.Fatal("query context err: ", err)
	}
	//打印日志
	log.Printf("delete ok!!!")
}

func init() {
	// Create DB pool
	// db, err := sql.Open("postgres", "host=10.2.0.104 port=5432 user=postgres password=lab_password dbname=postgres sslmode=disable")
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
}

func main() {
	defer db.Close()

	ContextSelectExample()
}
