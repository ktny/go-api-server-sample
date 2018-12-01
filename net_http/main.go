package main

import (
	"fmt"
	"net/http"
)

func main() {
	// リクエスト例
	// curl -d "key1=value1" "localhost:8080/hello/get?key2=value2"
	http.HandleFunc("/hello/get", getHello)

	// リクエスト例
	// curl -X POST "localhost:8080/hello/post"
	http.HandleFunc("/hello/post", postHello)

	http.ListenAndServe(":8080", nil)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	r.ParseForm()
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)

	params := r.Form
	fmt.Fprintf(w, "フォーム2：\n%v\n", params)
}

func postHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POSTのみ許可"))
		return
	}
	fmt.Fprintln(w, "hello world post")
}
