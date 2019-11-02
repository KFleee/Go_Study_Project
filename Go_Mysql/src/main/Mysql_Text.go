package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, _ := sql.Open("mysql", "root:LJH787807080886@/springt")
	result, _ := db.Exec(
    "INSERT INTO user (name, password) VALUES (?, ?)",
    "gopher",
    "sasd",
	)
	fmt.Printf("%v", result)
}