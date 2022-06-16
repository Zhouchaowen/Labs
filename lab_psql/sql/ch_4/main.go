package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	//_ "github.com/bmizerany/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func main() {
	// Create DB pool
	//db, err := sql.Open("postgres", "host=192.168.8.200 port=5432 user=postgres password=12345678 dbname=douyin sslmode=disable")
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	//执行更改操作
	_, err = db.Exec("DELETE FROM  users  where id=$1", 4)
	if err != nil {
		log.Fatal(err)
	}
	//打印日志
	log.Printf("delete ok!!!")

	//测试数据是否更改成功，执行具体的查询语句
	var myUser User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	//设置查询参数为4，即要更改数据的ID值
	err = db.QueryRow(userSql, 4).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	//输出查询结果
	fmt.Printf("hello email: %s, password: %s, welcome back!\n", myUser.Email, myUser.Password)
}
