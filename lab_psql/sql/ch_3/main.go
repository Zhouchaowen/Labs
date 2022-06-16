package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

type User struct {
	ID       int
	Email    string
	Password string
}

func TxSelectExample() {
	var myUser User

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin tx: ", err)
	}

	err = tx.QueryRow("SELECT id, email, password FROM users WHERE id = $1", 1).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to Prepare query: ", err)
		tx.Rollback()
	}
	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)
	tx.Commit()
}

func TxDeleteExample() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin tx: ", err)
	}

	stmt, err := tx.Prepare("DELETE FROM users  where id=$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, user := range users {
		if result, err := stmt.Exec(user.ID); err != nil {
			log.Fatal(err)
			tx.Rollback()
		} else {
			rowsAffected, _ := result.RowsAffected()
			lastInsertId, _ := result.LastInsertId()
			fmt.Printf("rowsAffected:%d, lastInsertId：%d\n", rowsAffected, lastInsertId)
		}
	}
	tx.Commit()
}

func TxInsertExample() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	stmt, err := db.Prepare("INSERT INTO users (id,email,password) VALUES($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, user := range users {
		if result, err := stmt.Exec(user.ID, user.Email, user.Password); err != nil {
			log.Fatal(err)
		} else {
			rowsAffected, _ := result.RowsAffected()
			lastInsertId, _ := result.LastInsertId()
			fmt.Printf("rowsAffected:%d, lastInsertId：%d\n", rowsAffected, lastInsertId)
		}
	}

}

func main() {
	// Create DB pool
	//db, err := sql.Open("postgres", "host=192.168.8.200 port=5432 user=postgres password=12345678 dbname=douyin sslmode=disable")
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	TxSelectExample()

}
