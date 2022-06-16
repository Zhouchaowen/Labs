package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
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

func PrepareSelectExample() {
	var myUser User

	stmt, err := db.Prepare("SELECT id, email, password FROM users WHERE id = $1")
	if err != nil {
		log.Fatal("Failed to Prepare query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal("Failed to Prepare query: ", err)
	}

	// 获取返回查询列
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("rows.Columns error: %v\n", err)
	}
	fmt.Printf("columns: %v\n", columns)

	// 获取返回查询列对应属性
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatalf("rows.ColumnTypes error: %v\n", err)
	}
	for _, v := range columnTypes {
		fmt.Printf("types: %v\n", v)
	}

	// 获取返回查询列的值
	for rows.Next() {
		err := rows.Scan(&myUser.ID, &myUser.Email, &myUser.Password)
		if err != nil {
			log.Fatalf("rows.Scan error: %v\n", err)
		}
		fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)
	}
}

func PrepareDeleteExample() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	stmt, err := db.Prepare("DELETE FROM users  where id=$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, user := range users {
		if result, err := stmt.Exec(user.ID); err != nil {
			log.Fatal(err)
		} else {
			rowsAffected, _ := result.RowsAffected()
			lastInsertId, _ := result.LastInsertId()
			fmt.Printf("rowsAffected:%d, lastInsertId：%d\n", rowsAffected, lastInsertId)
		}
	}

}

func PrepareInsertExample() {
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

// 重复利用 Prepare 后的 stmt 语句，可以有效提高 SQL 执行性能
func PrepareBenchmark() {
	var num int
	start := time.Now()
	for i := 0; i < 100; i++ {
		err = db.QueryRow("select count(*) from pg_stat_activity where datname = $1", "postgres").Scan(&num)
		if err != nil {
			log.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("got: %d in %s\n", num, elapsed)

	stmt, err := db.Prepare("select count(*) from pg_stat_activity where datname = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	start = time.Now()
	for i := 0; i < 100; i++ {
		err = stmt.QueryRow("postgres").Scan(&num)
	}
	elapsed = time.Since(start)
	fmt.Printf("got: %d in %s\n", num, elapsed)

}

func main() {
	// Create DB pool
	//db, err := sql.Open("postgres", "host=192.168.8.200 port=5432 user=postgres password=12345678 dbname=douyin sslmode=disable")
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	//PrepareSelectExample()
	PrepareDeleteExample()
	//PrepareInsertExample()
	//PrepareBenchmark()
}
