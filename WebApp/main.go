package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"log"
	"github.com/gorilla/mux"
	"coding_challenge_2/WebApp/p_db"
)

type Hello struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)                                  //GET
	r.HandleFunc("/users", users_handler).Methods("GET")        //GET
	r.HandleFunc("/users/{id}", userid_handler).Methods("POST") //GET
	r.HandleFunc("/users ", handler)                            //POST
	r.HandleFunc("/users/{id} ", handler)                       //PUT
	r.HandleFunc("/users/{id} ", handler)                       //DELETE

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil) //ポート8080で待機
}

//ハンドラ関数定義
func handler(w http.ResponseWriter, r *http.Request) {

	msg := Hello{"Hello World!!"}
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&msg)

	if err != nil { //エラー処理
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}

func users_handler(w http.ResponseWriter, r *http.Request) {
	member := p_db.DB_select()
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&member)

	if err != nil { //エラー処理
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}

func userid_handler(w http.ResponseWriter, r *http.Request) {
	member := p_db.DB_select_id("1")
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&member)

	if err != nil { //エラー処理
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}
