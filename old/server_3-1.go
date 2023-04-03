// ライブラリ以外は必ずpackage mainから始まる
package main

// 必要なパッケージの読み込み
// 宣言したパッケージのみ使用可能
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// まずmainパッケージが呼ばれる
func main() {
	// resp = http.Responseオブジェクト
	// サーバーからの情報を全て格納
	//
	resp, err := http.Get("http://localhost:18888")
	// エラー処理
	// 関数は戻り値としてエラーを返す
	// nil:ゼロ値を表す
	if err != nil {
		// panic: エラーを表示して処理を終了
		panic(err)
	}
	// 後処理
	// defer: 関数から抜けた後に実行する
	// ソケットからボディを読み込んだ後の処理
	// finallyのようなもの？
	defer resp.Body.Close()
	// ステータスコード
	defer fmt.Println(resp.Status)
	defer fmt.Println(resp.StatusCode)
	// レスポンスヘッダー
	// mapオブジェクト
	defer fmt.Println(resp.Header)
	// 特定項目の最初の要素のみをGETする
	defer log.Println(resp.Header.Get("Content-Length"))

	// 文字列をバイト列として読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// UTF-8に変換して画面に出力
	log.Println(string(body))
}

// // エラーチェックなしの最小構成
// func main() {
// 	resp, _ := http.Get("http://localhost:18888")
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	log.Println(string(body))
// }
