package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"log"
)

type Hello struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil) //ポート8080で待機
}

//ハンドラ関数定義
func handler(w http.ResponseWriter, r *http.Request) {

	msg := Hello{"Hello World!!"}
	json_indet := new(bytes.Buffer)
	json_string, err := json.Marshal(&msg)

	if err != nil {//エラー処理
		log.Fatal(err)
	}

	json.Indent(json_indet, json_string, "", "  ")     //jsonを整形
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, json_indet.String())
}