package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
	//_ "github.com/bmizerany/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Pass the request context to slowQuery(), so it can be used as the parent context.
	err := slowQuery(r.Context())
	if err != nil {
		serverError(w, err)
		return
	}

	fmt.Fprintln(w, "OK")
}

func slowQuery(ctx context.Context) error {
	// Create a new child context with a 5-second timeout, using the
	// provided ctx parameter as the parent.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Pass the child context (the one with the timeout) as the first
	// parameter to ExecContext().
	_, err := db.ExecContext(ctx, "SELECT pg_sleep(10)")
	return err
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %s", err.Error())
	http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
}

// 来源：https://www.alexedwards.net/blog/how-to-manage-database-timeouts-and-cancellations-in-go
func main() {
	// Create DB pool
	// db, err := sql.Open("postgres", "host=10.2.0.104 port=5432 user=postgres password=lab_password dbname=postgres sslmode=disable")
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", exampleHandler)

	log.Print("Listening...")
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatal(err)
	}

	// 启动程序后执行，检验查询超时 curl -i localhost:5000/
}
