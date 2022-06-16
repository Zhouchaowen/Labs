package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

type User struct {
	ID       int
	Email    string
	Password string
}

func Select() {
	// Create an empty user and make the sql query (using $1 for the parameter)
	var myUser User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	err := db.QueryRow(userSql, 1).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)
}

func Insert() {
	//创建一个用户，预要插入到数据库里
	var user User = User{ID: 4, Email: "110@qq.com", Password: "1234567890"}
	//执行插入操作
	_, err := db.Exec("INSERT INTO users (id,email,password) VALUES($1,$2,$3)",
		user.ID, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	//打印日志
	log.Printf("create ok!!!")

	//测试数据是否插入成功，执行具体的查询语句
	var myUser User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	//设置查询参数为4，即创建数据时的ID值
	err = db.QueryRow(userSql, 4).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	//输出查询结果
	fmt.Printf("Insert: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)
}

func Update() {
	//创建一个用户，预要通过主键更改到数据库里
	var user User = User{ID: 4, Email: "dong@qq.com", Password: "abcdedf120"}
	//执行更改操作
	_, err := db.Exec("UPDATE  users SET email=$1, password=$2 where id=$3", user.Email, user.Password, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	//打印日志
	log.Printf("update ok!!!")

	//测试数据是否更改成功，执行具体的查询语句
	var myUser User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	//设置查询参数为4，即要更改数据的ID值
	err = db.QueryRow(userSql, 4).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	//输出查询结果
	fmt.Printf("Update: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)
}

func Delete() {
	//执行更改操作
	_, err := db.Exec("DELETE FROM  users  where id=$1", 4)
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
	fmt.Printf("Delete: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)
}

func main() {
	// Create DB pool
	//db, err := sql.Open("postgres", "host=192.168.8.200 port=5432 user=postgres password=12345678 dbname=douyin sslmode=disable")
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	Select()
	Insert()
	Update()
	Delete()
}
