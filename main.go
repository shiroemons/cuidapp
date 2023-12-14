package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucsky/cuid"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// CUIDを生成
		cuidValue := cuid.New()

		// リクエストのContent-Typeに応じた処理
		if r.Header.Get("Content-Type") == "application/json" {
			// JSON形式でレスポンスを返す
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"cuid": cuidValue})
		} else {
			// プレーンテキストでレスポンスを返す
			fmt.Fprintln(w, cuidValue)
		}
	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
