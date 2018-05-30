package handler

import (
	"net/http"
	"encoding/json"
	"log"
	"coding_challenge_2/WebApp/p_db"
	"bytes"
	"fmt"
)

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

