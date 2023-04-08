package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// サーバープッシュ
	// http.Pusherにキャスト
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/style.css", nil)
	}
	// 通常のイベントハンドラの処理
}

func main() {
	// GoogleのHTTPプロトコルバージョンの確認
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Protocol Version: %s\n", resp.Proto)
}
