package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		// 必ずファイルを閉じる
		defer file.Close()

		// ファイルにNameを書き込む
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
