package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
}

func main() {
	test3()
}

func test3() {
	res, err := Db.Exec("update person set username = ? where user_id = ?", "stu002", 2)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("failed", err)
	}
	fmt.Println("update succ", row)

}

func test2() {
	var person []Person
	err := Db.Select(&person, "select * from person where user_id = ?", 2)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("select success", person)
}

func test1() {
	r, err := Db.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("insert succ", id)
}
