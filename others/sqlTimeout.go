package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var (
	db *sql.DB
)

func main() {
	age := 27
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	rows, err := db.QueryContext(ctx, "SELECT NAME FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //千万要记住关闭，将连接返回给进程池
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
}
