package p_db

import (
	"database/sql"
	_ "github.com/lib/pq" //read
	"time"
	"fmt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Request struct {
	DB_NAME           string `json:"name"`
	POSTGRES_PASSWORD string `json:"email"`
}

func checkErr(err error) {
	//エラーは全部これに集める
	if err != nil {
		panic(err)
	}
}

func DB_connect() *sql.DB {
	//dbへのコネクション関数
	conn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s ", "root", "root", "postgres", "webapp", "disable")
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	return db
}

func DB_select() []User {
	// 全権検索
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
	// id検索
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

func DB_insert(name string, email string) string {
	// データベースへの挿入
	db := DB_connect()
	defer db.Close()
	var last_id string
	query := "INSERT INTO users(name,email) VALUES($1,$2) RETURNING id"
	err := db.QueryRow(query, name, email).Scan(&last_id)
	checkErr(err)
	return last_id
}

func DB_update(id string, name string, email string) string {
	db := DB_connect()
	defer db.Close()
	var update_id string
	query := "UPDATE users SET name=$1,email=$2,updated_at=$3 WHERE id=$4 RETURNING id"
	err := db.QueryRow(query, name, email, time.Now(), id).Scan(&update_id)
	checkErr(err)
	return update_id
}

func DB_delete(id string) {
	db := DB_connect()
	defer db.Close()
	query := "DELETE FROM users WHERE id=$1"
	_, err := db.Exec(query, id)
	checkErr(err)
}
