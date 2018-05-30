package main

import (
	"database/sql"
	_ "github.com/lib/pq" //read
	"time"
	"fmt"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_HOST     = "150.95.134.161"
	DB_NAME     = "webapp"
	DB_SSLMODE  = "disable"
)

func checkErr(err error) {
	//エラーは全部これに集める
	if err != nil {
		panic(err)
	}
}

func DB_connect() *sql.DB {
	//dbへのコネクション関数
	conn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s ", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME, DB_SSLMODE)
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	return db
}

func DB_select() []User {
	db := DB_connect()
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM users`)
	checkErr(err)
	users := User{}
	member := []User{}
	for rows.Next() {

		rows.Scan(&users.ID, &users.Name, &users.Email, &users.CreatedAt, &users.UpdatedAt)
		checkErr(err)
		member = append(member, users)
	}
	return member
}

func DB_select_id(id string) []User {
	db := DB_connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", id)
	rows, err := db.Query(query)
	checkErr(err)
	users := User{}
	member := []User{}
	for rows.Next() {

		rows.Scan(&users.ID, &users.Name, &users.Email, &users.CreatedAt, &users.UpdatedAt)
		checkErr(err)
		member = append(member, users)
	}
	return member
}

func DB_insert() {
	db := DB_connect()
	defer db.Close()
	query, err := db.Prepare("INSERT INTO users(Name,Email) VALUES($1,$2)")
	data, err := query.Exec("test1", "test1.gmail")
	checkErr(err)
	fmt.Println(data)
}

func main() {
	DB_insert()
	fmt.Println(DB_select())
	fmt.Println(DB_select_id("2"))
}
