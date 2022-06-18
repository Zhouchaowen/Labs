package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const (
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func Select() {
	var user1 User
	// SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user1)
	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", user1.Email, user1.Password)

	var user2 User
	// SELECT * FROM users LIMIT 1;
	db.Take(&user2)
	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", user2.Email, user2.Password)

	var user3 User
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user3)
	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", user3.Email, user3.Password)

}

func main() {
	db, err = gorm.Open(postgres.Open(DB_DSN), &gorm.Config{})

	Select()
}
