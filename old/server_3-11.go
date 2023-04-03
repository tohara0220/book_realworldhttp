package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

// multipart/form-data形式でファイルの送信
func main() {
	// バイト列を格納するバッファを宣言
	var buffer bytes.Buffer
	// マルチパートを組み立てるライターを作る（これに格納していく）
	writer := multipart.NewWriter(&buffer)

	// 送信するファイルに任意のMIMEタイプを設定する
	// ファイルを開く
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		// ファイル読み込み失敗
		panic(err)
	}
	defer readFile.Close()

	// ファイル以外はWriteFieldを使って登録する
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)

	// // multipart/form-data形式でファイル送信
	// fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	// // ファイルを開く
	// readFile, err := os.Open("photo.jpg")
	// if err != nil {
	// 	// ファイル読み込み失敗
	// 	panic(err)
	// }
	// defer readFile.Close()

	// // ファイルの全コンテンツを、ファイル書き込み用のio.Writerにコピー
	// io.Copy(fileWriter, readFile)

	// マルチパートのio.Writerをクローズ => バッファにすべてを書き込み
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}

// func main() {
// 	// // os.Fileオブジェクトはio.Readerインターフェースを満たしている
// 	// // そのため、http.Post関数にそのまま渡せる
// 	// file, err := os.Open("main.go")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// // Content-Typeヘッダーの内容はhttp.Postメソッドの二つ目の引数
// 	// resp, err := http.Post("http://localhost:18888", "text/plain", file)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// プログラム内で生成したテキストはbytes.Buffer, string.Readerを使ってio.Readerする
// 	reader := strings.NewReader("テキスト")
// 	// Content-Typeヘッダーの内容はhttp.Postメソッドの二つ目の引数
// 	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
// 	if err != nil {
// 		panic(err)
// 	}

// 	log.Println("Status:", resp.Status)
// }

// // x-www-form-urlencoded形式のPOSTメソッドの送信
// func main() {
// 	values := url.Values{
// 		"test": {"value"},
// 	}
// 	// PostForm関数に渡す
// 	resp, err := http.PostForm("http://localhost:18888", values)
// 	if err != nil {
// 		// 送信失敗
// 		panic(err)
// 	}
// 	log.Println("Status: ", resp.Status)
// }

// // HEADメソッドでヘッダーを取得する
// // curl --head http://localhost:18888
// // HEADはBODYを取得しない
// func main() {
// 	resp, err := http.Head("http://localhost:18888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println("Status: ", resp.Status)
// 	log.Println("Headers: ", resp.Header)
// }

// // GETメソッドで送信する
// func main() {
// 	// クエリ文字列
// 	values := url.Values{
// 		"query": {"hello world"},
// 	}
// 	// クエリを文字列として追加する
// 	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
// 	if resp != nil {
// 		panic("error!")
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	log.Println(string(body))
// }
