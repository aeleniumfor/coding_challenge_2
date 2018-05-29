package main

import (
	"database/sql"
	_ "github.com/lib/pq" //read
	"time"
	"log"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	conn := "postgres://root:root@/150.95.134.161?sslmode=verify-full"
	db, err := sql.Open("postgres", conn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
