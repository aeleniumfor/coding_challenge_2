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

type User struct {
	Id      string
	Balance uint64
}

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)                                 //GET
	r.HandleFunc("/users", users_handler).Methods("GET")       //GET
	r.HandleFunc("/users/{id}", userid_handler).Methods("GET") //GET
	r.HandleFunc("/users", usercreate_handler).Methods("POST") //POST
	r.HandleFunc("/users/{id} ", handler)                      //PUT
	r.HandleFunc("/users/{id} ", handler)                      //DELETE

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil) //ポート8080で待機
}

//ハンドラ関数定義
func handler(w http.ResponseWriter, r *http.Request) {
	//r.HandleFunc("/", handler)
	msg := Hello{"Hello World!!"}
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&msg)

	if err != nil { //err
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}

func users_handler(w http.ResponseWriter, r *http.Request) {
	//r.HandleFunc("/users", users_handler).Methods("GET") //GET

	member := p_db.DB_select()
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&member)

	if err != nil { //err
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}

func userid_handler(w http.ResponseWriter, r *http.Request) {
	//r.HandleFunc("/users/{id}", userid_handler).Methods("GET") //GET

	parm := mux.Vars(r)["id"] //parmの取得
	member := p_db.DB_select_id(parm)
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&member)

	if err != nil { //err
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}

func usercreate_handler(w http.ResponseWriter, r *http.Request) {

	//encode post json data
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u Request
	err := decoder.Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

	id := p_db.DB_insert(u.Name, u.Email)
	member := p_db.DB_select_id(id)

	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&member)

	if err != nil { //err
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())

	//w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
}

func update_handler(w http.ResponseWriter, r *http.Request) {
	//parm := mux.Vars(r)["id"] //get parm
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u Request
	err := decoder.Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

}
