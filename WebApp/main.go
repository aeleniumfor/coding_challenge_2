package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"log"
	"github.com/gorilla/mux"
	"WebApp/p_db"
)

type Hello struct {
	Message string `json:"message"`
}

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//ハンドラ関数定義
func handler(w http.ResponseWriter, r *http.Request) {
	println(fmt.Sprintf("%s : %s", r.Method, r.URL))

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
	println(fmt.Sprintf("%s:%s", r.Method, r.URL))

	member := p_db.DB_select()
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&member)

	if err != nil { //err
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ") //jsonを整形
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}

func usercreate_handler(w http.ResponseWriter, r *http.Request) {
	println(fmt.Sprintf("%s:%s", r.Method, r.URL))
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
	json_string, err := json.Marshal(&member[0])

	if err != nil { //err
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ") //jsonを整形
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())

	//w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
}

func user_handler(w http.ResponseWriter, r *http.Request) {
	println(fmt.Sprintf("%s:%s", r.Method, r.URL))
	parm := mux.Vars(r)["id"] //get parm

	switch r.Method {
	case "PUT":
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var u Request
		err := decoder.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}
		if len(p_db.DB_select_id(parm)) != 0 {
			parm = p_db.DB_update(parm, u.Name, u.Email)
		}

		member := p_db.DB_select_id(parm)
		if len(member) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
			fmt.Fprint(w)
			break
		}

		json_indet := new(bytes.Buffer)
		json_string, err := json.Marshal(&member[0])

		if err != nil { //err
			log.Fatal(err)
		}

		json.Indent(json_indet, json_string, "", "  ") //jsonを整形
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
		fmt.Fprint(w, json_indet.String())
		break
	case "DELETE":

		member := p_db.DB_select_id(parm)
		if len(member) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
			fmt.Fprint(w)
			break
		}
		p_db.DB_delete(parm)
		json_indet := new(bytes.Buffer)
		json_string, err := json.Marshal(&member[0])

		if err != nil { //err
			log.Fatal(err)
		}
		json.Indent(json_indet, json_string, "", "  ") //jsonを整形
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
		fmt.Fprint(w, json_indet.String())
		break
	default:
		member := p_db.DB_select_id(parm)
		if len(member) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
			fmt.Fprint(w)
			break
		}
		json_indet := new(bytes.Buffer)
		json_string, err := json.Marshal(&member[0])

		if err != nil { //err
			log.Fatal(err)
		}

		json.Indent(json_indet, json_string, "", "  ") //jsonを整形
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
		fmt.Fprint(w, json_indet.String())
		break
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/users/{id}", user_handler) //GET
	r.HandleFunc("/users", users_handler).Methods("GET")
	r.HandleFunc("/users", usercreate_handler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil) //ポート8080で待機
}
