package main

import (
	"log"
	"net/http"
	// 組み込みライブラリとして実装されているクッキー機能
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	// cookieを保存する
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	// 保存可能なhttp.Clientインスタンスの作成
	client := http.Client{
		Jar: jar,
	}
	// 初回アクセスでクッキーを受信、2回目以降のアクセスでクッキーをサーバーに対して送信する
	for i := 0; i < 2; i++ {
		// http.GETの代わりにインスタンスのGetメソッドを利用
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}
