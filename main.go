package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ログを画面（標準出力）に出す
		fmt.Printf("[%s] リクエスト受信：\n", time.Now().Format(time.RFC3339))
		// ブラウザにレスポンスを返す
		fmt.Fprintf(w, "Hello, Docker! This is a simple Go server.")
	})

	fmt.Println("サーバーを8080番ポートで起動しています...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error：", err)
	}
}
