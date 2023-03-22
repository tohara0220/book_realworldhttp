package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// クライアントやcurlからアクセスがあったときに呼ばれる
// クライアントからのリクエストを受け取り、サーバーの処理結果を返す
func handler(w http.ResponseWriter, r *http.Request) {
	// 特に処理はせず、リクエストの内容をテキストで画面表示
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

// こちらが先に実行される
// HTTPサーバーを初期化する
func main() {
	var httpServer http.Server
	// トップのパスにアクセスがあったら"handler"関数を呼ぶ
	http.HandleFunc("/", handler)
	// 18888ポートで起動
	// HTTPのデフォルトポートは80だが、システムで使用中だったり一般ユーザー権限ではアクセスできないことがある
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
