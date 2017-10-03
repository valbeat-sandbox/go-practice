package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello world") // トップ
}

func UserHandler(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close() // 処理の最後にBodyを閉じる

	// GETされたらパラメータ取得
	if request.Method == "GET" {
		// パラメータ取得
		id, err := strconv.Atoi(request.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		// ファイル名はuser_{id}.txt
		fileName := fmt.Sprintf("user_%d.txt", id)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		user := &User{
			ID:   id,
			Name: string(data),
		}

		response, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			// コンテンツタイプをJSONに
			writer.Header().Set("Content-type", "application/json")
			// レスポンスとしてステータスコードを返す
			writer.WriteHeader(http.StatusCreated)
			// レスポンス返す
			fmt.Fprint(writer, string(response)) // トップ
		}()
	}

	// POSTされたらデータをテキストに保存
	if request.Method == "POST" {
		// リクエストボディをJSONに変換
		var user User
		decoder := json.NewDecoder(request.Body)
		// デコードの返り値はエラー
		err := decoder.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		// ファイル名はuser_{id}.txt
		fileName := fmt.Sprintf("user_%d.txt", user.ID)
		// ファイル作成 fileは*os.File構造体へのポインタ
		file, err := os.Create(fileName)
		if err != nil { // エラー処理
			log.Fatal(err)
		}
		// 必ずファイルを閉じる
		defer file.Close()

		// ファイルにNameを書き込む
		// WriteStringならbyte[]にキャストしなくても良い
		_, err = file.WriteString(user.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコードを返す
		writer.WriteHeader(http.StatusCreated)
	}
}

func main() {
	// ルーティング
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/users", UserHandler)
	// サーバー起動
	http.ListenAndServe(":3000", nil)
}
